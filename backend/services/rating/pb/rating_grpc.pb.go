// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: proto/rating.proto

package pb

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
	RatingService_HostRating_FullMethodName                = "/RatingService/HostRating"
	RatingService_RateHost_FullMethodName                  = "/RatingService/RateHost"
	RatingService_RemoveHostRating_FullMethodName          = "/RatingService/RemoveHostRating"
	RatingService_GetMyHostRating_FullMethodName           = "/RatingService/GetMyHostRating"
	RatingService_AccommodationRating_FullMethodName       = "/RatingService/AccommodationRating"
	RatingService_RateAccommodation_FullMethodName         = "/RatingService/RateAccommodation"
	RatingService_RemoveAccommodationRating_FullMethodName = "/RatingService/RemoveAccommodationRating"
	RatingService_GetMyAccommodationRating_FullMethodName  = "/RatingService/GetMyAccommodationRating"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	// Host rating
	HostRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	RateHost(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	RemoveHostRating(ctx context.Context, in *RemoveRatingRequest, opts ...grpc.CallOption) (*RemoveRatingResponse, error)
	GetMyHostRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*HostRatingList, error)
	// Accommodation rating
	AccommodationRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	RateAccommodation(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	RemoveAccommodationRating(ctx context.Context, in *RemoveRatingRequest, opts ...grpc.CallOption) (*RemoveRatingResponse, error)
	GetMyAccommodationRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*AccommodationRatingList, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) HostRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, RatingService_HostRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RateHost(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, RatingService_RateHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RemoveHostRating(ctx context.Context, in *RemoveRatingRequest, opts ...grpc.CallOption) (*RemoveRatingResponse, error) {
	out := new(RemoveRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_RemoveHostRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetMyHostRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*HostRatingList, error) {
	out := new(HostRatingList)
	err := c.cc.Invoke(ctx, RatingService_GetMyHostRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) AccommodationRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, RatingService_AccommodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RateAccommodation(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, RatingService_RateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RemoveAccommodationRating(ctx context.Context, in *RemoveRatingRequest, opts ...grpc.CallOption) (*RemoveRatingResponse, error) {
	out := new(RemoveRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_RemoveAccommodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetMyAccommodationRating(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*AccommodationRatingList, error) {
	out := new(AccommodationRatingList)
	err := c.cc.Invoke(ctx, RatingService_GetMyAccommodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	// Host rating
	HostRating(context.Context, *IdRequest) (*RatingResponse, error)
	RateHost(context.Context, *RateRequest) (*RateResponse, error)
	RemoveHostRating(context.Context, *RemoveRatingRequest) (*RemoveRatingResponse, error)
	GetMyHostRating(context.Context, *IdRequest) (*HostRatingList, error)
	// Accommodation rating
	AccommodationRating(context.Context, *IdRequest) (*RatingResponse, error)
	RateAccommodation(context.Context, *RateRequest) (*RateResponse, error)
	RemoveAccommodationRating(context.Context, *RemoveRatingRequest) (*RemoveRatingResponse, error)
	GetMyAccommodationRating(context.Context, *IdRequest) (*AccommodationRatingList, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) HostRating(context.Context, *IdRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostRating not implemented")
}
func (UnimplementedRatingServiceServer) RateHost(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateHost not implemented")
}
func (UnimplementedRatingServiceServer) RemoveHostRating(context.Context, *RemoveRatingRequest) (*RemoveRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHostRating not implemented")
}
func (UnimplementedRatingServiceServer) GetMyHostRating(context.Context, *IdRequest) (*HostRatingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyHostRating not implemented")
}
func (UnimplementedRatingServiceServer) AccommodationRating(context.Context, *IdRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) RateAccommodation(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) RemoveAccommodationRating(context.Context, *RemoveRatingRequest) (*RemoveRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) GetMyAccommodationRating(context.Context, *IdRequest) (*AccommodationRatingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_HostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).HostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_HostRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).HostRating(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_RateHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RateHost(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RemoveHostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RemoveHostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_RemoveHostRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RemoveHostRating(ctx, req.(*RemoveRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetMyHostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetMyHostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetMyHostRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetMyHostRating(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_AccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).AccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_AccommodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).AccommodationRating(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_RateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RateAccommodation(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RemoveAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RemoveAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_RemoveAccommodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RemoveAccommodationRating(ctx, req.(*RemoveRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetMyAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetMyAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetMyAccommodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetMyAccommodationRating(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HostRating",
			Handler:    _RatingService_HostRating_Handler,
		},
		{
			MethodName: "RateHost",
			Handler:    _RatingService_RateHost_Handler,
		},
		{
			MethodName: "RemoveHostRating",
			Handler:    _RatingService_RemoveHostRating_Handler,
		},
		{
			MethodName: "GetMyHostRating",
			Handler:    _RatingService_GetMyHostRating_Handler,
		},
		{
			MethodName: "AccommodationRating",
			Handler:    _RatingService_AccommodationRating_Handler,
		},
		{
			MethodName: "RateAccommodation",
			Handler:    _RatingService_RateAccommodation_Handler,
		},
		{
			MethodName: "RemoveAccommodationRating",
			Handler:    _RatingService_RemoveAccommodationRating_Handler,
		},
		{
			MethodName: "GetMyAccommodationRating",
			Handler:    _RatingService_GetMyAccommodationRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rating.proto",
}
