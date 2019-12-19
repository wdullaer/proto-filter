// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter.proto

package filter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type ValueFilter struct {
	Include              []string `protobuf:"bytes,1,rep,name=include" json:"include,omitempty"`
	Exclude              []string `protobuf:"bytes,2,rep,name=exclude" json:"exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueFilter) Reset()         { *m = ValueFilter{} }
func (m *ValueFilter) String() string { return proto.CompactTextString(m) }
func (*ValueFilter) ProtoMessage()    {}
func (*ValueFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{0}
}

func (m *ValueFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueFilter.Unmarshal(m, b)
}
func (m *ValueFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueFilter.Marshal(b, m, deterministic)
}
func (m *ValueFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueFilter.Merge(m, src)
}
func (m *ValueFilter) XXX_Size() int {
	return xxx_messageInfo_ValueFilter.Size(m)
}
func (m *ValueFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ValueFilter proto.InternalMessageInfo

func (m *ValueFilter) GetInclude() []string {
	if m != nil {
		return m.Include
	}
	return nil
}

func (m *ValueFilter) GetExclude() []string {
	if m != nil {
		return m.Exclude
	}
	return nil
}

var E_File = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*ValueFilter)(nil),
	Field:         61255,
	Name:          "filter.file",
	Tag:           "bytes,61255,opt,name=file",
	Filename:      "filter.proto",
}

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*ValueFilter)(nil),
	Field:         61255,
	Name:          "filter.service",
	Tag:           "bytes,61255,opt,name=service",
	Filename:      "filter.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*ValueFilter)(nil),
	Field:         61255,
	Name:          "filter.method",
	Tag:           "bytes,61255,opt,name=method",
	Filename:      "filter.proto",
}

var E_Enum = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*ValueFilter)(nil),
	Field:         61255,
	Name:          "filter.enum",
	Tag:           "bytes,61255,opt,name=enum",
	Filename:      "filter.proto",
}

var E_EnumValue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*ValueFilter)(nil),
	Field:         61255,
	Name:          "filter.enum_value",
	Tag:           "bytes,61255,opt,name=enum_value",
	Filename:      "filter.proto",
}

var E_Message = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*ValueFilter)(nil),
	Field:         61255,
	Name:          "filter.message",
	Tag:           "bytes,61255,opt,name=message",
	Filename:      "filter.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*ValueFilter)(nil),
	Field:         61255,
	Name:          "filter.field",
	Tag:           "bytes,61255,opt,name=field",
	Filename:      "filter.proto",
}

func init() {
	proto.RegisterType((*ValueFilter)(nil), "filter.ValueFilter")
	proto.RegisterExtension(E_File)
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Method)
	proto.RegisterExtension(E_Enum)
	proto.RegisterExtension(E_EnumValue)
	proto.RegisterExtension(E_Message)
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("filter.proto", fileDescriptor_1f5303cab7a20d6f) }

var fileDescriptor_1f5303cab7a20d6f = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xa9, 0xd5, 0x96, 0x4e, 0x3d, 0xad, 0x97, 0x20, 0x7e, 0x44, 0x4f, 0x3d, 0x6d, 0xc1,
	0xe3, 0xde, 0x3c, 0x58, 0x41, 0x08, 0x4a, 0x05, 0x3d, 0x4a, 0x4d, 0x26, 0x71, 0x61, 0x93, 0x0d,
	0xd9, 0xdd, 0xe2, 0x1b, 0xfa, 0x16, 0x3e, 0x8b, 0xec, 0x47, 0x40, 0xe8, 0x1e, 0xf6, 0x14, 0x26,
	0xbf, 0x99, 0xdf, 0x64, 0xfe, 0x81, 0xd3, 0x9a, 0x0b, 0x8d, 0x03, 0xed, 0x07, 0xa9, 0x25, 0x99,
	0xf9, 0xea, 0x3c, 0x6f, 0xa4, 0x6c, 0x04, 0xae, 0xdd, 0xdb, 0x4f, 0x53, 0xaf, 0x2b, 0x54, 0xe5,
	0xc0, 0x7b, 0x2d, 0x43, 0xe7, 0xed, 0x3d, 0x2c, 0xdf, 0x76, 0xc2, 0xe0, 0xc6, 0x0d, 0x90, 0x0c,
	0xe6, 0xbc, 0x2b, 0x85, 0xa9, 0x30, 0x9b, 0xe4, 0xd3, 0xd5, 0x62, 0x3b, 0x96, 0x96, 0xe0, 0xb7,
	0x27, 0x47, 0x9e, 0x84, 0x92, 0x3d, 0xc2, 0x71, 0xcd, 0x05, 0x92, 0x0b, 0xea, 0xb7, 0xd1, 0x71,
	0x1b, 0xdd, 0x70, 0x81, 0xcf, 0xbd, 0xe6, 0xb2, 0x53, 0xd9, 0xcf, 0xef, 0x34, 0x9f, 0xac, 0x96,
	0x77, 0x67, 0x34, 0x7c, 0xe9, 0xbf, 0xb5, 0x5b, 0x27, 0x60, 0x2f, 0x30, 0x57, 0x38, 0xec, 0x79,
	0x89, 0xe4, 0xfa, 0xc0, 0xf5, 0xea, 0x49, 0x92, 0x6e, 0xd4, 0xb0, 0x02, 0x66, 0x2d, 0xea, 0x2f,
	0x59, 0x91, 0xab, 0x03, 0x61, 0xe1, 0x40, 0x92, 0x2f, 0x48, 0xec, 0xa5, 0xd8, 0x99, 0x36, 0x72,
	0xe9, 0x43, 0x67, 0xda, 0xb4, 0x4b, 0xad, 0x80, 0xbd, 0x03, 0xd8, 0xe7, 0xc7, 0xde, 0x12, 0x72,
	0x13, 0xd5, 0xb9, 0xa9, 0x24, 0xe7, 0x02, 0xc7, 0x76, 0x1b, 0x61, 0x8b, 0x4a, 0xed, 0x9a, 0x58,
	0x84, 0x85, 0x27, 0x69, 0x11, 0x06, 0x0d, 0x7b, 0x82, 0x93, 0x9a, 0xa3, 0xa8, 0xc8, 0x65, 0xe4,
	0xf7, 0xa2, 0x48, 0x0b, 0xd0, 0x2b, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0xd5, 0x63, 0xbb,
	0xa5, 0x02, 0x00, 0x00,
}