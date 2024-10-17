// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: chat.proto

package __

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (ChatService_ConnectClient, error)
	FetchHistory(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*ChatHistory, error)
	AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	ReplyToComment(ctx context.Context, in *ReplyRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	FetchComments(ctx context.Context, in *FetchCommentsRequest, opts ...grpc.CallOption) (*FetchCommentsResponse, error)
	FetchUserComments(ctx context.Context, in *FetchUserCommentsRequest, opts ...grpc.CallOption) (*FetchUserCommentsResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (ChatService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/pb.ChatService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceConnectClient{stream}
	return x, nil
}

type ChatService_ConnectClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceConnectClient struct {
	grpc.ClientStream
}

func (x *chatServiceConnectClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) FetchHistory(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*ChatHistory, error) {
	out := new(ChatHistory)
	err := c.cc.Invoke(ctx, "/pb.ChatService/FetchHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/pb.ChatService/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ReplyToComment(ctx context.Context, in *ReplyRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/pb.ChatService/ReplyToComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) FetchComments(ctx context.Context, in *FetchCommentsRequest, opts ...grpc.CallOption) (*FetchCommentsResponse, error) {
	out := new(FetchCommentsResponse)
	err := c.cc.Invoke(ctx, "/pb.ChatService/FetchComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) FetchUserComments(ctx context.Context, in *FetchUserCommentsRequest, opts ...grpc.CallOption) (*FetchUserCommentsResponse, error) {
	out := new(FetchUserCommentsResponse)
	err := c.cc.Invoke(ctx, "/pb.ChatService/FetchUserComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Connect(ChatService_ConnectServer) error
	FetchHistory(context.Context, *ChatID) (*ChatHistory, error)
	AddComment(context.Context, *CommentRequest) (*CommentResponse, error)
	ReplyToComment(context.Context, *ReplyRequest) (*CommentResponse, error)
	FetchComments(context.Context, *FetchCommentsRequest) (*FetchCommentsResponse, error)
	FetchUserComments(context.Context, *FetchUserCommentsRequest) (*FetchUserCommentsResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Connect(ChatService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatServiceServer) FetchHistory(context.Context, *ChatID) (*ChatHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHistory not implemented")
}
func (UnimplementedChatServiceServer) AddComment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedChatServiceServer) ReplyToComment(context.Context, *ReplyRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyToComment not implemented")
}
func (UnimplementedChatServiceServer) FetchComments(context.Context, *FetchCommentsRequest) (*FetchCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchComments not implemented")
}
func (UnimplementedChatServiceServer) FetchUserComments(context.Context, *FetchUserCommentsRequest) (*FetchUserCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUserComments not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Connect(&chatServiceConnectServer{stream})
}

type ChatService_ConnectServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceConnectServer struct {
	grpc.ServerStream
}

func (x *chatServiceConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceConnectServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_FetchHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).FetchHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChatService/FetchHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).FetchHistory(ctx, req.(*ChatID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChatService/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ReplyToComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ReplyToComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChatService/ReplyToComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ReplyToComment(ctx, req.(*ReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_FetchComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).FetchComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChatService/FetchComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).FetchComments(ctx, req.(*FetchCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_FetchUserComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchUserCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).FetchUserComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChatService/FetchUserComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).FetchUserComments(ctx, req.(*FetchUserCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHistory",
			Handler:    _ChatService_FetchHistory_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _ChatService_AddComment_Handler,
		},
		{
			MethodName: "ReplyToComment",
			Handler:    _ChatService_ReplyToComment_Handler,
		},
		{
			MethodName: "FetchComments",
			Handler:    _ChatService_FetchComments_Handler,
		},
		{
			MethodName: "FetchUserComments",
			Handler:    _ChatService_FetchUserComments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
