// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

/*
Package cache is a generated protocol buffer package.

It is generated from these files:
	cache.proto

It has these top-level messages:
	RPCIDs
	RPCReply
	RPCBools
	RPCStorages
*/
package cache

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RPCIDs struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *RPCIDs) Reset()                    { *m = RPCIDs{} }
func (m *RPCIDs) String() string            { return proto.CompactTextString(m) }
func (*RPCIDs) ProtoMessage()               {}
func (*RPCIDs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RPCIDs) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type RPCReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *RPCReply) Reset()                    { *m = RPCReply{} }
func (m *RPCReply) String() string            { return proto.CompactTextString(m) }
func (*RPCReply) ProtoMessage()               {}
func (*RPCReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RPCReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RPCBools struct {
	Exists []bool `protobuf:"varint,1,rep,packed,name=exists" json:"exists,omitempty"`
}

func (m *RPCBools) Reset()                    { *m = RPCBools{} }
func (m *RPCBools) String() string            { return proto.CompactTextString(m) }
func (*RPCBools) ProtoMessage()               {}
func (*RPCBools) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RPCBools) GetExists() []bool {
	if m != nil {
		return m.Exists
	}
	return nil
}

type RPCStorages struct {
	Version uint32   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Hosts   []string `protobuf:"bytes,2,rep,name=hosts" json:"hosts,omitempty"`
}

func (m *RPCStorages) Reset()                    { *m = RPCStorages{} }
func (m *RPCStorages) String() string            { return proto.CompactTextString(m) }
func (*RPCStorages) ProtoMessage()               {}
func (*RPCStorages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RPCStorages) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RPCStorages) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func init() {
	proto.RegisterType((*RPCIDs)(nil), "RPCIDs")
	proto.RegisterType((*RPCReply)(nil), "RPCReply")
	proto.RegisterType((*RPCBools)(nil), "RPCBools")
	proto.RegisterType((*RPCStorages)(nil), "RPCStorages")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cache service

type CacheClient interface {
	UpdateStorage(ctx context.Context, in *RPCStorages, opts ...grpc.CallOption) (*RPCReply, error)
	Put(ctx context.Context, in *RPCIDs, opts ...grpc.CallOption) (*RPCReply, error)
	Get(ctx context.Context, in *RPCIDs, opts ...grpc.CallOption) (*RPCBools, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) UpdateStorage(ctx context.Context, in *RPCStorages, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := grpc.Invoke(ctx, "/Cache/UpdateStorage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Put(ctx context.Context, in *RPCIDs, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := grpc.Invoke(ctx, "/Cache/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Get(ctx context.Context, in *RPCIDs, opts ...grpc.CallOption) (*RPCBools, error) {
	out := new(RPCBools)
	err := grpc.Invoke(ctx, "/Cache/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheServer interface {
	UpdateStorage(context.Context, *RPCStorages) (*RPCReply, error)
	Put(context.Context, *RPCIDs) (*RPCReply, error)
	Get(context.Context, *RPCIDs) (*RPCBools, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_UpdateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCStorages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).UpdateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cache/UpdateStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).UpdateStorage(ctx, req.(*RPCStorages))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cache/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Put(ctx, req.(*RPCIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*RPCIDs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateStorage",
			Handler:    _Cache_UpdateStorage_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Cache_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache.proto",
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe2, 0x62, 0x0b, 0x0a, 0x70, 0xf6, 0x74,
	0x29, 0x16, 0x12, 0xe0, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x02,
	0x31, 0x95, 0x54, 0xb8, 0x38, 0x82, 0x02, 0x9c, 0x83, 0x52, 0x0b, 0x72, 0x2a, 0x85, 0x24, 0xb8,
	0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60,
	0x5c, 0x25, 0x25, 0xb0, 0x2a, 0xa7, 0xfc, 0xfc, 0x9c, 0x62, 0x21, 0x31, 0x2e, 0xb6, 0xd4, 0x8a,
	0xcc, 0xe2, 0x12, 0x88, 0x31, 0x1c, 0x41, 0x50, 0x9e, 0x92, 0x2d, 0x17, 0x77, 0x50, 0x80, 0x73,
	0x70, 0x49, 0x7e, 0x51, 0x62, 0x7a, 0x6a, 0x31, 0xc8, 0xb0, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc,
	0x3c, 0xb0, 0x61, 0xbc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x46, 0x3e, 0x48, 0x3f, 0x13,
	0xd8, 0x19, 0x10, 0x8e, 0x51, 0x2e, 0x17, 0xab, 0x33, 0xc8, 0xcd, 0x42, 0x5a, 0x5c, 0xbc, 0xa1,
	0x05, 0x29, 0x89, 0x25, 0xa9, 0x50, 0xa3, 0x84, 0x78, 0xf4, 0x90, 0xcc, 0x95, 0xe2, 0xd4, 0x83,
	0xb9, 0x57, 0x89, 0x41, 0x48, 0x9a, 0x8b, 0x39, 0xa0, 0xb4, 0x44, 0x88, 0x5d, 0x0f, 0xe2, 0x3f,
	0x0c, 0x49, 0xf7, 0x54, 0x74, 0x49, 0xb0, 0x1f, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x41, 0x63, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x57, 0x8b, 0xbb, 0x29, 0x01, 0x00, 0x00,
}
