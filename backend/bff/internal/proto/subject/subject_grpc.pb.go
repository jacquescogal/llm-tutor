// subject.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: subject.proto

package subject

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SubjectService_CreateSubject_FullMethodName           = "/subject.SubjectService/CreateSubject"
	SubjectService_GetSubjectById_FullMethodName          = "/subject.SubjectService/GetSubjectById"
	SubjectService_GetSubjectsByUserId_FullMethodName     = "/subject.SubjectService/GetSubjectsByUserId"
	SubjectService_GetSubjectsByNameSearch_FullMethodName = "/subject.SubjectService/GetSubjectsByNameSearch"
	SubjectService_UpdateSubject_FullMethodName           = "/subject.SubjectService/UpdateSubject"
	SubjectService_DeleteSubject_FullMethodName           = "/subject.SubjectService/DeleteSubject"
)

// SubjectServiceClient is the client API for SubjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubjectServiceClient interface {
	CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*CreateSubjectResponse, error)
	GetSubjectById(ctx context.Context, in *GetSubjectByIdRequest, opts ...grpc.CallOption) (*GetSubjectByIdResponse, error)
	GetSubjectsByUserId(ctx context.Context, in *GetSubjectsByUserIdRequest, opts ...grpc.CallOption) (*GetSubjectsByUserIdResponse, error)
	GetSubjectsByNameSearch(ctx context.Context, in *GetSubjectsByNameSearchRequest, opts ...grpc.CallOption) (*GetSubjectsByNameSearchResponse, error)
	UpdateSubject(ctx context.Context, in *UpdateSubjectRequest, opts ...grpc.CallOption) (*UpdateSubjectResponse, error)
	DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*DeleteSubjectResponse, error)
}

type subjectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubjectServiceClient(cc grpc.ClientConnInterface) SubjectServiceClient {
	return &subjectServiceClient{cc}
}

func (c *subjectServiceClient) CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*CreateSubjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSubjectResponse)
	err := c.cc.Invoke(ctx, SubjectService_CreateSubject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) GetSubjectById(ctx context.Context, in *GetSubjectByIdRequest, opts ...grpc.CallOption) (*GetSubjectByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubjectByIdResponse)
	err := c.cc.Invoke(ctx, SubjectService_GetSubjectById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) GetSubjectsByUserId(ctx context.Context, in *GetSubjectsByUserIdRequest, opts ...grpc.CallOption) (*GetSubjectsByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubjectsByUserIdResponse)
	err := c.cc.Invoke(ctx, SubjectService_GetSubjectsByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) GetSubjectsByNameSearch(ctx context.Context, in *GetSubjectsByNameSearchRequest, opts ...grpc.CallOption) (*GetSubjectsByNameSearchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubjectsByNameSearchResponse)
	err := c.cc.Invoke(ctx, SubjectService_GetSubjectsByNameSearch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) UpdateSubject(ctx context.Context, in *UpdateSubjectRequest, opts ...grpc.CallOption) (*UpdateSubjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSubjectResponse)
	err := c.cc.Invoke(ctx, SubjectService_UpdateSubject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*DeleteSubjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSubjectResponse)
	err := c.cc.Invoke(ctx, SubjectService_DeleteSubject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubjectServiceServer is the server API for SubjectService service.
// All implementations must embed UnimplementedSubjectServiceServer
// for forward compatibility.
type SubjectServiceServer interface {
	CreateSubject(context.Context, *CreateSubjectRequest) (*CreateSubjectResponse, error)
	GetSubjectById(context.Context, *GetSubjectByIdRequest) (*GetSubjectByIdResponse, error)
	GetSubjectsByUserId(context.Context, *GetSubjectsByUserIdRequest) (*GetSubjectsByUserIdResponse, error)
	GetSubjectsByNameSearch(context.Context, *GetSubjectsByNameSearchRequest) (*GetSubjectsByNameSearchResponse, error)
	UpdateSubject(context.Context, *UpdateSubjectRequest) (*UpdateSubjectResponse, error)
	DeleteSubject(context.Context, *DeleteSubjectRequest) (*DeleteSubjectResponse, error)
	mustEmbedUnimplementedSubjectServiceServer()
}

// UnimplementedSubjectServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSubjectServiceServer struct{}

func (UnimplementedSubjectServiceServer) CreateSubject(context.Context, *CreateSubjectRequest) (*CreateSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubject not implemented")
}
func (UnimplementedSubjectServiceServer) GetSubjectById(context.Context, *GetSubjectByIdRequest) (*GetSubjectByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubjectById not implemented")
}
func (UnimplementedSubjectServiceServer) GetSubjectsByUserId(context.Context, *GetSubjectsByUserIdRequest) (*GetSubjectsByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubjectsByUserId not implemented")
}
func (UnimplementedSubjectServiceServer) GetSubjectsByNameSearch(context.Context, *GetSubjectsByNameSearchRequest) (*GetSubjectsByNameSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubjectsByNameSearch not implemented")
}
func (UnimplementedSubjectServiceServer) UpdateSubject(context.Context, *UpdateSubjectRequest) (*UpdateSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubject not implemented")
}
func (UnimplementedSubjectServiceServer) DeleteSubject(context.Context, *DeleteSubjectRequest) (*DeleteSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubject not implemented")
}
func (UnimplementedSubjectServiceServer) mustEmbedUnimplementedSubjectServiceServer() {}
func (UnimplementedSubjectServiceServer) testEmbeddedByValue()                        {}

// UnsafeSubjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubjectServiceServer will
// result in compilation errors.
type UnsafeSubjectServiceServer interface {
	mustEmbedUnimplementedSubjectServiceServer()
}

func RegisterSubjectServiceServer(s grpc.ServiceRegistrar, srv SubjectServiceServer) {
	// If the following call pancis, it indicates UnimplementedSubjectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SubjectService_ServiceDesc, srv)
}

func _SubjectService_CreateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).CreateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_CreateSubject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).CreateSubject(ctx, req.(*CreateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_GetSubjectById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).GetSubjectById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_GetSubjectById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).GetSubjectById(ctx, req.(*GetSubjectByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_GetSubjectsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).GetSubjectsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_GetSubjectsByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).GetSubjectsByUserId(ctx, req.(*GetSubjectsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_GetSubjectsByNameSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectsByNameSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).GetSubjectsByNameSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_GetSubjectsByNameSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).GetSubjectsByNameSearch(ctx, req.(*GetSubjectsByNameSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_UpdateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).UpdateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_UpdateSubject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).UpdateSubject(ctx, req.(*UpdateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_DeleteSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).DeleteSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_DeleteSubject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).DeleteSubject(ctx, req.(*DeleteSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubjectService_ServiceDesc is the grpc.ServiceDesc for SubjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subject.SubjectService",
	HandlerType: (*SubjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubject",
			Handler:    _SubjectService_CreateSubject_Handler,
		},
		{
			MethodName: "GetSubjectById",
			Handler:    _SubjectService_GetSubjectById_Handler,
		},
		{
			MethodName: "GetSubjectsByUserId",
			Handler:    _SubjectService_GetSubjectsByUserId_Handler,
		},
		{
			MethodName: "GetSubjectsByNameSearch",
			Handler:    _SubjectService_GetSubjectsByNameSearch_Handler,
		},
		{
			MethodName: "UpdateSubject",
			Handler:    _SubjectService_UpdateSubject_Handler,
		},
		{
			MethodName: "DeleteSubject",
			Handler:    _SubjectService_DeleteSubject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subject.proto",
}
