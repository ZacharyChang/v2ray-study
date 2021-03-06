// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v2ray.com/v2ray-study/common/net/address.proto

package net

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

// Address of a network host. It may be either an IP address or a domain address.
type IPOrDomain struct {
	// Types that are valid to be assigned to Address:
	//	*IPOrDomain_Ip
	//	*IPOrDomain_Domain
	Address              isIPOrDomain_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IPOrDomain) Reset()         { *m = IPOrDomain{} }
func (m *IPOrDomain) String() string { return proto.CompactTextString(m) }
func (*IPOrDomain) ProtoMessage()    {}
func (*IPOrDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d6ad252b270fa4a, []int{0}
}

func (m *IPOrDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPOrDomain.Unmarshal(m, b)
}
func (m *IPOrDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPOrDomain.Marshal(b, m, deterministic)
}
func (m *IPOrDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPOrDomain.Merge(m, src)
}
func (m *IPOrDomain) XXX_Size() int {
	return xxx_messageInfo_IPOrDomain.Size(m)
}
func (m *IPOrDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_IPOrDomain.DiscardUnknown(m)
}

var xxx_messageInfo_IPOrDomain proto.InternalMessageInfo

type isIPOrDomain_Address interface {
	isIPOrDomain_Address()
}

type IPOrDomain_Ip struct {
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3,oneof"`
}

type IPOrDomain_Domain struct {
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3,oneof"`
}

func (*IPOrDomain_Ip) isIPOrDomain_Address() {}

func (*IPOrDomain_Domain) isIPOrDomain_Address() {}

func (m *IPOrDomain) GetAddress() isIPOrDomain_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IPOrDomain) GetIp() []byte {
	if x, ok := m.GetAddress().(*IPOrDomain_Ip); ok {
		return x.Ip
	}
	return nil
}

func (m *IPOrDomain) GetDomain() string {
	if x, ok := m.GetAddress().(*IPOrDomain_Domain); ok {
		return x.Domain
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IPOrDomain) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IPOrDomain_Ip)(nil),
		(*IPOrDomain_Domain)(nil),
	}
}

func init() {
	proto.RegisterType((*IPOrDomain)(nil), "v2ray.core.common.net.IPOrDomain")
}

func init() {
	proto.RegisterFile("v2ray.com/v2ray-study/common/net/address.proto", fileDescriptor_1d6ad252b270fa4a)
}

var fileDescriptor_1d6ad252b270fa4a = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xb3, 0x74, 0x8b, 0x4b, 0x4a, 0x53, 0x2a, 0xf5, 0x93,
	0xf3, 0x73, 0x73, 0xf3, 0xf3, 0xf4, 0xf3, 0x52, 0x4b, 0xf4, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b,
	0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x61, 0xea, 0x8b, 0x52, 0xf5, 0x20, 0x8a,
	0xf4, 0xf2, 0x52, 0x4b, 0x94, 0x9c, 0xb9, 0xb8, 0x3c, 0x03, 0xfc, 0x8b, 0x5c, 0xf2, 0x73, 0x13,
	0x33, 0xf3, 0x84, 0x04, 0xb8, 0x98, 0x32, 0x0b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x3c, 0x18,
	0x82, 0x98, 0x32, 0x0b, 0x84, 0x24, 0xb8, 0xd8, 0x52, 0xc0, 0x72, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x1e, 0x0c, 0x41, 0x50, 0xbe, 0x13, 0x27, 0x17, 0x3b, 0xd4, 0x06, 0x27, 0x2b, 0x2e, 0xc9,
	0xe4, 0xfc, 0x5c, 0x3d, 0xac, 0x36, 0x04, 0x30, 0x46, 0x31, 0xe7, 0xa5, 0x96, 0xac, 0x62, 0x12,
	0x0d, 0x33, 0x0a, 0x4a, 0xac, 0xd4, 0x73, 0x06, 0x49, 0x3b, 0x43, 0xa4, 0xfd, 0x52, 0x4b, 0x92,
	0xd8, 0xc0, 0xce, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x20, 0x54, 0xd9, 0xd3, 0xd0, 0x00,
	0x00, 0x00,
}
