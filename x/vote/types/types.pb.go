// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/vote/v1beta1/types.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// TalliedVote represents a vote for a poll with the accumulated stake of all
// validators voting for the same data
type TalliedVote struct {
	Tally  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=tally,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"tally"`
	Voters Voters                                 `protobuf:"bytes,2,rep,name=voters,proto3,castrepeated=Voters" json:"voters,omitempty"`
	Data   *types.Any                             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TalliedVote) Reset()         { *m = TalliedVote{} }
func (m *TalliedVote) String() string { return proto.CompactTextString(m) }
func (*TalliedVote) ProtoMessage()    {}
func (*TalliedVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_584be12bf9f97fd2, []int{0}
}
func (m *TalliedVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TalliedVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TalliedVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TalliedVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TalliedVote.Merge(m, src)
}
func (m *TalliedVote) XXX_Size() int {
	return m.Size()
}
func (m *TalliedVote) XXX_DiscardUnknown() {
	xxx_messageInfo_TalliedVote.DiscardUnknown(m)
}

var xxx_messageInfo_TalliedVote proto.InternalMessageInfo

type VoteRecord struct {
	Voter  github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=voter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"voter,omitempty"`
	IsLate bool                                          `protobuf:"varint,2,opt,name=is_late,json=isLate,proto3" json:"is_late,omitempty"`
}

func (m *VoteRecord) Reset()         { *m = VoteRecord{} }
func (m *VoteRecord) String() string { return proto.CompactTextString(m) }
func (*VoteRecord) ProtoMessage()    {}
func (*VoteRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_584be12bf9f97fd2, []int{1}
}
func (m *VoteRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRecord.Merge(m, src)
}
func (m *VoteRecord) XXX_Size() int {
	return m.Size()
}
func (m *VoteRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRecord.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRecord proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TalliedVote)(nil), "axelar.vote.v1beta1.TalliedVote")
	proto.RegisterType((*VoteRecord)(nil), "axelar.vote.v1beta1.VoteRecord")
}

func init() { proto.RegisterFile("axelar/vote/v1beta1/types.proto", fileDescriptor_584be12bf9f97fd2) }

