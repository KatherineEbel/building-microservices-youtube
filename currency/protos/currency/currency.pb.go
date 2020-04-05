// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

package currency

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

// Currencies is an enum which represents the allowed currencies for the API
type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_USD Currencies = 1
	Currencies_JPY Currencies = 2
	Currencies_BGN Currencies = 3
	Currencies_CZK Currencies = 4
	Currencies_DKK Currencies = 5
	Currencies_GBP Currencies = 6
	Currencies_HUF Currencies = 7
	Currencies_PLN Currencies = 8
	Currencies_RON Currencies = 9
	Currencies_SEK Currencies = 10
	Currencies_CHF Currencies = 11
	Currencies_ISK Currencies = 12
	Currencies_NOK Currencies = 13
	Currencies_HRK Currencies = 14
	Currencies_RUB Currencies = 15
	Currencies_TRY Currencies = 16
	Currencies_AUD Currencies = 17
	Currencies_BRL Currencies = 18
	Currencies_CAD Currencies = 19
	Currencies_CNY Currencies = 20
	Currencies_HKD Currencies = 21
	Currencies_IDR Currencies = 22
	Currencies_ILS Currencies = 23
	Currencies_INR Currencies = 24
	Currencies_KRW Currencies = 25
	Currencies_MXN Currencies = 26
	Currencies_MYR Currencies = 27
	Currencies_NZD Currencies = 28
	Currencies_PHP Currencies = 29
	Currencies_SGD Currencies = 30
	Currencies_THB Currencies = 31
	Currencies_ZAR Currencies = 32
)

var Currencies_name = map[int32]string{
	0:  "EUR",
	1:  "USD",
	2:  "JPY",
	3:  "BGN",
	4:  "CZK",
	5:  "DKK",
	6:  "GBP",
	7:  "HUF",
	8:  "PLN",
	9:  "RON",
	10: "SEK",
	11: "CHF",
	12: "ISK",
	13: "NOK",
	14: "HRK",
	15: "RUB",
	16: "TRY",
	17: "AUD",
	18: "BRL",
	19: "CAD",
	20: "CNY",
	21: "HKD",
	22: "IDR",
	23: "ILS",
	24: "INR",
	25: "KRW",
	26: "MXN",
	27: "MYR",
	28: "NZD",
	29: "PHP",
	30: "SGD",
	31: "THB",
	32: "ZAR",
}

var Currencies_value = map[string]int32{
	"EUR": 0,
	"USD": 1,
	"JPY": 2,
	"BGN": 3,
	"CZK": 4,
	"DKK": 5,
	"GBP": 6,
	"HUF": 7,
	"PLN": 8,
	"RON": 9,
	"SEK": 10,
	"CHF": 11,
	"ISK": 12,
	"NOK": 13,
	"HRK": 14,
	"RUB": 15,
	"TRY": 16,
	"AUD": 17,
	"BRL": 18,
	"CAD": 19,
	"CNY": 20,
	"HKD": 21,
	"IDR": 22,
	"ILS": 23,
	"INR": 24,
	"KRW": 25,
	"MXN": 26,
	"MYR": 27,
	"NZD": 28,
	"PHP": 29,
	"SGD": 30,
	"THB": 31,
	"ZAR": 32,
}

func (x Currencies) String() string {
	return proto.EnumName(Currencies_name, int32(x))
}

func (Currencies) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

