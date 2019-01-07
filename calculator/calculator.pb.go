// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CalculatorRequest struct {
	A                    float64  `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    float64  `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetA() float64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *CalculatorRequest) GetB() float64 {
	if m != nil {
		return m.B
	}
	return 0
}

type CalculatorResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Result               float64  `protobuf:"fixed64,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CalculatorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CalculatorResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xe9, 0x73, 0x09, 0x3a, 0xc3, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3c,
	0x5c, 0x8c, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x8c, 0x41, 0x8c, 0x89, 0x20, 0x5e, 0x92, 0x04,
	0x13, 0x84, 0x97, 0xa4, 0x14, 0xc3, 0x25, 0x84, 0xac, 0xa1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x15, 0xac, 0x8f, 0x33, 0x08, 0xc6,
	0x15, 0x12, 0xe1, 0x62, 0x4d, 0x2d, 0x2a, 0xca, 0x2f, 0x02, 0x9b, 0xc0, 0x19, 0x04, 0xe1, 0x08,
	0x89, 0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x48, 0x30, 0x83, 0x0d, 0x86, 0xf2, 0x8c,
	0x66, 0x30, 0x72, 0x71, 0x21, 0x8c, 0x17, 0x72, 0xe1, 0x62, 0x0e, 0x2e, 0xcd, 0x15, 0x92, 0xd5,
	0x43, 0xf2, 0x03, 0x86, 0x73, 0xa5, 0xe4, 0x70, 0x49, 0x43, 0x1d, 0xe7, 0xc2, 0xc5, 0xec, 0x92,
	0x59, 0x46, 0xa1, 0x29, 0x49, 0x6c, 0xe0, 0xc0, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0x30, 0x1c, 0x8f, 0x50, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	Div(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Div(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	Div(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sum(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Div(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Calculator_Sum_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _Calculator_Div_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
