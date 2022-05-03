// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// SimCardServiceClient is the client API for SimCardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimCardServiceClient interface {
	GetSimCards(ctx context.Context, in *GetSimCardsRequest, opts ...grpc.CallOption) (*GetSimCardsResponse, error)
}

type simCardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimCardServiceClient(cc grpc.ClientConnInterface) SimCardServiceClient {
	return &simCardServiceClient{cc}
}

func (c *simCardServiceClient) GetSimCards(ctx context.Context, in *GetSimCardsRequest, opts ...grpc.CallOption) (*GetSimCardsResponse, error) {
	out := new(GetSimCardsResponse)
	err := c.cc.Invoke(ctx, "/simcards.v1.SimCardService/GetSimCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimCardServiceServer is the server API for SimCardService service.
// All implementations must embed UnimplementedSimCardServiceServer
// for forward compatibility
type SimCardServiceServer interface {
	GetSimCards(context.Context, *GetSimCardsRequest) (*GetSimCardsResponse, error)
	mustEmbedUnimplementedSimCardServiceServer()
}

// UnimplementedSimCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSimCardServiceServer struct {
}

func (UnimplementedSimCardServiceServer) GetSimCards(context.Context, *GetSimCardsRequest) (*GetSimCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimCards not implemented")
}
func (UnimplementedSimCardServiceServer) mustEmbedUnimplementedSimCardServiceServer() {}

// UnsafeSimCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimCardServiceServer will
// result in compilation errors.
type UnsafeSimCardServiceServer interface {
	mustEmbedUnimplementedSimCardServiceServer()
}

func RegisterSimCardServiceServer(s grpc.ServiceRegistrar, srv SimCardServiceServer) {
	s.RegisterService(&SimCardService_ServiceDesc, srv)
}

func _SimCardService_GetSimCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCardServiceServer).GetSimCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simcards.v1.SimCardService/GetSimCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCardServiceServer).GetSimCards(ctx, req.(*GetSimCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimCardService_ServiceDesc is the grpc.ServiceDesc for SimCardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimCardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simcards.v1.SimCardService",
	HandlerType: (*SimCardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSimCards",
			Handler:    _SimCardService_GetSimCards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simcards/v1/simcards.proto",
}
