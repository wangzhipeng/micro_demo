// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello_word.proto

package hello_word

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

type CallRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_word_d33489e5be2810ae, []int{0}
}
func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (dst *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(dst, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CallResponse struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallResponse) Reset()         { *m = CallResponse{} }
func (m *CallResponse) String() string { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()    {}
func (*CallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_word_d33489e5be2810ae, []int{1}
}
func (m *CallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResponse.Unmarshal(m, b)
}
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResponse.Marshal(b, m, deterministic)
}
func (dst *CallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResponse.Merge(dst, src)
}
func (m *CallResponse) XXX_Size() int {
	return xxx_messageInfo_CallResponse.Size(m)
}
func (m *CallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallResponse proto.InternalMessageInfo

func (m *CallResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CallRequest)(nil), "CallRequest")
	proto.RegisterType((*CallResponse)(nil), "CallResponse")
}

func init() { proto.RegisterFile("hello_word.proto", fileDescriptor_hello_word_d33489e5be2810ae) }

var fileDescriptor_hello_word_d33489e5be2810ae = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x8f, 0x2f, 0xcf, 0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4, 0xe2,
	0x76, 0x4e, 0xcc, 0xc9, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x34, 0xb8, 0x78,
	0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0x25, 0x98, 0xc0, 0xca, 0x60, 0x5c, 0x23, 0x03, 0x2e, 0x76, 0xd7, 0x8a, 0xc4,
	0xdc, 0x82, 0x9c, 0x54, 0x21, 0x55, 0x2e, 0x16, 0x90, 0x26, 0x21, 0x1e, 0x3d, 0x24, 0xe3, 0xa5,
	0x78, 0xf5, 0x90, 0x4d, 0x52, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x70, 0x85, 0x10, 0xb3, 0x99, 0x00, 0x00, 0x00,
}
