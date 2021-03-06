// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
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
const _ = grpc.SupportPackageIsVersion7

// SecretStoreHealthsClient is the client API for SecretStoreHealths service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretStoreHealthsClient interface {
	// List lists an org's secret store healths
	List(ctx context.Context, in *SecretStoreHealthListRequest, opts ...grpc.CallOption) (*SecretStoreHealthListResponse, error)
	// Check pushes a healthcheck request for a secret store
	Check(ctx context.Context, in *SecretStoreHealthCheckRequest, opts ...grpc.CallOption) (*SecretStoreHealthCheckResponse, error)
}

type secretStoreHealthsClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretStoreHealthsClient(cc grpc.ClientConnInterface) SecretStoreHealthsClient {
	return &secretStoreHealthsClient{cc}
}

func (c *secretStoreHealthsClient) List(ctx context.Context, in *SecretStoreHealthListRequest, opts ...grpc.CallOption) (*SecretStoreHealthListResponse, error) {
	out := new(SecretStoreHealthListResponse)
	err := c.cc.Invoke(ctx, "/v1.SecretStoreHealths/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretStoreHealthsClient) Check(ctx context.Context, in *SecretStoreHealthCheckRequest, opts ...grpc.CallOption) (*SecretStoreHealthCheckResponse, error) {
	out := new(SecretStoreHealthCheckResponse)
	err := c.cc.Invoke(ctx, "/v1.SecretStoreHealths/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretStoreHealthsServer is the server API for SecretStoreHealths service.
// All implementations must embed UnimplementedSecretStoreHealthsServer
// for forward compatibility
type SecretStoreHealthsServer interface {
	// List lists an org's secret store healths
	List(context.Context, *SecretStoreHealthListRequest) (*SecretStoreHealthListResponse, error)
	// Check pushes a healthcheck request for a secret store
	Check(context.Context, *SecretStoreHealthCheckRequest) (*SecretStoreHealthCheckResponse, error)
	mustEmbedUnimplementedSecretStoreHealthsServer()
}

// UnimplementedSecretStoreHealthsServer must be embedded to have forward compatible implementations.
type UnimplementedSecretStoreHealthsServer struct {
}

func (UnimplementedSecretStoreHealthsServer) List(context.Context, *SecretStoreHealthListRequest) (*SecretStoreHealthListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSecretStoreHealthsServer) Check(context.Context, *SecretStoreHealthCheckRequest) (*SecretStoreHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedSecretStoreHealthsServer) mustEmbedUnimplementedSecretStoreHealthsServer() {}

// UnsafeSecretStoreHealthsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretStoreHealthsServer will
// result in compilation errors.
type UnsafeSecretStoreHealthsServer interface {
	mustEmbedUnimplementedSecretStoreHealthsServer()
}

func RegisterSecretStoreHealthsServer(s grpc.ServiceRegistrar, srv SecretStoreHealthsServer) {
	s.RegisterService(&_SecretStoreHealths_serviceDesc, srv)
}

func _SecretStoreHealths_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretStoreHealthListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretStoreHealthsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SecretStoreHealths/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretStoreHealthsServer).List(ctx, req.(*SecretStoreHealthListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretStoreHealths_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretStoreHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretStoreHealthsServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SecretStoreHealths/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretStoreHealthsServer).Check(ctx, req.(*SecretStoreHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretStoreHealths_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SecretStoreHealths",
	HandlerType: (*SecretStoreHealthsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _SecretStoreHealths_List_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _SecretStoreHealths_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret_store_healths.proto",
}
