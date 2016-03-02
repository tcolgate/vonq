// Code generated by protoc-gen-go.
// source: graphalg.proto
// DO NOT EDIT!

/*
Package graphalg is a generated protocol buffer package.

It is generated from these files:
	graphalg.proto

It has these top-level messages:
	Weight
	Message
	SendMessageRequest
	Error
	SendMessageResponse
	TestMessage
*/
package graphalg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import graph "github.com/tcolgate/vonq/graph"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Weight struct {
	Cost float64 `protobuf:"fixed64,1,opt,name=Cost" json:"Cost,omitempty"`
	Lsn  string  `protobuf:"bytes,2,opt,name=Lsn" json:"Lsn,omitempty"`
	Msn  string  `protobuf:"bytes,3,opt,name=Msn" json:"Msn,omitempty"`
}

func (m *Weight) Reset()                    { *m = Weight{} }
func (m *Weight) String() string            { return proto.CompactTextString(m) }
func (*Weight) ProtoMessage()               {}
func (*Weight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Payload *google_protobuf.Any `protobuf:"bytes,4,opt,name=Payload" json:"Payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetPayload() *google_protobuf.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SendMessageRequest struct {
	Gid *graph.GraphID     `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Aid *graph.AlgorithmID `protobuf:"bytes,2,opt,name=aid" json:"aid,omitempty"`
	Msg *Message           `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *SendMessageRequest) Reset()                    { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()               {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SendMessageRequest) GetGid() *graph.GraphID {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *SendMessageRequest) GetAid() *graph.AlgorithmID {
	if m != nil {
		return m.Aid
	}
	return nil
}

func (m *SendMessageRequest) GetMsg() *Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Error struct {
	Error int32  `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SendMessageResponse struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *SendMessageResponse) Reset()                    { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()               {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SendMessageResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type TestMessage struct {
	Value int32 `protobuf:"varint,1,opt,name=Value" json:"Value,omitempty"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Weight)(nil), "graphalg.Weight")
	proto.RegisterType((*Message)(nil), "graphalg.Message")
	proto.RegisterType((*SendMessageRequest)(nil), "graphalg.SendMessageRequest")
	proto.RegisterType((*Error)(nil), "graphalg.Error")
	proto.RegisterType((*SendMessageResponse)(nil), "graphalg.SendMessageResponse")
	proto.RegisterType((*TestMessage)(nil), "graphalg.TestMessage")
}

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x50, 0x5b, 0x4b, 0xc3, 0x30,
	0x14, 0xa6, 0xeb, 0x2e, 0x7a, 0xaa, 0x13, 0xa3, 0x0f, 0x63, 0xce, 0x0b, 0x15, 0x61, 0xf8, 0xd0,
	0x4a, 0xc5, 0x1f, 0x30, 0x54, 0x44, 0xd8, 0x40, 0x9c, 0xe8, 0x9b, 0x90, 0x6d, 0xc7, 0xac, 0xd0,
	0x35, 0x5d, 0x92, 0x0e, 0xfa, 0xef, 0x4d, 0xd2, 0xce, 0x4d, 0xc4, 0x87, 0x86, 0x7e, 0x7c, 0xe7,
	0xbb, 0x9c, 0x03, 0x6d, 0x26, 0x68, 0x36, 0xa7, 0x09, 0x0b, 0x32, 0xc1, 0x15, 0x27, 0x3b, 0x6b,
	0xdc, 0x0d, 0x59, 0xac, 0xe6, 0xf9, 0x24, 0x98, 0xf2, 0x45, 0xc8, 0x78, 0x42, 0x53, 0x16, 0xda,
	0x91, 0x49, 0xfe, 0x15, 0x66, 0xaa, 0xc8, 0x50, 0x86, 0x34, 0x2d, 0xcc, 0x57, 0x4a, 0xbb, 0xd7,
	0x5b, 0x02, 0x35, 0xe5, 0x09, 0xa3, 0x0a, 0xc3, 0x15, 0x4f, 0x97, 0xa1, 0xf5, 0x2c, 0xdf, 0x72,
	0xd6, 0xbf, 0x81, 0xe6, 0x07, 0xc6, 0x6c, 0xae, 0xc8, 0x1e, 0xd4, 0xef, 0xb9, 0x54, 0x1d, 0xe7,
	0xc2, 0xe9, 0x3b, 0xc4, 0x03, 0x77, 0x28, 0xd3, 0x4e, 0x4d, 0x83, 0x5d, 0x03, 0x46, 0x1a, 0xb8,
	0x06, 0x68, 0x45, 0x6b, 0x84, 0x52, 0x52, 0x86, 0xe4, 0x0a, 0x5a, 0x2f, 0xb4, 0x48, 0x38, 0x9d,
	0x75, 0xea, 0x9a, 0xf3, 0xa2, 0xe3, 0x80, 0x71, 0xce, 0x12, 0x0c, 0xd6, 0x05, 0x83, 0x41, 0x5a,
	0xf8, 0x02, 0xc8, 0x18, 0xd3, 0x59, 0xa5, 0x7a, 0xc5, 0x65, 0x8e, 0x52, 0x91, 0x13, 0x70, 0x59,
	0x3c, 0xb3, 0x71, 0x5e, 0xd4, 0x0e, 0xca, 0x52, 0x4f, 0xe6, 0x7d, 0x7e, 0x20, 0xe7, 0xe0, 0x52,
	0x4d, 0xd6, 0x2c, 0x49, 0x2a, 0x72, 0x90, 0x30, 0x2e, 0xf4, 0x6a, 0x0b, 0x3d, 0x70, 0x06, 0xee,
	0x42, 0x32, 0x5b, 0xc9, 0x8b, 0x0e, 0x83, 0x9f, 0xe3, 0x55, 0x21, 0xfe, 0x25, 0x34, 0x1e, 0x85,
	0xe0, 0x82, 0xec, 0x43, 0x03, 0xcd, 0x8f, 0x0d, 0x6a, 0x98, 0x55, 0x8c, 0xce, 0xee, 0xe5, 0xdf,
	0xc1, 0xd1, 0xaf, 0x62, 0x32, 0xe3, 0xa9, 0x44, 0xed, 0xbd, 0x25, 0xf1, 0xa2, 0x83, 0x8d, 0xbb,
	0xb5, 0xf4, 0x7b, 0xe0, 0xbd, 0xe9, 0x0d, 0xd6, 0x57, 0xd0, 0x09, 0xef, 0x34, 0xc9, 0xb1, 0x4c,
	0x88, 0x3e, 0xa1, 0x5d, 0x31, 0x63, 0x14, 0xab, 0x78, 0x8a, 0x64, 0x08, 0xde, 0x56, 0x0c, 0xe9,
	0x6d, 0xfc, 0xfe, 0x9e, 0xa5, 0x7b, 0xfa, 0x0f, 0x5b, 0x76, 0xeb, 0x3b, 0x93, 0xa6, 0xbd, 0xed,
	0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x07, 0xd7, 0x9a, 0x31, 0x02, 0x00, 0x00,
}
