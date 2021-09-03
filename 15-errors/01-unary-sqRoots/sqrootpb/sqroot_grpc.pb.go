// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sqrootpb

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

// SquareRootServiceClient is the client API for SquareRootService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SquareRootServiceClient interface {
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type squareRootServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSquareRootServiceClient(cc grpc.ClientConnInterface) SquareRootServiceClient {
	return &squareRootServiceClient{cc}
}

func (c *squareRootServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/sqroot.SquareRootService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquareRootServiceServer is the server API for SquareRootService service.
// All implementations must embed UnimplementedSquareRootServiceServer
// for forward compatibility
type SquareRootServiceServer interface {
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
	mustEmbedUnimplementedSquareRootServiceServer()
}

// UnimplementedSquareRootServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSquareRootServiceServer struct {
}

func (UnimplementedSquareRootServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}
func (UnimplementedSquareRootServiceServer) mustEmbedUnimplementedSquareRootServiceServer() {}

// UnsafeSquareRootServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SquareRootServiceServer will
// result in compilation errors.
type UnsafeSquareRootServiceServer interface {
	mustEmbedUnimplementedSquareRootServiceServer()
}

func RegisterSquareRootServiceServer(s grpc.ServiceRegistrar, srv SquareRootServiceServer) {
	s.RegisterService(&SquareRootService_ServiceDesc, srv)
}

func _SquareRootService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquareRootServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqroot.SquareRootService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquareRootServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SquareRootService_ServiceDesc is the grpc.ServiceDesc for SquareRootService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SquareRootService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sqroot.SquareRootService",
	HandlerType: (*SquareRootServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SquareRoot",
			Handler:    _SquareRootService_SquareRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sqroot.proto",
}
