// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: election.proto

package voting

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

const (
	ElectionService_CreateElection_FullMethodName  = "/voting.ElectionService/CreateElection"
	ElectionService_GetElectionInfo_FullMethodName = "/voting.ElectionService/GetElectionInfo"
	ElectionService_UpdateElection_FullMethodName  = "/voting.ElectionService/UpdateElection"
	ElectionService_DeleteElection_FullMethodName  = "/voting.ElectionService/DeleteElection"
)

// ElectionServiceClient is the client API for ElectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectionServiceClient interface {
	CreateElection(ctx context.Context, in *CreateElectionRequest, opts ...grpc.CallOption) (*ElectionResponse, error)
	GetElectionInfo(ctx context.Context, in *GetElectionInfoRequest, opts ...grpc.CallOption) (*ElectionResponse, error)
	UpdateElection(ctx context.Context, in *UpdateElectionRequest, opts ...grpc.CallOption) (*ElectionResponse, error)
	DeleteElection(ctx context.Context, in *DeleteElectionRequest, opts ...grpc.CallOption) (*Void2, error)
}

type electionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionServiceClient(cc grpc.ClientConnInterface) ElectionServiceClient {
	return &electionServiceClient{cc}
}

func (c *electionServiceClient) CreateElection(ctx context.Context, in *CreateElectionRequest, opts ...grpc.CallOption) (*ElectionResponse, error) {
	out := new(ElectionResponse)
	err := c.cc.Invoke(ctx, ElectionService_CreateElection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) GetElectionInfo(ctx context.Context, in *GetElectionInfoRequest, opts ...grpc.CallOption) (*ElectionResponse, error) {
	out := new(ElectionResponse)
	err := c.cc.Invoke(ctx, ElectionService_GetElectionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) UpdateElection(ctx context.Context, in *UpdateElectionRequest, opts ...grpc.CallOption) (*ElectionResponse, error) {
	out := new(ElectionResponse)
	err := c.cc.Invoke(ctx, ElectionService_UpdateElection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) DeleteElection(ctx context.Context, in *DeleteElectionRequest, opts ...grpc.CallOption) (*Void2, error) {
	out := new(Void2)
	err := c.cc.Invoke(ctx, ElectionService_DeleteElection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectionServiceServer is the server API for ElectionService service.
// All implementations must embed UnimplementedElectionServiceServer
// for forward compatibility
type ElectionServiceServer interface {
	CreateElection(context.Context, *CreateElectionRequest) (*ElectionResponse, error)
	GetElectionInfo(context.Context, *GetElectionInfoRequest) (*ElectionResponse, error)
	UpdateElection(context.Context, *UpdateElectionRequest) (*ElectionResponse, error)
	DeleteElection(context.Context, *DeleteElectionRequest) (*Void2, error)
	mustEmbedUnimplementedElectionServiceServer()
}

// UnimplementedElectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedElectionServiceServer struct {
}

func (UnimplementedElectionServiceServer) CreateElection(context.Context, *CreateElectionRequest) (*ElectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateElection not implemented")
}
func (UnimplementedElectionServiceServer) GetElectionInfo(context.Context, *GetElectionInfoRequest) (*ElectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetElectionInfo not implemented")
}
func (UnimplementedElectionServiceServer) UpdateElection(context.Context, *UpdateElectionRequest) (*ElectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateElection not implemented")
}
func (UnimplementedElectionServiceServer) DeleteElection(context.Context, *DeleteElectionRequest) (*Void2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteElection not implemented")
}
func (UnimplementedElectionServiceServer) mustEmbedUnimplementedElectionServiceServer() {}

// UnsafeElectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectionServiceServer will
// result in compilation errors.
type UnsafeElectionServiceServer interface {
	mustEmbedUnimplementedElectionServiceServer()
}

func RegisterElectionServiceServer(s grpc.ServiceRegistrar, srv ElectionServiceServer) {
	s.RegisterService(&ElectionService_ServiceDesc, srv)
}

func _ElectionService_CreateElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateElectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).CreateElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_CreateElection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).CreateElection(ctx, req.(*CreateElectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_GetElectionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetElectionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).GetElectionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_GetElectionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).GetElectionInfo(ctx, req.(*GetElectionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_UpdateElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateElectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).UpdateElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_UpdateElection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).UpdateElection(ctx, req.(*UpdateElectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_DeleteElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteElectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).DeleteElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_DeleteElection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).DeleteElection(ctx, req.(*DeleteElectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ElectionService_ServiceDesc is the grpc.ServiceDesc for ElectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voting.ElectionService",
	HandlerType: (*ElectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateElection",
			Handler:    _ElectionService_CreateElection_Handler,
		},
		{
			MethodName: "GetElectionInfo",
			Handler:    _ElectionService_GetElectionInfo_Handler,
		},
		{
			MethodName: "UpdateElection",
			Handler:    _ElectionService_UpdateElection_Handler,
		},
		{
			MethodName: "DeleteElection",
			Handler:    _ElectionService_DeleteElection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "election.proto",
}
