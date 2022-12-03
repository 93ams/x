// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: ddl.proto

package model

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

// DDLClient is the client API for DDL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DDLClient interface {
	CreateKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	AlterKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	DropKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	ListKeySpaces(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	AlterTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	DropTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	ListTables(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type dDLClient struct {
	cc grpc.ClientConnInterface
}

func NewDDLClient(cc grpc.ClientConnInterface) DDLClient {
	return &dDLClient{cc}
}

func (c *dDLClient) CreateKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/CreateKeySpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) AlterKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/AlterKeySpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) DropKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/DropKeySpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) ListKeySpaces(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/ListKeySpaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) GetKeySpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/GetKeySpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) AlterTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/AlterTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) DropTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/DropTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) ListTables(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/ListTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDLClient) GetTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cql.grpc.v1.DDL/GetTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DDLServer is the server API for DDL service.
// All implementations must embed UnimplementedDDLServer
// for forward compatibility
type DDLServer interface {
	CreateKeySpace(context.Context, *Empty) (*Empty, error)
	AlterKeySpace(context.Context, *Empty) (*Empty, error)
	DropKeySpace(context.Context, *Empty) (*Empty, error)
	ListKeySpaces(context.Context, *Empty) (*Empty, error)
	GetKeySpace(context.Context, *Empty) (*Empty, error)
	CreateTable(context.Context, *Empty) (*Empty, error)
	AlterTable(context.Context, *Empty) (*Empty, error)
	DropTable(context.Context, *Empty) (*Empty, error)
	ListTables(context.Context, *Empty) (*Empty, error)
	GetTable(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedDDLServer()
}

// UnimplementedDDLServer must be embedded to have forward compatible implementations.
type UnimplementedDDLServer struct {
}

func (UnimplementedDDLServer) CreateKeySpace(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeySpace not implemented")
}
func (UnimplementedDDLServer) AlterKeySpace(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterKeySpace not implemented")
}
func (UnimplementedDDLServer) DropKeySpace(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropKeySpace not implemented")
}
func (UnimplementedDDLServer) ListKeySpaces(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeySpaces not implemented")
}
func (UnimplementedDDLServer) GetKeySpace(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeySpace not implemented")
}
func (UnimplementedDDLServer) CreateTable(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedDDLServer) AlterTable(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterTable not implemented")
}
func (UnimplementedDDLServer) DropTable(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropTable not implemented")
}
func (UnimplementedDDLServer) ListTables(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTables not implemented")
}
func (UnimplementedDDLServer) GetTable(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTable not implemented")
}
func (UnimplementedDDLServer) mustEmbedUnimplementedDDLServer() {}

// UnsafeDDLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DDLServer will
// result in compilation errors.
type UnsafeDDLServer interface {
	mustEmbedUnimplementedDDLServer()
}

func RegisterDDLServer(s grpc.ServiceRegistrar, srv DDLServer) {
	s.RegisterService(&DDL_ServiceDesc, srv)
}

func _DDL_CreateKeySpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).CreateKeySpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/CreateKeySpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).CreateKeySpace(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_AlterKeySpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).AlterKeySpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/AlterKeySpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).AlterKeySpace(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_DropKeySpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).DropKeySpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/DropKeySpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).DropKeySpace(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_ListKeySpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).ListKeySpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/ListKeySpaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).ListKeySpaces(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_GetKeySpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).GetKeySpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/GetKeySpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).GetKeySpace(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).CreateTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_AlterTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).AlterTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/AlterTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).AlterTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_DropTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).DropTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/DropTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).DropTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).ListTables(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDL_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cql.grpc.v1.DDL/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServer).GetTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DDL_ServiceDesc is the grpc.ServiceDesc for DDL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DDL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cql.grpc.v1.DDL",
	HandlerType: (*DDLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKeySpace",
			Handler:    _DDL_CreateKeySpace_Handler,
		},
		{
			MethodName: "AlterKeySpace",
			Handler:    _DDL_AlterKeySpace_Handler,
		},
		{
			MethodName: "DropKeySpace",
			Handler:    _DDL_DropKeySpace_Handler,
		},
		{
			MethodName: "ListKeySpaces",
			Handler:    _DDL_ListKeySpaces_Handler,
		},
		{
			MethodName: "GetKeySpace",
			Handler:    _DDL_GetKeySpace_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _DDL_CreateTable_Handler,
		},
		{
			MethodName: "AlterTable",
			Handler:    _DDL_AlterTable_Handler,
		},
		{
			MethodName: "DropTable",
			Handler:    _DDL_DropTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _DDL_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _DDL_GetTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ddl.proto",
}
