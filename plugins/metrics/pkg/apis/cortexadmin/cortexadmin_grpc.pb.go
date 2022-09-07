// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/metrics/pkg/apis/cortexadmin/cortexadmin.proto

package cortexadmin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CortexAdminClient is the client API for CortexAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CortexAdminClient interface {
	AllUserStats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserIDStatsList, error)
	WriteMetrics(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	QueryRange(ctx context.Context, in *QueryRangeRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	GetRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	ListRules(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*QueryResponse, error)
	LoadRules(ctx context.Context, in *PostRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FlushBlocks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// list all metrics
	GetSeriesMetrics(ctx context.Context, in *SeriesRequest, opts ...grpc.CallOption) (*SeriesInfoList, error)
	GetMetricLabelSets(ctx context.Context, in *LabelRequest, opts ...grpc.CallOption) (*MetricLabels, error)
	GetCortexStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CortexStatus, error)
	GetCortexConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	ExtractRawSeries(ctx context.Context, in *MatcherRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type cortexAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewCortexAdminClient(cc grpc.ClientConnInterface) CortexAdminClient {
	return &cortexAdminClient{cc}
}

func (c *cortexAdminClient) AllUserStats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserIDStatsList, error) {
	out := new(UserIDStatsList)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/AllUserStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) WriteMetrics(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/WriteMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) QueryRange(ctx context.Context, in *QueryRangeRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/QueryRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) GetRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) ListRules(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/ListRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) LoadRules(ctx context.Context, in *PostRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/LoadRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) DeleteRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) FlushBlocks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/FlushBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) GetSeriesMetrics(ctx context.Context, in *SeriesRequest, opts ...grpc.CallOption) (*SeriesInfoList, error) {
	out := new(SeriesInfoList)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/GetSeriesMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) GetMetricLabelSets(ctx context.Context, in *LabelRequest, opts ...grpc.CallOption) (*MetricLabels, error) {
	out := new(MetricLabels)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/GetMetricLabelSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) GetCortexStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CortexStatus, error) {
	out := new(CortexStatus)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/GetCortexStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexAdminClient) GetCortexConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/cortexadmin.CortexAdmin/GetCortexConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CortexAdminServer is the server API for CortexAdmin service.
