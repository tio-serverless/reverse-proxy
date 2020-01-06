// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

package tio_control_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type JobStatus int32

const (
	JobStatus_Ready        JobStatus = 0
	JobStatus_BuildSucc    JobStatus = 1
	JobStatus_BuildFailed  JobStatus = 2
	JobStatus_BuildIng     JobStatus = 3
	JobStatus_DeployIng    JobStatus = 4
	JobStatus_DeploySuc    JobStatus = 5
	JobStatus_DeployFailed JobStatus = 6
)

var JobStatus_name = map[int32]string{
	0: "Ready",
	1: "BuildSucc",
	2: "BuildFailed",
	3: "BuildIng",
	4: "DeployIng",
	5: "DeploySuc",
	6: "DeployFailed",
}

var JobStatus_value = map[string]int32{
	"Ready":        0,
	"BuildSucc":    1,
	"BuildFailed":  2,
	"BuildIng":     3,
	"DeployIng":    4,
	"DeploySuc":    5,
	"DeployFailed": 6,
}

func (x JobStatus) String() string {
	return proto.EnumName(JobStatus_name, int32(x))
}

func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

type CommonRespCode int32

const (
	CommonRespCode_RespSucc  CommonRespCode = 0
	CommonRespCode_RespFaild CommonRespCode = -1
)

var CommonRespCode_name = map[int32]string{
	0:  "RespSucc",
	-1: "RespFaild",
}

var CommonRespCode_value = map[string]int32{
	"RespSucc":  0,
	"RespFaild": -1,
}

func (x CommonRespCode) String() string {
	return proto.EnumName(CommonRespCode_name, int32(x))
}

func (CommonRespCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

type TioReply struct {
	Code                 CommonRespCode `protobuf:"varint,1,opt,name=code,proto3,enum=CommonRespCode" json:"code,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TioReply) Reset()         { *m = TioReply{} }
func (m *TioReply) String() string { return proto.CompactTextString(m) }
func (*TioReply) ProtoMessage()    {}
func (*TioReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

func (m *TioReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioReply.Unmarshal(m, b)
}
func (m *TioReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioReply.Marshal(b, m, deterministic)
}
func (m *TioReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioReply.Merge(m, src)
}
func (m *TioReply) XXX_Size() int {
	return xxx_messageInfo_TioReply.Size(m)
}
func (m *TioReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioReply proto.InternalMessageInfo

func (m *TioReply) GetCode() CommonRespCode {
	if m != nil {
		return m.Code
	}
	return CommonRespCode_RespSucc
}

func (m *TioReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TioUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioUserRequest) Reset()         { *m = TioUserRequest{} }
func (m *TioUserRequest) String() string { return proto.CompactTextString(m) }
func (*TioUserRequest) ProtoMessage()    {}
func (*TioUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

func (m *TioUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioUserRequest.Unmarshal(m, b)
}
func (m *TioUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioUserRequest.Marshal(b, m, deterministic)
}
func (m *TioUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioUserRequest.Merge(m, src)
}
func (m *TioUserRequest) XXX_Size() int {
	return xxx_messageInfo_TioUserRequest.Size(m)
}
func (m *TioUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TioUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TioUserRequest proto.InternalMessageInfo

func (m *TioUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TioUserRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type TioUserReply struct {
	Code                 CommonRespCode `protobuf:"varint,1,opt,name=Code,proto3,enum=CommonRespCode" json:"Code,omitempty"`
	Token                *TioToken      `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	User                 *TioUserInfo   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TioUserReply) Reset()         { *m = TioUserReply{} }
func (m *TioUserReply) String() string { return proto.CompactTextString(m) }
func (*TioUserReply) ProtoMessage()    {}
func (*TioUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{2}
}

func (m *TioUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioUserReply.Unmarshal(m, b)
}
func (m *TioUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioUserReply.Marshal(b, m, deterministic)
}
func (m *TioUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioUserReply.Merge(m, src)
}
func (m *TioUserReply) XXX_Size() int {
	return xxx_messageInfo_TioUserReply.Size(m)
}
func (m *TioUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioUserReply proto.InternalMessageInfo

func (m *TioUserReply) GetCode() CommonRespCode {
	if m != nil {
		return m.Code
	}
	return CommonRespCode_RespSucc
}

func (m *TioUserReply) GetToken() *TioToken {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *TioUserReply) GetUser() *TioUserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type TioUserInfo struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioUserInfo) Reset()         { *m = TioUserInfo{} }
func (m *TioUserInfo) String() string { return proto.CompactTextString(m) }
func (*TioUserInfo) ProtoMessage()    {}
func (*TioUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{3}
}

func (m *TioUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioUserInfo.Unmarshal(m, b)
}
func (m *TioUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioUserInfo.Marshal(b, m, deterministic)
}
func (m *TioUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioUserInfo.Merge(m, src)
}
func (m *TioUserInfo) XXX_Size() int {
	return xxx_messageInfo_TioUserInfo.Size(m)
}
func (m *TioUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TioUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TioUserInfo proto.InternalMessageInfo

func (m *TioUserInfo) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type TioToken struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	Bucket               string   `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	CallBackUrl          string   `protobuf:"bytes,4,opt,name=callBackUrl,proto3" json:"callBackUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioToken) Reset()         { *m = TioToken{} }
func (m *TioToken) String() string { return proto.CompactTextString(m) }
func (*TioToken) ProtoMessage()    {}
func (*TioToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{4}
}

func (m *TioToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioToken.Unmarshal(m, b)
}
func (m *TioToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioToken.Marshal(b, m, deterministic)
}
func (m *TioToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioToken.Merge(m, src)
}
func (m *TioToken) XXX_Size() int {
	return xxx_messageInfo_TioToken.Size(m)
}
func (m *TioToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TioToken.DiscardUnknown(m)
}

var xxx_messageInfo_TioToken proto.InternalMessageInfo

func (m *TioToken) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *TioToken) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *TioToken) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *TioToken) GetCallBackUrl() string {
	if m != nil {
		return m.CallBackUrl
	}
	return ""
}

type TioAgentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioAgentRequest) Reset()         { *m = TioAgentRequest{} }
func (m *TioAgentRequest) String() string { return proto.CompactTextString(m) }
func (*TioAgentRequest) ProtoMessage()    {}
func (*TioAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{5}
}

func (m *TioAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioAgentRequest.Unmarshal(m, b)
}
func (m *TioAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioAgentRequest.Marshal(b, m, deterministic)
}
func (m *TioAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioAgentRequest.Merge(m, src)
}
func (m *TioAgentRequest) XXX_Size() int {
	return xxx_messageInfo_TioAgentRequest.Size(m)
}
func (m *TioAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TioAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TioAgentRequest proto.InternalMessageInfo

func (m *TioAgentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TioAgentReply struct {
	Code                 CommonRespCode `protobuf:"varint,1,opt,name=Code,proto3,enum=CommonRespCode" json:"Code,omitempty"`
	Address              string         `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TioAgentReply) Reset()         { *m = TioAgentReply{} }
func (m *TioAgentReply) String() string { return proto.CompactTextString(m) }
func (*TioAgentReply) ProtoMessage()    {}
func (*TioAgentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{6}
}

func (m *TioAgentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioAgentReply.Unmarshal(m, b)
}
func (m *TioAgentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioAgentReply.Marshal(b, m, deterministic)
}
func (m *TioAgentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioAgentReply.Merge(m, src)
}
func (m *TioAgentReply) XXX_Size() int {
	return xxx_messageInfo_TioAgentReply.Size(m)
}
func (m *TioAgentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioAgentReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioAgentReply proto.InternalMessageInfo

func (m *TioAgentReply) GetCode() CommonRespCode {
	if m != nil {
		return m.Code
	}
	return CommonRespCode_RespSucc
}

func (m *TioAgentReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type TioLogRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Flowing              bool     `protobuf:"varint,2,opt,name=flowing,proto3" json:"flowing,omitempty"`
	Stype                string   `protobuf:"bytes,3,opt,name=stype,proto3" json:"stype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioLogRequest) Reset()         { *m = TioLogRequest{} }
func (m *TioLogRequest) String() string { return proto.CompactTextString(m) }
func (*TioLogRequest) ProtoMessage()    {}
func (*TioLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{7}
}

func (m *TioLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioLogRequest.Unmarshal(m, b)
}
func (m *TioLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioLogRequest.Marshal(b, m, deterministic)
}
func (m *TioLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioLogRequest.Merge(m, src)
}
func (m *TioLogRequest) XXX_Size() int {
	return xxx_messageInfo_TioLogRequest.Size(m)
}
func (m *TioLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TioLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TioLogRequest proto.InternalMessageInfo

func (m *TioLogRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TioLogRequest) GetFlowing() bool {
	if m != nil {
		return m.Flowing
	}
	return false
}

func (m *TioLogRequest) GetStype() string {
	if m != nil {
		return m.Stype
	}
	return ""
}

type TioLogReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioLogReply) Reset()         { *m = TioLogReply{} }
func (m *TioLogReply) String() string { return proto.CompactTextString(m) }
func (*TioLogReply) ProtoMessage()    {}
func (*TioLogReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{8}
}

func (m *TioLogReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioLogReply.Unmarshal(m, b)
}
func (m *TioLogReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioLogReply.Marshal(b, m, deterministic)
}
func (m *TioLogReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioLogReply.Merge(m, src)
}
func (m *TioLogReply) XXX_Size() int {
	return xxx_messageInfo_TioLogReply.Size(m)
}
func (m *TioLogReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioLogReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioLogReply proto.InternalMessageInfo

func (m *TioLogReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("JobStatus", JobStatus_name, JobStatus_value)
	proto.RegisterEnum("CommonRespCode", CommonRespCode_name, CommonRespCode_value)
	proto.RegisterType((*TioReply)(nil), "TioReply")
	proto.RegisterType((*TioUserRequest)(nil), "TioUserRequest")
	proto.RegisterType((*TioUserReply)(nil), "TioUserReply")
	proto.RegisterType((*TioUserInfo)(nil), "TioUserInfo")
	proto.RegisterType((*TioToken)(nil), "TioToken")
	proto.RegisterType((*TioAgentRequest)(nil), "TioAgentRequest")
	proto.RegisterType((*TioAgentReply)(nil), "TioAgentReply")
	proto.RegisterType((*TioLogRequest)(nil), "TioLogRequest")
	proto.RegisterType((*TioLogReply)(nil), "TioLogReply")
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0xe2, 0x4f, 0x8d, 0x1d, 0x5b, 0x2c, 0x25, 0x88, 0x52, 0x88, 0x51, 0x29, 0x0d, 0x29,
	0x88, 0xd6, 0x85, 0x42, 0xa1, 0x97, 0xd8, 0xa5, 0x25, 0xad, 0xe9, 0x61, 0xed, 0x5c, 0x7a, 0x93,
	0x57, 0x13, 0xb1, 0x58, 0xd6, 0xa8, 0xda, 0x95, 0x83, 0x6f, 0xfd, 0xe7, 0x2d, 0xbb, 0x92, 0x6a,
	0xe7, 0x62, 0xea, 0x83, 0x99, 0xf7, 0x66, 0xf7, 0xcd, 0xbc, 0x27, 0x09, 0x06, 0x82, 0x32, 0xa5,
	0xc3, 0xbc, 0x20, 0x4d, 0xc1, 0x2d, 0xf4, 0x57, 0x92, 0x38, 0xe6, 0xe9, 0x9e, 0xbd, 0x84, 0xb6,
	0xa0, 0x18, 0x7d, 0x67, 0xe2, 0x5c, 0x8f, 0xa6, 0xe3, 0x70, 0x4e, 0xdb, 0x2d, 0x65, 0x1c, 0x55,
	0x3e, 0xa7, 0x18, 0xb9, 0x6d, 0x32, 0x0f, 0x5a, 0x5b, 0x95, 0xf8, 0xe7, 0x13, 0xe7, 0xda, 0xe5,
	0xa6, 0x0c, 0x3e, 0xc1, 0x68, 0x25, 0xe9, 0x5e, 0x61, 0xc1, 0xf1, 0x57, 0x89, 0x4a, 0x33, 0x06,
	0xed, 0x2c, 0xda, 0x56, 0x42, 0x2e, 0xb7, 0x35, 0xbb, 0x84, 0x6e, 0x1e, 0x29, 0xf5, 0x18, 0xd7,
	0x57, 0x6b, 0x14, 0xec, 0x60, 0xf8, 0xef, 0x76, 0xbd, 0xc4, 0xfc, 0xd4, 0x12, 0xe6, 0x9f, 0x5d,
	0x41, 0x47, 0xd3, 0x06, 0x33, 0xab, 0x35, 0x98, 0xba, 0xe1, 0x4a, 0xd2, 0xca, 0x10, 0xbc, 0xe2,
	0xd9, 0x04, 0xda, 0xa5, 0xc2, 0xc2, 0x6f, 0xd9, 0xfe, 0x30, 0xac, 0x47, 0xdc, 0x65, 0x0f, 0xc4,
	0x6d, 0x27, 0xb8, 0x82, 0xc1, 0x11, 0x69, 0x6c, 0x95, 0x32, 0xb6, 0x53, 0x3b, 0xdc, 0x94, 0xc1,
	0x6f, 0xc7, 0x46, 0x63, 0x65, 0xd9, 0x0b, 0x70, 0x23, 0x21, 0x50, 0xa9, 0xef, 0xb8, 0xaf, 0x6d,
	0x1d, 0x08, 0xd3, 0x55, 0x28, 0x0a, 0xd4, 0xa6, 0x5b, 0xd9, 0x3b, 0x10, 0xc6, 0xf9, 0xba, 0x14,
	0x1b, 0xd4, 0x76, 0x1b, 0x97, 0xd7, 0x88, 0x4d, 0x60, 0x20, 0xa2, 0x34, 0x9d, 0x45, 0x62, 0x73,
	0x5f, 0xa4, 0x7e, 0xdb, 0x36, 0x8f, 0xa9, 0xe0, 0x15, 0x8c, 0x57, 0x92, 0x6e, 0x13, 0xcc, 0xf4,
	0x89, 0x68, 0x83, 0x1f, 0x70, 0x71, 0x38, 0xf6, 0xdf, 0x19, 0xfa, 0xd0, 0x8b, 0xe2, 0xb8, 0x40,
	0xa5, 0xea, 0x95, 0x1b, 0x18, 0x2c, 0xad, 0xde, 0x82, 0x92, 0x53, 0xcf, 0xd3, 0x87, 0xde, 0x43,
	0x4a, 0x8f, 0x32, 0xab, 0xde, 0x85, 0x3e, 0x6f, 0x20, 0x7b, 0x06, 0x1d, 0xa5, 0xf7, 0x39, 0xd6,
	0x76, 0x2b, 0x10, 0xbc, 0xb6, 0x79, 0x5b, 0x51, 0xb3, 0xa2, 0x0f, 0xbd, 0x2d, 0x2a, 0x15, 0x25,
	0x8d, 0x6a, 0x03, 0x6f, 0x34, 0xb8, 0xdf, 0x68, 0xbd, 0xd4, 0x91, 0x2e, 0x15, 0x73, 0xa1, 0xc3,
	0x31, 0x8a, 0xf7, 0xde, 0x19, 0xbb, 0x00, 0x77, 0x56, 0xca, 0x34, 0x5e, 0x96, 0x42, 0x78, 0x0e,
	0x1b, 0xc3, 0xc0, 0xc2, 0x2f, 0x91, 0x4c, 0x31, 0xf6, 0xce, 0xd9, 0x10, 0xfa, 0x96, 0xb8, 0xcb,
	0x12, 0xaf, 0x65, 0x4e, 0x7f, 0xc6, 0x3c, 0xa5, 0xbd, 0x81, 0xed, 0x03, 0x5c, 0x96, 0xc2, 0xeb,
	0x30, 0x0f, 0x86, 0x15, 0xac, 0x6f, 0x77, 0x6f, 0x3e, 0xc0, 0xe8, 0x69, 0x4a, 0x46, 0xcf, 0xd4,
	0x76, 0xdc, 0x19, 0xbb, 0x04, 0xd7, 0x20, 0x73, 0x3e, 0xf6, 0xfe, 0x34, 0x3f, 0x67, 0xfa, 0x11,
	0x60, 0x41, 0xc9, 0x12, 0x8b, 0x9d, 0x14, 0xc8, 0xde, 0x40, 0xef, 0x2b, 0xea, 0x05, 0x25, 0x8a,
	0x8d, 0xc2, 0x27, 0x19, 0x3e, 0x1f, 0x86, 0x47, 0xf6, 0x83, 0xb3, 0xb7, 0xce, 0xcc, 0xfb, 0x39,
	0xd2, 0x92, 0x42, 0x41, 0x99, 0x2e, 0x28, 0x0d, 0x77, 0xef, 0xd6, 0x5d, 0xfb, 0x4d, 0xbe, 0xff,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x42, 0x81, 0xbd, 0xa2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogServiceClient interface {
	//    GetLogs 代理客户端读取其它服务的日志
	GetLogs(ctx context.Context, in *TioLogRequest, opts ...grpc.CallOption) (LogService_GetLogsClient, error)
}

type logServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogServiceClient(cc *grpc.ClientConn) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) GetLogs(ctx context.Context, in *TioLogRequest, opts ...grpc.CallOption) (LogService_GetLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogService_serviceDesc.Streams[0], "/LogService/GetLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceGetLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_GetLogsClient interface {
	Recv() (*TioLogReply, error)
	grpc.ClientStream
}

type logServiceGetLogsClient struct {
	grpc.ClientStream
}

func (x *logServiceGetLogsClient) Recv() (*TioLogReply, error) {
	m := new(TioLogReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
type LogServiceServer interface {
	//    GetLogs 代理客户端读取其它服务的日志
	GetLogs(*TioLogRequest, LogService_GetLogsServer) error
}

// UnimplementedLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (*UnimplementedLogServiceServer) GetLogs(req *TioLogRequest, srv LogService_GetLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TioLogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).GetLogs(m, &logServiceGetLogsServer{stream})
}

type LogService_GetLogsServer interface {
	Send(*TioLogReply) error
	grpc.ServerStream
}

type logServiceGetLogsServer struct {
	grpc.ServerStream
}

func (x *logServiceGetLogsServer) Send(m *TioLogReply) error {
	return x.ServerStream.SendMsg(m)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogs",
			Handler:       _LogService_GetLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "const.proto",
}