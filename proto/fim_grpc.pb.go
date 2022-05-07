// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/fim.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FimClient is the client API for Fim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FimClient interface {
	//This section applies to the agent
	GetStartupInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StartupInfo, error)
	CreateBaseline(ctx context.Context, opts ...grpc.CallOption) (Fim_CreateBaselineClient, error)
	UpdateBaseline(ctx context.Context, opts ...grpc.CallOption) (Fim_UpdateBaselineClient, error)
	ReportFsStatus(ctx context.Context, opts ...grpc.CallOption) (Fim_ReportFsStatusClient, error)
	ReportFsEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error)
	// This section applies to the client
	GetAgents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Fim_GetAgentsClient, error)
	GetAlertsByAgent(ctx context.Context, in *EndpointName, opts ...grpc.CallOption) (Fim_GetAlertsByAgentClient, error)
	CreateBaselineUpdateApproval(ctx context.Context, in *EndpointName, opts ...grpc.CallOption) (*Empty, error)
	CreateAgentEndpoint(ctx context.Context, in *AgentEndpoint, opts ...grpc.CallOption) (*Empty, error)
	CreateClientEndpoint(ctx context.Context, in *ClientEndpoint, opts ...grpc.CallOption) (*Empty, error)
	DeleteEndpoint(ctx context.Context, in *EndpointName, opts ...grpc.CallOption) (*Empty, error)
	UpdateEndpointWatchedPaths(ctx context.Context, in *AgentEndpoint, opts ...grpc.CallOption) (*Empty, error)
}

type fimClient struct {
	cc grpc.ClientConnInterface
}

func NewFimClient(cc grpc.ClientConnInterface) FimClient {
	return &fimClient{cc}
}

func (c *fimClient) GetStartupInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StartupInfo, error) {
	out := new(StartupInfo)
	err := c.cc.Invoke(ctx, "/fim.Fim/GetStartupInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimClient) CreateBaseline(ctx context.Context, opts ...grpc.CallOption) (Fim_CreateBaselineClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fim_ServiceDesc.Streams[0], "/fim.Fim/CreateBaseline", opts...)
	if err != nil {
		return nil, err
	}
	x := &fimCreateBaselineClient{stream}
	return x, nil
}

type Fim_CreateBaselineClient interface {
	Send(*FsObject) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type fimCreateBaselineClient struct {
	grpc.ClientStream
}

func (x *fimCreateBaselineClient) Send(m *FsObject) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fimCreateBaselineClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fimClient) UpdateBaseline(ctx context.Context, opts ...grpc.CallOption) (Fim_UpdateBaselineClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fim_ServiceDesc.Streams[1], "/fim.Fim/UpdateBaseline", opts...)
	if err != nil {
		return nil, err
	}
	x := &fimUpdateBaselineClient{stream}
	return x, nil
}

type Fim_UpdateBaselineClient interface {
	Send(*FsObject) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type fimUpdateBaselineClient struct {
	grpc.ClientStream
}

func (x *fimUpdateBaselineClient) Send(m *FsObject) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fimUpdateBaselineClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fimClient) ReportFsStatus(ctx context.Context, opts ...grpc.CallOption) (Fim_ReportFsStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fim_ServiceDesc.Streams[2], "/fim.Fim/ReportFsStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &fimReportFsStatusClient{stream}
	return x, nil
}

type Fim_ReportFsStatusClient interface {
	Send(*FsObject) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type fimReportFsStatusClient struct {
	grpc.ClientStream
}

func (x *fimReportFsStatusClient) Send(m *FsObject) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fimReportFsStatusClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fimClient) ReportFsEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fim/ReportFsEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimClient) GetAgents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Fim_GetAgentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fim_ServiceDesc.Streams[3], "/fim.Fim/GetAgents", opts...)
	if err != nil {
		return nil, err
	}
	x := &fimGetAgentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fim_GetAgentsClient interface {
	Recv() (*Agent, error)
	grpc.ClientStream
}

type fimGetAgentsClient struct {
	grpc.ClientStream
}

