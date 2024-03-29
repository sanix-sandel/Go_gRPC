// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package protofiles

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

type TransactionRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransactionRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransactionRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransactionResponse struct {
	Confirmation         bool     `protobuf:"varint,1,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetConfirmation() bool {
	if m != nil {
		return m.Confirmation
	}
	return false
}

func init() {
	proto.RegisterType((*TransactionRequest)(nil), "protofiles.TransactionRequest")
	proto.RegisterType((*TransactionResponse)(nil), "protofiles.TransactionResponse")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02,
	0x53, 0x69, 0x99, 0x39, 0xa9, 0xc5, 0x4a, 0x01, 0x5c, 0x42, 0x21, 0x08, 0x05, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x69, 0x45, 0xf9, 0xb9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x10, 0x1f, 0x17, 0x53, 0x49, 0xbe, 0x04, 0x13, 0x58, 0x84, 0xa9,
	0x24, 0x5f, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x59, 0x81, 0x51,
	0x83, 0x29, 0x08, 0xca, 0x53, 0xb2, 0xe4, 0x12, 0x46, 0x31, 0xb1, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x55, 0x48, 0x89, 0x8b, 0x27, 0x39, 0x3f, 0x2f, 0x2d, 0xb3, 0x28, 0x37, 0x11, 0x24, 0x0e, 0x36,
	0x9a, 0x23, 0x08, 0x45, 0xcc, 0x28, 0x83, 0x4b, 0xc0, 0x37, 0x3f, 0x2f, 0xb5, 0x12, 0x49, 0xbf,
	0x50, 0x08, 0x17, 0xbf, 0x6f, 0x62, 0x76, 0x2a, 0xb2, 0x90, 0x9c, 0x1e, 0xc2, 0x03, 0x7a, 0x98,
	0xae, 0x97, 0x92, 0xc7, 0x29, 0x0f, 0x71, 0x8b, 0x12, 0x43, 0x12, 0x1b, 0x58, 0x85, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x86, 0x5c, 0xac, 0x0f, 0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoneyTransactionClient is the client API for MoneyTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoneyTransactionClient interface {
	MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type moneyTransactionClient struct {
	cc *grpc.ClientConn
}

func NewMoneyTransactionClient(cc *grpc.ClientConn) MoneyTransactionClient {
	return &moneyTransactionClient{cc}
}

func (c *moneyTransactionClient) MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/protofiles.MoneyTransaction/MakeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoneyTransactionServer is the server API for MoneyTransaction service.
type MoneyTransactionServer interface {
	MakeTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
}

// UnimplementedMoneyTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedMoneyTransactionServer struct {
}

func (*UnimplementedMoneyTransactionServer) MakeTransaction(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeTransaction not implemented")
}

func RegisterMoneyTransactionServer(s *grpc.Server, srv MoneyTransactionServer) {
	s.RegisterService(&_MoneyTransaction_serviceDesc, srv)
}

func _MoneyTransaction_MakeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyTransactionServer).MakeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.MoneyTransaction/MakeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyTransactionServer).MakeTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoneyTransaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.MoneyTransaction",
	HandlerType: (*MoneyTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeTransaction",
			Handler:    _MoneyTransaction_MakeTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
