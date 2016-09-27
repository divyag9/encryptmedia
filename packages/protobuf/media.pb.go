// Code generated by protoc-gen-go.
// source: media.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	media.proto
	mediaEncrypted.proto

It has these top-level messages:
	Media
	MediaEncrypted
*/
package protobuf

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

type Media struct {
	Version            int64   `protobuf:"varint,1,opt,name=Version,json=version" json:"Version,omitempty"`
	GUID               string  `protobuf:"bytes,2,opt,name=GUID,json=gUID" json:"GUID,omitempty"`
	Client             string  `protobuf:"bytes,3,opt,name=Client,json=client" json:"Client,omitempty"`
	LoanType           string  `protobuf:"bytes,4,opt,name=LoanType,json=loanType" json:"LoanType,omitempty"`
	OrderNumber        string  `protobuf:"bytes,5,opt,name=OrderNumber,json=orderNumber" json:"OrderNumber,omitempty"`
	UserName           string  `protobuf:"bytes,6,opt,name=UserName,json=userName" json:"UserName,omitempty"`
	Latitude           float32 `protobuf:"fixed32,7,opt,name=Latitude,json=latitude" json:"Latitude,omitempty"`
	Longitude          float32 `protobuf:"fixed32,8,opt,name=Longitude,json=longitude" json:"Longitude,omitempty"`
	DateTaken          string  `protobuf:"bytes,9,opt,name=DateTaken,json=dateTaken" json:"DateTaken,omitempty"`
	DeviceModel        string  `protobuf:"bytes,10,opt,name=DeviceModel,json=deviceModel" json:"DeviceModel,omitempty"`
	DeviceOS           string  `protobuf:"bytes,11,opt,name=DeviceOS,json=deviceOS" json:"DeviceOS,omitempty"`
	DeviceOSVersion    string  `protobuf:"bytes,12,opt,name=DeviceOSVersion,json=deviceOSVersion" json:"DeviceOSVersion,omitempty"`
	FileName           string  `protobuf:"bytes,13,opt,name=FileName,json=fileName" json:"FileName,omitempty"`
	MimeType           string  `protobuf:"bytes,14,opt,name=MimeType,json=mimeType" json:"MimeType,omitempty"`
	Application        string  `protobuf:"bytes,15,opt,name=Application,json=application" json:"Application,omitempty"`
	ApplicationID      string  `protobuf:"bytes,16,opt,name=ApplicationID,json=applicationID" json:"ApplicationID,omitempty"`
	ApplicationVersion string  `protobuf:"bytes,17,opt,name=ApplicationVersion,json=applicationVersion" json:"ApplicationVersion,omitempty"`
	Bytes              []byte  `protobuf:"bytes,18,opt,name=Bytes,json=bytes,proto3" json:"Bytes,omitempty"`
}

func (m *Media) Reset()                    { *m = Media{} }
func (m *Media) String() string            { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()               {}
func (*Media) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Media)(nil), "protobuf.Media")
}

func init() { proto.RegisterFile("media.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0x41, 0x6b, 0xea, 0x40,
	0x14, 0x85, 0x51, 0x93, 0x98, 0x4c, 0xf4, 0xf9, 0xde, 0xe5, 0x51, 0x86, 0xd2, 0x85, 0x94, 0x2e,
	0x5c, 0xb9, 0xe9, 0x2f, 0x68, 0x2b, 0x2d, 0x05, 0xad, 0x60, 0xb5, 0xfb, 0x89, 0x73, 0x95, 0xa1,
	0x49, 0x26, 0xc4, 0x44, 0xf0, 0x47, 0xf5, 0x3f, 0x76, 0xe6, 0x4e, 0x82, 0x96, 0xae, 0x92, 0x73,
	0xbe, 0x33, 0x37, 0x67, 0x26, 0xc3, 0xe2, 0x0c, 0xa5, 0x12, 0xd3, 0xa2, 0xd4, 0x95, 0x86, 0x90,
	0x1e, 0x49, 0xbd, 0xbb, 0xfd, 0xf2, 0x98, 0xbf, 0xb0, 0x04, 0x38, 0xeb, 0x7f, 0x60, 0x79, 0x50,
	0x3a, 0xe7, 0x9d, 0x71, 0x67, 0xd2, 0x5b, 0xf5, 0x8f, 0x4e, 0x02, 0x30, 0xef, 0x65, 0xf3, 0x3a,
	0xe3, 0x5d, 0x63, 0x47, 0x2b, 0x6f, 0x6f, 0xde, 0xe1, 0x8a, 0x05, 0x4f, 0xa9, 0xc2, 0xbc, 0xe2,
	0x3d, 0x72, 0x83, 0x2d, 0x29, 0xb8, 0x66, 0xe1, 0x5c, 0x8b, 0x7c, 0x7d, 0x2a, 0x90, 0x7b, 0x44,
	0xc2, 0xb4, 0xd1, 0x30, 0x66, 0xf1, 0xb2, 0x94, 0x58, 0xbe, 0xd5, 0x59, 0x82, 0x25, 0xf7, 0x09,
	0xc7, 0xfa, 0x6c, 0xd9, 0xd5, 0x9b, 0x83, 0x51, 0x22, 0x43, 0x1e, 0xb8, 0xd5, 0x75, 0xa3, 0x69,
	0xb2, 0xa8, 0x54, 0x55, 0x4b, 0xe4, 0x7d, 0xc3, 0xba, 0x66, 0x72, 0xa3, 0xe1, 0x86, 0x45, 0x73,
	0x9d, 0xef, 0x1d, 0x0c, 0x09, 0x46, 0x69, 0x6b, 0x58, 0x3a, 0x13, 0x15, 0xae, 0xc5, 0x27, 0xe6,
	0x3c, 0xa2, 0xb1, 0x91, 0x6c, 0x0d, 0xdb, 0x6a, 0x86, 0x47, 0xb5, 0xc5, 0x85, 0x96, 0x98, 0x72,
	0xe6, 0x5a, 0xc9, 0xb3, 0x65, 0xbf, 0xec, 0x12, 0xcb, 0x77, 0x1e, 0xbb, 0x56, 0xb2, 0xd1, 0x30,
	0x61, 0xa3, 0x96, 0xb5, 0xa7, 0x37, 0xa0, 0xc8, 0x48, 0xfe, 0xb4, 0xed, 0x94, 0x67, 0x95, 0x22,
	0xed, 0x6d, 0xe8, 0xa6, 0xec, 0x1a, 0x6d, 0xd9, 0x42, 0x65, 0x48, 0xa7, 0xf6, 0xc7, 0xb1, 0xac,
	0xd1, 0xb6, 0xdf, 0x43, 0x51, 0xa4, 0x6a, 0x6b, 0x36, 0x6b, 0xa6, 0x8f, 0x5c, 0x3f, 0x71, 0xb6,
	0xe0, 0x8e, 0x0d, 0x2f, 0x12, 0xe6, 0x47, 0xfd, 0xa5, 0xcc, 0x50, 0x5c, 0x9a, 0x30, 0x65, 0x70,
	0x91, 0x6a, 0xcb, 0xfe, 0xa3, 0x28, 0x88, 0x5f, 0x04, 0xfe, 0x33, 0xff, 0xf1, 0x54, 0xe1, 0x81,
	0x83, 0x89, 0x0c, 0x56, 0x7e, 0x62, 0x45, 0x12, 0xd0, 0xcd, 0xb9, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x74, 0x7e, 0x95, 0x2a, 0x4f, 0x02, 0x00, 0x00,
}