func (x *fimGetAgentsClient) Recv() (*Agent, error) {
	m := new(Agent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fimClient) GetAlertsByAgent(ctx context.Context, in *EndpointName, opts ...grpc.CallOption) (Fim_GetAlertsByAgentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fim_ServiceDesc.Streams[4], "/fim.Fim/GetAlertsByAgent", opts...)
	if err != nil {
		return nil, err
	}
	x := &fimGetAlertsByAgentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fim_GetAlertsByAgentClient interface {
	Recv() (*Alert, error)
	grpc.ClientStream
}

type fimGetAlertsByAgentClient struct {
	grpc.ClientStream
}

func (x *fimGetAlertsByAgentClient) Recv() (*Alert, error) {
	m := new(Alert)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fimClient) CreateBaselineUpdateApproval(ctx context.Context, in *EndpointName, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fim/CreateBaselineUpdateApproval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimClient) CreateAgentEndpoint(ctx context.Context, in *AgentEndpoint, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fim/CreateAgentEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimClient) CreateClientEndpoint(ctx context.Context, in *ClientEndpoint, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fim/CreateClientEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimClient) DeleteEndpoint(ctx context.Context, in *EndpointName, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fim/DeleteEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimClient) UpdateEndpointWatchedPaths(ctx context.Context, in *AgentEndpoint, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fim/UpdateEndpointWatchedPaths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FimServer is the server API for Fim service.
// All implementations must embed UnimplementedFimServer
// for forward compatibility
type FimServer interface {
	//This section applies to the agent
	GetStartupInfo(context.Context, *Empty) (*StartupInfo, error)
	CreateBaseline(Fim_CreateBaselineServer) error
	UpdateBaseline(Fim_UpdateBaselineServer) error
	ReportFsStatus(Fim_ReportFsStatusServer) error
	ReportFsEvent(context.Context, *Event) (*Empty, error)
	// This section applies to the client
	GetAgents(*Empty, Fim_GetAgentsServer) error
	GetAlertsByAgent(*EndpointName, Fim_GetAlertsByAgentServer) error
	CreateBaselineUpdateApproval(context.Context, *EndpointName) (*Empty, error)
	CreateAgentEndpoint(context.Context, *AgentEndpoint) (*Empty, error)
	CreateClientEndpoint(context.Context, *ClientEndpoint) (*Empty, error)
	DeleteEndpoint(context.Context, *EndpointName) (*Empty, error)
	UpdateEndpointWatchedPaths(context.Context, *AgentEndpoint) (*Empty, error)
	mustEmbedUnimplementedFimServer()
}

// UnimplementedFimServer must be embedded to have forward compatible implementations.
type UnimplementedFimServer struct {
}

func (UnimplementedFimServer) GetStartupInfo(context.Context, *Empty) (*StartupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStartupInfo not implemented")
}
func (UnimplementedFimServer) CreateBaseline(Fim_CreateBaselineServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateBaseline not implemented")
}
func (UnimplementedFimServer) UpdateBaseline(Fim_UpdateBaselineServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateBaseline not implemented")
}
func (UnimplementedFimServer) ReportFsStatus(Fim_ReportFsStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportFsStatus not implemented")
}
func (UnimplementedFimServer) ReportFsEvent(context.Context, *Event) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportFsEvent not implemented")
}
func (UnimplementedFimServer) GetAgents(*Empty, Fim_GetAgentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAgents not implemented")
}
func (UnimplementedFimServer) GetAlertsByAgent(*EndpointName, Fim_GetAlertsByAgentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAlertsByAgent not implemented")
}
func (UnimplementedFimServer) CreateBaselineUpdateApproval(context.Context, *EndpointName) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBaselineUpdateApproval not implemented")
}
func (UnimplementedFimServer) CreateAgentEndpoint(context.Context, *AgentEndpoint) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentEndpoint not implemented")
}
func (UnimplementedFimServer) CreateClientEndpoint(context.Context, *ClientEndpoint) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientEndpoint not implemented")
}
func (UnimplementedFimServer) DeleteEndpoint(context.Context, *EndpointName) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEndpoint not implemented")
}
func (UnimplementedFimServer) UpdateEndpointWatchedPaths(context.Context, *AgentEndpoint) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEndpointWatchedPaths not implemented")
}
func (UnimplementedFimServer) mustEmbedUnimplementedFimServer() {}

// UnsafeFimServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FimServer will
// result in compilation errors.
type UnsafeFimServer interface {
	mustEmbedUnimplementedFimServer()
}

func RegisterFimServer(s grpc.ServiceRegistrar, srv FimServer) {
	s.RegisterService(&Fim_ServiceDesc, srv)
}

func _Fim_GetStartupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimServer).GetStartupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fim/GetStartupInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimServer).GetStartupInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fim_CreateBaseline_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FimServer).CreateBaseline(&fimCreateBaselineServer{stream})
}

type Fim_CreateBaselineServer interface {
	SendAndClose(*Empty) error
	Recv() (*FsObject, error)
	grpc.ServerStream
}

