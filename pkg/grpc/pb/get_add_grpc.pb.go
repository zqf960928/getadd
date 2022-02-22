// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GetAddClient is the client API for GetAdd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetAddClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
}

type getAddClient struct {
	cc grpc.ClientConnInterface
}

func NewGetAddClient(cc grpc.ClientConnInterface) GetAddClient {
	return &getAddClient{cc}
}

func (c *getAddClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/pb.GetAdd/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetAddServer is the server API for GetAdd service.
// All implementations must embed UnimplementedGetAddServer
// for forward compatibility
type GetAddServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	mustEmbedUnimplementedGetAddServer()
}

// UnimplementedGetAddServer must be embedded to have forward compatible implementations.
type UnimplementedGetAddServer struct {
}

func (UnimplementedGetAddServer) Add(context.Context, *AddRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedGetAddServer) mustEmbedUnimplementedGetAddServer() {}

// UnsafeGetAddServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetAddServer will
// result in compilation errors.
type UnsafeGetAddServer interface {
	mustEmbedUnimplementedGetAddServer()
}

func RegisterGetAddServer(s grpc.ServiceRegistrar, srv GetAddServer) {
	s.RegisterService(&_GetAdd_serviceDesc, srv)
}

func _GetAdd_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetAddServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GetAdd/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetAddServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetAdd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GetAdd",
	HandlerType: (*GetAddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _GetAdd_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "get_add/pkg/grpc/pb/get_add.proto",
}