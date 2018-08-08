// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldTypesEnum int32

const (
	FieldTypesEnum_CREDIT_CARD       FieldTypesEnum = 0
	FieldTypesEnum_CRYPTO            FieldTypesEnum = 1
	FieldTypesEnum_DATE_TIME         FieldTypesEnum = 2
	FieldTypesEnum_DOMAIN_NAME       FieldTypesEnum = 3
	FieldTypesEnum_EMAIL_ADDRESS     FieldTypesEnum = 4
	FieldTypesEnum_IBAN_CODE         FieldTypesEnum = 5
	FieldTypesEnum_IP_ADDRESS        FieldTypesEnum = 6
	FieldTypesEnum_NRP               FieldTypesEnum = 7
	FieldTypesEnum_LOCATION          FieldTypesEnum = 8
	FieldTypesEnum_PERSON            FieldTypesEnum = 9
	FieldTypesEnum_PHONE_NUMBER      FieldTypesEnum = 10
	FieldTypesEnum_US_BANK_NUMBER    FieldTypesEnum = 11
	FieldTypesEnum_US_DRIVER_LICENSE FieldTypesEnum = 12
	FieldTypesEnum_US_ITIN           FieldTypesEnum = 13
	FieldTypesEnum_US_PASSPORT       FieldTypesEnum = 14
	FieldTypesEnum_US_SSN            FieldTypesEnum = 15
)

var FieldTypesEnum_name = map[int32]string{
	0:  "CREDIT_CARD",
	1:  "CRYPTO",
	2:  "DATE_TIME",
	3:  "DOMAIN_NAME",
	4:  "EMAIL_ADDRESS",
	5:  "IBAN_CODE",
	6:  "IP_ADDRESS",
	7:  "NRP",
	8:  "LOCATION",
	9:  "PERSON",
	10: "PHONE_NUMBER",
	11: "US_BANK_NUMBER",
	12: "US_DRIVER_LICENSE",
	13: "US_ITIN",
	14: "US_PASSPORT",
	15: "US_SSN",
}
var FieldTypesEnum_value = map[string]int32{
	"CREDIT_CARD":       0,
	"CRYPTO":            1,
	"DATE_TIME":         2,
	"DOMAIN_NAME":       3,
	"EMAIL_ADDRESS":     4,
	"IBAN_CODE":         5,
	"IP_ADDRESS":        6,
	"NRP":               7,
	"LOCATION":          8,
	"PERSON":            9,
	"PHONE_NUMBER":      10,
	"US_BANK_NUMBER":    11,
	"US_DRIVER_LICENSE": 12,
	"US_ITIN":           13,
	"US_PASSPORT":       14,
	"US_SSN":            15,
}

func (x FieldTypesEnum) String() string {
	return proto.EnumName(FieldTypesEnum_name, int32(x))
}
func (FieldTypesEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_b7ed14342eac292c, []int{0}
}

type FieldTypes struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LanguageCode         string   `protobuf:"bytes,2,opt,name=languageCode,proto3" json:"languageCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldTypes) Reset()         { *m = FieldTypes{} }
func (m *FieldTypes) String() string { return proto.CompactTextString(m) }
func (*FieldTypes) ProtoMessage()    {}
func (*FieldTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b7ed14342eac292c, []int{0}
}
func (m *FieldTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldTypes.Unmarshal(m, b)
}
func (m *FieldTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldTypes.Marshal(b, m, deterministic)
}
func (dst *FieldTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldTypes.Merge(dst, src)
}
func (m *FieldTypes) XXX_Size() int {
	return xxx_messageInfo_FieldTypes.Size(m)
}
func (m *FieldTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldTypes.DiscardUnknown(m)
}

var xxx_messageInfo_FieldTypes proto.InternalMessageInfo

func (m *FieldTypes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldTypes) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

type AnalyzeResult struct {
	Text                 string      `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Field                *FieldTypes `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Probability          float32     `protobuf:"fixed32,3,opt,name=probability,proto3" json:"probability,omitempty"`
	Location             *Location   `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AnalyzeResult) Reset()         { *m = AnalyzeResult{} }
func (m *AnalyzeResult) String() string { return proto.CompactTextString(m) }
func (*AnalyzeResult) ProtoMessage()    {}
func (*AnalyzeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b7ed14342eac292c, []int{1}
}
func (m *AnalyzeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeResult.Unmarshal(m, b)
}
func (m *AnalyzeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeResult.Marshal(b, m, deterministic)
}
func (dst *AnalyzeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeResult.Merge(dst, src)
}
func (m *AnalyzeResult) XXX_Size() int {
	return xxx_messageInfo_AnalyzeResult.Size(m)
}
func (m *AnalyzeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeResult.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeResult proto.InternalMessageInfo

func (m *AnalyzeResult) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AnalyzeResult) GetField() *FieldTypes {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *AnalyzeResult) GetProbability() float32 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func (m *AnalyzeResult) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Location struct {
	Start                int32    `protobuf:"zigzag32,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int32    `protobuf:"zigzag32,2,opt,name=end,proto3" json:"end,omitempty"`
	Length               int32    `protobuf:"zigzag32,3,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b7ed14342eac292c, []int{2}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Location) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Location) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func init() {
	proto.RegisterType((*FieldTypes)(nil), "types.FieldTypes")
	proto.RegisterType((*AnalyzeResult)(nil), "types.AnalyzeResult")
	proto.RegisterType((*Location)(nil), "types.Location")
	proto.RegisterEnum("types.FieldTypesEnum", FieldTypesEnum_name, FieldTypesEnum_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_b7ed14342eac292c) }

