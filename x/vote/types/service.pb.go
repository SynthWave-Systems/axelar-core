// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/vote/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("axelar/vote/v1beta1/service.proto", fileDescriptor_030f863ebca64631) }
func init() {
	golang_proto.RegisterFile("axelar/vote/v1beta1/service.proto", fileDescriptor_030f863ebca64631)
}

var fileDescriptor_030f863ebca64631 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xb3, 0x22, 0x16, 0xe9, 0x8c, 0x62, 0x11, 0x8e, 0xc5, 0xbb, 0x52, 0x70, 0x87, 0x68,
	0x67, 0x69, 0x2f, 0x82, 0x82, 0x85, 0xdd, 0x26, 0x0c, 0xeb, 0xe2, 0xb9, 0x13, 0x77, 0xe7, 0x72,
	0xb1, 0xf5, 0x09, 0x04, 0x5f, 0xc8, 0xd2, 0xf2, 0xc0, 0xc6, 0x52, 0x12, 0x1f, 0x44, 0x72, 0x1b,
	0x0b, 0xe1, 0xb0, 0x9b, 0xe1, 0xfb, 0x66, 0xfe, 0x99, 0x74, 0xaa, 0x5b, 0x9c, 0x6b, 0x0f, 0x0d,
	0x31, 0x42, 0x53, 0x94, 0xc8, 0xba, 0x80, 0x80, 0xbe, 0xb1, 0x15, 0xaa, 0xda, 0x13, 0x53, 0xb6,
	0x17, 0x15, 0x35, 0x28, 0x6a, 0x54, 0xf2, 0x7d, 0x43, 0x86, 0xd6, 0x1c, 0x86, 0x2a, 0xaa, 0xf9,
	0xc4, 0x10, 0x99, 0x39, 0x82, 0xae, 0x2d, 0x68, 0xe7, 0x88, 0x35, 0x5b, 0x72, 0xe1, 0x97, 0x6e,
	0xca, 0xe2, 0x36, 0xd2, 0x93, 0x65, 0x9a, 0x5e, 0x04, 0x73, 0x1d, 0xa3, 0x33, 0x9b, 0x6e, 0xdf,
	0x10, 0x63, 0x76, 0xa8, 0x36, 0xa4, 0xab, 0x01, 0x5d, 0xe1, 0xe3, 0x02, 0x03, 0xe7, 0xd3, 0x7f,
	0x8c, 0x50, 0x93, 0x0b, 0x38, 0x9b, 0x3c, 0x7f, 0x7c, 0xbf, 0x6e, 0x1d, 0xcc, 0x76, 0xe1, 0xcf,
	0x05, 0xc4, 0x78, 0x26, 0x8e, 0xce, 0x2f, 0xdf, 0x3b, 0x29, 0x56, 0x9d, 0x14, 0x5f, 0x9d, 0x14,
	0x2f, 0xbd, 0x4c, 0xde, 0x7a, 0x29, 0x56, 0xbd, 0x4c, 0x3e, 0x7b, 0x99, 0xdc, 0x16, 0xc6, 0xf2,
	0xdd, 0xa2, 0x54, 0x15, 0x3d, 0x8c, 0xd3, 0x0e, 0x79, 0x49, 0xfe, 0x7e, 0xec, 0x8e, 0x2b, 0xf2,
	0x08, 0x6d, 0x5c, 0xc9, 0x4f, 0x35, 0x86, 0x72, 0x67, 0xfd, 0xd0, 0xe9, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x2d, 0xf7, 0x4b, 0x5c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, "/axelar.vote.v1beta1.MsgService/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) Vote(ctx context.Context, req *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.vote.v1beta1.MsgService/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.vote.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Vote",
			Handler:    _MsgService_Vote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/vote/v1beta1/service.proto",
}
