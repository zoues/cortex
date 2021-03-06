// Code generated by protoc-gen-go.
// source: httpgrpc.proto
// DO NOT EDIT!

/*
Package httpgrpc is a generated protocol buffer package.

It is generated from these files:
	httpgrpc.proto

It has these top-level messages:
	HTTPRequest
	HTTPResponse
	Header
*/
package httpgrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HTTPRequest struct {
	Method  string    `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Url     string    `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Headers []*Header `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty"`
	Body    []byte    `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HTTPRequest) Reset()                    { *m = HTTPRequest{} }
func (m *HTTPRequest) String() string            { return proto.CompactTextString(m) }
func (*HTTPRequest) ProtoMessage()               {}
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HTTPRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HTTPRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HTTPRequest) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type HTTPResponse struct {
	Code    int32     `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
	Headers []*Header `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
	Body    []byte    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HTTPResponse) Reset()                    { *m = HTTPResponse{} }
func (m *HTTPResponse) String() string            { return proto.CompactTextString(m) }
func (*HTTPResponse) ProtoMessage()               {}
func (*HTTPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HTTPResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HTTPResponse) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Header struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Header) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Header) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*HTTPRequest)(nil), "httpgrpc.HTTPRequest")
	proto.RegisterType((*HTTPResponse)(nil), "httpgrpc.HTTPResponse")
	proto.RegisterType((*Header)(nil), "httpgrpc.Header")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HTTP service

type HTTPClient interface {
	Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error)
}

type hTTPClient struct {
	cc *grpc.ClientConn
}

func NewHTTPClient(cc *grpc.ClientConn) HTTPClient {
	return &hTTPClient{cc}
}

func (c *hTTPClient) Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error) {
	out := new(HTTPResponse)
	err := grpc.Invoke(ctx, "/httpgrpc.HTTP/Handle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HTTP service

type HTTPServer interface {
	Handle(context.Context, *HTTPRequest) (*HTTPResponse, error)
}

func RegisterHTTPServer(s *grpc.Server, srv HTTPServer) {
	s.RegisterService(&_HTTP_serviceDesc, srv)
}

func _HTTP_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httpgrpc.HTTP/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPServer).Handle(ctx, req.(*HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "httpgrpc.HTTP",
	HandlerType: (*HTTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _HTTP_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httpgrpc.proto",
}

func init() { proto.RegisterFile("httpgrpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0x1d, 0x0c, 0xbd, 0x56, 0xa8, 0x3a, 0x89, 0xca, 0x62, 0x8a, 0x32, 0x45, 0x0c,
	0x1d, 0xc2, 0xc4, 0x88, 0x58, 0x32, 0x22, 0xab, 0x7f, 0x20, 0xc1, 0x27, 0x22, 0x11, 0x6a, 0x63,
	0x3b, 0xa0, 0xfe, 0x7b, 0x64, 0x3b, 0x85, 0x88, 0xa9, 0xdb, 0x7b, 0xe7, 0x27, 0x7f, 0xf7, 0x0e,
	0x6e, 0x7a, 0xef, 0xcd, 0x9b, 0x35, 0xaf, 0x3b, 0x63, 0xb5, 0xd7, 0x78, 0x7d, 0xf2, 0xe5, 0x37,
	0xac, 0x9a, 0xfd, 0xfe, 0x45, 0xd2, 0xe7, 0x48, 0xce, 0xe3, 0x16, 0xf8, 0x07, 0xf9, 0x5e, 0x2b,
	0x91, 0x15, 0x59, 0xb5, 0x94, 0x93, 0xc3, 0x0d, 0xb0, 0xd1, 0x0e, 0x62, 0x11, 0x87, 0x41, 0xe2,
	0x3d, 0x5c, 0xf5, 0xd4, 0x2a, 0xb2, 0x4e, 0xb0, 0x82, 0x55, 0xab, 0x7a, 0xb3, 0xfb, 0x85, 0x34,
	0xf1, 0x41, 0x9e, 0x02, 0x88, 0x90, 0x77, 0x5a, 0x1d, 0x45, 0x5e, 0x64, 0xd5, 0x5a, 0x46, 0x5d,
	0x76, 0xb0, 0x4e, 0x60, 0x67, 0xf4, 0xc1, 0x51, 0xc8, 0x3c, 0x6b, 0x45, 0x91, 0x7b, 0x29, 0xa3,
	0x9e, 0x33, 0x16, 0xe7, 0x32, 0xd8, 0x8c, 0x51, 0x03, 0x4f, 0xb1, 0xb0, 0xff, 0x3b, 0x1d, 0xa7,
	0x52, 0x41, 0x86, 0xa6, 0x5f, 0xed, 0x30, 0x52, 0xfa, 0x7a, 0x29, 0x27, 0x57, 0x3f, 0x41, 0x1e,
	0xf6, 0xc2, 0x47, 0xe0, 0x4d, 0x7b, 0x50, 0x03, 0xe1, 0xed, 0x0c, 0xfa, 0x77, 0xaa, 0xbb, 0xed,
	0xff, 0x71, 0x2a, 0x52, 0x5e, 0x74, 0x3c, 0x1e, 0xf9, 0xe1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x47,
	0x4e, 0x55, 0x95, 0x76, 0x01, 0x00, 0x00,
}