type fimCreateBaselineServer struct {
	grpc.ServerStream
}

func (x *fimCreateBaselineServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fimCreateBaselineServer) Recv() (*FsObject, error) {
	m := new(FsObject)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Fim_UpdateBaseline_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FimServer).UpdateBaseline(&fimUpdateBaselineServer{stream})
}

type Fim_UpdateBaselineServer interface {
	SendAndClose(*Empty) error
	Recv() (*FsObject, error)
	grpc.ServerStream
}

type fimUpdateBaselineServer struct {
	grpc.ServerStream
}

func (x *fimUpdateBaselineServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fimUpdateBaselineServer) Recv() (*FsObject, error) {
	m := new(FsObject)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Fim_ReportFsStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FimServer).ReportFsStatus(&fimReportFsStatusServer{stream})
}

type Fim_ReportFsStatusServer interface {
	SendAndClose(*Empty) error
	Recv() (*FsObject, error)
	grpc.ServerStream
}

type fimReportFsStatusServer struct {
	grpc.ServerStream
}

func (x *fimReportFsStatusServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fimReportFsStatusServer) Recv() (*FsObject, error) {
	m := new(FsObject)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Fim_ReportFsEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimServer).ReportFsEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fim/ReportFsEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimServer).ReportFsEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fim_GetAgents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FimServer).GetAgents(m, &fimGetAgentsServer{stream})
}

type Fim_GetAgentsServer interface {
	Send(*Agent) error
	grpc.ServerStream
}

type fimGetAgentsServer struct {
	grpc.ServerStream
}

func (x *fimGetAgentsServer) Send(m *Agent) error {
	return x.ServerStream.SendMsg(m)
}

func _Fim_GetAlertsByAgent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EndpointName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FimServer).GetAlertsByAgent(m, &fimGetAlertsByAgentServer{stream})
}

type Fim_GetAlertsByAgentServer interface {
	Send(*Alert) error
	grpc.ServerStream
}

type fimGetAlertsByAgentServer struct {
	grpc.ServerStream
}

func (x *fimGetAlertsByAgentServer) Send(m *Alert) error {
	return x.ServerStream.SendMsg(m)
}

func _Fim_CreateBaselineUpdateApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimServer).CreateBaselineUpdateApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fim/CreateBaselineUpdateApproval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimServer).CreateBaselineUpdateApproval(ctx, req.(*EndpointName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fim_CreateAgentEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimServer).CreateAgentEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fim/CreateAgentEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimServer).CreateAgentEndpoint(ctx, req.(*AgentEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fim_CreateClientEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimServer).CreateClientEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fim/CreateClientEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimServer).CreateClientEndpoint(ctx, req.(*ClientEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fim_DeleteEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimServer).DeleteEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fim/DeleteEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimServer).DeleteEndpoint(ctx, req.(*EndpointName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fim_UpdateEndpointWatchedPaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimServer).UpdateEndpointWatchedPaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fim/UpdateEndpointWatchedPaths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimServer).UpdateEndpointWatchedPaths(ctx, req.(*AgentEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

// Fim_ServiceDesc is the grpc.ServiceDesc for Fim service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fim_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fim.Fim",
	HandlerType: (*FimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStartupInfo",
			Handler:    _Fim_GetStartupInfo_Handler,
		},
		{
			MethodName: "ReportFsEvent",
			Handler:    _Fim_ReportFsEvent_Handler,
		},
		{
			MethodName: "CreateBaselineUpdateApproval",
			Handler:    _Fim_CreateBaselineUpdateApproval_Handler,
		},
		{
			MethodName: "CreateAgentEndpoint",
			Handler:    _Fim_CreateAgentEndpoint_Handler,
		},
		{
			MethodName: "CreateClientEndpoint",
			Handler:    _Fim_CreateClientEndpoint_Handler,
		},
		{
			MethodName: "DeleteEndpoint",
			Handler:    _Fim_DeleteEndpoint_Handler,
		},
		{
			MethodName: "UpdateEndpointWatchedPaths",
			Handler:    _Fim_UpdateEndpointWatchedPaths_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateBaseline",
			Handler:       _Fim_CreateBaseline_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateBaseline",
			Handler:       _Fim_UpdateBaseline_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReportFsStatus",
			Handler:       _Fim_ReportFsStatus_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetAgents",
			Handler:       _Fim_GetAgents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAlertsByAgent",
			Handler:       _Fim_GetAlertsByAgent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/fim.proto",
}