// All implementations must embed UnimplementedCortexAdminServer
// for forward compatibility
type CortexAdminServer interface {
	AllUserStats(context.Context, *emptypb.Empty) (*UserIDStatsList, error)
	WriteMetrics(context.Context, *WriteRequest) (*WriteResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	QueryRange(context.Context, *QueryRangeRequest) (*QueryResponse, error)
	GetRule(context.Context, *RuleRequest) (*QueryResponse, error)
	ListRules(context.Context, *Cluster) (*QueryResponse, error)
	LoadRules(context.Context, *PostRuleRequest) (*emptypb.Empty, error)
	DeleteRule(context.Context, *RuleRequest) (*emptypb.Empty, error)
	FlushBlocks(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// list all metrics
	GetSeriesMetrics(context.Context, *SeriesRequest) (*SeriesInfoList, error)
	GetMetricLabelSets(context.Context, *LabelRequest) (*MetricLabels, error)
	GetCortexStatus(context.Context, *emptypb.Empty) (*CortexStatus, error)
	GetCortexConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
	mustEmbedUnimplementedCortexAdminServer()
}

// UnimplementedCortexAdminServer must be embedded to have forward compatible implementations.
type UnimplementedCortexAdminServer struct {
}

func (UnimplementedCortexAdminServer) AllUserStats(context.Context, *emptypb.Empty) (*UserIDStatsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllUserStats not implemented")
}
func (UnimplementedCortexAdminServer) WriteMetrics(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMetrics not implemented")
}
func (UnimplementedCortexAdminServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedCortexAdminServer) QueryRange(context.Context, *QueryRangeRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRange not implemented")
}
func (UnimplementedCortexAdminServer) GetRule(context.Context, *RuleRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (UnimplementedCortexAdminServer) ListRules(context.Context, *Cluster) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRules not implemented")
}
func (UnimplementedCortexAdminServer) LoadRules(context.Context, *PostRuleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadRules not implemented")
}
func (UnimplementedCortexAdminServer) DeleteRule(context.Context, *RuleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (UnimplementedCortexAdminServer) FlushBlocks(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushBlocks not implemented")
}
func (UnimplementedCortexAdminServer) GetSeriesMetrics(context.Context, *SeriesRequest) (*SeriesInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeriesMetrics not implemented")
}
func (UnimplementedCortexAdminServer) GetMetricLabelSets(context.Context, *LabelRequest) (*MetricLabels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricLabelSets not implemented")
}
func (UnimplementedCortexAdminServer) GetCortexStatus(context.Context, *emptypb.Empty) (*CortexStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCortexStatus not implemented")
}
func (UnimplementedCortexAdminServer) GetCortexConfig(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCortexConfig not implemented")
}
func (UnimplementedCortexAdminServer) mustEmbedUnimplementedCortexAdminServer() {}

// UnsafeCortexAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CortexAdminServer will
// result in compilation errors.
type UnsafeCortexAdminServer interface {
	mustEmbedUnimplementedCortexAdminServer()
}

func RegisterCortexAdminServer(s grpc.ServiceRegistrar, srv CortexAdminServer) {
	s.RegisterService(&CortexAdmin_ServiceDesc, srv)
}

func _CortexAdmin_AllUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).AllUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/AllUserStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).AllUserStats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_WriteMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).WriteMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/WriteMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).WriteMetrics(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_QueryRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).QueryRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/QueryRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).QueryRange(ctx, req.(*QueryRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).GetRule(ctx, req.(*RuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/ListRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).ListRules(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_LoadRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).LoadRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/LoadRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).LoadRules(ctx, req.(*PostRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).DeleteRule(ctx, req.(*RuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_FlushBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).FlushBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/FlushBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).FlushBlocks(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_GetSeriesMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).GetSeriesMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/GetSeriesMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).GetSeriesMetrics(ctx, req.(*SeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_GetMetricLabelSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).GetMetricLabelSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/GetMetricLabelSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).GetMetricLabelSets(ctx, req.(*LabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_GetCortexStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).GetCortexStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/GetCortexStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).GetCortexStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexAdmin_GetCortexConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexAdminServer).GetCortexConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexadmin.CortexAdmin/GetCortexConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexAdminServer).GetCortexConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CortexAdmin_ServiceDesc is the grpc.ServiceDesc for CortexAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CortexAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cortexadmin.CortexAdmin",
	HandlerType: (*CortexAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllUserStats",
			Handler:    _CortexAdmin_AllUserStats_Handler,
		},
		{
			MethodName: "WriteMetrics",
			Handler:    _CortexAdmin_WriteMetrics_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _CortexAdmin_Query_Handler,
		},
		{
			MethodName: "QueryRange",
			Handler:    _CortexAdmin_QueryRange_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _CortexAdmin_GetRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _CortexAdmin_ListRules_Handler,
		},
		{
			MethodName: "LoadRules",
			Handler:    _CortexAdmin_LoadRules_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _CortexAdmin_DeleteRule_Handler,
		},
		{
			MethodName: "FlushBlocks",
			Handler:    _CortexAdmin_FlushBlocks_Handler,
		},
		{
			MethodName: "GetSeriesMetrics",
			Handler:    _CortexAdmin_GetSeriesMetrics_Handler,
		},
		{
			MethodName: "GetMetricLabelSets",
			Handler:    _CortexAdmin_GetMetricLabelSets_Handler,
		},
		{
			MethodName: "GetCortexStatus",
			Handler:    _CortexAdmin_GetCortexStatus_Handler,
		},
		{
			MethodName: "GetCortexConfig",
			Handler:    _CortexAdmin_GetCortexConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/pkg/apis/cortexadmin/cortexadmin.proto",
}
