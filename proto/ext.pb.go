// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ext.proto

package test_example

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

var E_Roles = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         4938,
	Name:          "test.example.roles",
	Tag:           "bytes,4938,opt,name=roles",
	Filename:      "proto/ext.proto",
}

func init() {
	proto.RegisterExtension(E_Roles)
}

func init() { proto.RegisterFile("proto/ext.proto", fileDescriptor_507b0bc5ae60d3fc) }

var fileDescriptor_507b0bc5ae60d3fc = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xad, 0x28, 0xd1, 0x03, 0xb3, 0x84, 0x78, 0x4a, 0x52, 0x8b, 0x4b, 0xf4, 0x52,
	0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0xa5, 0x14, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1,
	0x72, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x10,
	0xf5, 0x56, 0xa6, 0x5c, 0xac, 0x45, 0xf9, 0x39, 0xa9, 0xc5, 0x42, 0x72, 0x7a, 0x10, 0xb5, 0x7a,
	0x30, 0xb5, 0x7a, 0xbe, 0xa9, 0x25, 0x19, 0xf9, 0x29, 0xfe, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5,
	0x12, 0xa7, 0xd4, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0xaa, 0x93, 0xd8, 0xc0, 0xaa, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0x74, 0xd0, 0x39, 0x80, 0x00, 0x00, 0x00,
}
