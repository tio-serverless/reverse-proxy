// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy.proto

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

type ProxyEndpointRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyEndpointRequest) Reset()         { *m = ProxyEndpointRequest{} }
func (m *ProxyEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*ProxyEndpointRequest) ProtoMessage()    {}
func (*ProxyEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}

func (m *ProxyEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyEndpointRequest.Unmarshal(m, b)
}
func (m *ProxyEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyEndpointRequest.Marshal(b, m, deterministic)
}
func (m *ProxyEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyEndpointRequest.Merge(m, src)
}
func (m *ProxyEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_ProxyEndpointRequest.Size(m)
}
func (m *ProxyEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyEndpointRequest proto.InternalMessageInfo

func (m *ProxyEndpointRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProxyEndpointRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func init() {
	proto.RegisterType((*ProxyEndpointRequest)(nil), "ProxyEndpointRequest")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor_700b50b08ed8dbaf) }

var fileDescriptor_700b50b08ed8dbaf = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0xaf,
	0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x4e, 0xce, 0xcf, 0x2b, 0x2e, 0x81, 0x70,
	0x94, 0xdc, 0xb8, 0x44, 0x02, 0x40, 0x72, 0xae, 0x79, 0x29, 0x05, 0xf9, 0x99, 0x79, 0x25, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x14, 0x17, 0x47, 0x2a, 0x54, 0x99, 0x04, 0x13,
	0x58, 0x1c, 0xce, 0x37, 0x2a, 0xe6, 0xe2, 0x01, 0x9b, 0x13, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c,
	0x2a, 0x64, 0xc8, 0xc5, 0xed, 0x97, 0x5a, 0x0e, 0x33, 0x55, 0x48, 0x54, 0x0f, 0x9b, 0x2d, 0x52,
	0x9c, 0x7a, 0x21, 0x99, 0xf9, 0x41, 0xa9, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x42, 0x06, 0x5c, 0x5c,
	0x9e, 0x79, 0x59, 0xa9, 0xc9, 0x25, 0xc1, 0x95, 0x79, 0xc9, 0xc4, 0xe8, 0x70, 0x12, 0x88, 0xe2,
	0x2b, 0xc9, 0xcc, 0xd7, 0x4b, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x33, 0x4c, 0x62,
	0x03, 0xfb, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x3f, 0x7d, 0x3f, 0xf1, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyServiceClient interface {
	//     NewEndpoint Monitor通过此接口通知服务创建完成的通知
	NewEndpoint(ctx context.Context, in *ProxyEndpointRequest, opts ...grpc.CallOption) (*TioReply, error)
	//    InjectSync 通知Proxy更新缓存
	InjectSync(ctx context.Context, in *ProxyEndpointRequest, opts ...grpc.CallOption) (*TioReply, error)
}

type proxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewProxyServiceClient(cc *grpc.ClientConn) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) NewEndpoint(ctx context.Context, in *ProxyEndpointRequest, opts ...grpc.CallOption) (*TioReply, error) {
	out := new(TioReply)
	err := c.cc.Invoke(ctx, "/ProxyService/NewEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) InjectSync(ctx context.Context, in *ProxyEndpointRequest, opts ...grpc.CallOption) (*TioReply, error) {
	out := new(TioReply)
	err := c.cc.Invoke(ctx, "/ProxyService/InjectSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServiceServer is the server API for ProxyService service.
type ProxyServiceServer interface {
	//     NewEndpoint Monitor通过此接口通知服务创建完成的通知
	NewEndpoint(context.Context, *ProxyEndpointRequest) (*TioReply, error)
	//    InjectSync 通知Proxy更新缓存
	InjectSync(context.Context, *ProxyEndpointRequest) (*TioReply, error)
}

// UnimplementedProxyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProxyServiceServer struct {
}

func (*UnimplementedProxyServiceServer) NewEndpoint(ctx context.Context, req *ProxyEndpointRequest) (*TioReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewEndpoint not implemented")
}
func (*UnimplementedProxyServiceServer) InjectSync(ctx context.Context, req *ProxyEndpointRequest) (*TioReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InjectSync not implemented")
}

func RegisterProxyServiceServer(s *grpc.Server, srv ProxyServiceServer) {
	s.RegisterService(&_ProxyService_serviceDesc, srv)
}

func _ProxyService_NewEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).NewEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProxyService/NewEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).NewEndpoint(ctx, req.(*ProxyEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_InjectSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).InjectSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProxyService/InjectSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).InjectSync(ctx, req.(*ProxyEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewEndpoint",
			Handler:    _ProxyService_NewEndpoint_Handler,
		},
		{
			MethodName: "InjectSync",
			Handler:    _ProxyService_InjectSync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy.proto",
}
