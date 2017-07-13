// Code generated by protoc-gen-go. DO NOT EDIT.
// source: a.proto

/*
Package gameproto is a generated protocol buffer package.

It is generated from these files:
	a.proto
	base.proto
	test.proto

It has these top-level messages:
	Hello
	CliMsgHeader
	Person
	AddressBook
*/
package gameproto

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

type Hello struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Hello) Reset()                    { *m = Hello{} }
func (m *Hello) String() string            { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()               {}
func (*Hello) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hello) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "gameproto.Hello")
}

func init() { proto.RegisterFile("a.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 72 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x4f, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x4f, 0xcc, 0x4d, 0x05, 0x33, 0x95, 0xa4, 0xb9, 0x58, 0x3d,
	0x52, 0x73, 0x72, 0xf2, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0xec, 0x24, 0x36, 0xb0, 0x1a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x03, 0x9b, 0x9e, 0xce, 0x39, 0x00, 0x00, 0x00,
}
