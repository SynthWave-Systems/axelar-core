// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nexus/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/nexus/exported"
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

// GenesisState represents the genesis state
type GenesisState struct {
	Params          Params                        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Nonce           uint64                        `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Chains          []exported.Chain              `protobuf:"bytes,3,rep,name=chains,proto3" json:"chains"`
	ChainStates     []ChainState                  `protobuf:"bytes,4,rep,name=chain_states,json=chainStates,proto3" json:"chain_states"`
	LinkedAddresses []LinkedAddresses             `protobuf:"bytes,5,rep,name=linked_addresses,json=linkedAddresses,proto3" json:"linked_addresses"`
	Transfers       []exported.CrossChainTransfer `protobuf:"bytes,6,rep,name=transfers,proto3" json:"transfers"`
	Fee             exported.TransferFee          `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d58ead19bb1ba601, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "nexus.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("nexus/v1beta1/genesis.proto", fileDescriptor_d58ead19bb1ba601) }

var fileDescriptor_d58ead19bb1ba601 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0xef, 0xd2, 0x30,
	0x18, 0xc6, 0x37, 0x07, 0x33, 0x16, 0x8c, 0xa6, 0xc1, 0x64, 0xcc, 0x58, 0x09, 0x5e, 0x88, 0x89,
	0x5b, 0x80, 0x23, 0x27, 0x31, 0xd1, 0x8b, 0x51, 0x82, 0x9e, 0xbc, 0x90, 0xb2, 0xbd, 0x8c, 0x05,
	0x68, 0x97, 0xb6, 0xe8, 0xfc, 0x16, 0x7e, 0x1c, 0x3f, 0x02, 0x47, 0x8e, 0x9e, 0x8c, 0xc2, 0x17,
	0x31, 0xeb, 0x3a, 0xfc, 0x8f, 0x84, 0x5b, 0xdb, 0xe7, 0x79, 0x7e, 0x79, 0xde, 0xe6, 0x45, 0x4f,
	0x19, 0xe4, 0x7b, 0x19, 0x7e, 0x1d, 0x2e, 0x41, 0xd1, 0x61, 0x98, 0x00, 0x03, 0x99, 0xca, 0x20,
	0x13, 0x5c, 0x71, 0xfc, 0x50, 0x8b, 0x81, 0x11, 0xfd, 0x4e, 0xc2, 0x13, 0xae, 0x95, 0xb0, 0x38,
	0x95, 0x26, 0xdf, 0xaf, 0x13, 0x32, 0x2a, 0xe8, 0xce, 0x00, 0xfc, 0x7e, 0xa9, 0x41, 0x9e, 0x71,
	0xa1, 0x20, 0xbe, 0x98, 0xd4, 0xf7, 0x0c, 0x2a, 0x4f, 0xb7, 0x9e, 0xbf, 0x23, 0xf5, 0x7f, 0x3a,
	0xa8, 0xfd, 0xae, 0x6c, 0xf4, 0x49, 0x51, 0x05, 0x78, 0x8c, 0xdc, 0x92, 0xef, 0xd9, 0x3d, 0x7b,
	0xd0, 0x1a, 0x3d, 0x09, 0x6a, 0x0d, 0x83, 0x99, 0x16, 0xa7, 0x8d, 0xc3, 0xef, 0xe7, 0xd6, 0xdc,
	0x58, 0x71, 0x07, 0x35, 0x19, 0x67, 0x11, 0x78, 0xf7, 0x7a, 0xf6, 0xa0, 0x31, 0x2f, 0x2f, 0x78,
	0x82, 0xdc, 0x68, 0x4d, 0x53, 0x26, 0x3d, 0xa7, 0xe7, 0x0c, 0x5a, 0xa3, 0x67, 0x06, 0x55, 0x75,
	0xbd, 0x30, 0xdf, 0x14, 0xae, 0x0a, 0x59, 0x46, 0xf0, 0x14, 0xb5, 0xf5, 0x69, 0x21, 0x8b, 0x5a,
	0xd2, 0x6b, 0x68, 0x44, 0xf7, 0xaa, 0x8d, 0x4e, 0xea, 0xe2, 0x26, 0xde, 0x8a, 0x2e, 0x2f, 0x12,
	0x7f, 0x44, 0x8f, 0xb7, 0x29, 0xdb, 0x40, 0xbc, 0xa0, 0x71, 0x2c, 0x40, 0x4a, 0x90, 0x5e, 0x53,
	0x73, 0xc8, 0x15, 0xe7, 0xbd, 0xb6, 0xbd, 0xae, 0x5c, 0x06, 0xf6, 0x68, 0x5b, 0x7f, 0xc6, 0x1f,
	0xd0, 0x03, 0x25, 0x28, 0x93, 0x2b, 0x10, 0xd2, 0x73, 0x35, 0xe9, 0xe5, 0xcd, 0xa1, 0x04, 0x97,
	0x52, 0xf7, 0xfb, 0x6c, 0x22, 0x86, 0xfa, 0x1f, 0x81, 0x27, 0xc8, 0x59, 0x01, 0x78, 0xf7, 0xf5,
	0x4f, 0xbf, 0xb8, 0x45, 0xaa, 0xf2, 0x6f, 0xa1, 0x9a, 0xb2, 0x48, 0x4d, 0x67, 0x87, 0xbf, 0xc4,
	0x3a, 0x9c, 0x88, 0x7d, 0x3c, 0x11, 0xfb, 0xcf, 0x89, 0xd8, 0x3f, 0xce, 0xc4, 0x3a, 0x9e, 0x89,
	0xf5, 0xeb, 0x4c, 0xac, 0x2f, 0xa3, 0x24, 0x55, 0xeb, 0xfd, 0x32, 0x88, 0xf8, 0x2e, 0xa4, 0x39,
	0x6c, 0xa9, 0x60, 0xa0, 0xbe, 0x71, 0xb1, 0x31, 0xb7, 0x57, 0x11, 0x17, 0x10, 0xe6, 0x61, 0xb9,
	0x1a, 0x7a, 0x25, 0x96, 0xae, 0xde, 0x89, 0xf1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xdd,
	0x74, 0xf7, 0xb2, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Transfers) > 0 {
		for iNdEx := len(m.Transfers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transfers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LinkedAddresses) > 0 {
		for iNdEx := len(m.LinkedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LinkedAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ChainStates) > 0 {
		for iNdEx := len(m.ChainStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Nonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Nonce != 0 {
		n += 1 + sovGenesis(uint64(m.Nonce))
	}
	if len(m.Chains) > 0 {
		for _, e := range m.Chains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChainStates) > 0 {
		for _, e := range m.ChainStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LinkedAddresses) > 0 {
		for _, e := range m.LinkedAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Transfers) > 0 {
		for _, e := range m.Transfers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Fee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, exported.Chain{})
			if err := m.Chains[len(m.Chains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainStates = append(m.ChainStates, ChainState{})
			if err := m.ChainStates[len(m.ChainStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkedAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkedAddresses = append(m.LinkedAddresses, LinkedAddresses{})
			if err := m.LinkedAddresses[len(m.LinkedAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transfers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transfers = append(m.Transfers, exported.CrossChainTransfer{})
			if err := m.Transfers[len(m.Transfers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
