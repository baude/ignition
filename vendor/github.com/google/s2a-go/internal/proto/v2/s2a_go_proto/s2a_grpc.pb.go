// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/proto/v2/s2a.proto

package s2a_go_proto

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
	S2AService_SetUpSession_FullMethodName = "/s2a.proto.v2.S2AService/SetUpSession"
)

// S2AServiceClient is the client API for S2AService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S2AServiceClient interface {
	// SetUpSession is a bidirectional stream used by applications to offload
	// operations from the TLS handshake.
	SetUpSession(ctx context.Context, opts ...grpc.CallOption) (S2AService_SetUpSessionClient, error)
}

type s2AServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewS2AServiceClient(cc grpc.ClientConnInterface) S2AServiceClient {
	return &s2AServiceClient{cc}
}

func (c *s2AServiceClient) SetUpSession(ctx context.Context, opts ...grpc.CallOption) (S2AService_SetUpSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &S2AService_ServiceDesc.Streams[0], S2AService_SetUpSession_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &s2AServiceSetUpSessionClient{stream}
	return x, nil
}

type S2AService_SetUpSessionClient interface {
	Send(*SessionReq) error
	Recv() (*SessionResp, error)
	grpc.ClientStream
}

type s2AServiceSetUpSessionClient struct {
	grpc.ClientStream
}

func (x *s2AServiceSetUpSessionClient) Send(m *SessionReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *s2AServiceSetUpSessionClient) Recv() (*SessionResp, error) {
	m := new(SessionResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// S2AServiceServer is the server API for S2AService service.
// All implementations must embed UnimplementedS2AServiceServer
// for forward compatibility
type S2AServiceServer interface {
	// SetUpSession is a bidirectional stream used by applications to offload
	// operations from the TLS handshake.
	SetUpSession(S2AService_SetUpSessionServer) error
	mustEmbedUnimplementedS2AServiceServer()
}

// UnimplementedS2AServiceServer must be embedded to have forward compatible implementations.
type UnimplementedS2AServiceServer struct {
}

func (UnimplementedS2AServiceServer) SetUpSession(S2AService_SetUpSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method SetUpSession not implemented")
}
func (UnimplementedS2AServiceServer) mustEmbedUnimplementedS2AServiceServer() {}

// UnsafeS2AServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S2AServiceServer will
// result in compilation errors.
type UnsafeS2AServiceServer interface {
	mustEmbedUnimplementedS2AServiceServer()
}

func RegisterS2AServiceServer(s grpc.ServiceRegistrar, srv S2AServiceServer) {
	s.RegisterService(&S2AService_ServiceDesc, srv)
}

func _S2AService_SetUpSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(S2AServiceServer).SetUpSession(&s2AServiceSetUpSessionServer{stream})
}

type S2AService_SetUpSessionServer interface {
	Send(*SessionResp) error
	Recv() (*SessionReq, error)
	grpc.ServerStream
}

type s2AServiceSetUpSessionServer struct {
	grpc.ServerStream
}

func (x *s2AServiceSetUpSessionServer) Send(m *SessionResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *s2AServiceSetUpSessionServer) Recv() (*SessionReq, error) {
	m := new(SessionReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// S2AService_ServiceDesc is the grpc.ServiceDesc for S2AService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S2AService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "s2a.proto.v2.S2AService",
	HandlerType: (*S2AServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetUpSession",
			Handler:       _S2AService_SetUpSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/v2/s2a.proto",
}
