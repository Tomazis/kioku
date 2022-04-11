// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package srv_dba

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

// SrvDbaServiceClient is the client API for SrvDbaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvDbaServiceClient interface {
	//Kanji part
	GetKanjiByIdV1(ctx context.Context, in *GetKanjiByIdV1Request, opts ...grpc.CallOption) (*GetKanjiByIdV1Response, error)
	ListKanjiByLevelV1(ctx context.Context, in *ListKanjiByLevelV1Request, opts ...grpc.CallOption) (*ListKanjiV1Response, error)
	ListKanjiByIdsV1(ctx context.Context, in *ListKanjiByIdsV1Request, opts ...grpc.CallOption) (*ListKanjiV1Response, error)
	//Words part
	GetWordByIdV1(ctx context.Context, in *GetWordByIdV1Request, opts ...grpc.CallOption) (*GetWordByIdV1Response, error)
	ListWordsByLevelV1(ctx context.Context, in *ListWordsByLevelV1Request, opts ...grpc.CallOption) (*ListWordsV1Response, error)
	ListWordsByKanjiV1(ctx context.Context, in *ListWordsByKanjiV1Request, opts ...grpc.CallOption) (*ListWordsV1Response, error)
	ListWordsByIdsV1(ctx context.Context, in *ListWordsByIdsV1Request, opts ...grpc.CallOption) (*ListWordsV1Response, error)
	//KanjiProgress part
	GetKanjiProgressByIdV1(ctx context.Context, in *GetKanjiProgressByIdV1Request, opts ...grpc.CallOption) (*GetKanjiProgressByIdV1Response, error)
	ListKanjiProgressByTimeV1(ctx context.Context, in *ListKanjiProgressByTimeV1Request, opts ...grpc.CallOption) (*ListKanjiProgressV1Response, error)
	ListKanjiProgressByIdsV1(ctx context.Context, in *ListKanjiProgressByIdsV1Request, opts ...grpc.CallOption) (*ListKanjiProgressV1Response, error)
	ListKanjiProgressBySrsLevelV1(ctx context.Context, in *ListKanjiProgressBySrsLevelV1Request, opts ...grpc.CallOption) (*ListKanjiProgressV1Response, error)
	//WordProgress part
	GetWordProgressByIdV1(ctx context.Context, in *GetWordProgressByIdV1Request, opts ...grpc.CallOption) (*GetWordProgressByIdV1Response, error)
	ListWordProgressByTimeV1(ctx context.Context, in *ListWordProgressByTimeV1Request, opts ...grpc.CallOption) (*ListWordProgressV1Response, error)
	ListWordProgressByIdsV1(ctx context.Context, in *ListWordProgressByIdsV1Request, opts ...grpc.CallOption) (*ListWordProgressV1Response, error)
	ListWordProgressBySrsLevelV1(ctx context.Context, in *ListWordProgressBySrsLevelV1Request, opts ...grpc.CallOption) (*ListWordProgressV1Response, error)
}

type srvDbaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvDbaServiceClient(cc grpc.ClientConnInterface) SrvDbaServiceClient {
	return &srvDbaServiceClient{cc}
}

