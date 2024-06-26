// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: candidate.proto

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
	CandidateService_CreateCandidate_FullMethodName  = "/voting.CandidateService/CreateCandidate"
	CandidateService_GetCandidateInfo_FullMethodName = "/voting.CandidateService/GetCandidateInfo"
	CandidateService_UpdateCandidate_FullMethodName  = "/voting.CandidateService/UpdateCandidate"
	CandidateService_DeleteCandidate_FullMethodName  = "/voting.CandidateService/DeleteCandidate"
)

// CandidateServiceClient is the client API for CandidateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CandidateServiceClient interface {
	CreateCandidate(ctx context.Context, in *CreateCandidateRequest, opts ...grpc.CallOption) (*CandidateResponse, error)
	GetCandidateInfo(ctx context.Context, in *GetCandidateInfoRequest, opts ...grpc.CallOption) (*CandidateResponse, error)
	UpdateCandidate(ctx context.Context, in *UpdateCandidateRequest, opts ...grpc.CallOption) (*CandidateResponse, error)
	DeleteCandidate(ctx context.Context, in *DeleteCandidateRequest, opts ...grpc.CallOption) (*Void1, error)
}

type candidateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCandidateServiceClient(cc grpc.ClientConnInterface) CandidateServiceClient {
	return &candidateServiceClient{cc}
}

func (c *candidateServiceClient) CreateCandidate(ctx context.Context, in *CreateCandidateRequest, opts ...grpc.CallOption) (*CandidateResponse, error) {
	out := new(CandidateResponse)
	err := c.cc.Invoke(ctx, CandidateService_CreateCandidate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candidateServiceClient) GetCandidateInfo(ctx context.Context, in *GetCandidateInfoRequest, opts ...grpc.CallOption) (*CandidateResponse, error) {
	out := new(CandidateResponse)
	err := c.cc.Invoke(ctx, CandidateService_GetCandidateInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candidateServiceClient) UpdateCandidate(ctx context.Context, in *UpdateCandidateRequest, opts ...grpc.CallOption) (*CandidateResponse, error) {
	out := new(CandidateResponse)
	err := c.cc.Invoke(ctx, CandidateService_UpdateCandidate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candidateServiceClient) DeleteCandidate(ctx context.Context, in *DeleteCandidateRequest, opts ...grpc.CallOption) (*Void1, error) {
	out := new(Void1)
	err := c.cc.Invoke(ctx, CandidateService_DeleteCandidate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CandidateServiceServer is the server API for CandidateService service.
// All implementations must embed UnimplementedCandidateServiceServer
// for forward compatibility
type CandidateServiceServer interface {
	CreateCandidate(context.Context, *CreateCandidateRequest) (*CandidateResponse, error)
	GetCandidateInfo(context.Context, *GetCandidateInfoRequest) (*CandidateResponse, error)
	UpdateCandidate(context.Context, *UpdateCandidateRequest) (*CandidateResponse, error)
	DeleteCandidate(context.Context, *DeleteCandidateRequest) (*Void1, error)
	mustEmbedUnimplementedCandidateServiceServer()
}

// UnimplementedCandidateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCandidateServiceServer struct {
}

func (UnimplementedCandidateServiceServer) CreateCandidate(context.Context, *CreateCandidateRequest) (*CandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCandidate not implemented")
}
func (UnimplementedCandidateServiceServer) GetCandidateInfo(context.Context, *GetCandidateInfoRequest) (*CandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandidateInfo not implemented")
}
func (UnimplementedCandidateServiceServer) UpdateCandidate(context.Context, *UpdateCandidateRequest) (*CandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCandidate not implemented")
}
func (UnimplementedCandidateServiceServer) DeleteCandidate(context.Context, *DeleteCandidateRequest) (*Void1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCandidate not implemented")
}
func (UnimplementedCandidateServiceServer) mustEmbedUnimplementedCandidateServiceServer() {}

// UnsafeCandidateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CandidateServiceServer will
// result in compilation errors.
type UnsafeCandidateServiceServer interface {
	mustEmbedUnimplementedCandidateServiceServer()
}

func RegisterCandidateServiceServer(s grpc.ServiceRegistrar, srv CandidateServiceServer) {
	s.RegisterService(&CandidateService_ServiceDesc, srv)
}

func _CandidateService_CreateCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).CreateCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_CreateCandidate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).CreateCandidate(ctx, req.(*CreateCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandidateService_GetCandidateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandidateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).GetCandidateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_GetCandidateInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).GetCandidateInfo(ctx, req.(*GetCandidateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandidateService_UpdateCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).UpdateCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_UpdateCandidate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).UpdateCandidate(ctx, req.(*UpdateCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandidateService_DeleteCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).DeleteCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_DeleteCandidate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).DeleteCandidate(ctx, req.(*DeleteCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CandidateService_ServiceDesc is the grpc.ServiceDesc for CandidateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CandidateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voting.CandidateService",
	HandlerType: (*CandidateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCandidate",
			Handler:    _CandidateService_CreateCandidate_Handler,
		},
		{
			MethodName: "GetCandidateInfo",
			Handler:    _CandidateService_GetCandidateInfo_Handler,
		},
		{
			MethodName: "UpdateCandidate",
			Handler:    _CandidateService_UpdateCandidate_Handler,
		},
		{
			MethodName: "DeleteCandidate",
			Handler:    _CandidateService_DeleteCandidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "candidate.proto",
}
