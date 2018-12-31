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

// FieldTypes for Analyzing and Anonymizing
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
	FieldTypesEnum_UK_NHS            FieldTypesEnum = 16
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
	16: "UK_NHS",
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
	"UK_NHS":            16,
}

func (x FieldTypesEnum) String() string {
	return proto.EnumName(FieldTypesEnum_name, int32(x))
}
func (FieldTypesEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_1c74f1f93e5732a9, []int{0}
}

// FieldType strucy
type FieldTypes struct {
	// Field type name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Field type language code
	LanguageCode string `protobuf:"bytes,2,opt,name=languageCode,proto3" json:"languageCode,omitempty"`
	// The minScore will filter results which has lower certainty than the provided value.
	MinScore             string   `protobuf:"bytes,3,opt,name=minScore,proto3" json:"minScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldTypes) Reset()         { *m = FieldTypes{} }
func (m *FieldTypes) String() string { return proto.CompactTextString(m) }
func (*FieldTypes) ProtoMessage()    {}
func (*FieldTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_1c74f1f93e5732a9, []int{0}
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

func (m *FieldTypes) GetMinScore() string {
	if m != nil {
		return m.MinScore
	}
	return ""
}

// AnalyzeResult represents the Analyze serivce findings
type AnalyzeResult struct {
	// The sensitive text result
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The sensitive text type (supported types: FieldTypesEnum)
	Field *FieldTypes `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	// The score of the result
	Score float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	// The loaction in the text of the finding
	Location             *Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AnalyzeResult) Reset()         { *m = AnalyzeResult{} }
func (m *AnalyzeResult) String() string { return proto.CompactTextString(m) }
func (*AnalyzeResult) ProtoMessage()    {}
func (*AnalyzeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_1c74f1f93e5732a9, []int{1}
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

func (m *AnalyzeResult) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *AnalyzeResult) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

