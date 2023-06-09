// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlins/superfluid/genesis.proto

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

// GenesisState defines the module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// superfluid_assets defines the registered superfluid assets that have been
	// registered via governance.
	SuperfluidAssets []SuperfluidAsset `protobuf:"bytes,2,rep,name=superfluid_assets,json=superfluidAssets,proto3" json:"superfluid_assets"`
	// fury_equivalent_multipliers is the records of fury equivalent amount of
	// each superfluid registered pool, updated every epoch.
	FuryEquivalentMultipliers []FuryEquivalentMultiplierRecord `protobuf:"bytes,3,rep,name=fury_equivalent_multipliers,json=furyEquivalentMultipliers,proto3" json:"fury_equivalent_multipliers"`
	// intermediary_accounts is a secondary account for superfluid staking that
	// plays an intermediary role between validators and the delegators.
	IntermediaryAccounts          []SuperfluidIntermediaryAccount       `protobuf:"bytes,4,rep,name=intermediary_accounts,json=intermediaryAccounts,proto3" json:"intermediary_accounts"`
	IntemediaryAccountConnections []LockIdIntermediaryAccountConnection `protobuf:"bytes,5,rep,name=intemediary_account_connections,json=intemediaryAccountConnections,proto3" json:"intemediary_account_connections"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5256ebb7c83fff3, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSuperfluidAssets() []SuperfluidAsset {
	if m != nil {
		return m.SuperfluidAssets
	}
	return nil
}

func (m *GenesisState) GetFuryEquivalentMultipliers() []FuryEquivalentMultiplierRecord {
	if m != nil {
		return m.FuryEquivalentMultipliers
	}
	return nil
}

func (m *GenesisState) GetIntermediaryAccounts() []SuperfluidIntermediaryAccount {
	if m != nil {
		return m.IntermediaryAccounts
	}
	return nil
}

func (m *GenesisState) GetIntemediaryAccountConnections() []LockIdIntermediaryAccountConnection {
	if m != nil {
		return m.IntemediaryAccountConnections
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "merlins.superfluid.GenesisState")
}

func init() { proto.RegisterFile("merlins/superfluid/genesis.proto", fileDescriptor_d5256ebb7c83fff3) }

var fileDescriptor_d5256ebb7c83fff3 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0x87, 0x5b, 0x41, 0x0e, 0xc5, 0x83, 0x36, 0x98, 0x54, 0x8c, 0x85, 0xc8, 0x85, 0x8b, 0x6d,
	0xc0, 0xf8, 0xe7, 0x0a, 0xc6, 0x18, 0x12, 0x8d, 0x04, 0x12, 0x0f, 0x5e, 0x9a, 0xa5, 0xac, 0x75,
	0x63, 0xdb, 0xad, 0x3b, 0x5b, 0x02, 0x0f, 0xe0, 0x9d, 0xc7, 0xe2, 0xc8, 0xd1, 0x93, 0x31, 0xf0,
	0x22, 0xa6, 0xed, 0x5a, 0x50, 0xaa, 0xb7, 0x69, 0xe7, 0xfb, 0xcd, 0x37, 0x9b, 0x8c, 0x52, 0xa5,
	0xe0, 0x51, 0x20, 0x60, 0x42, 0x18, 0x60, 0xf6, 0xe4, 0x86, 0x64, 0x68, 0x3a, 0xd8, 0xc7, 0x40,
	0xc0, 0x08, 0x18, 0xe5, 0x54, 0x55, 0x05, 0x61, 0xac, 0x88, 0x72, 0xc9, 0xa1, 0x0e, 0x8d, 0xdb,
	0x66, 0x54, 0x25, 0x64, 0xb9, 0x96, 0x31, 0x6b, 0x55, 0x0a, 0xa8, 0x92, 0x01, 0x05, 0x88, 0x21,
	0x4f, 0xf8, 0x8e, 0xa7, 0x79, 0x65, 0xe7, 0x26, 0xd9, 0xa0, 0xcf, 0x11, 0xc7, 0xea, 0xa5, 0x52,
	0x48, 0x00, 0x4d, 0xae, 0xca, 0xf5, 0x62, 0xb3, 0x6c, 0x6c, 0x6e, 0x64, 0x74, 0x63, 0xa2, 0x9d,
	0x9f, 0x7d, 0x54, 0xa4, 0x9e, 0xe0, 0xd5, 0x07, 0x65, 0x6f, 0x85, 0x58, 0x08, 0x00, 0x73, 0xd0,
	0xb6, 0xaa, 0xb9, 0x7a, 0xb1, 0x59, 0xcb, 0x1a, 0xd2, 0x4f, 0xcb, 0x56, 0xc4, 0x8a, 0x69, 0xbb,
	0xf0, 0xf3, 0x37, 0xa8, 0x63, 0xe5, 0x30, 0x4a, 0x5b, 0xf8, 0x35, 0x24, 0x23, 0xe4, 0x62, 0x9f,
	0x5b, 0x5e, 0xe8, 0x72, 0x12, 0xb8, 0x04, 0x33, 0xd0, 0x72, 0xb1, 0xa1, 0x99, 0x65, 0xb8, 0x07,
	0x8f, 0x5e, 0xa7, 0xa9, 0xbb, 0x34, 0xd4, 0xc3, 0x36, 0x65, 0x43, 0x21, 0x3c, 0xa0, 0x7f, 0x50,
	0xa0, 0xba, 0xca, 0x3e, 0xf1, 0x39, 0x66, 0x1e, 0x1e, 0x12, 0xc4, 0x26, 0x16, 0xb2, 0x6d, 0x1a,
	0xfa, 0x1c, 0xb4, 0x7c, 0xec, 0x6c, 0xfc, 0xff, 0xaa, 0xce, 0x5a, 0xb4, 0x95, 0x24, 0x85, 0xb2,
	0x44, 0x36, 0x5b, 0xa0, 0xbe, 0xc9, 0x4a, 0x25, 0x6a, 0xfc, 0xb2, 0x59, 0x36, 0xf5, 0x7d, 0x6c,
	0x73, 0x42, 0x7d, 0xd0, 0xb6, 0x63, 0xf1, 0x45, 0x96, 0xf8, 0x96, 0xda, 0x2f, 0x9d, 0x2c, 0xe9,
	0x55, 0x9a, 0x17, 0xfa, 0xa3, 0x35, 0xcb, 0x06, 0x03, 0xed, 0xee, 0x6c, 0xa1, 0xcb, 0xf3, 0x85,
	0x2e, 0x7f, 0x2e, 0x74, 0x79, 0xba, 0xd4, 0xa5, 0xf9, 0x52, 0x97, 0xde, 0x97, 0xba, 0xf4, 0x78,
	0xee, 0x10, 0xfe, 0x1c, 0x0e, 0x0c, 0x9b, 0x7a, 0xa6, 0xd8, 0xe0, 0xc4, 0x45, 0x03, 0xf8, 0xfe,
	0x30, 0x47, 0x8d, 0x33, 0x73, 0xbc, 0x7e, 0x6b, 0x7c, 0x12, 0x60, 0x18, 0x14, 0xe2, 0x5b, 0x3b,
	0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x72, 0x55, 0x84, 0xff, 0x02, 0x00, 0x00,
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
	if len(m.IntemediaryAccountConnections) > 0 {
		for iNdEx := len(m.IntemediaryAccountConnections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IntemediaryAccountConnections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.IntermediaryAccounts) > 0 {
		for iNdEx := len(m.IntermediaryAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IntermediaryAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.FuryEquivalentMultipliers) > 0 {
		for iNdEx := len(m.FuryEquivalentMultipliers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FuryEquivalentMultipliers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SuperfluidAssets) > 0 {
		for iNdEx := len(m.SuperfluidAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SuperfluidAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
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
	if len(m.SuperfluidAssets) > 0 {
		for _, e := range m.SuperfluidAssets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FuryEquivalentMultipliers) > 0 {
		for _, e := range m.FuryEquivalentMultipliers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IntermediaryAccounts) > 0 {
		for _, e := range m.IntermediaryAccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IntemediaryAccountConnections) > 0 {
		for _, e := range m.IntemediaryAccountConnections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidAssets", wireType)
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
			m.SuperfluidAssets = append(m.SuperfluidAssets, SuperfluidAsset{})
			if err := m.SuperfluidAssets[len(m.SuperfluidAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuryEquivalentMultipliers", wireType)
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
			m.FuryEquivalentMultipliers = append(m.FuryEquivalentMultipliers, FuryEquivalentMultiplierRecord{})
			if err := m.FuryEquivalentMultipliers[len(m.FuryEquivalentMultipliers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntermediaryAccounts", wireType)
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
			m.IntermediaryAccounts = append(m.IntermediaryAccounts, SuperfluidIntermediaryAccount{})
			if err := m.IntermediaryAccounts[len(m.IntermediaryAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntemediaryAccountConnections", wireType)
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
			m.IntemediaryAccountConnections = append(m.IntemediaryAccountConnections, LockIdIntermediaryAccountConnection{})
			if err := m.IntemediaryAccountConnections[len(m.IntemediaryAccountConnections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
