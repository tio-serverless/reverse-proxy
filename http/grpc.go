package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	tio_control_v1 "github.com/tio-serverless/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func start(si *svcImplement, port int) {

	p := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", p)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	tio_control_v1.RegisterProxyServiceServer(s, si)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
