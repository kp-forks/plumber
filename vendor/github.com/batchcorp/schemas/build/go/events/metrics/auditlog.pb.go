// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/metrics/auditlog.proto

package metrics

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AuditLog_Resource int32

const (
	AuditLog_UNSET       AuditLog_Resource = 0
	AuditLog_COLLECTION  AuditLog_Resource = 1
	AuditLog_REPLAY      AuditLog_Resource = 2
	AuditLog_STORAGE     AuditLog_Resource = 3
	AuditLog_DESTINATION AuditLog_Resource = 4
	AuditLog_SCHEMA      AuditLog_Resource = 5
	AuditLog_SEARCH      AuditLog_Resource = 6
	AuditLog_TASK        AuditLog_Resource = 7
	AuditLog_BILLING     AuditLog_Resource = 8
)

var AuditLog_Resource_name = map[int32]string{
	0: "UNSET",
	1: "COLLECTION",
	2: "REPLAY",
	3: "STORAGE",
	4: "DESTINATION",
	5: "SCHEMA",
	6: "SEARCH",
	7: "TASK",
	8: "BILLING",
}

var AuditLog_Resource_value = map[string]int32{
	"UNSET":       0,
	"COLLECTION":  1,
	"REPLAY":      2,
	"STORAGE":     3,
	"DESTINATION": 4,
	"SCHEMA":      5,
	"SEARCH":      6,
	"TASK":        7,
	"BILLING":     8,
}

func (x AuditLog_Resource) String() string {
	return proto.EnumName(AuditLog_Resource_name, int32(x))
}

func (AuditLog_Resource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae2ce3e11a0fe8d5, []int{0, 0}
}

type AuditLog_Type int32

const (
	AuditLog_UNKNOWN AuditLog_Type = 0
	AuditLog_INFO    AuditLog_Type = 1
	AuditLog_WARNING AuditLog_Type = 2
	AuditLog_ERROR   AuditLog_Type = 3
	AuditLog_FATAL   AuditLog_Type = 4
)

var AuditLog_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "FATAL",
}

var AuditLog_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
	"FATAL":   4,
}

func (x AuditLog_Type) String() string {
	return proto.EnumName(AuditLog_Type_name, int32(x))
}

func (AuditLog_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae2ce3e11a0fe8d5, []int{0, 1}
}

// Used for internal error reporting using mlib library and metrics service.
type AuditLog struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               string            `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Resource             AuditLog_Resource `protobuf:"varint,3,opt,name=resource,proto3,enum=metrics.AuditLog_Resource" json:"resource,omitempty"`
	Source               string            `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Value                string            `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            int64             `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Count                int64             `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	Type                 AuditLog_Type     `protobuf:"varint,8,opt,name=type,proto3,enum=metrics.AuditLog_Type" json:"type,omitempty"`
	Metadata             []byte            `protobuf:"bytes,9,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuditLog) Reset()         { *m = AuditLog{} }
func (m *AuditLog) String() string { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()    {}
func (*AuditLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae2ce3e11a0fe8d5, []int{0}
}

func (m *AuditLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditLog.Unmarshal(m, b)
}
func (m *AuditLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLog.Marshal(b, m, deterministic)
}
func (m *AuditLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLog.Merge(m, src)
}
func (m *AuditLog) XXX_Size() int {
	return xxx_messageInfo_AuditLog.Size(m)
}
func (m *AuditLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLog.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLog proto.InternalMessageInfo

func (m *AuditLog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuditLog) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *AuditLog) GetResource() AuditLog_Resource {
	if m != nil {
		return m.Resource
	}
	return AuditLog_UNSET
}

func (m *AuditLog) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AuditLog) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *AuditLog) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AuditLog) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *AuditLog) GetType() AuditLog_Type {
	if m != nil {
		return m.Type
	}
	return AuditLog_UNKNOWN
}

func (m *AuditLog) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("metrics.AuditLog_Resource", AuditLog_Resource_name, AuditLog_Resource_value)
	proto.RegisterEnum("metrics.AuditLog_Type", AuditLog_Type_name, AuditLog_Type_value)
	proto.RegisterType((*AuditLog)(nil), "metrics.AuditLog")
}

func init() { proto.RegisterFile("events/metrics/auditlog.proto", fileDescriptor_ae2ce3e11a0fe8d5) }

