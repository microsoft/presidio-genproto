// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasink.proto

package types

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

// The data sink supported destenation types
type DatasinkTypesEnum int32

const (
	DatasinkTypesEnum_mysql         DatasinkTypesEnum = 0
	DatasinkTypesEnum_mssql         DatasinkTypesEnum = 1
	DatasinkTypesEnum_postgres      DatasinkTypesEnum = 2
	DatasinkTypesEnum_sqlite3       DatasinkTypesEnum = 3
	DatasinkTypesEnum_oracle        DatasinkTypesEnum = 4
	DatasinkTypesEnum_kafka         DatasinkTypesEnum = 5
	DatasinkTypesEnum_eventhub      DatasinkTypesEnum = 6
	DatasinkTypesEnum_kinesis       DatasinkTypesEnum = 7
	DatasinkTypesEnum_s3            DatasinkTypesEnum = 8
	DatasinkTypesEnum_azureblob     DatasinkTypesEnum = 9
	DatasinkTypesEnum_googlestorage DatasinkTypesEnum = 10
)

var DatasinkTypesEnum_name = map[int32]string{
	0:  "mysql",
	1:  "mssql",
	2:  "postgres",
	3:  "sqlite3",
	4:  "oracle",
	5:  "kafka",
	6:  "eventhub",
	7:  "kinesis",
	8:  "s3",
	9:  "azureblob",
	10: "googlestorage",
}
var DatasinkTypesEnum_value = map[string]int32{
	"mysql":         0,
	"mssql":         1,
	"postgres":      2,
	"sqlite3":       3,
	"oracle":        4,
	"kafka":         5,
	"eventhub":      6,
	"kinesis":       7,
	"s3":            8,
	"azureblob":     9,
	"googlestorage": 10,
}

func (x DatasinkTypesEnum) String() string {
	return proto.EnumName(DatasinkTypesEnum_name, int32(x))
}
func (DatasinkTypesEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_datasink_f6dd47e2bb530188, []int{0}
}

// DatasinkRequest represents the request to the data-sink service via GRPC
type DatasinkRequest struct {
	// Array of the analyzed results
	AnalyzeResults []*AnalyzeResult `protobuf:"bytes,1,rep,name=analyzeResults" json:"analyzeResults,omitempty"`
	// The anonymized text
	AnonymizeResult *AnonymizeResponse `protobuf:"bytes,2,opt,name=anonymizeResult" json:"anonymizeResult,omitempty"`
	// The path where the anonymized text is located
	Path                 string   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasinkRequest) Reset()         { *m = DatasinkRequest{} }