func (c *srvDbaServiceClient) GetKanjiByIdV1(ctx context.Context, in *GetKanjiByIdV1Request, opts ...grpc.CallOption) (*GetKanjiByIdV1Response, error) {
	out := new(GetKanjiByIdV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/GetKanjiByIdV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListKanjiByLevelV1(ctx context.Context, in *ListKanjiByLevelV1Request, opts ...grpc.CallOption) (*ListKanjiV1Response, error) {
	out := new(ListKanjiV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiByLevelV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListKanjiByIdsV1(ctx context.Context, in *ListKanjiByIdsV1Request, opts ...grpc.CallOption) (*ListKanjiV1Response, error) {
	out := new(ListKanjiV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiByIdsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) GetWordByIdV1(ctx context.Context, in *GetWordByIdV1Request, opts ...grpc.CallOption) (*GetWordByIdV1Response, error) {
	out := new(GetWordByIdV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/GetWordByIdV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListWordsByLevelV1(ctx context.Context, in *ListWordsByLevelV1Request, opts ...grpc.CallOption) (*ListWordsV1Response, error) {
	out := new(ListWordsV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListWordsByLevelV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListWordsByKanjiV1(ctx context.Context, in *ListWordsByKanjiV1Request, opts ...grpc.CallOption) (*ListWordsV1Response, error) {
	out := new(ListWordsV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListWordsByKanjiV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListWordsByIdsV1(ctx context.Context, in *ListWordsByIdsV1Request, opts ...grpc.CallOption) (*ListWordsV1Response, error) {
	out := new(ListWordsV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListWordsByIdsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) GetKanjiProgressByIdV1(ctx context.Context, in *GetKanjiProgressByIdV1Request, opts ...grpc.CallOption) (*GetKanjiProgressByIdV1Response, error) {
	out := new(GetKanjiProgressByIdV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/GetKanjiProgressByIdV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListKanjiProgressByTimeV1(ctx context.Context, in *ListKanjiProgressByTimeV1Request, opts ...grpc.CallOption) (*ListKanjiProgressV1Response, error) {
	out := new(ListKanjiProgressV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiProgressByTimeV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListKanjiProgressByIdsV1(ctx context.Context, in *ListKanjiProgressByIdsV1Request, opts ...grpc.CallOption) (*ListKanjiProgressV1Response, error) {
	out := new(ListKanjiProgressV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiProgressByIdsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListKanjiProgressBySrsLevelV1(ctx context.Context, in *ListKanjiProgressBySrsLevelV1Request, opts ...grpc.CallOption) (*ListKanjiProgressV1Response, error) {
	out := new(ListKanjiProgressV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiProgressBySrsLevelV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) GetWordProgressByIdV1(ctx context.Context, in *GetWordProgressByIdV1Request, opts ...grpc.CallOption) (*GetWordProgressByIdV1Response, error) {
	out := new(GetWordProgressByIdV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/GetWordProgressByIdV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListWordProgressByTimeV1(ctx context.Context, in *ListWordProgressByTimeV1Request, opts ...grpc.CallOption) (*ListWordProgressV1Response, error) {
	out := new(ListWordProgressV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListWordProgressByTimeV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListWordProgressByIdsV1(ctx context.Context, in *ListWordProgressByIdsV1Request, opts ...grpc.CallOption) (*ListWordProgressV1Response, error) {
	out := new(ListWordProgressV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListWordProgressByIdsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvDbaServiceClient) ListWordProgressBySrsLevelV1(ctx context.Context, in *ListWordProgressBySrsLevelV1Request, opts ...grpc.CallOption) (*ListWordProgressV1Response, error) {
	out := new(ListWordProgressV1Response)
	err := c.cc.Invoke(ctx, "/kioku.server.srv_dba.v1.SrvDbaService/ListWordProgressBySrsLevelV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvDbaServiceServer is the server API for SrvDbaService service.
// All implementations must embed UnimplementedSrvDbaServiceServer
// for forward compatibility
type SrvDbaServiceServer interface {
	//Kanji part
	GetKanjiByIdV1(context.Context, *GetKanjiByIdV1Request) (*GetKanjiByIdV1Response, error)
	ListKanjiByLevelV1(context.Context, *ListKanjiByLevelV1Request) (*ListKanjiV1Response, error)
	ListKanjiByIdsV1(context.Context, *ListKanjiByIdsV1Request) (*ListKanjiV1Response, error)
	//Words part
	GetWordByIdV1(context.Context, *GetWordByIdV1Request) (*GetWordByIdV1Response, error)
	ListWordsByLevelV1(context.Context, *ListWordsByLevelV1Request) (*ListWordsV1Response, error)
	ListWordsByKanjiV1(context.Context, *ListWordsByKanjiV1Request) (*ListWordsV1Response, error)
	ListWordsByIdsV1(context.Context, *ListWordsByIdsV1Request) (*ListWordsV1Response, error)
	//KanjiProgress part
	GetKanjiProgressByIdV1(context.Context, *GetKanjiProgressByIdV1Request) (*GetKanjiProgressByIdV1Response, error)
	ListKanjiProgressByTimeV1(context.Context, *ListKanjiProgressByTimeV1Request) (*ListKanjiProgressV1Response, error)
	ListKanjiProgressByIdsV1(context.Context, *ListKanjiProgressByIdsV1Request) (*ListKanjiProgressV1Response, error)
	ListKanjiProgressBySrsLevelV1(context.Context, *ListKanjiProgressBySrsLevelV1Request) (*ListKanjiProgressV1Response, error)
	//WordProgress part
	GetWordProgressByIdV1(context.Context, *GetWordProgressByIdV1Request) (*GetWordProgressByIdV1Response, error)
	ListWordProgressByTimeV1(context.Context, *ListWordProgressByTimeV1Request) (*ListWordProgressV1Response, error)
	ListWordProgressByIdsV1(context.Context, *ListWordProgressByIdsV1Request) (*ListWordProgressV1Response, error)
	ListWordProgressBySrsLevelV1(context.Context, *ListWordProgressBySrsLevelV1Request) (*ListWordProgressV1Response, error)
	mustEmbedUnimplementedSrvDbaServiceServer()
}

// UnimplementedSrvDbaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSrvDbaServiceServer struct {
}

func (UnimplementedSrvDbaServiceServer) GetKanjiByIdV1(context.Context, *GetKanjiByIdV1Request) (*GetKanjiByIdV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKanjiByIdV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListKanjiByLevelV1(context.Context, *ListKanjiByLevelV1Request) (*ListKanjiV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKanjiByLevelV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListKanjiByIdsV1(context.Context, *ListKanjiByIdsV1Request) (*ListKanjiV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKanjiByIdsV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) GetWordByIdV1(context.Context, *GetWordByIdV1Request) (*GetWordByIdV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWordByIdV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListWordsByLevelV1(context.Context, *ListWordsByLevelV1Request) (*ListWordsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWordsByLevelV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListWordsByKanjiV1(context.Context, *ListWordsByKanjiV1Request) (*ListWordsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWordsByKanjiV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListWordsByIdsV1(context.Context, *ListWordsByIdsV1Request) (*ListWordsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWordsByIdsV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) GetKanjiProgressByIdV1(context.Context, *GetKanjiProgressByIdV1Request) (*GetKanjiProgressByIdV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKanjiProgressByIdV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListKanjiProgressByTimeV1(context.Context, *ListKanjiProgressByTimeV1Request) (*ListKanjiProgressV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKanjiProgressByTimeV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListKanjiProgressByIdsV1(context.Context, *ListKanjiProgressByIdsV1Request) (*ListKanjiProgressV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKanjiProgressByIdsV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListKanjiProgressBySrsLevelV1(context.Context, *ListKanjiProgressBySrsLevelV1Request) (*ListKanjiProgressV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKanjiProgressBySrsLevelV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) GetWordProgressByIdV1(context.Context, *GetWordProgressByIdV1Request) (*GetWordProgressByIdV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWordProgressByIdV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListWordProgressByTimeV1(context.Context, *ListWordProgressByTimeV1Request) (*ListWordProgressV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWordProgressByTimeV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListWordProgressByIdsV1(context.Context, *ListWordProgressByIdsV1Request) (*ListWordProgressV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWordProgressByIdsV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) ListWordProgressBySrsLevelV1(context.Context, *ListWordProgressBySrsLevelV1Request) (*ListWordProgressV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWordProgressBySrsLevelV1 not implemented")
}
func (UnimplementedSrvDbaServiceServer) mustEmbedUnimplementedSrvDbaServiceServer() {}

// UnsafeSrvDbaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvDbaServiceServer will
// result in compilation errors.
type UnsafeSrvDbaServiceServer interface {
	mustEmbedUnimplementedSrvDbaServiceServer()
}

func RegisterSrvDbaServiceServer(s grpc.ServiceRegistrar, srv SrvDbaServiceServer) {
	s.RegisterService(&SrvDbaService_ServiceDesc, srv)
}

func _SrvDbaService_GetKanjiByIdV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKanjiByIdV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).GetKanjiByIdV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/GetKanjiByIdV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).GetKanjiByIdV1(ctx, req.(*GetKanjiByIdV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListKanjiByLevelV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKanjiByLevelV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListKanjiByLevelV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiByLevelV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListKanjiByLevelV1(ctx, req.(*ListKanjiByLevelV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListKanjiByIdsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKanjiByIdsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListKanjiByIdsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiByIdsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListKanjiByIdsV1(ctx, req.(*ListKanjiByIdsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_GetWordByIdV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWordByIdV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).GetWordByIdV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/GetWordByIdV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).GetWordByIdV1(ctx, req.(*GetWordByIdV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListWordsByLevelV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordsByLevelV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListWordsByLevelV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListWordsByLevelV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListWordsByLevelV1(ctx, req.(*ListWordsByLevelV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListWordsByKanjiV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordsByKanjiV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListWordsByKanjiV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListWordsByKanjiV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListWordsByKanjiV1(ctx, req.(*ListWordsByKanjiV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListWordsByIdsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordsByIdsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListWordsByIdsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListWordsByIdsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListWordsByIdsV1(ctx, req.(*ListWordsByIdsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_GetKanjiProgressByIdV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKanjiProgressByIdV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).GetKanjiProgressByIdV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/GetKanjiProgressByIdV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).GetKanjiProgressByIdV1(ctx, req.(*GetKanjiProgressByIdV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListKanjiProgressByTimeV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKanjiProgressByTimeV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListKanjiProgressByTimeV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiProgressByTimeV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListKanjiProgressByTimeV1(ctx, req.(*ListKanjiProgressByTimeV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListKanjiProgressByIdsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKanjiProgressByIdsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListKanjiProgressByIdsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiProgressByIdsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListKanjiProgressByIdsV1(ctx, req.(*ListKanjiProgressByIdsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListKanjiProgressBySrsLevelV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKanjiProgressBySrsLevelV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListKanjiProgressBySrsLevelV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListKanjiProgressBySrsLevelV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListKanjiProgressBySrsLevelV1(ctx, req.(*ListKanjiProgressBySrsLevelV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_GetWordProgressByIdV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWordProgressByIdV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).GetWordProgressByIdV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/GetWordProgressByIdV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).GetWordProgressByIdV1(ctx, req.(*GetWordProgressByIdV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListWordProgressByTimeV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordProgressByTimeV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListWordProgressByTimeV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListWordProgressByTimeV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListWordProgressByTimeV1(ctx, req.(*ListWordProgressByTimeV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListWordProgressByIdsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordProgressByIdsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListWordProgressByIdsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListWordProgressByIdsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListWordProgressByIdsV1(ctx, req.(*ListWordProgressByIdsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvDbaService_ListWordProgressBySrsLevelV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordProgressBySrsLevelV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvDbaServiceServer).ListWordProgressBySrsLevelV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kioku.server.srv_dba.v1.SrvDbaService/ListWordProgressBySrsLevelV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvDbaServiceServer).ListWordProgressBySrsLevelV1(ctx, req.(*ListWordProgressBySrsLevelV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvDbaService_ServiceDesc is the grpc.ServiceDesc for SrvDbaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvDbaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kioku.server.srv_dba.v1.SrvDbaService",
	HandlerType: (*SrvDbaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKanjiByIdV1",
			Handler:    _SrvDbaService_GetKanjiByIdV1_Handler,
		},
		{
			MethodName: "ListKanjiByLevelV1",
			Handler:    _SrvDbaService_ListKanjiByLevelV1_Handler,
		},
		{
			MethodName: "ListKanjiByIdsV1",
			Handler:    _SrvDbaService_ListKanjiByIdsV1_Handler,
		},
		{
			MethodName: "GetWordByIdV1",
			Handler:    _SrvDbaService_GetWordByIdV1_Handler,
		},
		{
			MethodName: "ListWordsByLevelV1",
			Handler:    _SrvDbaService_ListWordsByLevelV1_Handler,
		},
		{
			MethodName: "ListWordsByKanjiV1",
			Handler:    _SrvDbaService_ListWordsByKanjiV1_Handler,
		},
		{
			MethodName: "ListWordsByIdsV1",
			Handler:    _SrvDbaService_ListWordsByIdsV1_Handler,
		},
		{
			MethodName: "GetKanjiProgressByIdV1",
			Handler:    _SrvDbaService_GetKanjiProgressByIdV1_Handler,
		},
		{
			MethodName: "ListKanjiProgressByTimeV1",
			Handler:    _SrvDbaService_ListKanjiProgressByTimeV1_Handler,
		},
		{
			MethodName: "ListKanjiProgressByIdsV1",
			Handler:    _SrvDbaService_ListKanjiProgressByIdsV1_Handler,
		},
		{
			MethodName: "ListKanjiProgressBySrsLevelV1",
			Handler:    _SrvDbaService_ListKanjiProgressBySrsLevelV1_Handler,
		},
		{
			MethodName: "GetWordProgressByIdV1",
			Handler:    _SrvDbaService_GetWordProgressByIdV1_Handler,
		},
		{
			MethodName: "ListWordProgressByTimeV1",
			Handler:    _SrvDbaService_ListWordProgressByTimeV1_Handler,
		},
		{
			MethodName: "ListWordProgressByIdsV1",
			Handler:    _SrvDbaService_ListWordProgressByIdsV1_Handler,
		},
		{
			MethodName: "ListWordProgressBySrsLevelV1",
			Handler:    _SrvDbaService_ListWordProgressBySrsLevelV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kioku/srv-dba/v1/srv-dba.proto",
}
