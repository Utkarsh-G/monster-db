// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: monster.proto

package monster

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

// MonsterServiceClient is the client API for MonsterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonsterServiceClient interface {
	GetMonster(ctx context.Context, in *MonsterName, opts ...grpc.CallOption) (*Monster, error)
}

type monsterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonsterServiceClient(cc grpc.ClientConnInterface) MonsterServiceClient {
	return &monsterServiceClient{cc}
}

func (c *monsterServiceClient) GetMonster(ctx context.Context, in *MonsterName, opts ...grpc.CallOption) (*Monster, error) {
	out := new(Monster)
	err := c.cc.Invoke(ctx, "/mmonster.MonsterService/GetMonster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonsterServiceServer is the server API for MonsterService service.
// All implementations should embed UnimplementedMonsterServiceServer
// for forward compatibility
type MonsterServiceServer interface {
	GetMonster(context.Context, *MonsterName) (*Monster, error)
}

// UnimplementedMonsterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMonsterServiceServer struct {
}

func (UnimplementedMonsterServiceServer) GetMonster(context.Context, *MonsterName) (*Monster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonster not implemented")
}

// UnsafeMonsterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonsterServiceServer will
// result in compilation errors.
type UnsafeMonsterServiceServer interface {
	mustEmbedUnimplementedMonsterServiceServer()
}

func RegisterMonsterServiceServer(s grpc.ServiceRegistrar, srv MonsterServiceServer) {
	s.RegisterService(&MonsterService_ServiceDesc, srv)
}

func _MonsterService_GetMonster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonsterName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonsterServiceServer).GetMonster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mmonster.MonsterService/GetMonster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonsterServiceServer).GetMonster(ctx, req.(*MonsterName))
	}
	return interceptor(ctx, in, info, handler)
}

// MonsterService_ServiceDesc is the grpc.ServiceDesc for MonsterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonsterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mmonster.MonsterService",
	HandlerType: (*MonsterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMonster",
			Handler:    _MonsterService_GetMonster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monster.proto",
}