var fileDescriptor_584be12bf9f97fd2 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0xeb, 0xd3, 0x30,
	0x18, 0xc7, 0x9b, 0xcd, 0x55, 0xc9, 0x76, 0xaa, 0x03, 0xbb, 0x1d, 0xda, 0xb2, 0x83, 0x14, 0xa1,
	0x09, 0xd5, 0x57, 0xb0, 0x22, 0x88, 0xa0, 0x28, 0x65, 0xec, 0xe0, 0x65, 0xa4, 0x6d, 0xec, 0xca,
	0xb2, 0x66, 0x24, 0xd9, 0x5c, 0xdf, 0x85, 0xaf, 0xc3, 0xb3, 0x2f, 0x62, 0x08, 0xc2, 0x8e, 0xe2,
	0x61, 0xea, 0xf6, 0x2e, 0x3c, 0x49, 0x9b, 0x08, 0x1e, 0x3c, 0xfc, 0x4e, 0x7d, 0xfe, 0x7c, 0x9e,
	0x7e, 0x9f, 0x6f, 0x1e, 0xe8, 0x93, 0x23, 0x65, 0x44, 0xe0, 0x03, 0x57, 0x14, 0x1f, 0xe2, 0x8c,
	0x2a, 0x12, 0x63, 0xd5, 0xec, 0xa8, 0x44, 0x3b, 0xc1, 0x15, 0x77, 0x1e, 0x6a, 0x00, 0xb5, 0x00,
	0x32, 0xc0, 0x74, 0x52, 0x72, 0x5e, 0x32, 0x8a, 0x3b, 0x24, 0xdb, 0xbf, 0xc7, 0xa4, 0x6e, 0x34,
	0x3f, 0x1d, 0x97, 0xbc, 0xe4, 0x5d, 0x88, 0xdb, 0xc8, 0x54, 0x27, 0x39, 0x97, 0x5b, 0x2e, 0x57,
	0xba, 0xa1, 0x13, 0xdd, 0x9a, 0x7d, 0x05, 0x70, 0xb8, 0x20, 0x8c, 0x55, 0xb4, 0x58, 0x72, 0x45,
	0x9d, 0xe7, 0x70, 0xa0, 0x08, 0x63, 0x8d, 0x0b, 0x02, 0x10, 0x8e, 0x12, 0x74, 0xba, 0xf8, 0xd6,
	0xf7, 0x8b, 0xff, 0xb8, 0xac, 0xd4, 0x7a, 0x9f, 0xa1, 0x9c, 0x6f, 0xcd, 0xbc, 0xf9, 0x44, 0xb2,
	0xd8, 0x98, 0x8d, 0x5f, 0xd6, 0x2a, 0xd5, 0xc3, 0xce, 0x0c, 0xda, 0xed, 0xc6, 0x42, 0xba, 0xbd,
	0xa0, 0x1f, 0x8e, 0x12, 0xf8, 0xe9, 0x87, 0x6f, 0x2f, 0xbb, 0x4a, 0x6a, 0x3a, 0xce, 0x02, 0xde,
	0x2b, 0x88, 0x22, 0x6e, 0x3f, 0x00, 0xe1, 0xf0, 0xe9, 0x18, 0x69, 0x53, 0xe8, 0xaf, 0x29, 0x34,
	0xaf, 0x9b, 0xe4, 0xc9, 0x97, 0xcf, 0xd1, 0x7f, 0xa5, 0x0b, 0x9a, 0xe3, 0xb7, 0x2d, 0xf9, 0x9a,
	0x08, 0xb9, 0x26, 0x8c, 0x8a, 0xb4, 0xfb, 0xdb, 0xac, 0x86, 0xb0, 0xd5, 0x49, 0x69, 0xce, 0x45,
	0xe1, 0xbc, 0x80, 0x83, 0x4e, 0xcd, 0xb8, 0x89, 0x7f, 0x5f, 0xfc, 0xe8, 0x0e, 0x4e, 0x96, 0x84,
	0xcd, 0x8b, 0x42, 0x50, 0x29, 0x53, 0x3d, 0xef, 0x3c, 0x82, 0xf7, 0x2b, 0xb9, 0x62, 0x44, 0x51,
	0xb7, 0x17, 0x80, 0xf0, 0x41, 0x6a, 0x57, 0xf2, 0x15, 0x51, 0x34, 0x79, 0x73, 0xfa, 0xe5, 0x59,
	0xa7, 0xab, 0x07, 0xce, 0x57, 0x0f, 0xfc, 0xbc, 0x7a, 0xe0, 0xe3, 0xcd, 0xb3, 0xce, 0x37, 0xcf,
	0xfa, 0x76, 0xf3, 0xac, 0x77, 0xf1, 0x3f, 0x62, 0xfa, 0x92, 0x35, 0x55, 0x1f, 0xb8, 0xd8, 0x98,
	0x2c, 0xca, 0xb9, 0xa0, 0xf8, 0xa8, 0xef, 0xdf, 0x69, 0x67, 0x76, 0xf7, 0x00, 0xcf, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x3d, 0x50, 0x19, 0x1c, 0x1b, 0x02, 0x00, 0x00,
}

func (m *TalliedVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TalliedVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TalliedVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.Tally.Size()
		i -= size
		if _, err := m.Tally.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VoteRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsLate {
		i--
		if m.IsLate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TalliedVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tally.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Voters) > 0 {
		for _, b := range m.Voters {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *VoteRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.IsLate {
		n += 2
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TalliedVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: TalliedVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TalliedVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tally", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters, make([]byte, postIndex-iNdEx))
			copy(m.Voters[len(m.Voters)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *VoteRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: VoteRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.IsLate = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