var fileDescriptor_common_b7ed14342eac292c = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xcd, 0xae, 0xd2, 0x40,
	0x18, 0xb5, 0xfc, 0xf3, 0xb5, 0xc0, 0xf4, 0x8b, 0x1a, 0x96, 0x84, 0x8d, 0x37, 0x9a, 0xb0, 0xd0,
	0x27, 0x18, 0xda, 0x31, 0x8e, 0xd2, 0x69, 0x33, 0xd3, 0x9a, 0xb8, 0x9a, 0x94, 0x7b, 0x47, 0x24,
	0x29, 0x2d, 0x81, 0x92, 0x88, 0x6f, 0xe0, 0x33, 0xf8, 0xb2, 0x66, 0x0a, 0x5c, 0x74, 0x77, 0x7e,
	0xbe, 0x39, 0xe7, 0x2c, 0x06, 0xbc, 0xc7, 0x6a, 0xb7, 0xab, 0xca, 0xc5, 0xfe, 0x50, 0xd5, 0x15,
	0x76, 0xeb, 0xf3, 0xde, 0x1c, 0xe7, 0x21, 0xc0, 0xc7, 0xad, 0x29, 0x9e, 0x52, 0xcb, 0x10, 0xa1,
	0x53, 0xe6, 0x3b, 0x33, 0x75, 0x66, 0xce, 0xc3, 0x50, 0x36, 0x18, 0xe7, 0xe0, 0x15, 0x79, 0xb9,
	0x39, 0xe5, 0x1b, 0x13, 0x54, 0x4f, 0x66, 0xda, 0x6a, 0xbc, 0xff, 0xb4, 0xf9, 0x1f, 0x07, 0x46,
	0xb4, 0xcc, 0x8b, 0xf3, 0x2f, 0x23, 0xcd, 0xf1, 0x54, 0xd4, 0x36, 0xa9, 0x36, 0x3f, 0xeb, 0x5b,
	0x92, 0xc5, 0xf8, 0x06, 0xba, 0xdf, 0x6d, 0x57, 0x13, 0xe1, 0xbe, 0xf7, 0x17, 0xcd, 0x84, 0xc5,
	0xbd, 0x5f, 0x5e, 0x7c, 0x9c, 0x81, 0xbb, 0x3f, 0x54, 0xeb, 0x7c, 0xbd, 0x2d, 0xb6, 0xf5, 0x79,
	0xda, 0x9e, 0x39, 0x0f, 0x2d, 0xf9, 0xaf, 0x84, 0xef, 0x60, 0x50, 0x54, 0x8f, 0x79, 0xbd, 0xad,
	0xca, 0x69, 0xa7, 0x49, 0x9b, 0x5c, 0xd3, 0x56, 0x57, 0x59, 0x3e, 0x1f, 0xcc, 0x3f, 0xc3, 0xe0,
	0xa6, 0xe2, 0x4b, 0xe8, 0x1e, 0xeb, 0xfc, 0x70, 0x19, 0xe6, 0xcb, 0x0b, 0x41, 0x02, 0x6d, 0x53,
	0x5e, 0x76, 0xf9, 0xd2, 0x42, 0x7c, 0x0d, 0xbd, 0xc2, 0x94, 0x9b, 0xfa, 0x47, 0xd3, 0xee, 0xcb,
	0x2b, 0x7b, 0xfb, 0xbb, 0x05, 0xe3, 0xfb, 0x60, 0x56, 0x9e, 0x76, 0x38, 0x01, 0x37, 0x90, 0x2c,
	0xe4, 0xa9, 0x0e, 0xa8, 0x0c, 0xc9, 0x0b, 0x04, 0xe8, 0x05, 0xf2, 0x5b, 0x92, 0xc6, 0xc4, 0xc1,
	0x11, 0x0c, 0x43, 0x9a, 0x32, 0x9d, 0xf2, 0x88, 0x91, 0x96, 0xbd, 0x0d, 0xe3, 0x88, 0x72, 0xa1,
	0x05, 0x8d, 0x18, 0x69, 0xa3, 0x0f, 0x23, 0x16, 0x51, 0xbe, 0xd2, 0x34, 0x0c, 0x25, 0x53, 0x8a,
	0x74, 0xec, 0x13, 0xbe, 0xa4, 0x42, 0x07, 0x71, 0xc8, 0x48, 0x17, 0xc7, 0x00, 0x3c, 0x79, 0xb6,
	0x7b, 0xd8, 0x87, 0xb6, 0x90, 0x09, 0xe9, 0xa3, 0x07, 0x83, 0x55, 0x1c, 0xd0, 0x94, 0xc7, 0x82,
	0x0c, 0x6c, 0x69, 0xc2, 0xa4, 0x8a, 0x05, 0x19, 0x22, 0x01, 0x2f, 0xf9, 0x14, 0x0b, 0xa6, 0x45,
	0x16, 0x2d, 0x99, 0x24, 0x80, 0x08, 0xe3, 0x4c, 0xe9, 0x25, 0x15, 0x5f, 0x6e, 0x9a, 0x8b, 0xaf,
	0xc0, 0xcf, 0x94, 0x0e, 0x25, 0xff, 0xca, 0xa4, 0x5e, 0xf1, 0x80, 0x09, 0xc5, 0x88, 0x87, 0x2e,
	0xf4, 0x33, 0xa5, 0x79, 0xca, 0x05, 0x19, 0xd9, 0xbd, 0x99, 0xd2, 0x09, 0x55, 0x2a, 0x89, 0x65,
	0x4a, 0xc6, 0xb6, 0x26, 0x53, 0x5a, 0x29, 0x41, 0x26, 0xeb, 0x5e, 0xf3, 0x93, 0x3e, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x6f, 0xf8, 0x58, 0x5c, 0x59, 0x02, 0x00, 0x00,
}
