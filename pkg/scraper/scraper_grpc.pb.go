// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: scraper.proto

package scraper_v1

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

// ScraperClient is the client API for Scraper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScraperClient interface {
	Scrap(ctx context.Context, in *ScrapUrl, opts ...grpc.CallOption) (*ScrapData, error)
}

type scraperClient struct {
	cc grpc.ClientConnInterface
}

func NewScraperClient(cc grpc.ClientConnInterface) ScraperClient {
	return &scraperClient{cc}
}

func (c *scraperClient) Scrap(ctx context.Context, in *ScrapUrl, opts ...grpc.CallOption) (*ScrapData, error) {
	out := new(ScrapData)
	err := c.cc.Invoke(ctx, "/Scraper/Scrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScraperServer is the server API for Scraper service.
// All implementations must embed UnimplementedScraperServer
// for forward compatibility
type ScraperServer interface {
	Scrap(context.Context, *ScrapUrl) (*ScrapData, error)
	mustEmbedUnimplementedScraperServer()
}

// UnimplementedScraperServer must be embedded to have forward compatible implementations.
type UnimplementedScraperServer struct {
}

func (UnimplementedScraperServer) Scrap(context.Context, *ScrapUrl) (*ScrapData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scrap not implemented")
}
func (UnimplementedScraperServer) mustEmbedUnimplementedScraperServer() {}

// UnsafeScraperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScraperServer will
// result in compilation errors.
type UnsafeScraperServer interface {
	mustEmbedUnimplementedScraperServer()
}

func RegisterScraperServer(s grpc.ServiceRegistrar, srv ScraperServer) {
	s.RegisterService(&Scraper_ServiceDesc, srv)
}

func _Scraper_Scrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScraperServer).Scrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Scraper/Scrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScraperServer).Scrap(ctx, req.(*ScrapUrl))
	}
	return interceptor(ctx, in, info, handler)
}

// Scraper_ServiceDesc is the grpc.ServiceDesc for Scraper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scraper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Scraper",
	HandlerType: (*ScraperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Scrap",
			Handler:    _Scraper_Scrap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scraper.proto",
}
