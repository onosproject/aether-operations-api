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

// SiteServiceClient is the client API for SiteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiteServiceClient interface {
	GetSites(ctx context.Context, in *GetSitesRequest, opts ...grpc.CallOption) (*GetSitesResponse, error)
}

type siteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSiteServiceClient(cc grpc.ClientConnInterface) SiteServiceClient {
	return &siteServiceClient{cc}
}

func (c *siteServiceClient) GetSites(ctx context.Context, in *GetSitesRequest, opts ...grpc.CallOption) (*GetSitesResponse, error) {
	out := new(GetSitesResponse)
	err := c.cc.Invoke(ctx, "/v1.SiteService/GetSites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiteServiceServer is the server API for SiteService service.
// All implementations must embed UnimplementedSiteServiceServer
// for forward compatibility
type SiteServiceServer interface {
	GetSites(context.Context, *GetSitesRequest) (*GetSitesResponse, error)
	mustEmbedUnimplementedSiteServiceServer()
}

// UnimplementedSiteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSiteServiceServer struct {
}

func (UnimplementedSiteServiceServer) GetSites(context.Context, *GetSitesRequest) (*GetSitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSites not implemented")
}
func (UnimplementedSiteServiceServer) mustEmbedUnimplementedSiteServiceServer() {}

// UnsafeSiteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiteServiceServer will
// result in compilation errors.
type UnsafeSiteServiceServer interface {
	mustEmbedUnimplementedSiteServiceServer()
}

func RegisterSiteServiceServer(s grpc.ServiceRegistrar, srv SiteServiceServer) {
	s.RegisterService(&SiteService_ServiceDesc, srv)
}

func _SiteService_GetSites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteServiceServer).GetSites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SiteService/GetSites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteServiceServer).GetSites(ctx, req.(*GetSitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SiteService_ServiceDesc is the grpc.ServiceDesc for SiteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SiteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SiteService",
	HandlerType: (*SiteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSites",
			Handler:    _SiteService_GetSites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sites.proto",
}
