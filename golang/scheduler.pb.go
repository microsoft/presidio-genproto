// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

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

// ScannerCronJobApiRequest represents the request to the API HTTP service
type ScannerCronJobApiRequest struct {
	// The scanner cron job template id
	ScannerCronJobTemplateId string `protobuf:"bytes,1,opt,name=ScannerCronJobTemplateId,json=scannerCronJobTemplateId,proto3" json:"ScannerCronJobTemplateId,omitempty"`
	// The scanner cron job requeset represents the  the scanning job details [optional parameter]
	ScannerCronJobRequest *ScannerCronJobRequest `protobuf:"bytes,2,opt,name=scannerCronJobRequest,proto3" json:"scannerCronJobRequest,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *ScannerCronJobApiRequest) Reset()         { *m = ScannerCronJobApiRequest{} }
func (m *ScannerCronJobApiRequest) String() string { return proto.CompactTextString(m) }
func (*ScannerCronJobApiRequest) ProtoMessage()    {}
func (*ScannerCronJobApiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scheduler_e211eb632181759c, []int{0}
}
func (m *ScannerCronJobApiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScannerCronJobApiRequest.Unmarshal(m, b)
}
func (m *ScannerCronJobApiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScannerCronJobApiRequest.Marshal(b, m, deterministic)
}
func (dst *ScannerCronJobApiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScannerCronJobApiRequest.Merge(dst, src)
}
func (m *ScannerCronJobApiRequest) XXX_Size() int {
	return xxx_messageInfo_ScannerCronJobApiRequest.Size(m)
}
func (m *ScannerCronJobApiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScannerCronJobApiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScannerCronJobApiRequest proto.InternalMessageInfo

func (m *ScannerCronJobApiRequest) GetScannerCronJobTemplateId() string {
	if m != nil {
		return m.ScannerCronJobTemplateId
	}
	return ""
}

func (m *ScannerCronJobApiRequest) GetScannerCronJobRequest() *ScannerCronJobRequest {
	if m != nil {
		return m.ScannerCronJobRequest
	}
	return nil
}

// ScannerCronJobRequest represents the request to the scheduler service via GRPC
type ScannerCronJobRequest struct {
	// Scanner cronjob name
	Name string `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// The trigger for a new job to start
	Trigger *Trigger `protobuf:"bytes,2,opt,name=trigger,proto3" json:"trigger,omitempty"`
	// The scanner request represents the scanning details
	ScanRequest          *ScanRequest `protobuf:"bytes,3,opt,name=scanRequest,proto3" json:"scanRequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScannerCronJobRequest) Reset()         { *m = ScannerCronJobRequest{} }
func (m *ScannerCronJobRequest) String() string { return proto.CompactTextString(m) }
func (*ScannerCronJobRequest) ProtoMessage()    {}
func (*ScannerCronJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scheduler_e211eb632181759c, []int{1}
}
func (m *ScannerCronJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScannerCronJobRequest.Unmarshal(m, b)
}
func (m *ScannerCronJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScannerCronJobRequest.Marshal(b, m, deterministic)
}
func (dst *ScannerCronJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScannerCronJobRequest.Merge(dst, src)
}
func (m *ScannerCronJobRequest) XXX_Size() int {
	return xxx_messageInfo_ScannerCronJobRequest.Size(m)
}
func (m *ScannerCronJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScannerCronJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScannerCronJobRequest proto.InternalMessageInfo

func (m *ScannerCronJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScannerCronJobRequest) GetTrigger() *Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *ScannerCronJobRequest) GetScanRequest() *ScanRequest {
	if m != nil {
		return m.ScanRequest
	}
	return nil
}

type ScannerCronJobResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScannerCronJobResponse) Reset()         { *m = ScannerCronJobResponse{} }
func (m *ScannerCronJobResponse) String() string { return proto.CompactTextString(m) }
func (*ScannerCronJobResponse) ProtoMessage()    {}
func (*ScannerCronJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scheduler_e211eb632181759c, []int{2}
}
func (m *ScannerCronJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScannerCronJobResponse.Unmarshal(m, b)
}
func (m *ScannerCronJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScannerCronJobResponse.Marshal(b, m, deterministic)
}
func (dst *ScannerCronJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScannerCronJobResponse.Merge(dst, src)
}
func (m *ScannerCronJobResponse) XXX_Size() int {
	return xxx_messageInfo_ScannerCronJobResponse.Size(m)
}
func (m *ScannerCronJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScannerCronJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScannerCronJobResponse proto.InternalMessageInfo

// StreamsJobApiRequest represents the request to the API HTTP service
type StreamsJobApiRequest struct {
	// The streams job template id
	StreamsJobTemplateId string `protobuf:"bytes,1,opt,name=StreamsJobTemplateId,json=streamsJobTemplateId,proto3" json:"StreamsJobTemplateId,omitempty"`
	// The streams request that contains the full streamsRequest
	StreamsJobRequest    *StreamsJobRequest `protobuf:"bytes,2,opt,name=streamsJobRequest,proto3" json:"streamsJobRequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StreamsJobApiRequest) Reset()         { *m = StreamsJobApiRequest{} }
func (m *StreamsJobApiRequest) String() string { return proto.CompactTextString(m) }
func (*StreamsJobApiRequest) ProtoMessage()    {}
func (*StreamsJobApiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scheduler_e211eb632181759c, []int{3}
}
func (m *StreamsJobApiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamsJobApiRequest.Unmarshal(m, b)
}
func (m *StreamsJobApiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamsJobApiRequest.Marshal(b, m, deterministic)
}
func (dst *StreamsJobApiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamsJobApiRequest.Merge(dst, src)
}
func (m *StreamsJobApiRequest) XXX_Size() int {
	return xxx_messageInfo_StreamsJobApiRequest.Size(m)
}
func (m *StreamsJobApiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamsJobApiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamsJobApiRequest proto.InternalMessageInfo

func (m *StreamsJobApiRequest) GetStreamsJobTemplateId() string {
	if m != nil {
		return m.StreamsJobTemplateId
	}
	return ""
}

func (m *StreamsJobApiRequest) GetStreamsJobRequest() *StreamsJobRequest {
	if m != nil {
		return m.StreamsJobRequest
	}
	return nil
}

// StreamsJobRequest represents the request to the scheduler service via GRPC
type StreamsJobRequest struct {
	// The streams job name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The streams requeset that hold the streaming details
	StreamsRequest       *StreamRequest `protobuf:"bytes,2,opt,name=streamsRequest,proto3" json:"streamsRequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamsJobRequest) Reset()         { *m = StreamsJobRequest{} }
func (m *StreamsJobRequest) String() string { return proto.CompactTextString(m) }
func (*StreamsJobRequest) ProtoMessage()    {}
func (*StreamsJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scheduler_e211eb632181759c, []int{4}
}
func (m *StreamsJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamsJobRequest.Unmarshal(m, b)
}
func (m *StreamsJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamsJobRequest.Marshal(b, m, deterministic)
}
func (dst *StreamsJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamsJobRequest.Merge(dst, src)
}
func (m *StreamsJobRequest) XXX_Size() int {
	return xxx_messageInfo_StreamsJobRequest.Size(m)
}
func (m *StreamsJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamsJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamsJobRequest proto.InternalMessageInfo

func (m *StreamsJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamsJobRequest) GetStreamsRequest() *StreamRequest {
	if m != nil {
		return m.StreamsRequest
	}
	return nil
}

type StreamsJobResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamsJobResponse) Reset()         { *m = StreamsJobResponse{} }
func (m *StreamsJobResponse) String() string { return proto.CompactTextString(m) }
func (*StreamsJobResponse) ProtoMessage()    {}
func (*StreamsJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scheduler_e211eb632181759c, []int{5}
}
func (m *StreamsJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamsJobResponse.Unmarshal(m, b)
}
func (m *StreamsJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamsJobResponse.Marshal(b, m, deterministic)
}
func (dst *StreamsJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamsJobResponse.Merge(dst, src)
}
func (m *StreamsJobResponse) XXX_Size() int {
	return xxx_messageInfo_StreamsJobResponse.Size(m)
}
func (m *StreamsJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamsJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamsJobResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScannerCronJobApiRequest)(nil), "types.ScannerCronJobApiRequest")
	proto.RegisterType((*ScannerCronJobRequest)(nil), "types.ScannerCronJobRequest")
	proto.RegisterType((*ScannerCronJobResponse)(nil), "types.ScannerCronJobResponse")
	proto.RegisterType((*StreamsJobApiRequest)(nil), "types.StreamsJobApiRequest")
	proto.RegisterType((*StreamsJobRequest)(nil), "types.StreamsJobRequest")
	proto.RegisterType((*StreamsJobResponse)(nil), "types.StreamsJobResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchedulerServiceClient is the client API for SchedulerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerServiceClient interface {
	// ApplyStream method will trigger a new scanning cron job and will return if it was triggered successfully
	ApplyStream(ctx context.Context, in *StreamsJobRequest, opts ...grpc.CallOption) (*StreamsJobResponse, error)
	// Apply method will trigger a new scanning cron job and will return if it was triggered successfully
	ApplyScan(ctx context.Context, in *ScannerCronJobRequest, opts ...grpc.CallOption) (*ScannerCronJobResponse, error)
}

type schedulerServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerServiceClient(cc *grpc.ClientConn) SchedulerServiceClient {
	return &schedulerServiceClient{cc}
}

func (c *schedulerServiceClient) ApplyStream(ctx context.Context, in *StreamsJobRequest, opts ...grpc.CallOption) (*StreamsJobResponse, error) {
	out := new(StreamsJobResponse)
	err := c.cc.Invoke(ctx, "/types.SchedulerService/ApplyStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) ApplyScan(ctx context.Context, in *ScannerCronJobRequest, opts ...grpc.CallOption) (*ScannerCronJobResponse, error) {
	out := new(ScannerCronJobResponse)
	err := c.cc.Invoke(ctx, "/types.SchedulerService/ApplyScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServiceServer is the server API for SchedulerService service.
type SchedulerServiceServer interface {
	// ApplyStream method will trigger a new scanning cron job and will return if it was triggered successfully
	ApplyStream(context.Context, *StreamsJobRequest) (*StreamsJobResponse, error)
	// Apply method will trigger a new scanning cron job and will return if it was triggered successfully
	ApplyScan(context.Context, *ScannerCronJobRequest) (*ScannerCronJobResponse, error)
}

func RegisterSchedulerServiceServer(s *grpc.Server, srv SchedulerServiceServer) {
	s.RegisterService(&_SchedulerService_serviceDesc, srv)
}

func _SchedulerService_ApplyStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamsJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ApplyStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.SchedulerService/ApplyStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ApplyStream(ctx, req.(*StreamsJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_ApplyScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScannerCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ApplyScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.SchedulerService/ApplyScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ApplyScan(ctx, req.(*ScannerCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchedulerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.SchedulerService",
	HandlerType: (*SchedulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyStream",
			Handler:    _SchedulerService_ApplyStream_Handler,
		},
		{
			MethodName: "ApplyScan",
			Handler:    _SchedulerService_ApplyScan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor_scheduler_e211eb632181759c) }

var fileDescriptor_scheduler_e211eb632181759c = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x7d, 0xf3, 0x1e, 0x4f, 0xc3, 0xc5, 0xa0, 0xdc, 0x14, 0x53, 0x1b, 0x4d, 0x48, 0x57, 0xac,
	0x58, 0xa0, 0x2b, 0xe3, 0x86, 0x68, 0x4c, 0x64, 0xe1, 0xa2, 0xe5, 0x07, 0x86, 0x72, 0x83, 0x24,
	0x74, 0x3a, 0xce, 0x0c, 0x26, 0x7c, 0x83, 0x3b, 0xbf, 0xc1, 0x95, 0x5f, 0x69, 0x68, 0x67, 0x62,
	0x5b, 0x8a, 0xbb, 0xf6, 0x9e, 0x73, 0xee, 0x39, 0x3d, 0x9d, 0x81, 0x53, 0x9d, 0xbc, 0xd0, 0x62,
	0xb3, 0x26, 0x35, 0x92, 0x2a, 0x33, 0x19, 0xfe, 0x37, 0x5b, 0x49, 0x3a, 0x38, 0xd1, 0x46, 0x11,
	0x4f, 0x8b, 0x61, 0x00, 0x3a, 0xe1, 0xc2, 0x3e, 0x77, 0x0d, 0xa5, 0x72, 0xcd, 0x0d, 0x15, 0xef,
	0xe1, 0x17, 0x03, 0x3f, 0x4e, 0xb8, 0x10, 0xa4, 0xee, 0x55, 0x26, 0xa6, 0xd9, 0x7c, 0x22, 0x57,
	0x11, 0xbd, 0x6e, 0x48, 0x1b, 0xbc, 0xad, 0x63, 0x33, 0x2b, 0x7e, 0x5a, 0xf8, 0x6c, 0xc0, 0x86,
	0xed, 0xc8, 0xd7, 0x07, 0x70, 0x8c, 0xa0, 0x5f, 0xc5, 0xec, 0x52, 0xff, 0xef, 0x80, 0x0d, 0x3b,
	0xe3, 0xcb, 0x51, 0x9e, 0x74, 0x14, 0x37, 0x71, 0xa2, 0x66, 0x69, 0xf8, 0xce, 0xa0, 0xdf, 0x28,
	0x40, 0x84, 0xd6, 0x33, 0x4f, 0xc9, 0xa6, 0x6a, 0x09, 0x9e, 0x12, 0x0e, 0xe1, 0xd8, 0xa8, 0xd5,
	0x72, 0x49, 0xca, 0x7a, 0x76, 0xad, 0xe7, 0xac, 0x98, 0x46, 0x0e, 0xc6, 0x1b, 0xe8, 0xec, 0x0c,
	0x5d, 0xc2, 0x7f, 0x39, 0x1b, 0x4b, 0x09, 0x5d, 0xae, 0x32, 0x2d, 0xf4, 0xe1, 0xbc, 0x1e, 0x46,
	0xcb, 0x4c, 0x68, 0x0a, 0x3f, 0x18, 0x78, 0x71, 0xfe, 0x07, 0x74, 0xb5, 0xd0, 0x71, 0x79, 0xbe,
	0x57, 0xa6, 0xa7, 0x1b, 0x30, 0x7c, 0x84, 0xde, 0xcf, 0xbc, 0x5a, 0xa2, 0xef, 0x22, 0xd6, 0xf1,
	0x68, 0x5f, 0x12, 0x12, 0xf4, 0xf6, 0x78, 0xbb, 0xde, 0x44, 0xbd, 0xb7, 0x3b, 0xe8, 0x5a, 0x75,
	0xd5, 0xcd, 0xab, 0xb8, 0x39, 0xa7, 0x1a, 0x37, 0xf4, 0x00, 0xcb, 0x36, 0x45, 0x23, 0xe3, 0x4f,
	0x06, 0x67, 0xb1, 0x3b, 0xab, 0x31, 0xa9, 0xb7, 0x55, 0x42, 0xf8, 0x00, 0x9d, 0x89, 0x94, 0xeb,
	0x6d, 0xc1, 0xc7, 0x83, 0x5f, 0x13, 0x5c, 0x34, 0x20, 0xb6, 0xea, 0x3f, 0x38, 0x85, 0x76, 0xb1,
	0x25, 0xe1, 0x02, 0x7f, 0x3d, 0x56, 0xc1, 0xd5, 0x01, 0xd4, 0xed, 0x9a, 0x1f, 0xe5, 0x97, 0xe2,
	0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x10, 0x7e, 0x31, 0x2a, 0x58, 0x03, 0x00, 0x00,
}
