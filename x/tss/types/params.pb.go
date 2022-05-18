// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/tss/v1beta1/params.proto

package types

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
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

// Params is the parameter set for this module
type Params struct {
	// KeyRequirements defines the requirement for each key role
	KeyRequirements []exported.KeyRequirement `protobuf:"bytes,1,rep,name=key_requirements,json=keyRequirements,proto3" json:"key_requirements"`
	// SuspendDurationInBlocks defines the number of blocks a
	// validator is disallowed to participate in any TSS ceremony after
	// committing a malicious behaviour during signing
	SuspendDurationInBlocks int64 `protobuf:"varint,2,opt,name=suspend_duration_in_blocks,json=suspendDurationInBlocks,proto3" json:"suspend_duration_in_blocks,omitempty"`
	// HeartBeatPeriodInBlocks defines the time period in blocks for tss to
	// emit the event asking validators to send their heartbeats
	HeartbeatPeriodInBlocks          int64           `protobuf:"varint,3,opt,name=heartbeat_period_in_blocks,json=heartbeatPeriodInBlocks,proto3" json:"heartbeat_period_in_blocks,omitempty"`
	MaxMissedBlocksPerWindow         utils.Threshold `protobuf:"bytes,4,opt,name=max_missed_blocks_per_window,json=maxMissedBlocksPerWindow,proto3" json:"max_missed_blocks_per_window"`
	UnbondingLockingKeyRotationCount int64           `protobuf:"varint,5,opt,name=unbonding_locking_key_rotation_count,json=unbondingLockingKeyRotationCount,proto3" json:"unbonding_locking_key_rotation_count,omitempty"`
	ExternalMultisigThreshold        utils.Threshold `protobuf:"bytes,6,opt,name=external_multisig_threshold,json=externalMultisigThreshold,proto3" json:"external_multisig_threshold"`
	MaxSignQueueSize                 int64           `protobuf:"varint,7,opt,name=max_sign_queue_size,json=maxSignQueueSize,proto3" json:"max_sign_queue_size,omitempty"`
	MaxSimultaneousSignShares        int64           `protobuf:"varint,8,opt,name=max_simultaneous_sign_shares,json=maxSimultaneousSignShares,proto3" json:"max_simultaneous_sign_shares,omitempty"`
	TssSignedBlocksWindow            int64           `protobuf:"varint,9,opt,name=tss_signed_blocks_window,json=tssSignedBlocksWindow,proto3" json:"tss_signed_blocks_window,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_58e1981dcb992056, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "axelar.tss.v1beta1.Params")
}

func init() { proto.RegisterFile("axelar/tss/v1beta1/params.proto", fileDescriptor_58e1981dcb992056) }

var fileDescriptor_58e1981dcb992056 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x63, 0xd2, 0x06, 0x70, 0x17, 0x54, 0x06, 0x84, 0x1b, 0x90, 0x13, 0xa1, 0x4a, 0x44,
	0x42, 0xb5, 0x69, 0x59, 0xb0, 0x60, 0x81, 0x14, 0xd8, 0x20, 0x68, 0x15, 0x12, 0x24, 0x24, 0x84,
	0x34, 0x1a, 0xc7, 0x57, 0xce, 0x28, 0xf6, 0x8c, 0x3b, 0x77, 0x4c, 0x9c, 0x3e, 0x05, 0x6f, 0xc1,
	0xab, 0x64, 0xd9, 0x25, 0x2b, 0x04, 0xc9, 0x8b, 0x20, 0x8f, 0xc7, 0x4e, 0x58, 0x76, 0x95, 0x9f,
	0xfb, 0x9d, 0xe3, 0x73, 0xcf, 0x78, 0xec, 0x1e, 0x2d, 0x20, 0xa1, 0x32, 0x50, 0x88, 0xc1, 0xf7,
	0xd3, 0x10, 0x14, 0x3d, 0x0d, 0x32, 0x2a, 0x69, 0x8a, 0x7e, 0x26, 0x85, 0x12, 0x8e, 0x53, 0x01,
	0xbe, 0x42, 0xf4, 0x0d, 0xd0, 0x7d, 0x10, 0x8b, 0x58, 0xe8, 0x71, 0x50, 0x7e, 0xab, 0xc8, 0xee,
	0xb1, 0xb1, 0xca, 0x15, 0x4b, 0xb6, 0x66, 0x6a, 0x26, 0x01, 0x67, 0x22, 0x89, 0x0c, 0xf5, 0x6c,
	0xe7, 0x81, 0x50, 0x64, 0x42, 0x2a, 0x88, 0xb6, 0xf0, 0x32, 0x03, 0xf3, 0xe0, 0xa7, 0x3f, 0xf7,
	0xed, 0xce, 0x48, 0x27, 0x71, 0xbe, 0xd9, 0x87, 0x73, 0x58, 0x12, 0x09, 0x97, 0x39, 0x93, 0x90,
	0x02, 0x57, 0xe8, 0x5a, 0xfd, 0xf6, 0xe0, 0xe0, 0xec, 0xb9, 0xbf, 0x13, 0xaf, 0xb6, 0xab, 0x73,
	0xfa, 0x1f, 0x60, 0x39, 0xde, 0x6a, 0x86, 0x7b, 0xab, 0xdf, 0xbd, 0xd6, 0xf8, 0xde, 0xfc, 0xbf,
	0x7f, 0xd1, 0x79, 0x6d, 0x77, 0x31, 0xc7, 0x0c, 0x78, 0x44, 0xa2, 0x5c, 0x52, 0xc5, 0x04, 0x27,
	0x8c, 0x93, 0x30, 0x11, 0xd3, 0x39, 0xba, 0xb7, 0xfa, 0xd6, 0xa0, 0x3d, 0x7e, 0x64, 0x88, 0x77,
	0x06, 0x78, 0xcf, 0x87, 0x7a, 0x5c, 0x8a, 0x67, 0x40, 0xa5, 0x0a, 0x81, 0x2a, 0x92, 0x81, 0x64,
	0x22, 0xda, 0x11, 0xb7, 0x2b, 0x71, 0x43, 0x8c, 0x34, 0xd0, 0x88, 0xc1, 0x7e, 0x92, 0xd2, 0x82,
	0xa4, 0x0c, 0x11, 0x22, 0xa3, 0x29, 0x4d, 0xc8, 0x82, 0xf1, 0x48, 0x2c, 0xdc, 0xbd, 0xbe, 0x35,
	0x38, 0x38, 0xeb, 0xd5, 0x3b, 0xea, 0x62, 0x9b, 0xe5, 0x3e, 0xd7, 0xc5, 0x9a, 0xbd, 0xdc, 0x94,
	0x16, 0xe7, 0xda, 0xa9, 0x72, 0x1f, 0x81, 0xfc, 0xa2, 0x6d, 0x9c, 0x0b, 0xfb, 0x38, 0xe7, 0xa1,
	0xe0, 0x11, 0xe3, 0x31, 0x29, 0x67, 0xe5, 0xa7, 0x2e, 0x54, 0xa8, 0x6a, 0xdd, 0xa9, 0xc8, 0xb9,
	0x72, 0xf7, 0x75, 0xda, 0x7e, 0xc3, 0x7e, 0xac, 0xd0, 0xb2, 0x45, 0x03, 0xbe, 0x2d, 0x39, 0x07,
	0xec, 0xc7, 0x50, 0x28, 0x90, 0x9c, 0x26, 0x24, 0xcd, 0x13, 0xc5, 0x90, 0xc5, 0xa4, 0x39, 0x67,
	0xb7, 0x73, 0x93, 0xd4, 0x47, 0xb5, 0xd3, 0xb9, 0x31, 0x6a, 0x00, 0xe7, 0xc4, 0xbe, 0x5f, 0xb6,
	0x83, 0x2c, 0xe6, 0xe4, 0x32, 0x87, 0x1c, 0x08, 0xb2, 0x2b, 0x70, 0x6f, 0xeb, 0x94, 0x87, 0x29,
	0x2d, 0x26, 0x2c, 0xe6, 0x9f, 0xca, 0xc1, 0x84, 0x5d, 0x81, 0xf3, 0xa6, 0x2a, 0x13, 0x59, 0x19,
	0x89, 0x72, 0x10, 0x39, 0x56, 0x5a, 0x9c, 0x51, 0x09, 0xe8, 0xde, 0xd1, 0xba, 0x23, 0xad, 0xdb,
	0x22, 0xa5, 0xc7, 0x44, 0x03, 0xce, 0x2b, 0xdb, 0x55, 0x58, 0x69, 0xb6, 0xa7, 0x61, 0x4e, 0xe2,
	0xae, 0x16, 0x3f, 0x54, 0xa8, 0x05, 0x75, 0xc5, 0x55, 0xbf, 0xc3, 0x8b, 0xd5, 0x5f, 0xaf, 0xb5,
	0x5a, 0x7b, 0xd6, 0xf5, 0xda, 0xb3, 0xfe, 0xac, 0x3d, 0xeb, 0xc7, 0xc6, 0x6b, 0x5d, 0x6f, 0xbc,
	0xd6, 0xaf, 0x8d, 0xd7, 0xfa, 0xfa, 0x22, 0x66, 0x6a, 0x96, 0x87, 0xfe, 0x54, 0xa4, 0x41, 0x55,
	0x09, 0x07, 0xb5, 0x10, 0x72, 0x6e, 0x7e, 0x9d, 0x4c, 0x85, 0x84, 0xa0, 0xd0, 0x17, 0x42, 0xbf,
	0xff, 0x61, 0x47, 0x5f, 0x80, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0xd1, 0xcc, 0xbd,
	0x9c, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TssSignedBlocksWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TssSignedBlocksWindow))
		i--
		dAtA[i] = 0x48
	}
	if m.MaxSimultaneousSignShares != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSimultaneousSignShares))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxSignQueueSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSignQueueSize))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.ExternalMultisigThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.UnbondingLockingKeyRotationCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnbondingLockingKeyRotationCount))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.MaxMissedBlocksPerWindow.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.HeartbeatPeriodInBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HeartbeatPeriodInBlocks))
		i--
		dAtA[i] = 0x18
	}
	if m.SuspendDurationInBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SuspendDurationInBlocks))
		i--
		dAtA[i] = 0x10
	}
	if len(m.KeyRequirements) > 0 {
		for iNdEx := len(m.KeyRequirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyRequirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.KeyRequirements) > 0 {
		for _, e := range m.KeyRequirements {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.SuspendDurationInBlocks != 0 {
		n += 1 + sovParams(uint64(m.SuspendDurationInBlocks))
	}
	if m.HeartbeatPeriodInBlocks != 0 {
		n += 1 + sovParams(uint64(m.HeartbeatPeriodInBlocks))
	}
	l = m.MaxMissedBlocksPerWindow.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.UnbondingLockingKeyRotationCount != 0 {
		n += 1 + sovParams(uint64(m.UnbondingLockingKeyRotationCount))
	}
	l = m.ExternalMultisigThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxSignQueueSize != 0 {
		n += 1 + sovParams(uint64(m.MaxSignQueueSize))
	}
	if m.MaxSimultaneousSignShares != 0 {
		n += 1 + sovParams(uint64(m.MaxSimultaneousSignShares))
	}
	if m.TssSignedBlocksWindow != 0 {
		n += 1 + sovParams(uint64(m.TssSignedBlocksWindow))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyRequirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyRequirements = append(m.KeyRequirements, exported.KeyRequirement{})
			if err := m.KeyRequirements[len(m.KeyRequirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuspendDurationInBlocks", wireType)
			}
			m.SuspendDurationInBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuspendDurationInBlocks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatPeriodInBlocks", wireType)
			}
			m.HeartbeatPeriodInBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartbeatPeriodInBlocks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMissedBlocksPerWindow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxMissedBlocksPerWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingLockingKeyRotationCount", wireType)
			}
			m.UnbondingLockingKeyRotationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingLockingKeyRotationCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalMultisigThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExternalMultisigThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSignQueueSize", wireType)
			}
			m.MaxSignQueueSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSignQueueSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSimultaneousSignShares", wireType)
			}
			m.MaxSimultaneousSignShares = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSimultaneousSignShares |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssSignedBlocksWindow", wireType)
			}
			m.TssSignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TssSignedBlocksWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
