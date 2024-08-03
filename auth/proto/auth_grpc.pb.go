// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: auth.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_SendOtp_FullMethodName            = "/proto.AuthService/SendOtp"
	AuthService_VerifyOtp_FullMethodName          = "/proto.AuthService/VerifyOtp"
	AuthService_IssueSessionToken_FullMethodName  = "/proto.AuthService/IssueSessionToken"
	AuthService_RevokeSessionToken_FullMethodName = "/proto.AuthService/RevokeSessionToken"
	AuthService_VerifyChallenge_FullMethodName    = "/proto.AuthService/VerifyChallenge"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The AuthService definition for managing authentication
type AuthServiceClient interface {
	// Sends an OTP to the user
	SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error)
	// Verifies the OTP provided by the user
	VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...grpc.CallOption) (*VerifyOtpResponse, error)
	// Issues a session token upon successful authentication
	IssueSessionToken(ctx context.Context, in *IssueSessionTokenRequest, opts ...grpc.CallOption) (*IssueSessionTokenResponse, error)
	// Revokes a session token
	RevokeSessionToken(ctx context.Context, in *RevokeSessionTokenRequest, opts ...grpc.CallOption) (*RevokeSessionTokenResponse, error)
	// Verifies a signed challenge from the identity service
	VerifyChallenge(ctx context.Context, in *VerifyChallengeRequest, opts ...grpc.CallOption) (*VerifyChallengeResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendOtpResponse)
	err := c.cc.Invoke(ctx, AuthService_SendOtp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...grpc.CallOption) (*VerifyOtpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyOtpResponse)
	err := c.cc.Invoke(ctx, AuthService_VerifyOtp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IssueSessionToken(ctx context.Context, in *IssueSessionTokenRequest, opts ...grpc.CallOption) (*IssueSessionTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IssueSessionTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_IssueSessionToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RevokeSessionToken(ctx context.Context, in *RevokeSessionTokenRequest, opts ...grpc.CallOption) (*RevokeSessionTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeSessionTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_RevokeSessionToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyChallenge(ctx context.Context, in *VerifyChallengeRequest, opts ...grpc.CallOption) (*VerifyChallengeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyChallengeResponse)
	err := c.cc.Invoke(ctx, AuthService_VerifyChallenge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
//
// The AuthService definition for managing authentication
type AuthServiceServer interface {
	// Sends an OTP to the user
	SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error)
	// Verifies the OTP provided by the user
	VerifyOtp(context.Context, *VerifyOtpRequest) (*VerifyOtpResponse, error)
	// Issues a session token upon successful authentication
	IssueSessionToken(context.Context, *IssueSessionTokenRequest) (*IssueSessionTokenResponse, error)
	// Revokes a session token
	RevokeSessionToken(context.Context, *RevokeSessionTokenRequest) (*RevokeSessionTokenResponse, error)
	// Verifies a signed challenge from the identity service
	VerifyChallenge(context.Context, *VerifyChallengeRequest) (*VerifyChallengeResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOtp not implemented")
}
func (UnimplementedAuthServiceServer) VerifyOtp(context.Context, *VerifyOtpRequest) (*VerifyOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOtp not implemented")
}
func (UnimplementedAuthServiceServer) IssueSessionToken(context.Context, *IssueSessionTokenRequest) (*IssueSessionTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueSessionToken not implemented")
}
func (UnimplementedAuthServiceServer) RevokeSessionToken(context.Context, *RevokeSessionTokenRequest) (*RevokeSessionTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSessionToken not implemented")
}
func (UnimplementedAuthServiceServer) VerifyChallenge(context.Context, *VerifyChallengeRequest) (*VerifyChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyChallenge not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_SendOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SendOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SendOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SendOtp(ctx, req.(*SendOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyOtp(ctx, req.(*VerifyOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IssueSessionToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueSessionTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IssueSessionToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_IssueSessionToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IssueSessionToken(ctx, req.(*IssueSessionTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RevokeSessionToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSessionTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RevokeSessionToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RevokeSessionToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RevokeSessionToken(ctx, req.(*RevokeSessionTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyChallenge(ctx, req.(*VerifyChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOtp",
			Handler:    _AuthService_SendOtp_Handler,
		},
		{
			MethodName: "VerifyOtp",
			Handler:    _AuthService_VerifyOtp_Handler,
		},
		{
			MethodName: "IssueSessionToken",
			Handler:    _AuthService_IssueSessionToken_Handler,
		},
		{
			MethodName: "RevokeSessionToken",
			Handler:    _AuthService_RevokeSessionToken_Handler,
		},
		{
			MethodName: "VerifyChallenge",
			Handler:    _AuthService_VerifyChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
