package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var si *svcImplement

func main() {
	setupLog()

	si, err := newSI()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Service Init Finish! ")
	go start(si, 8000)
	setupRoute(si, si.routeChan)
}

func setupLog() {
	switch strings.ToLower("TIO_PROXY_LOG") {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func setupRoute(loader dataLoader, r chan string) {
	router := mux.NewRouter()
	go func(router *mux.Router, r chan string) {
		for {
			select {
			case u := <-r:
				logrus.Infof("Register New Path %s", u)
				router.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
					logrus.Debugf("New Root Request %s", r.URL.String())
					Proxy(loader, w, r)
				})
			}
		}
	}(router, r)
	// router.HandleFunc("/{url}", func(w http.ResponseWriter, r *http.Request) {
	// 	logrus.Debugf("New URL Request %s", r.URL.String())
	// 	Proxy(loader, w, r)
	// })

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.Debugf("New Root Request %s", r.URL.String())
		Proxy(loader, w, r)
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Debug("Server Start ---")
	logrus.Println(srv.ListenAndServe())
}