// The location in the text of the finding
type Location struct {
	// The location start
	Start int32 `protobuf:"zigzag32,1,opt,name=start,proto3" json:"start,omitempty"`
	// The location end
	End int32 `protobuf:"zigzag32,2,opt,name=end,proto3" json:"end,omitempty"`
	// The location length
	Length               int32    `protobuf:"zigzag32,3,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_1c74f1f93e5732a9, []int{2}
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

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_1c74f1f93e5732a9) }

var fileDescriptor_common_1c74f1f93e5732a9 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0xa5, 0xed, 0xf4, 0x75, 0x9b, 0xb6, 0xce, 0x15, 0xa0, 0x8a, 0x15, 0xea, 0x06, 0x04, 0x52,
	0x17, 0xf0, 0x05, 0x6e, 0x62, 0x34, 0x66, 0x1a, 0x27, 0xb2, 0x13, 0x24, 0x56, 0x26, 0x74, 0x4c,
	0xa9, 0x94, 0xc7, 0xa8, 0x4d, 0x25, 0x86, 0x9f, 0x60, 0xc5, 0xff, 0x22, 0x3b, 0x6d, 0x47, 0xec,
	0xce, 0xe3, 0xfa, 0x9c, 0xb3, 0x30, 0x78, 0xdb, 0xba, 0x2c, 0xeb, 0x6a, 0xf5, 0x70, 0xa8, 0x9b,
	0x1a, 0xfb, 0xcd, 0xe3, 0x83, 0x39, 0x2e, 0xbf, 0x01, 0x7c, 0xda, 0x9b, 0xe2, 0x3e, 0xb5, 0x0c,
	0x11, 0x6e, 0xaa, 0xbc, 0x34, 0x8b, 0xce, 0xeb, 0xce, 0xdb, 0xb1, 0x74, 0x18, 0x97, 0xe0, 0x15,
	0x79, 0xb5, 0x3b, 0xe5, 0x3b, 0x13, 0xd4, 0xf7, 0x66, 0xd1, 0x75, 0xde, 0x7f, 0x1a, 0xbe, 0x82,
	0x51, 0xb9, 0xaf, 0xd4, 0xb6, 0x3e, 0x98, 0x45, 0xcf, 0xf9, 0x57, 0xbe, 0xfc, 0xd3, 0x81, 0x29,
	0xad, 0xf2, 0xe2, 0xf1, 0xb7, 0x91, 0xe6, 0x78, 0x2a, 0x1a, 0xdb, 0xd2, 0x98, 0x5f, 0xcd, 0xa5,
	0xc5, 0x62, 0x7c, 0x03, 0xfd, 0x1f, 0x76, 0x87, 0x8b, 0x9f, 0x7c, 0xf0, 0x57, 0x6e, 0xde, 0xea,
	0x69, 0x9b, 0x6c, 0x7d, 0x7c, 0x0e, 0xfd, 0xe3, 0xb5, 0xa7, 0x2b, 0x5b, 0x82, 0xef, 0x61, 0x54,
	0xd4, 0xdb, 0xbc, 0xd9, 0xd7, 0xd5, 0xe2, 0xc6, 0x25, 0xcc, 0xcf, 0x09, 0x9b, 0xb3, 0x2c, 0xaf,
	0x07, 0xcb, 0xcf, 0x30, 0xba, 0xa8, 0x2e, 0xae, 0xc9, 0x0f, 0xed, 0x18, 0x5f, 0xb6, 0x04, 0x09,
	0xf4, 0x4c, 0xd5, 0x6e, 0xf1, 0xa5, 0x85, 0xf8, 0x12, 0x06, 0x85, 0xa9, 0x76, 0xcd, 0x4f, 0xd7,
	0xeb, 0xcb, 0x33, 0x7b, 0xf7, 0xb7, 0x0b, 0xb3, 0xa7, 0x91, 0xac, 0x3a, 0x95, 0x38, 0x87, 0x49,
	0x20, 0x59, 0xc8, 0x53, 0x1d, 0x50, 0x19, 0x92, 0x67, 0x08, 0x30, 0x08, 0xe4, 0xd7, 0x24, 0x8d,
	0x49, 0x07, 0xa7, 0x30, 0x0e, 0x69, 0xca, 0x74, 0xca, 0x23, 0x46, 0xba, 0xf6, 0x36, 0x8c, 0x23,
	0xca, 0x85, 0x16, 0x34, 0x62, 0xa4, 0x87, 0x3e, 0x4c, 0x59, 0x44, 0xf9, 0x46, 0xd3, 0x30, 0x94,
	0x4c, 0x29, 0x72, 0x63, 0x9f, 0xf0, 0x35, 0x15, 0x3a, 0x88, 0x43, 0x46, 0xfa, 0x38, 0x03, 0xe0,
	0xc9, 0xd5, 0x1e, 0xe0, 0x10, 0x7a, 0x42, 0x26, 0x64, 0x88, 0x1e, 0x8c, 0x36, 0x71, 0x40, 0x53,
	0x1e, 0x0b, 0x32, 0xb2, 0xa5, 0x09, 0x93, 0x2a, 0x16, 0x64, 0x8c, 0x04, 0xbc, 0xe4, 0x36, 0x16,
	0x4c, 0x8b, 0x2c, 0x5a, 0x33, 0x49, 0x00, 0x11, 0x66, 0x99, 0xd2, 0x6b, 0x2a, 0xee, 0x2e, 0xda,
	0x04, 0x5f, 0x80, 0x9f, 0x29, 0x1d, 0x4a, 0xfe, 0x85, 0x49, 0xbd, 0xe1, 0x01, 0x13, 0x8a, 0x11,
	0x0f, 0x27, 0x30, 0xcc, 0x94, 0xe6, 0x29, 0x17, 0x64, 0x6a, 0xf7, 0x66, 0x4a, 0x27, 0x54, 0xa9,
	0x24, 0x96, 0x29, 0x99, 0xd9, 0x9a, 0x4c, 0x69, 0xa5, 0x04, 0x99, 0x3b, 0x7c, 0xa7, 0xc5, 0xad,
	0x22, 0xe4, 0xfb, 0xc0, 0xfd, 0xb2, 0x8f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xac, 0x84,
	0x8e, 0x75, 0x02, 0x00, 0x00,
}
