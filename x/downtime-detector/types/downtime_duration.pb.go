// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlins/downtime-detector/v1beta1/downtime_duration.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Downtime int32

const (
	Downtime_DURATION_30S  Downtime = 0
	Downtime_DURATION_1M   Downtime = 1
	Downtime_DURATION_2M   Downtime = 2
	Downtime_DURATION_3M   Downtime = 3
	Downtime_DURATION_4M   Downtime = 4
	Downtime_DURATION_5M   Downtime = 5
	Downtime_DURATION_10M  Downtime = 6
	Downtime_DURATION_20M  Downtime = 7
	Downtime_DURATION_30M  Downtime = 8
	Downtime_DURATION_40M  Downtime = 9
	Downtime_DURATION_50M  Downtime = 10
	Downtime_DURATION_1H   Downtime = 11
	Downtime_DURATION_1_5H Downtime = 12
	Downtime_DURATION_2H   Downtime = 13
	Downtime_DURATION_2_5H Downtime = 14
	Downtime_DURATION_3H   Downtime = 15
	Downtime_DURATION_4H   Downtime = 16
	Downtime_DURATION_5H   Downtime = 17
	Downtime_DURATION_6H   Downtime = 18
	Downtime_DURATION_9H   Downtime = 19
	Downtime_DURATION_12H  Downtime = 20
	Downtime_DURATION_18H  Downtime = 21
	Downtime_DURATION_24H  Downtime = 22
	Downtime_DURATION_36H  Downtime = 23
	Downtime_DURATION_48H  Downtime = 24
)

var Downtime_name = map[int32]string{
	0:  "DURATION_30S",
	1:  "DURATION_1M",
	2:  "DURATION_2M",
	3:  "DURATION_3M",
	4:  "DURATION_4M",
	5:  "DURATION_5M",
	6:  "DURATION_10M",
	7:  "DURATION_20M",
	8:  "DURATION_30M",
	9:  "DURATION_40M",
	10: "DURATION_50M",
	11: "DURATION_1H",
	12: "DURATION_1_5H",
	13: "DURATION_2H",
	14: "DURATION_2_5H",
	15: "DURATION_3H",
	16: "DURATION_4H",
	17: "DURATION_5H",
	18: "DURATION_6H",
	19: "DURATION_9H",
	20: "DURATION_12H",
	21: "DURATION_18H",
	22: "DURATION_24H",
	23: "DURATION_36H",
	24: "DURATION_48H",
}

var Downtime_value = map[string]int32{
	"DURATION_30S":  0,
	"DURATION_1M":   1,
	"DURATION_2M":   2,
	"DURATION_3M":   3,
	"DURATION_4M":   4,
	"DURATION_5M":   5,
	"DURATION_10M":  6,
	"DURATION_20M":  7,
	"DURATION_30M":  8,
	"DURATION_40M":  9,
	"DURATION_50M":  10,
	"DURATION_1H":   11,
	"DURATION_1_5H": 12,
	"DURATION_2H":   13,
	"DURATION_2_5H": 14,
	"DURATION_3H":   15,
	"DURATION_4H":   16,
	"DURATION_5H":   17,
	"DURATION_6H":   18,
	"DURATION_9H":   19,
	"DURATION_12H":  20,
	"DURATION_18H":  21,
	"DURATION_24H":  22,
	"DURATION_36H":  23,
	"DURATION_48H":  24,
}

func (x Downtime) String() string {
	return proto.EnumName(Downtime_name, int32(x))
}

func (Downtime) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21a1969f22fb2a7e, []int{0}
}

func init() {
	proto.RegisterEnum("merlins.downtimedetector.v1beta1.Downtime", Downtime_name, Downtime_value)
}

func init() {
	proto.RegisterFile("merlins/downtime-detector/v1beta1/downtime_duration.proto", fileDescriptor_21a1969f22fb2a7e)
}

