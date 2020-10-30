// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package darkbase

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DarkbaseServiceClient is the client API for DarkbaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DarkbaseServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	BatchGet(ctx context.Context, in *BatchGetRequest, opts ...grpc.CallOption) (*BatchGetResponse, error)
	BatchPut(ctx context.Context, in *BatchPutRequest, opts ...grpc.CallOption) (*BatchPutRequest, error)
}

type darkbaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDarkbaseServiceClient(cc grpc.ClientConnInterface) DarkbaseServiceClient {
	return &darkbaseServiceClient{cc}
}

func (c *darkbaseServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/DarkbaseService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darkbaseServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/DarkbaseService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darkbaseServiceClient) BatchGet(ctx context.Context, in *BatchGetRequest, opts ...grpc.CallOption) (*BatchGetResponse, error) {
	out := new(BatchGetResponse)
	err := c.cc.Invoke(ctx, "/DarkbaseService/BatchGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darkbaseServiceClient) BatchPut(ctx context.Context, in *BatchPutRequest, opts ...grpc.CallOption) (*BatchPutRequest, error) {
	out := new(BatchPutRequest)
	err := c.cc.Invoke(ctx, "/DarkbaseService/BatchPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DarkbaseServiceServer is the server API for DarkbaseService service.
// All implementations must embed UnimplementedDarkbaseServiceServer
// for forward compatibility
type DarkbaseServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	BatchGet(context.Context, *BatchGetRequest) (*BatchGetResponse, error)
	BatchPut(context.Context, *BatchPutRequest) (*BatchPutRequest, error)
	mustEmbedUnimplementedDarkbaseServiceServer()
}

// UnimplementedDarkbaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDarkbaseServiceServer struct {
}

func (UnimplementedDarkbaseServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDarkbaseServiceServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedDarkbaseServiceServer) BatchGet(context.Context, *BatchGetRequest) (*BatchGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedDarkbaseServiceServer) BatchPut(context.Context, *BatchPutRequest) (*BatchPutRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchPut not implemented")
}
func (UnimplementedDarkbaseServiceServer) mustEmbedUnimplementedDarkbaseServiceServer() {}

// UnsafeDarkbaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DarkbaseServiceServer will
// result in compilation errors.
type UnsafeDarkbaseServiceServer interface {
	mustEmbedUnimplementedDarkbaseServiceServer()
}

func RegisterDarkbaseServiceServer(s *grpc.Server, srv DarkbaseServiceServer) {
	s.RegisterService(&_DarkbaseService_serviceDesc, srv)
}

func _DarkbaseService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarkbaseServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DarkbaseService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarkbaseServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DarkbaseService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarkbaseServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DarkbaseService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarkbaseServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DarkbaseService_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarkbaseServiceServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DarkbaseService/BatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarkbaseServiceServer).BatchGet(ctx, req.(*BatchGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DarkbaseService_BatchPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarkbaseServiceServer).BatchPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DarkbaseService/BatchPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarkbaseServiceServer).BatchPut(ctx, req.(*BatchPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DarkbaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DarkbaseService",
	HandlerType: (*DarkbaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DarkbaseService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _DarkbaseService_Put_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _DarkbaseService_BatchGet_Handler,
		},
		{
			MethodName: "BatchPut",
			Handler:    _DarkbaseService_BatchPut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "darkbase/darkbase.proto",
}
