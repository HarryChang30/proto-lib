// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/options.proto

package options

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_HttpMode = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50056,
	Name:          "httpMode",
	Tag:           "bytes,50056,opt,name=httpMode",
	Filename:      "proto/options.proto",
}

var E_Agregator = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50062,
	Name:          "agregator",
	Tag:           "bytes,50062,opt,name=agregator",
	Filename:      "proto/options.proto",
}

var E_Withelist = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50067,
	Name:          "withelist",
	Tag:           "bytes,50067,opt,name=withelist",
	Filename:      "proto/options.proto",
}

var E_Foo = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50001,
	Name:          "foo",
	Tag:           "bytes,50001,opt,name=foo",
	Filename:      "proto/options.proto",
}

var E_IsRepo = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50057,
	Name:          "isRepo",
	Tag:           "varint,50057,opt,name=isRepo",
	Filename:      "proto/options.proto",
}

var E_IsEs = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50064,
	Name:          "isEs",
	Tag:           "varint,50064,opt,name=isEs",
	Filename:      "proto/options.proto",
}

var E_IgnoreFieldDb = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50058,
	Name:          "ignoreFieldDb",
	Tag:           "varint,50058,opt,name=ignoreFieldDb",
	Filename:      "proto/options.proto",
}

var E_IsPrimaryKey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50059,
	Name:          "isPrimaryKey",
	Tag:           "varint,50059,opt,name=isPrimaryKey",
	Filename:      "proto/options.proto",
}

var E_Required = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50060,
	Name:          "required",
	Tag:           "varint,50060,opt,name=required",
	Filename:      "proto/options.proto",
}

var E_RequiredType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50061,
	Name:          "required_type",
	Tag:           "bytes,50061,opt,name=required_type,json=requiredType",
	Filename:      "proto/options.proto",
}

var E_Fulltext = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50063,
	Name:          "fulltext",
	Tag:           "varint,50063,opt,name=fulltext",
	Filename:      "proto/options.proto",
}

var E_ErrDesc = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50065,
	Name:          "errDesc",
	Tag:           "bytes,50065,opt,name=errDesc",
	Filename:      "proto/options.proto",
}

var E_ForeignKey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50066,
	Name:          "foreignKey",
	Tag:           "bytes,50066,opt,name=foreignKey",
	Filename:      "proto/options.proto",
}

func init() {
	proto.RegisterExtension(E_HttpMode)
	proto.RegisterExtension(E_Agregator)
	proto.RegisterExtension(E_Withelist)
	proto.RegisterExtension(E_Foo)
	proto.RegisterExtension(E_IsRepo)
	proto.RegisterExtension(E_IsEs)
	proto.RegisterExtension(E_IgnoreFieldDb)
	proto.RegisterExtension(E_IsPrimaryKey)
	proto.RegisterExtension(E_Required)
	proto.RegisterExtension(E_RequiredType)
	proto.RegisterExtension(E_Fulltext)
	proto.RegisterExtension(E_ErrDesc)
	proto.RegisterExtension(E_ForeignKey)
}

func init() { proto.RegisterFile("proto/options.proto", fileDescriptor_options_81053feb3b46270a) }

var fileDescriptor_options_81053feb3b46270a = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x07, 0xf0, 0x18, 0x0c, 0xc2, 0x06, 0x2e, 0x78, 0x31, 0x26, 0x2a, 0x47, 0x4f, 0xe5, 0x40,
	0x3c, 0xb0, 0x1a, 0x3d, 0x08, 0x5e, 0x0c, 0xd1, 0x10, 0xef, 0xa6, 0xd0, 0xe9, 0xb2, 0x49, 0x65,
	0xd6, 0xd9, 0x25, 0xda, 0x17, 0xf0, 0xfb, 0xdb, 0x27, 0xf3, 0x8d, 0x0c, 0xed, 0x6e, 0x95, 0x68,
	0xb2, 0xdc, 0x3a, 0xe9, 0xff, 0xf7, 0x9f, 0x49, 0xcb, 0x56, 0x15, 0xa1, 0xc1, 0x16, 0x2a, 0x23,
	0x71, 0xa2, 0x83, 0x6c, 0x5a, 0x6f, 0x0a, 0x44, 0x91, 0x40, 0x2b, 0x9b, 0x86, 0xd3, 0xb8, 0x15,
	0x81, 0x1e, 0x91, 0x54, 0x06, 0x29, 0x4f, 0xf0, 0x3d, 0x56, 0x19, 0x1b, 0xa3, 0xfa, 0x18, 0x41,
	0x63, 0x33, 0xc8, 0xe3, 0x81, 0x8b, 0x07, 0x7d, 0x30, 0x63, 0x8c, 0x4e, 0xf2, 0xce, 0xb5, 0xbb,
	0xdb, 0x52, 0x73, 0x69, 0xbb, 0x3a, 0x28, 0x04, 0xdf, 0x67, 0xd5, 0x50, 0x10, 0x88, 0xd0, 0x20,
	0x79, 0xf9, 0x8b, 0xe5, 0x3f, 0x64, 0xe6, 0xaf, 0xa4, 0x19, 0x43, 0x22, 0xb5, 0xf1, 0xfa, 0x4f,
	0xe7, 0x0b, 0xc2, 0xdb, 0xac, 0x14, 0x23, 0x36, 0xb6, 0xfe, 0x91, 0x5a, 0x87, 0x02, 0x1c, 0xfd,
	0xba, 0xc9, 0xe9, 0x2c, 0xcd, 0x3b, 0xac, 0x2c, 0xf5, 0x00, 0xd4, 0x02, 0xee, 0x3e, 0x5b, 0x59,
	0x19, 0x58, 0xc0, 0x77, 0xd8, 0xb2, 0xd4, 0x3d, 0xed, 0x87, 0x6f, 0x16, 0x66, 0x71, 0xde, 0x63,
	0x75, 0x29, 0x26, 0x48, 0x70, 0x24, 0x21, 0x89, 0xba, 0xc3, 0xc6, 0xc6, 0x1f, 0x9f, 0xbd, 0x71,
	0xfa, 0xc1, 0xea, 0x79, 0xc5, 0x0f, 0x59, 0x4d, 0xea, 0x53, 0x92, 0x17, 0x21, 0xa5, 0xc7, 0x90,
	0xfa, 0x5a, 0x1e, 0x6d, 0xcb, 0x1c, 0xe2, 0xbb, 0xac, 0x42, 0x70, 0x39, 0x95, 0x04, 0x91, 0xaf,
	0xe0, 0xc9, 0x16, 0x14, 0x80, 0x77, 0x59, 0xdd, 0x3d, 0x9f, 0x9b, 0x54, 0x81, 0xaf, 0xe1, 0xd9,
	0xfe, 0xb2, 0x9a, 0x53, 0x67, 0xa9, 0x82, 0xd9, 0x09, 0xf1, 0x34, 0x49, 0x0c, 0x5c, 0x1b, 0x5f,
	0xc1, 0xab, 0x3b, 0xc1, 0x01, 0xde, 0x61, 0x2b, 0x40, 0xd4, 0x05, 0x3d, 0xf2, 0xd9, 0x77, 0xbb,
	0xdc, 0xe5, 0xf9, 0x01, 0x63, 0x31, 0x12, 0x48, 0x31, 0x59, 0xe0, 0xeb, 0x7d, 0x58, 0xfd, 0x8b,
	0x0c, 0xcb, 0x59, 0xb4, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x10, 0x4a, 0x35, 0xa3, 0x6c, 0x03,
	0x00, 0x00,
}
