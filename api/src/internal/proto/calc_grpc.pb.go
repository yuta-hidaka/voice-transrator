// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: calc.proto

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

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServiceClient interface {
	Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	CalcServerStream(ctx context.Context, in *CalcCalcServerStreamRequest, opts ...grpc.CallOption) (CalcService_CalcServerStreamClient, error)
	CalcClientStream(ctx context.Context, opts ...grpc.CallOption) (CalcService_CalcClientStreamClient, error)
	CalcBiDirectionStream(ctx context.Context, opts ...grpc.CallOption) (CalcService_CalcBiDirectionStreamClient, error)
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/translator.CalcService/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) CalcServerStream(ctx context.Context, in *CalcCalcServerStreamRequest, opts ...grpc.CallOption) (CalcService_CalcServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[0], "/translator.CalcService/CalcServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceCalcServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_CalcServerStreamClient interface {
	Recv() (*CalcCalcServerStreamResponse, error)
	grpc.ClientStream
}

type calcServiceCalcServerStreamClient struct {
	grpc.ClientStream
}

func (x *calcServiceCalcServerStreamClient) Recv() (*CalcCalcServerStreamResponse, error) {
	m := new(CalcCalcServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) CalcClientStream(ctx context.Context, opts ...grpc.CallOption) (CalcService_CalcClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[1], "/translator.CalcService/CalcClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceCalcClientStreamClient{stream}
	return x, nil
}

type CalcService_CalcClientStreamClient interface {
	Send(*CalcCalcServerStreamRequest) error
	CloseAndRecv() (*CalcCalcClientStreamResponse, error)
	grpc.ClientStream
}

type calcServiceCalcClientStreamClient struct {
	grpc.ClientStream
}

func (x *calcServiceCalcClientStreamClient) Send(m *CalcCalcServerStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceCalcClientStreamClient) CloseAndRecv() (*CalcCalcClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalcCalcClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) CalcBiDirectionStream(ctx context.Context, opts ...grpc.CallOption) (CalcService_CalcBiDirectionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[2], "/translator.CalcService/CalcBiDirectionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceCalcBiDirectionStreamClient{stream}
	return x, nil
}

type CalcService_CalcBiDirectionStreamClient interface {
	Send(*CalcCalcServerStreamRequest) error
	Recv() (*CalcCalcServerStreamResponse, error)
	grpc.ClientStream
}

type calcServiceCalcBiDirectionStreamClient struct {
	grpc.ClientStream
}

func (x *calcServiceCalcBiDirectionStreamClient) Send(m *CalcCalcServerStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceCalcBiDirectionStreamClient) Recv() (*CalcCalcServerStreamResponse, error) {
	m := new(CalcCalcServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/translator.CalcService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
// All implementations must embed UnimplementedCalcServiceServer
// for forward compatibility
type CalcServiceServer interface {
	Calc(context.Context, *CalcRequest) (*CalcResponse, error)
	CalcServerStream(*CalcCalcServerStreamRequest, CalcService_CalcServerStreamServer) error
	CalcClientStream(CalcService_CalcClientStreamServer) error
	CalcBiDirectionStream(CalcService_CalcBiDirectionStreamServer) error
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
	mustEmbedUnimplementedCalcServiceServer()
}

// UnimplementedCalcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (UnimplementedCalcServiceServer) Calc(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calc not implemented")
}
func (UnimplementedCalcServiceServer) CalcServerStream(*CalcCalcServerStreamRequest, CalcService_CalcServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CalcServerStream not implemented")
}
func (UnimplementedCalcServiceServer) CalcClientStream(CalcService_CalcClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CalcClientStream not implemented")
}
func (UnimplementedCalcServiceServer) CalcBiDirectionStream(CalcService_CalcBiDirectionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CalcBiDirectionStream not implemented")
}
func (UnimplementedCalcServiceServer) Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedCalcServiceServer) mustEmbedUnimplementedCalcServiceServer() {}

// UnsafeCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServiceServer will
// result in compilation errors.
type UnsafeCalcServiceServer interface {
	mustEmbedUnimplementedCalcServiceServer()
}

func RegisterCalcServiceServer(s grpc.ServiceRegistrar, srv CalcServiceServer) {
	s.RegisterService(&CalcService_ServiceDesc, srv)
}

func _CalcService_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.CalcService/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Calc(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_CalcServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalcCalcServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).CalcServerStream(m, &calcServiceCalcServerStreamServer{stream})
}

type CalcService_CalcServerStreamServer interface {
	Send(*CalcCalcServerStreamResponse) error
	grpc.ServerStream
}

type calcServiceCalcServerStreamServer struct {
	grpc.ServerStream
}

func (x *calcServiceCalcServerStreamServer) Send(m *CalcCalcServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcService_CalcClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).CalcClientStream(&calcServiceCalcClientStreamServer{stream})
}

type CalcService_CalcClientStreamServer interface {
	SendAndClose(*CalcCalcClientStreamResponse) error
	Recv() (*CalcCalcServerStreamRequest, error)
	grpc.ServerStream
}

type calcServiceCalcClientStreamServer struct {
	grpc.ServerStream
}

func (x *calcServiceCalcClientStreamServer) SendAndClose(m *CalcCalcClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceCalcClientStreamServer) Recv() (*CalcCalcServerStreamRequest, error) {
	m := new(CalcCalcServerStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_CalcBiDirectionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).CalcBiDirectionStream(&calcServiceCalcBiDirectionStreamServer{stream})
}

type CalcService_CalcBiDirectionStreamServer interface {
	Send(*CalcCalcServerStreamResponse) error
	Recv() (*CalcCalcServerStreamRequest, error)
	grpc.ServerStream
}

type calcServiceCalcBiDirectionStreamServer struct {
	grpc.ServerStream
}

func (x *calcServiceCalcBiDirectionStreamServer) Send(m *CalcCalcServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceCalcBiDirectionStreamServer) Recv() (*CalcCalcServerStreamRequest, error) {
	m := new(CalcCalcServerStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.CalcService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalcService_ServiceDesc is the grpc.ServiceDesc for CalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "translator.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _CalcService_Calc_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _CalcService_Sqrt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalcServerStream",
			Handler:       _CalcService_CalcServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CalcClientStream",
			Handler:       _CalcService_CalcClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CalcBiDirectionStream",
			Handler:       _CalcService_CalcBiDirectionStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calc.proto",
}