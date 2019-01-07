// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v2ray.com/v2ray-study/common/log/log.proto

package log

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

type Severity int32

const (
	Severity_Unknown Severity = 0
	Severity_Error   Severity = 1
	Severity_Warning Severity = 2
	Severity_Info    Severity = 3
	Severity_Debug   Severity = 4
)

var Severity_name = map[int32]string{
	0: "Unknown",
	1: "Error",
	2: "Warning",
	3: "Info",
	4: "Debug",
}

var Severity_value = map[string]int32{
	"Unknown": 0,
	"Error":   1,
	"Warning": 2,
	"Info":    3,
	"Debug":   4,
}

func (x Severity) String() string {
	return proto.EnumName(Severity_name, int32(x))
}

func (Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3530f35d355cc9ca, []int{0}
}

func init() {
	proto.RegisterEnum("v2ray.core.common.log.Severity", Severity_name, Severity_value)
}

func init() {
	proto.RegisterFile("v2ray.com/v2ray-study/common/log/log.proto", fileDescriptor_3530f35d355cc9ca)
}

var fileDescriptor_3530f35d355cc9ca = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xb3, 0x74, 0x8b, 0x4b, 0x4a, 0x53, 0x2a, 0xf5, 0x93,
	0xf3, 0x73, 0x73, 0xf3, 0xf3, 0xf4, 0x73, 0xf2, 0xd3, 0x41, 0x58, 0xaf, 0xa0, 0x28, 0xbf, 0x24,
	0x5f, 0x48, 0x14, 0xa6, 0xb6, 0x28, 0x55, 0x0f, 0xa2, 0x40, 0x2f, 0x27, 0x3f, 0x5d, 0xcb, 0x85,
	0x8b, 0x23, 0x38, 0xb5, 0x2c, 0xb5, 0x28, 0xb3, 0xa4, 0x52, 0x88, 0x9b, 0x8b, 0x3d, 0x34, 0x2f,
	0x3b, 0x2f, 0xbf, 0x3c, 0x4f, 0x80, 0x41, 0x88, 0x93, 0x8b, 0xd5, 0xb5, 0xa8, 0x28, 0xbf, 0x48,
	0x80, 0x11, 0x24, 0x1e, 0x9e, 0x58, 0x94, 0x97, 0x99, 0x97, 0x2e, 0xc0, 0x24, 0xc4, 0xc1, 0xc5,
	0xe2, 0x99, 0x97, 0x96, 0x2f, 0xc0, 0x0c, 0x52, 0xe1, 0x92, 0x9a, 0x54, 0x9a, 0x2e, 0xc0, 0xe2,
	0x64, 0xc5, 0x25, 0x99, 0x9c, 0x9f, 0xab, 0x87, 0xd5, 0x8a, 0x00, 0xc6, 0x28, 0xe6, 0x9c, 0xfc,
	0xf4, 0x55, 0x4c, 0xa2, 0x61, 0x46, 0x41, 0x89, 0x95, 0x7a, 0xce, 0x20, 0x69, 0x67, 0x88, 0xb4,
	0x4f, 0x7e, 0x7a, 0x12, 0x1b, 0xd8, 0x7d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x30,
	0x78, 0x25, 0xcd, 0x00, 0x00, 0x00,
}