var fileDescriptor_21a1969f22fb2a7e = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3d, 0x6f, 0xe2, 0x40,
	0x10, 0x86, 0xed, 0xe3, 0x8e, 0xe3, 0x0c, 0x1c, 0x83, 0x8f, 0xfb, 0xa2, 0xf0, 0x5d, 0x1d, 0x09,
	0xaf, 0x6d, 0x30, 0x82, 0x22, 0x45, 0x22, 0x8a, 0x4d, 0xb1, 0x89, 0x94, 0x0f, 0x45, 0x4a, 0x63,
	0xd9, 0xe0, 0x38, 0x96, 0x30, 0x8b, 0xf0, 0x42, 0xc2, 0xbf, 0xc8, 0x6f, 0x4a, 0x95, 0x92, 0x32,
	0x65, 0x04, 0x7f, 0x24, 0xc2, 0x1f, 0x44, 0x83, 0xd2, 0x79, 0x9e, 0x99, 0xd7, 0xfb, 0xbe, 0xa3,
	0x51, 0xfa, 0x3c, 0x8e, 0x78, 0x1c, 0xc6, 0x64, 0xc4, 0xef, 0x27, 0x22, 0x8c, 0xfc, 0xd6, 0xc8,
	0x17, 0xfe, 0x50, 0xf0, 0x19, 0x59, 0x98, 0x9e, 0x2f, 0x5c, 0x73, 0xd7, 0x71, 0x46, 0xf3, 0x99,
	0x2b, 0x42, 0x3e, 0xd1, 0xa7, 0x33, 0x2e, 0xb8, 0xfa, 0x3f, 0x93, 0xea, 0xf9, 0x40, 0xae, 0xd4,
	0x33, 0x65, 0xb3, 0x11, 0xf0, 0x80, 0x27, 0xc3, 0x64, 0xfb, 0x95, 0xea, 0x9a, 0x7f, 0x03, 0xce,
	0x83, 0xb1, 0x4f, 0x92, 0xca, 0x9b, 0xdf, 0x12, 0x77, 0xb2, 0xcc, 0x5b, 0xc3, 0xe4, 0x9f, 0x4e,
	0xaa, 0x49, 0x8b, 0xac, 0xa5, 0xed, 0xab, 0xb0, 0x9b, 0xe6, 0xbf, 0xfd, 0xfe, 0xd6, 0x51, 0x2c,
	0xdc, 0x68, 0x9a, 0x0e, 0x1c, 0x3c, 0x15, 0x94, 0xd2, 0x20, 0x73, 0xaa, 0x82, 0x52, 0x19, 0x5c,
	0x9d, 0x1f, 0x5d, 0x9e, 0x9c, 0x9d, 0x3a, 0x6d, 0xe3, 0x02, 0x24, 0xb5, 0xa6, 0x94, 0x77, 0xc4,
	0x64, 0x20, 0x23, 0x60, 0x31, 0xf8, 0x84, 0x40, 0x9b, 0x41, 0x01, 0x81, 0x0e, 0x83, 0xcf, 0x08,
	0xd8, 0x0c, 0xbe, 0xa0, 0x67, 0x4c, 0x83, 0x41, 0x11, 0x11, 0xcb, 0x60, 0xf0, 0x75, 0xcf, 0x0a,
	0x83, 0x12, 0x22, 0x1d, 0x83, 0xc1, 0x37, 0x44, 0x6c, 0x83, 0x81, 0x82, 0xed, 0x52, 0x28, 0xab,
	0x75, 0xa5, 0xfa, 0x0e, 0x1c, 0x9b, 0x42, 0x05, 0x27, 0xa0, 0x50, 0x45, 0x33, 0xd6, 0x76, 0xe6,
	0x3b, 0x0e, 0x45, 0xa1, 0x86, 0x43, 0x51, 0x00, 0x1c, 0x8a, 0x42, 0x1d, 0x81, 0x2e, 0x05, 0x15,
	0x81, 0x3e, 0x85, 0x1f, 0x38, 0xb6, 0x45, 0xa1, 0x81, 0x49, 0x8f, 0xc2, 0x4f, 0xbc, 0x88, 0x0e,
	0x85, 0x5f, 0x78, 0x11, 0x5d, 0x0a, 0xbf, 0xf1, 0x22, 0x7a, 0x14, 0xfe, 0x1c, 0x5f, 0x3f, 0xaf,
	0x35, 0x79, 0xb5, 0xd6, 0xe4, 0xd7, 0xb5, 0x26, 0x3f, 0x6e, 0x34, 0x69, 0xb5, 0xd1, 0xa4, 0x97,
	0x8d, 0x26, 0xdd, 0x1c, 0x06, 0xa1, 0xb8, 0x9b, 0x7b, 0xfa, 0x90, 0x47, 0x24, 0x3b, 0xcc, 0xd6,
	0xd8, 0xf5, 0xe2, 0xbc, 0x20, 0x0b, 0xd3, 0x26, 0x0f, 0x1f, 0x9c, 0xb9, 0x58, 0x4e, 0xfd, 0xd8,
	0x2b, 0x26, 0x47, 0xd2, 0x7e, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x51, 0xea, 0x63, 0x10, 0x03,
	0x00, 0x00,
}