func (m *DatasinkRequest) String() string { return proto.CompactTextString(m) }
func (*DatasinkRequest) ProtoMessage()    {}
func (*DatasinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datasink_f6dd47e2bb530188, []int{0}
}
func (m *DatasinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasinkRequest.Unmarshal(m, b)
}
func (m *DatasinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasinkRequest.Marshal(b, m, deterministic)
}
func (dst *DatasinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasinkRequest.Merge(dst, src)
}
func (m *DatasinkRequest) XXX_Size() int {
	return xxx_messageInfo_DatasinkRequest.Size(m)
}
func (m *DatasinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatasinkRequest proto.InternalMessageInfo

func (m *DatasinkRequest) GetAnalyzeResults() []*AnalyzeResult {
	if m != nil {
		return m.AnalyzeResults
	}
	return nil
}

func (m *DatasinkRequest) GetAnonymizeResult() *AnonymizeResponse {
	if m != nil {
		return m.AnonymizeResult
	}
	return nil
}

func (m *DatasinkRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// DatasinkResponse represents the response from the data-sink service via GRPC
type DatasinkResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasinkResponse) Reset()         { *m = DatasinkResponse{} }
func (m *DatasinkResponse) String() string { return proto.CompactTextString(m) }
func (*DatasinkResponse) ProtoMessage()    {}
func (*DatasinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datasink_f6dd47e2bb530188, []int{1}
}
func (m *DatasinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasinkResponse.Unmarshal(m, b)
}
func (m *DatasinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasinkResponse.Marshal(b, m, deterministic)
}
func (dst *DatasinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasinkResponse.Merge(dst, src)
}
func (m *DatasinkResponse) XXX_Size() int {
	return xxx_messageInfo_DatasinkResponse.Size(m)
}
func (m *DatasinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatasinkResponse proto.InternalMessageInfo

// CompletionMessage represents an indication to the data sink service that the scanning job is done service via GRPC
type CompletionMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompletionMessage) Reset()         { *m = CompletionMessage{} }
func (m *CompletionMessage) String() string { return proto.CompactTextString(m) }
func (*CompletionMessage) ProtoMessage()    {}
func (*CompletionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_datasink_f6dd47e2bb530188, []int{2}
}
func (m *CompletionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompletionMessage.Unmarshal(m, b)
}
func (m *CompletionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompletionMessage.Marshal(b, m, deterministic)
}
func (dst *CompletionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompletionMessage.Merge(dst, src)
}
func (m *CompletionMessage) XXX_Size() int {
	return xxx_messageInfo_CompletionMessage.Size(m)
}
func (m *CompletionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CompletionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CompletionMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DatasinkRequest)(nil), "types.DatasinkRequest")
	proto.RegisterType((*DatasinkResponse)(nil), "types.DatasinkResponse")
	proto.RegisterType((*CompletionMessage)(nil), "types.CompletionMessage")
	proto.RegisterEnum("types.DatasinkTypesEnum", DatasinkTypesEnum_name, DatasinkTypesEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatasinkServiceClient is the client API for DatasinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatasinkServiceClient interface {
	// Apply method will execute on the given request and return whether the result where written successfully to the destination
	Apply(ctx context.Context, in *DatasinkRequest, opts ...grpc.CallOption) (*DatasinkResponse, error)
	// Init the data sink service with the provided data sink template
	Init(ctx context.Context, in *DatasinkTemplate, opts ...grpc.CallOption) (*DatasinkResponse, error)
	// Completion method for indicating that the scanning job is done
	Completion(ctx context.Context, in *CompletionMessage, opts ...grpc.CallOption) (*DatasinkResponse, error)
}

type datasinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewDatasinkServiceClient(cc *grpc.ClientConn) DatasinkServiceClient {
	return &datasinkServiceClient{cc}
}

func (c *datasinkServiceClient) Apply(ctx context.Context, in *DatasinkRequest, opts ...grpc.CallOption) (*DatasinkResponse, error) {
	out := new(DatasinkResponse)
	err := c.cc.Invoke(ctx, "/types.DatasinkService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasinkServiceClient) Init(ctx context.Context, in *DatasinkTemplate, opts ...grpc.CallOption) (*DatasinkResponse, error) {
	out := new(DatasinkResponse)
	err := c.cc.Invoke(ctx, "/types.DatasinkService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasinkServiceClient) Completion(ctx context.Context, in *CompletionMessage, opts ...grpc.CallOption) (*DatasinkResponse, error) {
	out := new(DatasinkResponse)
	err := c.cc.Invoke(ctx, "/types.DatasinkService/Completion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasinkServiceServer is the server API for DatasinkService service.
type DatasinkServiceServer interface {
	// Apply method will execute on the given request and return whether the result where written successfully to the destination
	Apply(context.Context, *DatasinkRequest) (*DatasinkResponse, error)
	// Init the data sink service with the provided data sink template
	Init(context.Context, *DatasinkTemplate) (*DatasinkResponse, error)
	// Completion method for indicating that the scanning job is done
	Completion(context.Context, *CompletionMessage) (*DatasinkResponse, error)
}

func RegisterDatasinkServiceServer(s *grpc.Server, srv DatasinkServiceServer) {
	s.RegisterService(&_DatasinkService_serviceDesc, srv)
}

func _DatasinkService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasinkServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.DatasinkService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasinkServiceServer).Apply(ctx, req.(*DatasinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasinkService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasinkTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasinkServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.DatasinkService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasinkServiceServer).Init(ctx, req.(*DatasinkTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasinkService_Completion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasinkServiceServer).Completion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.DatasinkService/Completion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasinkServiceServer).Completion(ctx, req.(*CompletionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.DatasinkService",
	HandlerType: (*DatasinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _DatasinkService_Apply_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _DatasinkService_Init_Handler,
		},
		{
			MethodName: "Completion",
			Handler:    _DatasinkService_Completion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datasink.proto",
}

func init() { proto.RegisterFile("datasink.proto", fileDescriptor_datasink_f6dd47e2bb530188) }

var fileDescriptor_datasink_f6dd47e2bb530188 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3d, 0x6e, 0xdb, 0x40,
	0x10, 0x85, 0x4d, 0xfd, 0xd9, 0x1c, 0xd9, 0xf2, 0x6a, 0x12, 0x24, 0x84, 0x2a, 0x41, 0x95, 0x90,
	0x42, 0x85, 0xd4, 0x19, 0x69, 0x94, 0x9f, 0x22, 0x45, 0x1a, 0xc6, 0x17, 0x58, 0x2a, 0x13, 0x7a,
	0xc1, 0xe5, 0xee, 0x8a, 0xb3, 0x34, 0x40, 0x9d, 0x26, 0x5d, 0x2e, 0x93, 0x43, 0x05, 0xa4, 0xe8,
	0xd0, 0xa0, 0x01, 0x77, 0xc3, 0x37, 0xf3, 0xcd, 0x10, 0xef, 0x2d, 0xcc, 0x7e, 0x4a, 0x2f, 0x59,
	0x99, 0x6c, 0xe3, 0x0a, 0xeb, 0x2d, 0x8e, 0x7d, 0xe5, 0x88, 0x17, 0xd7, 0x07, 0x9b, 0xe7, 0xd6,
	0x9c, 0xc5, 0xc5, 0xcc, 0x53, 0xee, 0xb4, 0xf4, 0xd4, 0x7e, 0xdf, 0x4a, 0x63, 0x4d, 0x95, 0xab,
	0x53, 0x2b, 0xac, 0xfe, 0x04, 0x70, 0xfb, 0xa5, 0x5d, 0x14, 0xd3, 0xb1, 0x24, 0xf6, 0xf8, 0x11,
	0x66, 0xd2, 0x48, 0x5d, 0x9d, 0x28, 0x26, 0x2e, 0xb5, 0xe7, 0x28, 0x58, 0x0e, 0xd7, 0xd3, 0xed,
	0xdb, 0x4d, 0x73, 0x62, 0xb3, 0x7f, 0xde, 0x8c, 0x7b, 0xb3, 0xf8, 0x09, 0xba, 0x23, 0x67, 0x2d,
	0x1a, 0x2c, 0x83, 0xf5, 0x74, 0x1b, 0xfd, 0xc7, 0xbb, 0xae, 0xb3, 0x86, 0x29, 0xee, 0x03, 0x88,
	0x30, 0x72, 0xd2, 0x3f, 0x44, 0xc3, 0x65, 0xb0, 0x0e, 0xe3, 0xa6, 0x5e, 0x21, 0x88, 0xee, 0x47,
	0xcf, 0xe0, 0xea, 0x0d, 0xcc, 0x3f, 0xdb, 0xdc, 0x69, 0xf2, 0xca, 0x9a, 0xef, 0xc4, 0x2c, 0x53,
	0xfa, 0xf0, 0x3b, 0x80, 0xf9, 0xd3, 0xe4, 0x7d, 0x7d, 0xf1, 0xab, 0x29, 0x73, 0x0c, 0x61, 0x9c,
	0x57, 0x7c, 0xd4, 0xe2, 0xa2, 0x29, 0xb9, 0x2e, 0x03, 0xbc, 0x86, 0x2b, 0x67, 0xd9, 0xa7, 0x05,
	0xb1, 0x18, 0xe0, 0x14, 0x2e, 0xf9, 0xa8, 0x95, 0xa7, 0x9d, 0x18, 0x22, 0xc0, 0xc4, 0x16, 0xf2,
	0xa0, 0x49, 0x8c, 0x6a, 0x22, 0x93, 0xbf, 0x32, 0x29, 0xc6, 0x35, 0x41, 0x8f, 0x64, 0xfc, 0x43,
	0x99, 0x88, 0x49, 0x4d, 0x64, 0xca, 0x10, 0x2b, 0x16, 0x97, 0x38, 0x81, 0x01, 0xef, 0xc4, 0x15,
	0xde, 0x40, 0x28, 0x4f, 0x65, 0x41, 0x89, 0xb6, 0x89, 0x08, 0x71, 0x0e, 0x37, 0xa9, 0xb5, 0xa9,
	0x26, 0xf6, 0xb6, 0x90, 0x29, 0x09, 0xd8, 0xfe, 0x7d, 0xe6, 0xfa, 0x0f, 0x2a, 0x1e, 0xd5, 0x81,
	0xf0, 0x0e, 0xc6, 0x7b, 0xe7, 0x74, 0x85, 0xef, 0x5a, 0x9f, 0x7a, 0xb1, 0x2c, 0xde, 0xbf, 0xd0,
	0x5b, 0x17, 0x2e, 0xf0, 0x0e, 0x46, 0xdf, 0x8c, 0xf2, 0xd8, 0x1f, 0xb9, 0x6f, 0xd3, 0x7f, 0x8d,
	0xdd, 0x03, 0x74, 0x1e, 0xe2, 0x53, 0x48, 0x2f, 0x6c, 0x7d, 0x65, 0x45, 0x32, 0x69, 0xde, 0xd2,
	0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x1a, 0xa5, 0x99, 0x93, 0x02, 0x00, 0x00,
}
