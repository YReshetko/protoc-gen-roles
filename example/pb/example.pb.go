// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/example.proto

package pb

import (
	fmt "fmt"
	_ "github.com/YReshetko/protoc-gen-roles/proto"
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

type Request struct {
	// request body
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c78cffa5d645ba4, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Response struct {
	// response body
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c78cffa5d645ba4, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "test.example.Request")
	proto.RegisterType((*Response)(nil), "test.example.Response")
}

func init() { proto.RegisterFile("example/example.proto", fileDescriptor_1c78cffa5d645ba4) }

var fileDescriptor_1c78cffa5d645ba4 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd1, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x5a, 0x24, 0x0a, 0xa7, 0x02, 0xc5, 0xa8, 0x15, 0x2a, 0x0c, 0x28, 0x13, 0x43,
	0x9b, 0x48, 0x20, 0x26, 0xa6, 0x22, 0x16, 0x84, 0x58, 0x1c, 0x15, 0x04, 0x1b, 0x49, 0xaf, 0x49,
	0xd4, 0xc6, 0x67, 0xec, 0x0b, 0xf4, 0xc1, 0xfa, 0x1a, 0x2c, 0x7d, 0x22, 0xe4, 0x24, 0x08, 0x04,
	0x13, 0x30, 0xd9, 0xfa, 0x6d, 0x7d, 0x3e, 0xdf, 0x41, 0x17, 0x17, 0x4f, 0xb9, 0x9e, 0x63, 0x50,
	0xaf, 0xbe, 0x36, 0xc4, 0x24, 0xda, 0x8c, 0x96, 0xfd, 0x3a, 0xeb, 0x9f, 0x27, 0x19, 0xa7, 0x45,
	0xe4, 0xc7, 0x94, 0x07, 0x0f, 0x12, 0x6d, 0x8a, 0x3c, 0xa3, 0xa0, 0xbc, 0x19, 0x0f, 0x13, 0x54,
	0x43, 0x43, 0x73, 0xb4, 0x55, 0x10, 0xe0, 0x82, 0x2b, 0xc4, 0x3b, 0x84, 0x96, 0xc4, 0xe7, 0x02,
	0x2d, 0x8b, 0x0e, 0xac, 0xe7, 0x36, 0x39, 0x68, 0x1c, 0x37, 0x4e, 0xb6, 0xa4, 0xdb, 0x7a, 0x47,
	0xb0, 0x29, 0xd1, 0x6a, 0x52, 0x16, 0x7f, 0x9e, 0x9e, 0xbe, 0x35, 0xa1, 0x15, 0xa2, 0x79, 0xc9,
	0x62, 0x14, 0x37, 0x00, 0x21, 0xe5, 0x78, 0x8b, 0x9c, 0xd2, 0x44, 0x74, 0xfd, 0xaf, 0xa5, 0xf9,
	0xf5, 0x03, 0xfd, 0xde, 0xf7, 0xb8, 0xa2, 0xbd, 0x9d, 0xd5, 0xb2, 0x09, 0xa3, 0x49, 0x9e, 0xa9,
	0xc1, 0xd8, 0xa2, 0x11, 0x12, 0xf6, 0x1c, 0x36, 0x52, 0xc4, 0x29, 0x9a, 0x7f, 0x98, 0x61, 0xa1,
	0xd1, 0x94, 0xb0, 0xb8, 0x83, 0x7d, 0x67, 0x5e, 0x65, 0xd3, 0x29, 0x1a, 0x54, 0xfc, 0x37, 0xb5,
	0xb3, 0x5a, 0x36, 0xdb, 0xae, 0xc6, 0xc1, 0x58, 0xcd, 0x14, 0xbd, 0x2a, 0x71, 0x0d, 0xbd, 0xcf,
	0x8f, 0xdf, 0x67, 0x9c, 0x52, 0xc1, 0xd2, 0x75, 0xfa, 0xb7, 0xf4, 0xda, 0xe5, 0xee, 0xe3, 0xf6,
	0xc7, 0xa0, 0x75, 0x74, 0xa1, 0xa3, 0x68, 0xa3, 0x1c, 0xd1, 0xd9, 0x7b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x50, 0xe6, 0x16, 0x3b, 0x00, 0x02, 0x00, 0x00,
}
