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

type LogRequest struct {
	Time    int64  `protobuf:"varint,1,opt,name=Time,json=time" json:"Time,omitempty"`
	NodeID  string `protobuf:"bytes,2,opt,name=NodeID,json=nodeID" json:"NodeID,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,json=message" json:"Message,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto1.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LogResponse struct {
}

func (m *LogResponse) Reset()                    { *m = LogResponse{} }
func (m *LogResponse) String() string            { return proto1.CompactTextString(m) }
func (*LogResponse) ProtoMessage()               {}
func (*LogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NodeUpdateRequest struct {
	Time   int64  `protobuf:"varint,1,opt,name=Time,json=time" json:"Time,omitempty"`
	NodeID string `protobuf:"bytes,2,opt,name=NodeID,json=nodeID" json:"NodeID,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=Status,json=status" json:"Status,omitempty"`
}

func (m *NodeUpdateRequest) Reset()                    { *m = NodeUpdateRequest{} }
func (m *NodeUpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*NodeUpdateRequest) ProtoMessage()               {}
func (*NodeUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type NodeUpdateResponse struct {
}

func (m *NodeUpdateResponse) Reset()                    { *m = NodeUpdateResponse{} }
func (m *NodeUpdateResponse) String() string            { return proto1.CompactTextString(m) }
func (*NodeUpdateResponse) ProtoMessage()               {}
func (*NodeUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type EdgeUpdateRequest struct {
	Time     int64  `protobuf:"varint,1,opt,name=Time,json=time" json:"Time,omitempty"`
	NodeID   string `protobuf:"bytes,2,opt,name=NodeID,json=nodeID" json:"NodeID,omitempty"`
	EdgeName string `protobuf:"bytes,3,opt,name=EdgeName,json=edgeName" json:"EdgeName,omitempty"`
	Status   string `protobuf:"bytes,4,opt,name=Status,json=status" json:"Status,omitempty"`
}

func (m *EdgeUpdateRequest) Reset()                    { *m = EdgeUpdateRequest{} }
func (m *EdgeUpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*EdgeUpdateRequest) ProtoMessage()               {}
func (*EdgeUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type EdgeUpdateResponse struct {
}

func (m *EdgeUpdateResponse) Reset()                    { *m = EdgeUpdateResponse{} }
func (m *EdgeUpdateResponse) String() string            { return proto1.CompactTextString(m) }
func (*EdgeUpdateResponse) ProtoMessage()               {}
func (*EdgeUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type EdgeMessageRequest struct {
	Time     int64  `protobuf:"varint,1,opt,name=Time,json=time" json:"Time,omitempty"`
	NodeID   string `protobuf:"bytes,2,opt,name=NodeID,json=nodeID" json:"NodeID,omitempty"`
	EdgeName string `protobuf:"bytes,3,opt,name=EdgeName,json=edgeName" json:"EdgeName,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=Message,json=message" json:"Message,omitempty"`
}

func (m *EdgeMessageRequest) Reset()                    { *m = EdgeMessageRequest{} }
func (m *EdgeMessageRequest) String() string            { return proto1.CompactTextString(m) }
func (*EdgeMessageRequest) ProtoMessage()               {}
func (*EdgeMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0x3f, 0x4f, 0xfb, 0x30,
	0x10, 0x55, 0x9b, 0xfc, 0xd2, 0xfe, 0xae, 0x65, 0xc8, 0xf1, 0x47, 0xc6, 0x13, 0xca, 0xc4, 0x80,
	0x3a, 0xc0, 0x27, 0x40, 0x2a, 0x03, 0x12, 0x74, 0x48, 0x8b, 0x98, 0x4d, 0x73, 0x8a, 0x3a, 0xa4,
	0x0e, 0xb1, 0xdb, 0xcf, 0xc0, 0xc7, 0x26, 0xb1, 0x4d, 0xe3, 0x90, 0x4e, 0xc0, 0x94, 0xdc, 0x9d,
	0xfd, 0xde, 0xf3, 0x7b, 0x07, 0x53, 0x5d, 0x89, 0x35, 0x55, 0xb3, 0xb2, 0x92, 0x5a, 0xe2, 0x3f,
	0xf3, 0x49, 0x52, 0x80, 0x27, 0x99, 0xa7, 0xf4, 0xbe, 0x23, 0xa5, 0x11, 0x21, 0x5c, 0x6d, 0x0a,
	0x62, 0x83, 0xab, 0xc1, 0x75, 0x90, 0x86, 0xba, 0xfe, 0xc7, 0x0b, 0x88, 0x16, 0x32, 0xa3, 0xc7,
	0x39, 0x1b, 0xd6, 0xdd, 0xff, 0x69, 0xb4, 0x35, 0x15, 0x32, 0x18, 0x3d, 0x93, 0x52, 0x22, 0x27,
	0x16, 0x98, 0xc1, 0xa8, 0xb0, 0x65, 0x72, 0x02, 0x13, 0x83, 0xa9, 0x4a, 0xb9, 0x55, 0x94, 0xbc,
	0x42, 0xdc, 0x00, 0xbc, 0x94, 0x99, 0xd0, 0xf4, 0x13, 0xa6, 0xba, 0xbf, 0xd4, 0x42, 0xef, 0x94,
	0x23, 0x8a, 0x94, 0xa9, 0x92, 0x33, 0x40, 0x1f, 0xd8, 0xd1, 0x29, 0x88, 0x1f, 0xb2, 0xfc, 0x17,
	0x74, 0x1c, 0xc6, 0x0d, 0xc0, 0x42, 0x14, 0x5f, 0x2f, 0x1b, 0x93, 0xab, 0x3d, 0x29, 0xe1, 0x77,
	0x29, 0x3e, 0xa9, 0x93, 0xb2, 0xb7, 0x5d, 0x67, 0xd3, 0x5f, 0x6b, 0xf1, 0x02, 0x08, 0xbb, 0x01,
	0x9c, 0xc3, 0x69, 0x87, 0xd7, 0xca, 0xb9, 0xfd, 0x18, 0xc2, 0x74, 0xd5, 0xec, 0xc0, 0x92, 0xaa,
	0xfd, 0x66, 0x4d, 0x78, 0x03, 0x41, 0x1d, 0x14, 0xc6, 0x76, 0x25, 0x66, 0xed, 0x22, 0x70, 0xf4,
	0x5b, 0xf6, 0x3a, 0xde, 0x03, 0xb4, 0x76, 0x23, 0x73, 0x27, 0x7a, 0xd1, 0xf2, 0xcb, 0x23, 0x93,
	0x16, 0xa2, 0xb5, 0xe9, 0x00, 0xd1, 0x8b, 0xeb, 0x00, 0xd1, 0xf7, 0x14, 0xe7, 0x30, 0xf1, 0xde,
	0x86, 0xfe, 0xc9, 0xae, 0xcf, 0x9c, 0x1f, 0x1b, 0x59, 0x94, 0xb7, 0xc8, 0x8c, 0xee, 0x3e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0x09, 0xc8, 0x29, 0x14, 0x03, 0x00, 0x00,
}
