// Code generated by protoc-gen-go.
// source: tracer.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	tracer.proto

It has these top-level messages:
	LogRequest
	LogResponse
	NodeUpdateRequest
	NodeUpdateResponse
	EdgeUpdateRequest
	EdgeUpdateResponse
	EdgeMessageRequest
	EdgeMessageResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import graph "github.com/tcolgate/radia/graph"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type EdgeMessageRequest_Dir int32

const (
	EdgeMessageRequest_IN    EdgeMessageRequest_Dir = 0
	EdgeMessageRequest_OUT   EdgeMessageRequest_Dir = 1
	EdgeMessageRequest_QUEUE EdgeMessageRequest_Dir = 2
)

var EdgeMessageRequest_Dir_name = map[int32]string{
	0: "IN",
	1: "OUT",
	2: "QUEUE",
}
var EdgeMessageRequest_Dir_value = map[string]int32{
	"IN":    0,
	"OUT":   1,
	"QUEUE": 2,
}

func (x EdgeMessageRequest_Dir) String() string {
	return proto1.EnumName(EdgeMessageRequest_Dir_name, int32(x))
}
func (EdgeMessageRequest_Dir) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type LogRequest struct {
	Time    int64              `protobuf:"varint,1,opt,name=Time" json:"Time,omitempty"`
	Gid     *graph.GraphID     `protobuf:"bytes,2,opt,name=Gid" json:"Gid,omitempty"`
	Aid     *graph.AlgorithmID `protobuf:"bytes,3,opt,name=Aid" json:"Aid,omitempty"`
	NodeID  string             `protobuf:"bytes,4,opt,name=NodeID" json:"NodeID,omitempty"`
	Message string             `protobuf:"bytes,5,opt,name=Message" json:"Message,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto1.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogRequest) GetGid() *graph.GraphID {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *LogRequest) GetAid() *graph.AlgorithmID {
	if m != nil {
		return m.Aid
	}
	return nil
}

type LogResponse struct {
}

func (m *LogResponse) Reset()                    { *m = LogResponse{} }
func (m *LogResponse) String() string            { return proto1.CompactTextString(m) }
func (*LogResponse) ProtoMessage()               {}
func (*LogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NodeUpdateRequest struct {
	Time   int64              `protobuf:"varint,1,opt,name=Time" json:"Time,omitempty"`
	Gid    *graph.GraphID     `protobuf:"bytes,2,opt,name=Gid" json:"Gid,omitempty"`
	Aid    *graph.AlgorithmID `protobuf:"bytes,3,opt,name=Aid" json:"Aid,omitempty"`
	NodeID string             `protobuf:"bytes,4,opt,name=NodeID" json:"NodeID,omitempty"`
	Status string             `protobuf:"bytes,5,opt,name=Status" json:"Status,omitempty"`
}

func (m *NodeUpdateRequest) Reset()                    { *m = NodeUpdateRequest{} }
func (m *NodeUpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*NodeUpdateRequest) ProtoMessage()               {}
func (*NodeUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NodeUpdateRequest) GetGid() *graph.GraphID {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *NodeUpdateRequest) GetAid() *graph.AlgorithmID {
	if m != nil {
		return m.Aid
	}
	return nil
}

type NodeUpdateResponse struct {
}

func (m *NodeUpdateResponse) Reset()                    { *m = NodeUpdateResponse{} }
func (m *NodeUpdateResponse) String() string            { return proto1.CompactTextString(m) }
func (*NodeUpdateResponse) ProtoMessage()               {}
func (*NodeUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type EdgeUpdateRequest struct {
	Time     int64              `protobuf:"varint,1,opt,name=Time" json:"Time,omitempty"`
	Gid      *graph.GraphID     `protobuf:"bytes,2,opt,name=Gid" json:"Gid,omitempty"`
	Aid      *graph.AlgorithmID `protobuf:"bytes,3,opt,name=Aid" json:"Aid,omitempty"`
	NodeID   string             `protobuf:"bytes,4,opt,name=NodeID" json:"NodeID,omitempty"`
	EdgeName string             `protobuf:"bytes,5,opt,name=EdgeName" json:"EdgeName,omitempty"`
	Status   string             `protobuf:"bytes,6,opt,name=Status" json:"Status,omitempty"`
}

func (m *EdgeUpdateRequest) Reset()                    { *m = EdgeUpdateRequest{} }
func (m *EdgeUpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*EdgeUpdateRequest) ProtoMessage()               {}
func (*EdgeUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EdgeUpdateRequest) GetGid() *graph.GraphID {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *EdgeUpdateRequest) GetAid() *graph.AlgorithmID {
	if m != nil {
		return m.Aid
	}
	return nil
}

type EdgeUpdateResponse struct {
}

func (m *EdgeUpdateResponse) Reset()                    { *m = EdgeUpdateResponse{} }
func (m *EdgeUpdateResponse) String() string            { return proto1.CompactTextString(m) }
func (*EdgeUpdateResponse) ProtoMessage()               {}
func (*EdgeUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type EdgeMessageRequest struct {
	Time      int64                  `protobuf:"varint,1,opt,name=Time" json:"Time,omitempty"`
	Gid       *graph.GraphID         `protobuf:"bytes,2,opt,name=Gid" json:"Gid,omitempty"`
	Aid       *graph.AlgorithmID     `protobuf:"bytes,3,opt,name=Aid" json:"Aid,omitempty"`
	NodeID    string                 `protobuf:"bytes,4,opt,name=NodeID" json:"NodeID,omitempty"`
	EdgeName  string                 `protobuf:"bytes,5,opt,name=EdgeName" json:"EdgeName,omitempty"`
	Message   string                 `protobuf:"bytes,6,opt,name=Message" json:"Message,omitempty"`
	Direction EdgeMessageRequest_Dir `protobuf:"varint,7,opt,name=Direction,enum=proto.EdgeMessageRequest_Dir" json:"Direction,omitempty"`
}

func (m *EdgeMessageRequest) Reset()                    { *m = EdgeMessageRequest{} }
func (m *EdgeMessageRequest) String() string            { return proto1.CompactTextString(m) }
func (*EdgeMessageRequest) ProtoMessage()               {}
func (*EdgeMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EdgeMessageRequest) GetGid() *graph.GraphID {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *EdgeMessageRequest) GetAid() *graph.AlgorithmID {
	if m != nil {
		return m.Aid
	}
	return nil
}

type EdgeMessageResponse struct {
}

func (m *EdgeMessageResponse) Reset()                    { *m = EdgeMessageResponse{} }
func (m *EdgeMessageResponse) String() string            { return proto1.CompactTextString(m) }
func (*EdgeMessageResponse) ProtoMessage()               {}
func (*EdgeMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto1.RegisterType((*LogRequest)(nil), "proto.LogRequest")
	proto1.RegisterType((*LogResponse)(nil), "proto.LogResponse")
	proto1.RegisterType((*NodeUpdateRequest)(nil), "proto.NodeUpdateRequest")
	proto1.RegisterType((*NodeUpdateResponse)(nil), "proto.NodeUpdateResponse")
	proto1.RegisterType((*EdgeUpdateRequest)(nil), "proto.EdgeUpdateRequest")
	proto1.RegisterType((*EdgeUpdateResponse)(nil), "proto.EdgeUpdateResponse")
	proto1.RegisterType((*EdgeMessageRequest)(nil), "proto.EdgeMessageRequest")
	proto1.RegisterType((*EdgeMessageResponse)(nil), "proto.EdgeMessageResponse")
	proto1.RegisterEnum("proto.EdgeMessageRequest_Dir", EdgeMessageRequest_Dir_name, EdgeMessageRequest_Dir_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for TraceService service

type TraceServiceClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	NodeUpdate(ctx context.Context, in *NodeUpdateRequest, opts ...grpc.CallOption) (*NodeUpdateResponse, error)
	EdgeUpdate(ctx context.Context, in *EdgeUpdateRequest, opts ...grpc.CallOption) (*EdgeUpdateResponse, error)
	EdgeMessage(ctx context.Context, in *EdgeMessageRequest, opts ...grpc.CallOption) (*EdgeMessageResponse, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := grpc.Invoke(ctx, "/proto.TraceService/Log", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) NodeUpdate(ctx context.Context, in *NodeUpdateRequest, opts ...grpc.CallOption) (*NodeUpdateResponse, error) {
	out := new(NodeUpdateResponse)
	err := grpc.Invoke(ctx, "/proto.TraceService/NodeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) EdgeUpdate(ctx context.Context, in *EdgeUpdateRequest, opts ...grpc.CallOption) (*EdgeUpdateResponse, error) {
	out := new(EdgeUpdateResponse)
	err := grpc.Invoke(ctx, "/proto.TraceService/EdgeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) EdgeMessage(ctx context.Context, in *EdgeMessageRequest, opts ...grpc.CallOption) (*EdgeMessageResponse, error) {
	out := new(EdgeMessageResponse)
	err := grpc.Invoke(ctx, "/proto.TraceService/EdgeMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TraceService service

type TraceServiceServer interface {
	Log(context.Context, *LogRequest) (*LogResponse, error)
	NodeUpdate(context.Context, *NodeUpdateRequest) (*NodeUpdateResponse, error)
	EdgeUpdate(context.Context, *EdgeUpdateRequest) (*EdgeUpdateResponse, error)
	EdgeMessage(context.Context, *EdgeMessageRequest) (*EdgeMessageResponse, error)
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TraceServiceServer).Log(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _TraceService_NodeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NodeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TraceServiceServer).NodeUpdate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _TraceService_EdgeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EdgeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TraceServiceServer).EdgeUpdate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _TraceService_EdgeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EdgeMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TraceServiceServer).EdgeMessage(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _TraceService_Log_Handler,
		},
		{
			MethodName: "NodeUpdate",
			Handler:    _TraceService_NodeUpdate_Handler,
		},
		{
			MethodName: "EdgeUpdate",
			Handler:    _TraceService_EdgeUpdate_Handler,
		},
		{
			MethodName: "EdgeMessage",
			Handler:    _TraceService_EdgeMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x93, 0xdd, 0x8e, 0x9a, 0x40,
	0x14, 0xc7, 0x0b, 0x28, 0xd6, 0xa3, 0xb5, 0x78, 0xda, 0x26, 0x48, 0xd3, 0xb4, 0xe5, 0xaa, 0x69,
	0x1a, 0x6c, 0xec, 0x13, 0x98, 0x60, 0x8c, 0x49, 0x6b, 0xd3, 0x2a, 0x0f, 0x80, 0x30, 0x41, 0x12,
	0x71, 0x10, 0x46, 0x93, 0x5e, 0xf5, 0x76, 0x6f, 0xf7, 0x2d, 0xf7, 0x31, 0x76, 0x60, 0x66, 0x85,
	0x0d, 0xee, 0xad, 0x7b, 0x33, 0x93, 0x39, 0x9f, 0x3f, 0xce, 0xf9, 0x03, 0x7d, 0x96, 0xf9, 0x01,
	0xc9, 0x9c, 0x34, 0xa3, 0x8c, 0x62, 0xbb, 0xbc, 0xac, 0xaf, 0x51, 0xcc, 0xb6, 0xc7, 0x8d, 0x13,
	0xd0, 0x64, 0xcc, 0x02, 0xba, 0x8b, 0x7c, 0x46, 0xc6, 0x27, 0xba, 0x3f, 0x8c, 0xa3, 0xcc, 0x4f,
	0xb7, 0xe2, 0x14, 0x29, 0xf6, 0x3f, 0x80, 0x9f, 0x34, 0xfa, 0x4b, 0x0e, 0x47, 0x92, 0x33, 0xec,
	0x43, 0x6b, 0x1d, 0x27, 0xc4, 0x54, 0x3e, 0x29, 0x5f, 0x34, 0x7c, 0x0f, 0xda, 0x3c, 0x0e, 0x4d,
	0x95, 0x3f, 0x7a, 0x93, 0x81, 0x23, 0xd2, 0xe6, 0xc5, 0xb9, 0x70, 0xf1, 0x23, 0x68, 0x53, 0xee,
	0xd4, 0x4a, 0x27, 0x4a, 0xe7, 0x74, 0x17, 0xd1, 0x8c, 0x37, 0x4f, 0x78, 0xc0, 0x00, 0xf4, 0x25,
	0x0d, 0xc9, 0xc2, 0x35, 0x5b, 0x3c, 0xa6, 0x8b, 0xaf, 0xa1, 0xf3, 0x8b, 0xe4, 0xb9, 0x1f, 0x11,
	0xb3, 0x5d, 0x18, 0xec, 0x57, 0xd0, 0x2b, 0x5b, 0xe7, 0x29, 0xdd, 0xe7, 0xc4, 0xfe, 0x0f, 0xc3,
	0x22, 0xde, 0x4b, 0x43, 0x8e, 0x7b, 0x0d, 0x20, 0xfe, 0x5e, 0x31, 0x9f, 0x1d, 0x73, 0xc9, 0xf3,
	0x16, 0xb0, 0x0e, 0x20, 0xb1, 0x6e, 0x15, 0x18, 0xce, 0xc2, 0xe8, 0x8a, 0x5c, 0x06, 0xbc, 0x2c,
	0x1a, 0x2e, 0xfd, 0x44, 0x4e, 0xaa, 0x46, 0xaa, 0x3f, 0x90, 0xd6, 0x91, 0x24, 0xe9, 0x9d, 0x22,
	0xcc, 0x72, 0xca, 0xcf, 0x83, 0x5a, 0xdb, 0x72, 0xc9, 0x8a, 0xdf, 0xa1, 0xeb, 0xc6, 0x19, 0x09,
	0x58, 0x4c, 0xf7, 0x66, 0x87, 0x9b, 0x06, 0x93, 0x0f, 0x42, 0x7b, 0x4e, 0x13, 0xd6, 0xe1, 0xa1,
	0xf6, 0x67, 0xd0, 0xf8, 0x85, 0x3a, 0xa8, 0x8b, 0xa5, 0xf1, 0x02, 0x3b, 0xa0, 0xfd, 0xf6, 0xd6,
	0x86, 0x82, 0x5d, 0x68, 0xff, 0xf1, 0x66, 0xde, 0xcc, 0x50, 0xed, 0x77, 0xf0, 0xe6, 0x51, 0xb2,
	0x98, 0xc0, 0xe4, 0x46, 0x85, 0xfe, 0xba, 0xf8, 0x21, 0x56, 0x24, 0x3b, 0xc5, 0x01, 0xc1, 0x6f,
	0xa0, 0x71, 0x89, 0xe1, 0x50, 0x36, 0xac, 0x94, 0x6e, 0x61, 0xdd, 0x24, 0xd2, 0x71, 0x0a, 0x50,
	0x09, 0x00, 0x4d, 0x19, 0xd1, 0x10, 0xa5, 0x35, 0xba, 0xe0, 0xa9, 0x4a, 0x54, 0x9b, 0x39, 0x97,
	0x68, 0xe8, 0xe7, 0x5c, 0xa2, 0xb9, 0x46, 0x74, 0xa1, 0x57, 0xfb, 0x36, 0x1c, 0x3d, 0x39, 0x2c,
	0xcb, 0xba, 0xe4, 0x12, 0x55, 0x36, 0x7a, 0xe9, 0xfa, 0x71, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x49,
	0x21, 0xb4, 0xa8, 0x21, 0x04, 0x00, 0x00,
}
