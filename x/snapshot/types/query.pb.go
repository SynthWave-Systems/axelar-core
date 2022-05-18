// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/snapshot/v1beta1/query.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryValidatorsResponse struct {
	Validators []*QueryValidatorsResponse_Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *QueryValidatorsResponse) Reset()         { *m = QueryValidatorsResponse{} }
func (m *QueryValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsResponse) ProtoMessage()    {}
func (*QueryValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_044ba28efe73d39e, []int{0}
}
func (m *QueryValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse.Merge(m, src)
}
func (m *QueryValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse proto.InternalMessageInfo

type QueryValidatorsResponse_TssIllegibilityInfo struct {
	Tombstoned            bool `protobuf:"varint,1,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	Jailed                bool `protobuf:"varint,2,opt,name=jailed,proto3" json:"jailed,omitempty"`
	MissedTooManyBlocks   bool `protobuf:"varint,3,opt,name=missed_too_many_blocks,json=missedTooManyBlocks,proto3" json:"missed_too_many_blocks,omitempty"`
	NoProxyRegistered     bool `protobuf:"varint,4,opt,name=no_proxy_registered,json=noProxyRegistered,proto3" json:"no_proxy_registered,omitempty"`
	TssSuspended          bool `protobuf:"varint,5,opt,name=tss_suspended,json=tssSuspended,proto3" json:"tss_suspended,omitempty"`
	ProxyInsuficientFunds bool `protobuf:"varint,6,opt,name=proxy_insuficient_funds,json=proxyInsuficientFunds,proto3" json:"proxy_insuficient_funds,omitempty"`
	StaleTssHeartbeat     bool `protobuf:"varint,7,opt,name=stale_tss_heartbeat,json=staleTssHeartbeat,proto3" json:"stale_tss_heartbeat,omitempty"`
}

func (m *QueryValidatorsResponse_TssIllegibilityInfo) Reset() {
	*m = QueryValidatorsResponse_TssIllegibilityInfo{}
}
func (m *QueryValidatorsResponse_TssIllegibilityInfo) String() string {
	return proto.CompactTextString(m)
}
func (*QueryValidatorsResponse_TssIllegibilityInfo) ProtoMessage() {}
func (*QueryValidatorsResponse_TssIllegibilityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_044ba28efe73d39e, []int{0, 0}
}
func (m *QueryValidatorsResponse_TssIllegibilityInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse_TssIllegibilityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse_TssIllegibilityInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse_TssIllegibilityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse_TssIllegibilityInfo.Merge(m, src)
}
func (m *QueryValidatorsResponse_TssIllegibilityInfo) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse_TssIllegibilityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse_TssIllegibilityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse_TssIllegibilityInfo proto.InternalMessageInfo

type QueryValidatorsResponse_Validator struct {
	OperatorAddress     string                                      `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	Moniker             string                                      `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	TssIllegibilityInfo QueryValidatorsResponse_TssIllegibilityInfo `protobuf:"bytes,3,opt,name=tss_illegibility_info,json=tssIllegibilityInfo,proto3" json:"tss_illegibility_info"`
}

func (m *QueryValidatorsResponse_Validator) Reset()         { *m = QueryValidatorsResponse_Validator{} }
func (m *QueryValidatorsResponse_Validator) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsResponse_Validator) ProtoMessage()    {}
func (*QueryValidatorsResponse_Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_044ba28efe73d39e, []int{0, 1}
}
func (m *QueryValidatorsResponse_Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse_Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse_Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse_Validator.Merge(m, src)
}
func (m *QueryValidatorsResponse_Validator) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse_Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse_Validator proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryValidatorsResponse)(nil), "axelar.snapshot.v1beta1.QueryValidatorsResponse")
	proto.RegisterType((*QueryValidatorsResponse_TssIllegibilityInfo)(nil), "axelar.snapshot.v1beta1.QueryValidatorsResponse.TssIllegibilityInfo")
	proto.RegisterType((*QueryValidatorsResponse_Validator)(nil), "axelar.snapshot.v1beta1.QueryValidatorsResponse.Validator")
}

func init() {
	proto.RegisterFile("axelar/snapshot/v1beta1/query.proto", fileDescriptor_044ba28efe73d39e)
}

var fileDescriptor_044ba28efe73d39e = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x73, 0xd3, 0x4c,
	0x10, 0xc6, 0xad, 0xfc, 0x71, 0x5e, 0x5f, 0x5e, 0x06, 0x38, 0x93, 0x58, 0xe3, 0x42, 0x64, 0x48,
	0x13, 0x0a, 0xa4, 0x49, 0x32, 0x50, 0xd0, 0x91, 0x61, 0x18, 0x5c, 0x30, 0x03, 0x8a, 0x87, 0x22,
	0x8d, 0xe6, 0x64, 0xad, 0xed, 0xc3, 0xf2, 0xad, 0xb8, 0x3d, 0x07, 0xab, 0xe1, 0x03, 0x50, 0xf1,
	0x8d, 0x68, 0x5d, 0xba, 0xa4, 0x62, 0xc0, 0xfe, 0x22, 0x8c, 0x4e, 0x96, 0xe3, 0x22, 0x29, 0xe8,
	0x6e, 0x9f, 0xdf, 0x73, 0x7b, 0x3b, 0xcf, 0x4a, 0xec, 0x58, 0x4c, 0x21, 0x15, 0x3a, 0x20, 0x25,
	0x32, 0x1a, 0xa2, 0x09, 0xae, 0x4f, 0x63, 0x30, 0xe2, 0x34, 0xf8, 0x3c, 0x01, 0x9d, 0xfb, 0x99,
	0x46, 0x83, 0xbc, 0x55, 0x9a, 0xfc, 0xca, 0xe4, 0xaf, 0x4c, 0xed, 0x47, 0x03, 0x1c, 0xa0, 0xf5,
	0x04, 0xc5, 0xa9, 0xb4, 0x3f, 0xf9, 0xb6, 0xcb, 0x5a, 0x1f, 0x8a, 0xeb, 0x1f, 0x45, 0x2a, 0x13,
	0x61, 0x50, 0x53, 0x08, 0x94, 0xa1, 0x22, 0xe0, 0x57, 0x8c, 0x5d, 0xaf, 0x55, 0xd7, 0x39, 0xda,
	0x3e, 0xd9, 0x3f, 0x7b, 0xe9, 0xdf, 0xd1, 0xdf, 0xbf, 0xa3, 0x8b, 0xbf, 0x96, 0xc2, 0x8d, 0x6e,
	0xed, 0x1f, 0x5b, 0xac, 0xd9, 0x25, 0xea, 0xa4, 0x29, 0x0c, 0x64, 0x2c, 0x53, 0x69, 0xf2, 0x8e,
	0xea, 0x23, 0xf7, 0x18, 0x33, 0x38, 0x8e, 0xc9, 0xa0, 0x82, 0xc4, 0x75, 0x8e, 0x9c, 0x93, 0xff,
	0xc2, 0x0d, 0x85, 0x1f, 0xb2, 0xfa, 0x27, 0x21, 0x53, 0x48, 0xdc, 0x2d, 0xcb, 0x56, 0x15, 0x3f,
	0x67, 0x87, 0x63, 0x49, 0x04, 0x49, 0x64, 0x10, 0xa3, 0xb1, 0x50, 0x79, 0x14, 0xa7, 0xd8, 0x1b,
	0x91, 0xbb, 0x6d, 0x7d, 0xcd, 0x92, 0x76, 0x11, 0xdf, 0x09, 0x95, 0x5f, 0x58, 0xc4, 0x7d, 0xd6,
	0x54, 0x18, 0x65, 0x1a, 0xa7, 0x79, 0xa4, 0x61, 0x20, 0xc9, 0x80, 0x86, 0xc4, 0xdd, 0xb1, 0x37,
	0x1e, 0x2a, 0x7c, 0x5f, 0x90, 0x70, 0x0d, 0xf8, 0x31, 0xbb, 0x67, 0x88, 0x22, 0x9a, 0x50, 0x06,
	0x2a, 0x81, 0xc4, 0xdd, 0xb5, 0xce, 0xff, 0x0d, 0xd1, 0x65, 0xa5, 0xf1, 0x17, 0xac, 0x55, 0x76,
	0x94, 0x8a, 0x26, 0x7d, 0xd9, 0x93, 0xa0, 0x4c, 0xd4, 0x9f, 0xa8, 0x84, 0xdc, 0xba, 0xb5, 0x1f,
	0x58, 0xdc, 0xb9, 0xa1, 0x6f, 0x0a, 0x58, 0x0c, 0x43, 0x46, 0xa4, 0x10, 0x15, 0x4f, 0x0c, 0x41,
	0x68, 0x13, 0x83, 0x30, 0xee, 0x5e, 0x39, 0x8c, 0x45, 0x5d, 0xa2, 0xb7, 0x15, 0x68, 0xcf, 0x1d,
	0xd6, 0x58, 0x67, 0xcb, 0x9f, 0xb2, 0x07, 0x98, 0x81, 0x2e, 0xce, 0x91, 0x48, 0x12, 0x0d, 0x44,
	0x36, 0xbd, 0x46, 0x78, 0xbf, 0xd2, 0x5f, 0x95, 0x32, 0x77, 0xd9, 0xde, 0x18, 0x95, 0x1c, 0x81,
	0xb6, 0x19, 0x36, 0xc2, 0xaa, 0xe4, 0x5f, 0xd9, 0x41, 0xf1, 0xb8, 0xdc, 0x58, 0x4a, 0x24, 0x55,
	0x1f, 0x6d, 0x86, 0xfb, 0x67, 0xaf, 0xff, 0x79, 0xf7, 0xb7, 0x6c, 0xf8, 0x62, 0x67, 0xf6, 0xeb,
	0x71, 0x2d, 0x6c, 0x9a, 0x5b, 0xd0, 0xe5, 0xec, 0x8f, 0x57, 0x9b, 0x2d, 0x3c, 0x67, 0xbe, 0xf0,
	0x9c, 0xdf, 0x0b, 0xcf, 0xf9, 0xbe, 0xf4, 0x6a, 0xf3, 0xa5, 0x57, 0xfb, 0xb9, 0xf4, 0x6a, 0x57,
	0xcf, 0x07, 0xd2, 0x0c, 0x27, 0xb1, 0xdf, 0xc3, 0x71, 0x50, 0x0e, 0xa2, 0xc0, 0x7c, 0x41, 0x3d,
	0x5a, 0x55, 0xcf, 0x7a, 0xa8, 0x21, 0x98, 0xde, 0xfc, 0x1e, 0x26, 0xcf, 0x80, 0xe2, 0xba, 0xfd,
	0xd0, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x47, 0xfb, 0xa5, 0x3e, 0x03, 0x00, 0x00,
}

func (m *QueryValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorsResponse_TssIllegibilityInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse_TssIllegibilityInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse_TssIllegibilityInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StaleTssHeartbeat {
		i--
		if m.StaleTssHeartbeat {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.ProxyInsuficientFunds {
		i--
		if m.ProxyInsuficientFunds {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.TssSuspended {
		i--
		if m.TssSuspended {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.NoProxyRegistered {
		i--
		if m.NoProxyRegistered {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.MissedTooManyBlocks {
		i--
		if m.MissedTooManyBlocks {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorsResponse_Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse_Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse_Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TssIllegibilityInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryValidatorsResponse_TssIllegibilityInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tombstoned {
		n += 2
	}
	if m.Jailed {
		n += 2
	}
	if m.MissedTooManyBlocks {
		n += 2
	}
	if m.NoProxyRegistered {
		n += 2
	}
	if m.TssSuspended {
		n += 2
	}
	if m.ProxyInsuficientFunds {
		n += 2
	}
	if m.StaleTssHeartbeat {
		n += 2
	}
	return n
}

func (m *QueryValidatorsResponse_Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.TssIllegibilityInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &QueryValidatorsResponse_Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryValidatorsResponse_TssIllegibilityInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TssIllegibilityInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TssIllegibilityInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedTooManyBlocks", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MissedTooManyBlocks = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoProxyRegistered", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoProxyRegistered = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssSuspended", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TssSuspended = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProxyInsuficientFunds", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProxyInsuficientFunds = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaleTssHeartbeat", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StaleTssHeartbeat = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryValidatorsResponse_Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssIllegibilityInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TssIllegibilityInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
