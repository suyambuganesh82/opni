// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/metrics/apis/node/node.proto

package node

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NodeMetricsCapability_Sync_FullMethodName = "/node.metrics.NodeMetricsCapability/Sync"
)

// NodeMetricsCapabilityClient is the client API for NodeMetricsCapability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeMetricsCapabilityClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type nodeMetricsCapabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeMetricsCapabilityClient(cc grpc.ClientConnInterface) NodeMetricsCapabilityClient {
	return &nodeMetricsCapabilityClient{cc}
}

func (c *nodeMetricsCapabilityClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, NodeMetricsCapability_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeMetricsCapabilityServer is the server API for NodeMetricsCapability service.
// All implementations must embed UnimplementedNodeMetricsCapabilityServer
// for forward compatibility
type NodeMetricsCapabilityServer interface {
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	mustEmbedUnimplementedNodeMetricsCapabilityServer()
}

// UnimplementedNodeMetricsCapabilityServer must be embedded to have forward compatible implementations.
type UnimplementedNodeMetricsCapabilityServer struct {
}

func (UnimplementedNodeMetricsCapabilityServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedNodeMetricsCapabilityServer) mustEmbedUnimplementedNodeMetricsCapabilityServer() {}

// UnsafeNodeMetricsCapabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeMetricsCapabilityServer will
// result in compilation errors.
type UnsafeNodeMetricsCapabilityServer interface {
	mustEmbedUnimplementedNodeMetricsCapabilityServer()
}

func RegisterNodeMetricsCapabilityServer(s grpc.ServiceRegistrar, srv NodeMetricsCapabilityServer) {
	s.RegisterService(&NodeMetricsCapability_ServiceDesc, srv)
}

func _NodeMetricsCapability_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeMetricsCapabilityServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeMetricsCapability_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeMetricsCapabilityServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeMetricsCapability_ServiceDesc is the grpc.ServiceDesc for NodeMetricsCapability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeMetricsCapability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.metrics.NodeMetricsCapability",
	HandlerType: (*NodeMetricsCapabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _NodeMetricsCapability_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/apis/node/node.proto",
}

const (
	NodeConfiguration_GetDefaultConfiguration_FullMethodName = "/node.metrics.NodeConfiguration/GetDefaultConfiguration"
	NodeConfiguration_SetDefaultConfiguration_FullMethodName = "/node.metrics.NodeConfiguration/SetDefaultConfiguration"
	NodeConfiguration_GetNodeConfiguration_FullMethodName    = "/node.metrics.NodeConfiguration/GetNodeConfiguration"
	NodeConfiguration_SetNodeConfiguration_FullMethodName    = "/node.metrics.NodeConfiguration/SetNodeConfiguration"
)

// NodeConfigurationClient is the client API for NodeConfiguration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeConfigurationClient interface {
	GetDefaultConfiguration(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MetricsCapabilitySpec, error)
	SetDefaultConfiguration(ctx context.Context, in *MetricsCapabilitySpec, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetNodeConfiguration(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*MetricsCapabilitySpec, error)
	SetNodeConfiguration(ctx context.Context, in *NodeConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nodeConfigurationClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeConfigurationClient(cc grpc.ClientConnInterface) NodeConfigurationClient {
	return &nodeConfigurationClient{cc}
}

func (c *nodeConfigurationClient) GetDefaultConfiguration(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MetricsCapabilitySpec, error) {
	out := new(MetricsCapabilitySpec)
	err := c.cc.Invoke(ctx, NodeConfiguration_GetDefaultConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeConfigurationClient) SetDefaultConfiguration(ctx context.Context, in *MetricsCapabilitySpec, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NodeConfiguration_SetDefaultConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeConfigurationClient) GetNodeConfiguration(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*MetricsCapabilitySpec, error) {
	out := new(MetricsCapabilitySpec)
	err := c.cc.Invoke(ctx, NodeConfiguration_GetNodeConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeConfigurationClient) SetNodeConfiguration(ctx context.Context, in *NodeConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NodeConfiguration_SetNodeConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeConfigurationServer is the server API for NodeConfiguration service.
// All implementations must embed UnimplementedNodeConfigurationServer
// for forward compatibility
type NodeConfigurationServer interface {
	GetDefaultConfiguration(context.Context, *emptypb.Empty) (*MetricsCapabilitySpec, error)
	SetDefaultConfiguration(context.Context, *MetricsCapabilitySpec) (*emptypb.Empty, error)
	GetNodeConfiguration(context.Context, *v1.Reference) (*MetricsCapabilitySpec, error)
	SetNodeConfiguration(context.Context, *NodeConfigRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNodeConfigurationServer()
}

// UnimplementedNodeConfigurationServer must be embedded to have forward compatible implementations.
type UnimplementedNodeConfigurationServer struct {
}

func (UnimplementedNodeConfigurationServer) GetDefaultConfiguration(context.Context, *emptypb.Empty) (*MetricsCapabilitySpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultConfiguration not implemented")
}
func (UnimplementedNodeConfigurationServer) SetDefaultConfiguration(context.Context, *MetricsCapabilitySpec) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultConfiguration not implemented")
}
func (UnimplementedNodeConfigurationServer) GetNodeConfiguration(context.Context, *v1.Reference) (*MetricsCapabilitySpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeConfiguration not implemented")
}
func (UnimplementedNodeConfigurationServer) SetNodeConfiguration(context.Context, *NodeConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNodeConfiguration not implemented")
}
func (UnimplementedNodeConfigurationServer) mustEmbedUnimplementedNodeConfigurationServer() {}

// UnsafeNodeConfigurationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeConfigurationServer will
// result in compilation errors.
type UnsafeNodeConfigurationServer interface {
	mustEmbedUnimplementedNodeConfigurationServer()
}

func RegisterNodeConfigurationServer(s grpc.ServiceRegistrar, srv NodeConfigurationServer) {
	s.RegisterService(&NodeConfiguration_ServiceDesc, srv)
}

func _NodeConfiguration_GetDefaultConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeConfigurationServer).GetDefaultConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeConfiguration_GetDefaultConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeConfigurationServer).GetDefaultConfiguration(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeConfiguration_SetDefaultConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsCapabilitySpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeConfigurationServer).SetDefaultConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeConfiguration_SetDefaultConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeConfigurationServer).SetDefaultConfiguration(ctx, req.(*MetricsCapabilitySpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeConfiguration_GetNodeConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeConfigurationServer).GetNodeConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeConfiguration_GetNodeConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeConfigurationServer).GetNodeConfiguration(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeConfiguration_SetNodeConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeConfigurationServer).SetNodeConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeConfiguration_SetNodeConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeConfigurationServer).SetNodeConfiguration(ctx, req.(*NodeConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeConfiguration_ServiceDesc is the grpc.ServiceDesc for NodeConfiguration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeConfiguration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.metrics.NodeConfiguration",
	HandlerType: (*NodeConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDefaultConfiguration",
			Handler:    _NodeConfiguration_GetDefaultConfiguration_Handler,
		},
		{
			MethodName: "SetDefaultConfiguration",
			Handler:    _NodeConfiguration_SetDefaultConfiguration_Handler,
		},
		{
			MethodName: "GetNodeConfiguration",
			Handler:    _NodeConfiguration_GetNodeConfiguration_Handler,
		},
		{
			MethodName: "SetNodeConfiguration",
			Handler:    _NodeConfiguration_SetNodeConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/apis/node/node.proto",
}