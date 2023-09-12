// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: lock_service.proto

package proto

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
	LockManagerService_Lock_FullMethodName   = "/lock_manager.LockManagerService/Lock"
	LockManagerService_Unlock_FullMethodName = "/lock_manager.LockManagerService/Unlock"
)

// LockManagerServiceClient is the client API for LockManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LockManagerServiceClient interface {
	// Lock locks resource and returns token that can be used to unlock it.
	Lock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockRes, error)
	// Unlock unlocks resource only if token fits.
	Unlock(ctx context.Context, in *UnlockReq, opts ...grpc.CallOption) (*UnlockRes, error)
}

type lockManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLockManagerServiceClient(cc grpc.ClientConnInterface) LockManagerServiceClient {
	return &lockManagerServiceClient{cc}
}

func (c *lockManagerServiceClient) Lock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockRes, error) {
	out := new(LockRes)
	err := c.cc.Invoke(ctx, LockManagerService_Lock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockManagerServiceClient) Unlock(ctx context.Context, in *UnlockReq, opts ...grpc.CallOption) (*UnlockRes, error) {
	out := new(UnlockRes)
	err := c.cc.Invoke(ctx, LockManagerService_Unlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockManagerServiceServer is the server API for LockManagerService service.
// All implementations must embed UnimplementedLockManagerServiceServer
// for forward compatibility
type LockManagerServiceServer interface {
	// Lock locks resource and returns token that can be used to unlock it.
	Lock(context.Context, *LockReq) (*LockRes, error)
	// Unlock unlocks resource only if token fits.
	Unlock(context.Context, *UnlockReq) (*UnlockRes, error)
	mustEmbedUnimplementedLockManagerServiceServer()
}

// UnimplementedLockManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLockManagerServiceServer struct {
}

func (UnimplementedLockManagerServiceServer) Lock(context.Context, *LockReq) (*LockRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (UnimplementedLockManagerServiceServer) Unlock(context.Context, *UnlockReq) (*UnlockRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}
func (UnimplementedLockManagerServiceServer) mustEmbedUnimplementedLockManagerServiceServer() {}

// UnsafeLockManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LockManagerServiceServer will
// result in compilation errors.
type UnsafeLockManagerServiceServer interface {
	mustEmbedUnimplementedLockManagerServiceServer()
}

func RegisterLockManagerServiceServer(s grpc.ServiceRegistrar, srv LockManagerServiceServer) {
	s.RegisterService(&LockManagerService_ServiceDesc, srv)
}

func _LockManagerService_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockManagerServiceServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LockManagerService_Lock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockManagerServiceServer).Lock(ctx, req.(*LockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockManagerService_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockManagerServiceServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LockManagerService_Unlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockManagerServiceServer).Unlock(ctx, req.(*UnlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LockManagerService_ServiceDesc is the grpc.ServiceDesc for LockManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LockManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lock_manager.LockManagerService",
	HandlerType: (*LockManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lock",
			Handler:    _LockManagerService_Lock_Handler,
		},
		{
			MethodName: "Unlock",
			Handler:    _LockManagerService_Unlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lock_service.proto",
}
