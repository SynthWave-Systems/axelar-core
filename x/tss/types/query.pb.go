// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/query.proto

package types

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
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

type VoteStatus int32

const (
	Unspecified VoteStatus = 0
	Pending     VoteStatus = 1
	Decided     VoteStatus = 2
)

var VoteStatus_name = map[int32]string{
	0: "VOTE_STATUS_UNSPECIFIED",
	1: "VOTE_STATUS_PENDING",
	2: "VOTE_STATUS_DECIDED",
}

var VoteStatus_value = map[string]int32{
	"VOTE_STATUS_UNSPECIFIED": 0,
	"VOTE_STATUS_PENDING":     1,
	"VOTE_STATUS_DECIDED":     2,
}

func (x VoteStatus) String() string {
	return proto.EnumName(VoteStatus_name, int32(x))
}

func (VoteStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9e98857940a4a89, []int{0}
}

type QuerySigResponse struct {
	VoteStatus VoteStatus `protobuf:"varint,1,opt,name=vote_status,json=voteStatus,proto3,enum=tss.v1beta1.VoteStatus" json:"vote_status,omitempty"`
	Signature  *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *QuerySigResponse) Reset()         { *m = QuerySigResponse{} }
func (m *QuerySigResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigResponse) ProtoMessage()    {}
func (*QuerySigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e98857940a4a89, []int{0}
}
func (m *QuerySigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigResponse.Merge(m, src)
}
func (m *QuerySigResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigResponse proto.InternalMessageInfo

type Signature struct {
	R []byte `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	S []byte `protobuf:"bytes,3,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e98857940a4a89, []int{1}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return m.Size()
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

type QueryKeyResponse struct {
	VoteStatus VoteStatus       `protobuf:"varint,1,opt,name=vote_status,json=voteStatus,proto3,enum=tss.v1beta1.VoteStatus" json:"vote_status,omitempty"`
	Role       exported.KeyRole `protobuf:"varint,2,opt,name=role,proto3,enum=tss.exported.v1beta1.KeyRole" json:"role,omitempty"`
}

func (m *QueryKeyResponse) Reset()         { *m = QueryKeyResponse{} }
func (m *QueryKeyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryKeyResponse) ProtoMessage()    {}
func (*QueryKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e98857940a4a89, []int{2}
}
func (m *QueryKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryKeyResponse.Merge(m, src)
}
func (m *QueryKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryKeyResponse proto.InternalMessageInfo

type QueryRecoveryResponse struct {
	PartyUids          []string `protobuf:"bytes,1,rep,name=party_uids,json=partyUids,proto3" json:"party_uids,omitempty"`
	PartyShareCounts   []uint32 `protobuf:"varint,2,rep,packed,name=party_share_counts,json=partyShareCounts,proto3" json:"party_share_counts,omitempty"`
	Threshold          int32    `protobuf:"varint,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	ShareRecoveryInfos [][]byte `protobuf:"bytes,4,rep,name=share_recovery_infos,json=shareRecoveryInfos,proto3" json:"share_recovery_infos,omitempty"`
}

func (m *QueryRecoveryResponse) Reset()         { *m = QueryRecoveryResponse{} }
func (m *QueryRecoveryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRecoveryResponse) ProtoMessage()    {}
func (*QueryRecoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e98857940a4a89, []int{3}
}
func (m *QueryRecoveryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRecoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRecoveryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRecoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecoveryResponse.Merge(m, src)
}
func (m *QueryRecoveryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRecoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecoveryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("tss.v1beta1.VoteStatus", VoteStatus_name, VoteStatus_value)
	proto.RegisterType((*QuerySigResponse)(nil), "tss.v1beta1.QuerySigResponse")
	proto.RegisterType((*Signature)(nil), "tss.v1beta1.Signature")
	proto.RegisterType((*QueryKeyResponse)(nil), "tss.v1beta1.QueryKeyResponse")
	proto.RegisterType((*QueryRecoveryResponse)(nil), "tss.v1beta1.QueryRecoveryResponse")
}

func init() { proto.RegisterFile("tss/v1beta1/query.proto", fileDescriptor_b9e98857940a4a89) }

var fileDescriptor_b9e98857940a4a89 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xf6, 0x26, 0xe5, 0x27, 0x9b, 0x50, 0xa2, 0xa5, 0x90, 0x28, 0xa2, 0x96, 0x15, 0x21, 0x11,
	0xa1, 0x12, 0x37, 0x81, 0x03, 0x57, 0x88, 0x0d, 0x8a, 0x90, 0x42, 0xb0, 0x93, 0x1e, 0xb8, 0x58,
	0x8e, 0x3d, 0x75, 0x2c, 0x82, 0xd7, 0xec, 0xae, 0x43, 0x72, 0x42, 0xe2, 0x84, 0x7a, 0xea, 0x0b,
	0xf4, 0x04, 0x07, 0x1e, 0x80, 0x87, 0xe8, 0xb1, 0x47, 0x8e, 0x90, 0xbc, 0x08, 0xf2, 0x3a, 0x3f,
	0x85, 0x73, 0x4f, 0xde, 0xf9, 0x7e, 0x66, 0x3e, 0x8f, 0x06, 0x57, 0x04, 0xe7, 0xfa, 0xb4, 0x35,
	0x02, 0xe1, 0xb6, 0xf4, 0x8f, 0x09, 0xb0, 0x79, 0x33, 0x66, 0x54, 0x50, 0x52, 0x14, 0x9c, 0x37,
	0x57, 0x44, 0x6d, 0x2f, 0xa0, 0x01, 0x95, 0xb8, 0x9e, 0xbe, 0x32, 0x49, 0x4d, 0x4b, 0xbd, 0x30,
	0x8b, 0x29, 0x13, 0xe0, 0x6f, 0x9a, 0x88, 0x79, 0x0c, 0x3c, 0x53, 0xd4, 0xbf, 0x20, 0x5c, 0x7e,
	0x9b, 0x36, 0xb5, 0xc3, 0xc0, 0x02, 0x1e, 0xd3, 0x88, 0x03, 0x79, 0x86, 0x8b, 0x53, 0x2a, 0xc0,
	0xe1, 0xc2, 0x15, 0x09, 0xaf, 0x22, 0x0d, 0x35, 0x76, 0xdb, 0x95, 0xe6, 0xa5, 0x79, 0xcd, 0x23,
	0x2a, 0xc0, 0x96, 0xb4, 0x85, 0xa7, 0x9b, 0x37, 0x79, 0x8a, 0x0b, 0x3c, 0x0c, 0x22, 0x57, 0x24,
	0x0c, 0xaa, 0x39, 0x0d, 0x35, 0x8a, 0xed, 0x7b, 0xff, 0xf8, 0xec, 0x35, 0x6b, 0x6d, 0x85, 0xf5,
	0x87, 0xb8, 0xb0, 0xc1, 0x49, 0x09, 0x23, 0x26, 0xad, 0x25, 0x0b, 0xb1, 0xb4, 0xe2, 0xd5, 0x7c,
	0x56, 0xf1, 0xfa, 0xe7, 0x55, 0xd8, 0xd7, 0x30, 0xbf, 0x82, 0xb0, 0x2d, 0xbc, 0xc3, 0xe8, 0x24,
	0xcb, 0xb9, 0xdb, 0xde, 0x97, 0x96, 0xf5, 0xb2, 0x36, 0xde, 0x74, 0x14, 0x9d, 0x80, 0x25, 0xa5,
	0xf5, 0x9f, 0x08, 0xdf, 0x95, 0x09, 0x2c, 0xf0, 0xe8, 0x54, 0x7e, 0x57, 0x31, 0xf6, 0x31, 0x8e,
	0x5d, 0x26, 0xe6, 0x4e, 0x12, 0xfa, 0x69, 0x8a, 0x7c, 0xa3, 0x60, 0x15, 0x24, 0x32, 0x0c, 0x7d,
	0x4e, 0x0e, 0x30, 0xc9, 0x68, 0x3e, 0x76, 0x19, 0x38, 0x1e, 0x4d, 0x22, 0xc1, 0xab, 0x39, 0x2d,
	0xdf, 0xb8, 0x65, 0x95, 0x25, 0x63, 0xa7, 0x44, 0x47, 0xe2, 0xe4, 0x3e, 0x2e, 0x88, 0x31, 0x03,
	0x3e, 0xa6, 0x13, 0x5f, 0xfe, 0xfd, 0x35, 0x6b, 0x0b, 0x90, 0x43, 0xbc, 0x97, 0x75, 0x61, 0xab,
	0x10, 0x4e, 0x18, 0x1d, 0x53, 0x5e, 0xdd, 0xd1, 0xf2, 0x8d, 0x92, 0x45, 0x24, 0xb7, 0xce, 0xd7,
	0x4d, 0x99, 0x47, 0xa7, 0x08, 0xe3, 0xed, 0x12, 0xc8, 0x01, 0xae, 0x1c, 0xbd, 0x19, 0x98, 0x8e,
	0x3d, 0x78, 0x3e, 0x18, 0xda, 0xce, 0xb0, 0x67, 0xf7, 0xcd, 0x4e, 0xf7, 0x65, 0xd7, 0x34, 0xca,
	0x4a, 0xed, 0xf6, 0xc9, 0x99, 0x56, 0x1c, 0x46, 0x3c, 0x06, 0x2f, 0x3c, 0x0e, 0xc1, 0x27, 0x0f,
	0xf0, 0x9d, 0xcb, 0xea, 0xbe, 0xd9, 0x33, 0xba, 0xbd, 0x57, 0x65, 0x54, 0x2b, 0x9e, 0x9c, 0x69,
	0x37, 0xfa, 0x10, 0xf9, 0x61, 0x14, 0xfc, 0xaf, 0x32, 0xcc, 0x4e, 0xd7, 0x30, 0x8d, 0x72, 0x2e,
	0x53, 0x19, 0xe0, 0x85, 0x3e, 0xf8, 0xb5, 0x9b, 0x5f, 0xbf, 0xa9, 0xca, 0x8f, 0xef, 0x2a, 0x7a,
	0xd1, 0x3b, 0xff, 0xa3, 0x2a, 0xe7, 0x0b, 0x15, 0x5d, 0x2c, 0x54, 0xf4, 0x7b, 0xa1, 0xa2, 0xd3,
	0xa5, 0xaa, 0x5c, 0x2c, 0x55, 0xe5, 0xd7, 0x52, 0x55, 0xde, 0x1d, 0x06, 0xa1, 0x18, 0x27, 0xa3,
	0xa6, 0x47, 0x3f, 0xe8, 0xee, 0x0c, 0x26, 0x2e, 0x8b, 0x40, 0x7c, 0xa2, 0xec, 0xfd, 0xaa, 0x7a,
	0xec, 0x51, 0x06, 0xfa, 0x4c, 0x4f, 0xef, 0x5b, 0x9e, 0xf3, 0xe8, 0xba, 0xbc, 0xe7, 0x27, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x88, 0xdf, 0xf4, 0x2f, 0x03, 0x00, 0x00,
}

func (m *QuerySigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Signature != nil {
		{
			size, err := m.Signature.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.VoteStatus != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.VoteStatus))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Signature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Signature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.S) > 0 {
		i -= len(m.S)
		copy(dAtA[i:], m.S)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.S)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.R) > 0 {
		i -= len(m.R)
		copy(dAtA[i:], m.R)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.R)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *QueryKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Role != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x10
	}
	if m.VoteStatus != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.VoteStatus))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryRecoveryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRecoveryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRecoveryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShareRecoveryInfos) > 0 {
		for iNdEx := len(m.ShareRecoveryInfos) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ShareRecoveryInfos[iNdEx])
			copy(dAtA[i:], m.ShareRecoveryInfos[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.ShareRecoveryInfos[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Threshold != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PartyShareCounts) > 0 {
		dAtA3 := make([]byte, len(m.PartyShareCounts)*10)
		var j2 int
		for _, num := range m.PartyShareCounts {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintQuery(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PartyUids) > 0 {
		for iNdEx := len(m.PartyUids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PartyUids[iNdEx])
			copy(dAtA[i:], m.PartyUids[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.PartyUids[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *QuerySigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VoteStatus != 0 {
		n += 1 + sovQuery(uint64(m.VoteStatus))
	}
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *Signature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.R)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.S)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VoteStatus != 0 {
		n += 1 + sovQuery(uint64(m.VoteStatus))
	}
	if m.Role != 0 {
		n += 1 + sovQuery(uint64(m.Role))
	}
	return n
}

func (m *QueryRecoveryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PartyUids) > 0 {
		for _, s := range m.PartyUids {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.PartyShareCounts) > 0 {
		l = 0
		for _, e := range m.PartyShareCounts {
			l += sovQuery(uint64(e))
		}
		n += 1 + sovQuery(uint64(l)) + l
	}
	if m.Threshold != 0 {
		n += 1 + sovQuery(uint64(m.Threshold))
	}
	if len(m.ShareRecoveryInfos) > 0 {
		for _, b := range m.ShareRecoveryInfos {
			l = len(b)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySigResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteStatus", wireType)
			}
			m.VoteStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteStatus |= VoteStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			if m.Signature == nil {
				m.Signature = &Signature{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Signature) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Signature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.R = append(m.R[:0], dAtA[iNdEx:postIndex]...)
			if m.R == nil {
				m.R = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.S = append(m.S[:0], dAtA[iNdEx:postIndex]...)
			if m.S == nil {
				m.S = []byte{}
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
func (m *QueryKeyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteStatus", wireType)
			}
			m.VoteStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteStatus |= VoteStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= exported.KeyRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryRecoveryResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRecoveryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRecoveryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyUids", wireType)
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
			m.PartyUids = append(m.PartyUids, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PartyShareCounts = append(m.PartyShareCounts, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthQuery
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthQuery
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PartyShareCounts) == 0 {
					m.PartyShareCounts = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PartyShareCounts = append(m.PartyShareCounts, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyShareCounts", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareRecoveryInfos", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareRecoveryInfos = append(m.ShareRecoveryInfos, make([]byte, postIndex-iNdEx))
			copy(m.ShareRecoveryInfos[len(m.ShareRecoveryInfos)-1], dAtA[iNdEx:postIndex])
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