// RateRequest defines the request for a GetRate call
type RateRequest struct {
	// Base is the base currency code for the rate
	Base Currencies `protobuf:"varint,1,opt,name=Base,json=base,proto3,enum=Currencies" json:"Base,omitempty"`
	// Destination is the destination currency code for the rate
	Destination          Currencies `protobuf:"varint,2,opt,name=Destination,json=destination,proto3,enum=Currencies" json:"Destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RateRequest) Reset()         { *m = RateRequest{} }
func (m *RateRequest) String() string { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()    {}
func (*RateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

func (m *RateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateRequest.Unmarshal(m, b)
}
func (m *RateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateRequest.Marshal(b, m, deterministic)
}
func (m *RateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateRequest.Merge(m, src)
}
func (m *RateRequest) XXX_Size() int {
	return xxx_messageInfo_RateRequest.Size(m)
}
func (m *RateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateRequest proto.InternalMessageInfo

func (m *RateRequest) GetBase() Currencies {
	if m != nil {
		return m.Base
	}
	return Currencies_EUR
}

func (m *RateRequest) GetDestination() Currencies {
	if m != nil {
		return m.Destination
	}
	return Currencies_EUR
}

// RateResponse is the response from a GetRate call, it contains
// rate which is a floating point number and can be used to convert between the
// two currencies specified in the request.
type RateResponse struct {
	// Base is the base currency code for the rate
	Base Currencies `protobuf:"varint,1,opt,name=Base,json=base,proto3,enum=Currencies" json:"Base,omitempty"`
	// Destination is the destination currency code for the rate
	Destination Currencies `protobuf:"varint,2,opt,name=Destination,json=destination,proto3,enum=Currencies" json:"Destination,omitempty"`
	// Rate of exchange between the two currencies
	Rate                 float64  `protobuf:"fixed64,3,opt,name=Rate,json=rate,proto3" json:"Rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateResponse) Reset()         { *m = RateResponse{} }
func (m *RateResponse) String() string { return proto.CompactTextString(m) }
func (*RateResponse) ProtoMessage()    {}
func (*RateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{1}
}

func (m *RateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateResponse.Unmarshal(m, b)
}
func (m *RateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateResponse.Marshal(b, m, deterministic)
}
func (m *RateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateResponse.Merge(m, src)
}
func (m *RateResponse) XXX_Size() int {
	return xxx_messageInfo_RateResponse.Size(m)
}
func (m *RateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateResponse proto.InternalMessageInfo

func (m *RateResponse) GetBase() Currencies {
	if m != nil {
		return m.Base
	}
	return Currencies_EUR
}

func (m *RateResponse) GetDestination() Currencies {
	if m != nil {
		return m.Destination
	}
	return Currencies_EUR
}

func (m *RateResponse) GetRate() float64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func init() {
	proto.RegisterEnum("Currencies", Currencies_name, Currencies_value)
	proto.RegisterType((*RateRequest)(nil), "RateRequest")
	proto.RegisterType((*RateResponse)(nil), "RateResponse")
}

func init() {
	proto.RegisterFile("currency.proto", fileDescriptor_d3dc60ed002193ea)
}

var fileDescriptor_d3dc60ed002193ea = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd2, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x05, 0x60, 0x9c, 0x98, 0xa6, 0x8c, 0xd3, 0x70, 0x18, 0xfe, 0x42, 0xf9, 0x69, 0xd4, 0x05,
	0x8a, 0x90, 0xb0, 0xa0, 0x3c, 0x41, 0xdc, 0x69, 0x6d, 0x98, 0xe0, 0x5a, 0x63, 0x2c, 0x70, 0x24,
	0x16, 0x76, 0x98, 0x85, 0x37, 0x76, 0xf1, 0x38, 0x8b, 0x3e, 0x18, 0xef, 0x87, 0xee, 0x2d, 0x52,
	0x41, 0x48, 0xac, 0xd8, 0x7d, 0xd2, 0xdc, 0x7b, 0x74, 0x34, 0x33, 0x62, 0xb6, 0xdd, 0xf5, 0xbd,
	0x6d, 0xb7, 0x57, 0xe1, 0x65, 0xdf, 0x0d, 0xdd, 0xf1, 0x57, 0x11, 0x98, 0x6a, 0xb0, 0xc6, 0x7e,
	0xdf, 0x59, 0x37, 0xc8, 0x23, 0xe1, 0x47, 0x95, 0xb3, 0x73, 0x6f, 0xe1, 0x2d, 0x67, 0x27, 0x41,
	0x78, 0x7a, 0x3d, 0xdd, 0x58, 0x67, 0xfc, 0xba, 0x72, 0x56, 0xbe, 0x16, 0x81, 0xb2, 0x6e, 0x68,
	0xda, 0x6a, 0x68, 0xba, 0x76, 0x3e, 0xfa, 0x7b, 0x2e, 0xf8, 0x76, 0x73, 0x7e, 0xdc, 0x8b, 0xe9,
	0x75, 0xbc, 0xbb, 0xec, 0x5a, 0x67, 0xff, 0x77, 0xbe, 0x94, 0xc2, 0xa7, 0xfc, 0xf9, 0x78, 0xe1,
	0x2d, 0x3d, 0xe3, 0xf7, 0xd5, 0x60, 0x5f, 0xfd, 0x18, 0x09, 0x71, 0x33, 0x2f, 0x27, 0x62, 0x7c,
	0x56, 0x18, 0xdc, 0x22, 0x14, 0xb9, 0x82, 0x47, 0xf8, 0x90, 0x95, 0x18, 0x11, 0xa2, 0x38, 0xc5,
	0x98, 0x70, 0xba, 0xd1, 0xf0, 0x09, 0x4a, 0x6b, 0xdc, 0x26, 0xc4, 0x51, 0x86, 0x3d, 0x42, 0x52,
	0x9c, 0x63, 0x42, 0xc8, 0xd6, 0x29, 0xf6, 0x09, 0xe6, 0x22, 0xc5, 0x1d, 0x42, 0x7e, 0xa6, 0x21,
	0x78, 0x3d, 0x39, 0x47, 0x40, 0x78, 0x9f, 0x6b, 0x4c, 0x09, 0xe9, 0x85, 0xc6, 0x01, 0xaf, 0x1b,
	0x8d, 0x19, 0x6f, 0x15, 0x11, 0xee, 0x12, 0x3e, 0x99, 0x12, 0x20, 0xac, 0x0a, 0x85, 0x7b, 0x5c,
	0xc3, 0xac, 0x21, 0x39, 0x67, 0xa5, 0x70, 0x9f, 0x91, 0x96, 0x78, 0xc0, 0xeb, 0x5a, 0xe1, 0x21,
	0x27, 0x2b, 0x83, 0x47, 0x8c, 0x75, 0x8e, 0xc7, 0x8c, 0xd4, 0x60, 0x4e, 0xd0, 0xe6, 0x33, 0x9e,
	0x10, 0x3e, 0x7e, 0x49, 0x71, 0xc8, 0x28, 0x0d, 0x9e, 0x72, 0x8d, 0x8d, 0xc2, 0x33, 0x2e, 0x9f,
	0x64, 0x78, 0xce, 0x9d, 0x63, 0x85, 0x17, 0x5c, 0x23, 0x89, 0x70, 0x44, 0xd8, 0xac, 0x0c, 0x16,
	0x27, 0x56, 0xec, 0xff, 0xba, 0xb6, 0x2b, 0xf9, 0x52, 0x4c, 0x62, 0x3b, 0xd0, 0xd5, 0xca, 0x69,
	0xf8, 0xdb, 0x07, 0x39, 0x3c, 0x08, 0xff, 0x78, 0xcf, 0xb7, 0x62, 0x96, 0xef, 0x6a, 0xb7, 0xed,
	0x9b, 0xda, 0xd2, 0x81, 0xfb, 0xe7, 0xf8, 0xd2, 0x7b, 0xe3, 0xd5, 0x7b, 0xfc, 0xf1, 0xde, 0xfd,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x65, 0x47, 0x37, 0x74, 0x8a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyClient interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	SubscribeRates(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeRatesClient, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) SubscribeRates(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeRatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Currency_serviceDesc.Streams[0], "/Currency/SubscribeRates", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencySubscribeRatesClient{stream}
	return x, nil
}

type Currency_SubscribeRatesClient interface {
	Send(*RateRequest) error
	Recv() (*RateResponse, error)
	grpc.ClientStream
}

type currencySubscribeRatesClient struct {
	grpc.ClientStream
}

func (x *currencySubscribeRatesClient) Send(m *RateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *currencySubscribeRatesClient) Recv() (*RateResponse, error) {
	m := new(RateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServer is the server API for Currency service.
type CurrencyServer interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
	SubscribeRates(Currency_SubscribeRatesServer) error
}

// UnimplementedCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetRate(ctx context.Context, req *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (*UnimplementedCurrencyServer) SubscribeRates(srv Currency_SubscribeRatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeRates not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv CurrencyServer) {
	s.RegisterService(&_Currency_serviceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_SubscribeRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CurrencyServer).SubscribeRates(&currencySubscribeRatesServer{stream})
}

type Currency_SubscribeRatesServer interface {
	Send(*RateResponse) error
	Recv() (*RateRequest, error)
	grpc.ServerStream
}

type currencySubscribeRatesServer struct {
	grpc.ServerStream
}

func (x *currencySubscribeRatesServer) Send(m *RateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *currencySubscribeRatesServer) Recv() (*RateRequest, error) {
	m := new(RateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Currency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeRates",
			Handler:       _Currency_SubscribeRates_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "currency.proto",
}
