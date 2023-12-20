// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/pb/product.proto

package pb

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

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stock                int64    `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	Price                int64    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c53c2f7d36d479c, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetStock() int64 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *CreateRequest) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type CreateResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c53c2f7d36d479c, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FindOneData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Stock                int64    `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Price                int64    `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindOneData) Reset()         { *m = FindOneData{} }
func (m *FindOneData) String() string { return proto.CompactTextString(m) }
func (*FindOneData) ProtoMessage()    {}
func (*FindOneData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c53c2f7d36d479c, []int{2}
}

func (m *FindOneData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneData.Unmarshal(m, b)
}
func (m *FindOneData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneData.Marshal(b, m, deterministic)
}
func (m *FindOneData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneData.Merge(m, src)
}
func (m *FindOneData) XXX_Size() int {
	return xxx_messageInfo_FindOneData.Size(m)
}
func (m *FindOneData) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneData.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneData proto.InternalMessageInfo

func (m *FindOneData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FindOneData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FindOneData) GetStock() int64 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *FindOneData) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type FindOneRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindOneRequest) Reset()         { *m = FindOneRequest{} }
func (m *FindOneRequest) String() string { return proto.CompactTextString(m) }
func (*FindOneRequest) ProtoMessage()    {}
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c53c2f7d36d479c, []int{3}
}

func (m *FindOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneRequest.Unmarshal(m, b)
}
func (m *FindOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneRequest.Marshal(b, m, deterministic)
}
func (m *FindOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneRequest.Merge(m, src)
}
func (m *FindOneRequest) XXX_Size() int {
	return xxx_messageInfo_FindOneRequest.Size(m)
}
func (m *FindOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneRequest proto.InternalMessageInfo

func (m *FindOneRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FindOneResponse struct {
	Status               int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *FindOneData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FindOneResponse) Reset()         { *m = FindOneResponse{} }
func (m *FindOneResponse) String() string { return proto.CompactTextString(m) }
func (*FindOneResponse) ProtoMessage()    {}
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c53c2f7d36d479c, []int{4}
}

func (m *FindOneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneResponse.Unmarshal(m, b)
}
func (m *FindOneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneResponse.Marshal(b, m, deterministic)
}
func (m *FindOneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneResponse.Merge(m, src)
}
func (m *FindOneResponse) XXX_Size() int {
	return xxx_messageInfo_FindOneResponse.Size(m)
}
func (m *FindOneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneResponse proto.InternalMessageInfo

func (m *FindOneResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FindOneResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *FindOneResponse) GetData() *FindOneData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DecreaseStockRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecreaseStockRequest) Reset()         { *m = DecreaseStockRequest{} }
func (m *DecreaseStockRequest) String() string { return proto.CompactTextString(m) }
func (*DecreaseStockRequest) ProtoMessage()    {}
func (*DecreaseStockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c53c2f7d36d479c, []int{5}
}

func (m *DecreaseStockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecreaseStockRequest.Unmarshal(m, b)
}
func (m *DecreaseStockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecreaseStockRequest.Marshal(b, m, deterministic)
}
func (m *DecreaseStockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecreaseStockRequest.Merge(m, src)
}
func (m *DecreaseStockRequest) XXX_Size() int {
	return xxx_messageInfo_DecreaseStockRequest.Size(m)
}
func (m *DecreaseStockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecreaseStockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecreaseStockRequest proto.InternalMessageInfo

func (m *DecreaseStockRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DecreaseStockRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type DecreaseStockResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecreaseStockResponse) Reset()         { *m = DecreaseStockResponse{} }
func (m *DecreaseStockResponse) String() string { return proto.CompactTextString(m) }
func (*DecreaseStockResponse) ProtoMessage()    {}
func (*DecreaseStockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c53c2f7d36d479c, []int{6}
}

func (m *DecreaseStockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecreaseStockResponse.Unmarshal(m, b)
}
func (m *DecreaseStockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecreaseStockResponse.Marshal(b, m, deterministic)
}
func (m *DecreaseStockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecreaseStockResponse.Merge(m, src)
}
func (m *DecreaseStockResponse) XXX_Size() int {
	return xxx_messageInfo_DecreaseStockResponse.Size(m)
}
func (m *DecreaseStockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecreaseStockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecreaseStockResponse proto.InternalMessageInfo

func (m *DecreaseStockResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DecreaseStockResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "product.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "product.CreateResponse")
	proto.RegisterType((*FindOneData)(nil), "product.FindOneData")
	proto.RegisterType((*FindOneRequest)(nil), "product.FindOneRequest")
	proto.RegisterType((*FindOneResponse)(nil), "product.FindOneResponse")
	proto.RegisterType((*DecreaseStockRequest)(nil), "product.DecreaseStockRequest")
	proto.RegisterType((*DecreaseStockResponse)(nil), "product.DecreaseStockResponse")
}

func init() { proto.RegisterFile("internal/pb/product.proto", fileDescriptor_5c53c2f7d36d479c) }

var fileDescriptor_5c53c2f7d36d479c = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4b, 0x4f, 0x02, 0x31,
	0x10, 0x76, 0x1f, 0x80, 0x0e, 0x61, 0x49, 0x1a, 0x84, 0x85, 0x44, 0x43, 0x7a, 0xe2, 0x04, 0x09,
	0x1e, 0x4d, 0x4c, 0x54, 0x34, 0xf1, 0x22, 0x64, 0xb9, 0x99, 0x18, 0x53, 0xb6, 0x73, 0x68, 0xd4,
	0xed, 0xda, 0x16, 0xff, 0xae, 0x7f, 0xc5, 0xd0, 0x7d, 0x84, 0xc7, 0x72, 0xe1, 0xb6, 0x5f, 0x67,
	0xe6, 0x7b, 0x4c, 0xb7, 0xd0, 0x17, 0x89, 0x41, 0x95, 0xb0, 0xaf, 0x49, 0xba, 0x9a, 0xa4, 0x4a,
	0xf2, 0x75, 0x6c, 0xc6, 0xa9, 0x92, 0x46, 0x92, 0x46, 0x0e, 0xe9, 0x1c, 0x5a, 0x8f, 0x0a, 0x99,
	0xc1, 0x08, 0x7f, 0xd6, 0xa8, 0x0d, 0x21, 0xe0, 0x27, 0xec, 0x1b, 0x43, 0x67, 0xe8, 0x8c, 0x2e,
	0x22, 0xfb, 0x4d, 0x3a, 0x50, 0xd3, 0x46, 0xc6, 0x9f, 0xa1, 0x37, 0x74, 0x46, 0x5e, 0x94, 0x81,
	0xcd, 0x69, 0xaa, 0x44, 0x8c, 0xa1, 0x9f, 0x9d, 0x5a, 0x40, 0x5f, 0x21, 0x28, 0x08, 0x75, 0x2a,
	0x13, 0x8d, 0xa4, 0x0b, 0x75, 0x6d, 0x98, 0x59, 0x6b, 0xcb, 0xe9, 0x45, 0x39, 0xda, 0xcc, 0xa3,
	0x52, 0x52, 0x85, 0xae, 0x95, 0xca, 0x00, 0x09, 0xc0, 0x15, 0x3c, 0x17, 0x72, 0x05, 0xa7, 0xef,
	0xd0, 0x7c, 0x16, 0x09, 0x9f, 0x27, 0x38, 0x63, 0x86, 0xe5, 0x65, 0xa7, 0x28, 0x97, 0x76, 0xdd,
	0x2a, 0xbb, 0x7e, 0xa5, 0xdd, 0xda, 0xb6, 0xdd, 0x21, 0x04, 0x39, 0x7d, 0xb1, 0x80, 0x3d, 0x05,
	0x2a, 0xa0, 0x5d, 0x76, 0x9c, 0x94, 0x68, 0x04, 0x3e, 0x67, 0x86, 0xd9, 0x4c, 0xcd, 0x69, 0x67,
	0x5c, 0xdc, 0xc4, 0x56, 0xac, 0xc8, 0x76, 0xd0, 0x7b, 0xe8, 0xcc, 0x30, 0x56, 0xc8, 0x34, 0x2e,
	0x37, 0x9e, 0x8f, 0x58, 0x22, 0x7d, 0x38, 0x97, 0x8a, 0xa3, 0xfa, 0x10, 0xdc, 0x4a, 0x79, 0x51,
	0xc3, 0xe2, 0x17, 0x4e, 0x9f, 0xe0, 0x72, 0x8f, 0xe2, 0x14, 0xcf, 0xd3, 0x3f, 0x07, 0x82, 0x45,
	0xe6, 0x73, 0x89, 0xea, 0x57, 0xc4, 0x48, 0x6e, 0xa1, 0x9e, 0x5d, 0x2c, 0xe9, 0x96, 0x11, 0x76,
	0x7e, 0x9d, 0x41, 0xef, 0xe0, 0x3c, 0xd3, 0xa6, 0x67, 0xe4, 0x0e, 0x1a, 0x79, 0x5c, 0xd2, 0xdb,
	0x5f, 0x40, 0x31, 0x1e, 0x1e, 0x16, 0xca, 0xf9, 0x05, 0xb4, 0x76, 0x62, 0x91, 0xab, 0xb2, 0xb9,
	0x6a, 0x63, 0x83, 0xeb, 0x63, 0xe5, 0x82, 0xf1, 0xa1, 0xfd, 0xd6, 0x1a, 0x4f, 0xb6, 0x1e, 0xc8,
	0xaa, 0x6e, 0x5f, 0xc6, 0xcd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x48, 0x04, 0x99, 0x36,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error)
	DecreaseStock(ctx context.Context, in *DecreaseStockRequest, opts ...grpc.CallOption) (*DecreaseStockResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DecreaseStock(ctx context.Context, in *DecreaseStockRequest, opts ...grpc.CallOption) (*DecreaseStockResponse, error) {
	out := new(DecreaseStockResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/DecreaseStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error)
	DecreaseStock(context.Context, *DecreaseStockRequest) (*DecreaseStockResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProductServiceServer) FindOne(ctx context.Context, req *FindOneRequest) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (*UnimplementedProductServiceServer) DecreaseStock(ctx context.Context, req *DecreaseStockRequest) (*DecreaseStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseStock not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindOne(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DecreaseStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecreaseStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DecreaseStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/DecreaseStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DecreaseStock(ctx, req.(*DecreaseStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductService_Create_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _ProductService_FindOne_Handler,
		},
		{
			MethodName: "DecreaseStock",
			Handler:    _ProductService_DecreaseStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pb/product.proto",
}