var fileDescriptor_ae2ce3e11a0fe8d5 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6b, 0x9c, 0x40,
	0x14, 0xc7, 0xe3, 0x8f, 0x55, 0xf7, 0xa5, 0x6c, 0x87, 0xa1, 0xa4, 0x12, 0x5a, 0x58, 0xf6, 0xb4,
	0xf4, 0xa0, 0xd0, 0x96, 0x9c, 0x3b, 0xd9, 0x9a, 0x44, 0x62, 0xc6, 0x32, 0x1a, 0x42, 0x7b, 0x29,
	0xb3, 0x3a, 0xb8, 0xc2, 0x9a, 0x11, 0x1d, 0x03, 0xb9, 0xf4, 0xdf, 0xec, 0xbf, 0x53, 0xc6, 0x35,
	0x5b, 0x4a, 0x6f, 0xf3, 0xfd, 0xcc, 0xe7, 0x3d, 0xbe, 0x87, 0x07, 0xef, 0xc5, 0x93, 0x78, 0x54,
	0x7d, 0xd8, 0x08, 0xd5, 0xd5, 0x45, 0x1f, 0xf2, 0xa1, 0xac, 0xd5, 0x5e, 0x56, 0x41, 0xdb, 0x49,
	0x25, 0xb1, 0x3b, 0xf1, 0xd5, 0x6f, 0x0b, 0x3c, 0xa2, 0xff, 0x12, 0x59, 0xe1, 0x05, 0x98, 0x75,
	0xe9, 0x1b, 0x4b, 0x63, 0x3d, 0x67, 0x66, 0x5d, 0xe2, 0xb7, 0xe0, 0x2a, 0xc1, 0x9b, 0x9f, 0x75,
	0xe9, 0x9b, 0x23, 0x74, 0x74, 0x8c, 0x4b, 0x7c, 0x01, 0x5e, 0x27, 0x7a, 0x39, 0x74, 0x85, 0xf0,
	0xad, 0xa5, 0xb1, 0x5e, 0x7c, 0x3c, 0x0f, 0xa6, 0x8d, 0xc1, 0xcb, 0xb6, 0x80, 0x4d, 0x06, 0x3b,
	0xba, 0xf8, 0x0c, 0x9c, 0x69, 0xca, 0x3e, 0xec, 0x9b, 0xf8, 0x1b, 0x98, 0x3d, 0xf1, 0xfd, 0x20,
	0xfc, 0xd9, 0x88, 0x0f, 0x01, 0xbf, 0x83, 0xb9, 0xaa, 0x1b, 0xd1, 0x2b, 0xde, 0xb4, 0xbe, 0xb3,
	0x34, 0xd6, 0x16, 0xfb, 0x0b, 0xf4, 0x4c, 0x21, 0x87, 0x47, 0xe5, 0xbb, 0xe3, 0xcf, 0x21, 0xe0,
	0x0f, 0x60, 0xab, 0xe7, 0x56, 0xf8, 0xde, 0xd8, 0xea, 0xec, 0xff, 0x56, 0xf9, 0x73, 0x2b, 0xd8,
	0xe8, 0xe0, 0x73, 0xf0, 0xee, 0x84, 0xe2, 0x25, 0x57, 0xdc, 0x9f, 0x2f, 0x8d, 0xf5, 0x2b, 0x76,
	0xcc, 0xab, 0x5f, 0xe0, 0xbd, 0xf4, 0xc7, 0x73, 0x98, 0xdd, 0xd3, 0x2c, 0xca, 0xd1, 0x09, 0x5e,
	0x00, 0x6c, 0xd2, 0x24, 0x89, 0x36, 0x79, 0x9c, 0x52, 0x64, 0x60, 0x00, 0x87, 0x45, 0xdf, 0x12,
	0xf2, 0x1d, 0x99, 0xf8, 0x14, 0xdc, 0x2c, 0x4f, 0x19, 0xb9, 0x8e, 0x90, 0x85, 0x5f, 0xc3, 0xe9,
	0xd7, 0x28, 0xcb, 0x63, 0x4a, 0x46, 0xd3, 0xd6, 0x66, 0xb6, 0xb9, 0x89, 0xee, 0x08, 0x9a, 0x8d,
	0xef, 0x88, 0xb0, 0xcd, 0x0d, 0x72, 0xb0, 0x07, 0x76, 0x4e, 0xb2, 0x5b, 0xe4, 0xea, 0xf9, 0xcb,
	0x38, 0x49, 0x62, 0x7a, 0x8d, 0xbc, 0xd5, 0x17, 0xb0, 0x75, 0x53, 0x0d, 0xef, 0xe9, 0x2d, 0x4d,
	0x1f, 0x28, 0x3a, 0xd1, 0x6e, 0x4c, 0xaf, 0x52, 0x64, 0x68, 0xfc, 0x40, 0x18, 0xd5, 0xae, 0xa9,
	0xfb, 0x45, 0x8c, 0xa5, 0x0c, 0x59, 0xfa, 0x79, 0x45, 0x72, 0x92, 0x20, 0xfb, 0xf2, 0xe2, 0xc7,
	0xe7, 0xaa, 0x56, 0xbb, 0x61, 0x1b, 0x14, 0xb2, 0x09, 0xb7, 0x5c, 0x15, 0xbb, 0x42, 0x76, 0x6d,
	0xd8, 0x17, 0x3b, 0xd1, 0xf0, 0x3e, 0xdc, 0x0e, 0xf5, 0xbe, 0x0c, 0x2b, 0x19, 0xfe, 0x7b, 0x29,
	0x5b, 0x67, 0xbc, 0x90, 0x4f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0xee, 0x93, 0x8b, 0x42,
	0x02, 0x00, 0x00,
}