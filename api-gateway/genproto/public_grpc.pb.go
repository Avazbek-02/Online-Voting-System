// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: public.proto

package genproto

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
	PublicService_CreatePublic_FullMethodName  = "/protos.PublicService/CreatePublic"
	PublicService_GetPublicInfo_FullMethodName = "/protos.PublicService/GetPublicInfo"
	PublicService_UpdatePublic_FullMethodName  = "/protos.PublicService/UpdatePublic"
	PublicService_DeletePublic_FullMethodName  = "/protos.PublicService/DeletePublic"
)

// PublicServiceClient is the client API for PublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicServiceClient interface {
	CreatePublic(ctx context.Context, in *CreatePublicRequest, opts ...grpc.CallOption) (*PublicResponse, error)
	GetPublicInfo(ctx context.Context, in *GetPublicInfoRequest, opts ...grpc.CallOption) (*PublicResponse, error)
	UpdatePublic(ctx context.Context, in *UpdatePublicRequest, opts ...grpc.CallOption) (*PublicResponse, error)
	DeletePublic(ctx context.Context, in *DeletePublicRequest, opts ...grpc.CallOption) (*Void, error)
}

type publicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicServiceClient(cc grpc.ClientConnInterface) PublicServiceClient {
	return &publicServiceClient{cc}
}

func (c *publicServiceClient) CreatePublic(ctx context.Context, in *CreatePublicRequest, opts ...grpc.CallOption) (*PublicResponse, error) {
	out := new(PublicResponse)
	err := c.cc.Invoke(ctx, PublicService_CreatePublic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) GetPublicInfo(ctx context.Context, in *GetPublicInfoRequest, opts ...grpc.CallOption) (*PublicResponse, error) {
	out := new(PublicResponse)
	err := c.cc.Invoke(ctx, PublicService_GetPublicInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) UpdatePublic(ctx context.Context, in *UpdatePublicRequest, opts ...grpc.CallOption) (*PublicResponse, error) {
	out := new(PublicResponse)
	err := c.cc.Invoke(ctx, PublicService_UpdatePublic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) DeletePublic(ctx context.Context, in *DeletePublicRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PublicService_DeletePublic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicServiceServer is the server API for PublicService service.
// All implementations must embed UnimplementedPublicServiceServer
// for forward compatibility
type PublicServiceServer interface {
	CreatePublic(context.Context, *CreatePublicRequest) (*PublicResponse, error)
	GetPublicInfo(context.Context, *GetPublicInfoRequest) (*PublicResponse, error)
	UpdatePublic(context.Context, *UpdatePublicRequest) (*PublicResponse, error)
	DeletePublic(context.Context, *DeletePublicRequest) (*Void, error)
	mustEmbedUnimplementedPublicServiceServer()
}

// UnimplementedPublicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicServiceServer struct {
}

func (UnimplementedPublicServiceServer) CreatePublic(context.Context, *CreatePublicRequest) (*PublicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublic not implemented")
}
func (UnimplementedPublicServiceServer) GetPublicInfo(context.Context, *GetPublicInfoRequest) (*PublicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicInfo not implemented")
}
func (UnimplementedPublicServiceServer) UpdatePublic(context.Context, *UpdatePublicRequest) (*PublicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublic not implemented")
}
func (UnimplementedPublicServiceServer) DeletePublic(context.Context, *DeletePublicRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePublic not implemented")
}
func (UnimplementedPublicServiceServer) mustEmbedUnimplementedPublicServiceServer() {}

// UnsafePublicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicServiceServer will
// result in compilation errors.
type UnsafePublicServiceServer interface {
	mustEmbedUnimplementedPublicServiceServer()
}

func RegisterPublicServiceServer(s grpc.ServiceRegistrar, srv PublicServiceServer) {
	s.RegisterService(&PublicService_ServiceDesc, srv)
}

func _PublicService_CreatePublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).CreatePublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_CreatePublic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).CreatePublic(ctx, req.(*CreatePublicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_GetPublicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).GetPublicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_GetPublicInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).GetPublicInfo(ctx, req.(*GetPublicInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_UpdatePublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).UpdatePublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_UpdatePublic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).UpdatePublic(ctx, req.(*UpdatePublicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_DeletePublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePublicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).DeletePublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_DeletePublic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).DeletePublic(ctx, req.(*DeletePublicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicService_ServiceDesc is the grpc.ServiceDesc for PublicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PublicService",
	HandlerType: (*PublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePublic",
			Handler:    _PublicService_CreatePublic_Handler,
		},
		{
			MethodName: "GetPublicInfo",
			Handler:    _PublicService_GetPublicInfo_Handler,
		},
		{
			MethodName: "UpdatePublic",
			Handler:    _PublicService_UpdatePublic_Handler,
		},
		{
			MethodName: "DeletePublic",
			Handler:    _PublicService_DeletePublic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public.proto",
}
