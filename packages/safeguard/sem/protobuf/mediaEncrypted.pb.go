// Code generated by protoc-gen-go.
// source: mediaEncrypted.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MediaEncrypted struct {
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
	EncryptedBytes     []byte  `protobuf:"bytes,18,opt,name=EncryptedBytes,json=encryptedBytes,proto3" json:"EncryptedBytes,omitempty"`
	SymmetricKey       []byte  `protobuf:"bytes,19,opt,name=SymmetricKey,json=symmetricKey,proto3" json:"SymmetricKey,omitempty"`
}

func (m *MediaEncrypted) Reset()                    { *m = MediaEncrypted{} }
func (m *MediaEncrypted) String() string            { return proto.CompactTextString(m) }
func (*MediaEncrypted) ProtoMessage()               {}
func (*MediaEncrypted) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*MediaEncrypted)(nil), "protobuf.MediaEncrypted")
}

func init() { proto.RegisterFile("mediaEncrypted.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x6e, 0xe2, 0x30,
	0x14, 0x85, 0x05, 0x64, 0x42, 0x62, 0x42, 0x98, 0xb9, 0x33, 0x1a, 0x59, 0xa3, 0x59, 0x20, 0x54,
	0x55, 0xac, 0xd8, 0xf4, 0x09, 0xda, 0xd2, 0x56, 0x55, 0xa1, 0x48, 0xfc, 0x74, 0x6f, 0xe2, 0x0b,
	0xb2, 0x9a, 0x3f, 0x25, 0x0e, 0x52, 0xde, 0xb9, 0x0f, 0x51, 0xc7, 0x4e, 0x4a, 0x50, 0x57, 0xc9,
	0xf9, 0xce, 0xf1, 0xcd, 0xb1, 0x63, 0xf2, 0x27, 0x42, 0x2e, 0xd8, 0x43, 0x1c, 0x64, 0x65, 0x2a,
	0x91, 0xcf, 0xd2, 0x2c, 0x91, 0x09, 0x38, 0xfa, 0xb1, 0x2f, 0x0e, 0x93, 0x0f, 0x8b, 0xf8, 0xcb,
	0x8b, 0x08, 0x50, 0xd2, 0x7f, 0xc3, 0x2c, 0x17, 0x49, 0x4c, 0x3b, 0xe3, 0xce, 0xb4, 0xb7, 0xee,
	0x9f, 0x8c, 0x04, 0x20, 0xd6, 0xd3, 0xee, 0x79, 0x4e, 0xbb, 0x0a, 0xbb, 0x6b, 0xeb, 0xa8, 0xde,
	0xe1, 0x2f, 0xb1, 0xef, 0x43, 0x81, 0xb1, 0xa4, 0x3d, 0x4d, 0xed, 0x40, 0x2b, 0xf8, 0x47, 0x9c,
	0x45, 0xc2, 0xe2, 0x6d, 0x99, 0x22, 0xb5, 0xb4, 0xe3, 0x84, 0xb5, 0x86, 0x31, 0x19, 0xac, 0x32,
	0x8e, 0xd9, 0x6b, 0x11, 0xed, 0x31, 0xa3, 0x3f, 0xb4, 0x3d, 0x48, 0xce, 0xa8, 0x5a, 0xbd, 0xcb,
	0x95, 0x62, 0x11, 0x52, 0xdb, 0xac, 0x2e, 0x6a, 0xad, 0x27, 0x33, 0x29, 0x64, 0xc1, 0x91, 0xf6,
	0x95, 0xd7, 0x55, 0x93, 0x6b, 0x0d, 0xff, 0x89, 0xbb, 0x48, 0xe2, 0xa3, 0x31, 0x1d, 0x6d, 0xba,
	0x61, 0x03, 0x2a, 0x77, 0xce, 0x24, 0x6e, 0xd9, 0x3b, 0xc6, 0xd4, 0xd5, 0x63, 0x5d, 0xde, 0x80,
	0xaa, 0xd5, 0x1c, 0x4f, 0x22, 0xc0, 0x65, 0xc2, 0x31, 0xa4, 0xc4, 0xb4, 0xe2, 0x67, 0x54, 0x7d,
	0xd9, 0x24, 0x56, 0x1b, 0x3a, 0x30, 0xad, 0x78, 0xad, 0x61, 0x4a, 0x46, 0x8d, 0xd7, 0x9c, 0x9e,
	0xa7, 0x23, 0x23, 0x7e, 0x89, 0xab, 0x29, 0x8f, 0x22, 0x44, 0xbd, 0xb7, 0xa1, 0x99, 0x72, 0xa8,
	0x75, 0xe5, 0x2d, 0x45, 0x84, 0xfa, 0xd4, 0x7c, 0xe3, 0x45, 0xb5, 0xae, 0xfa, 0xdd, 0xa6, 0x69,
	0x28, 0x02, 0xb5, 0x59, 0x35, 0x7d, 0x64, 0xfa, 0xb1, 0x33, 0x82, 0x2b, 0x32, 0x6c, 0x25, 0xd4,
	0x8f, 0xfa, 0xa9, 0x33, 0x43, 0xd6, 0x86, 0x30, 0x23, 0xd0, 0x4a, 0x35, 0x65, 0x7f, 0xe9, 0x28,
	0xb0, 0x6f, 0x0e, 0x5c, 0x13, 0xff, 0xeb, 0x72, 0xdc, 0x95, 0x12, 0x73, 0x0a, 0x2a, 0xeb, 0xad,
	0x7d, 0xbc, 0xa0, 0x30, 0x21, 0xde, 0xa6, 0x8c, 0x22, 0x94, 0x99, 0x08, 0x5e, 0xb0, 0xa4, 0xbf,
	0x75, 0xca, 0xcb, 0x5b, 0x6c, 0x6f, 0xeb, 0x8b, 0x77, 0xf3, 0x19, 0x00, 0x00, 0xff, 0xff, 0x75,
	0x67, 0x28, 0xfc, 0x97, 0x02, 0x00, 0x00,
}
