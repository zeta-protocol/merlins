// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlins/gamm/v1beta1/gov.proto

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

// ReplaceMigrationRecordsProposal is a gov Content type for updating the
// migration records. If a ReplaceMigrationRecordsProposal passes, the
// proposal’s records override the existing MigrationRecords set in the module.
// Each record specifies a single connection between a single balancer pool and
// a single concentrated pool.
type ReplaceMigrationRecordsProposal struct {
	Title       string                           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Records     []BalancerToConcentratedPoolLink `protobuf:"bytes,3,rep,name=records,proto3" json:"records"`
}

func (m *ReplaceMigrationRecordsProposal) Reset()      { *m = ReplaceMigrationRecordsProposal{} }
func (*ReplaceMigrationRecordsProposal) ProtoMessage() {}
func (*ReplaceMigrationRecordsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f31b9a6c0dbbdfa3, []int{0}
}
func (m *ReplaceMigrationRecordsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplaceMigrationRecordsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReplaceMigrationRecordsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplaceMigrationRecordsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceMigrationRecordsProposal.Merge(m, src)
}
func (m *ReplaceMigrationRecordsProposal) XXX_Size() int {
	return m.Size()
}
func (m *ReplaceMigrationRecordsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceMigrationRecordsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceMigrationRecordsProposal proto.InternalMessageInfo

// For example: if the existing DistrRecords were:
// [(Balancer 1, CL 5), (Balancer 2, CL 6), (Balancer 3, CL 7)]
// And an UpdateMigrationRecordsProposal includes
// [(Balancer 2, CL 0), (Balancer 3, CL 4), (Balancer 4, CL 10)]
// This would leave Balancer 1 record, delete Balancer 2 record,
// Edit Balancer 3 record, and Add Balancer 4 record
// The result MigrationRecords in state would be:
// [(Balancer 1, CL 5), (Balancer 3, CL 4), (Balancer 4, CL 10)]
type UpdateMigrationRecordsProposal struct {
	Title       string                           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Records     []BalancerToConcentratedPoolLink `protobuf:"bytes,3,rep,name=records,proto3" json:"records"`
}

func (m *UpdateMigrationRecordsProposal) Reset()      { *m = UpdateMigrationRecordsProposal{} }
func (*UpdateMigrationRecordsProposal) ProtoMessage() {}
func (*UpdateMigrationRecordsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f31b9a6c0dbbdfa3, []int{1}
}
func (m *UpdateMigrationRecordsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateMigrationRecordsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateMigrationRecordsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateMigrationRecordsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMigrationRecordsProposal.Merge(m, src)
}
func (m *UpdateMigrationRecordsProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateMigrationRecordsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMigrationRecordsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMigrationRecordsProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReplaceMigrationRecordsProposal)(nil), "merlins.gamm.v1beta1.ReplaceMigrationRecordsProposal")
	proto.RegisterType((*UpdateMigrationRecordsProposal)(nil), "merlins.gamm.v1beta1.UpdateMigrationRecordsProposal")
}

func init() { proto.RegisterFile("merlins/gamm/v1beta1/gov.proto", fileDescriptor_f31b9a6c0dbbdfa3) }

var fileDescriptor_f31b9a6c0dbbdfa3 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe3, 0xaf, 0x1f, 0x20, 0x52, 0xa6, 0xa8, 0x43, 0xd5, 0xc1, 0xa9, 0x3a, 0x75, 0x21,
	0xa6, 0xfc, 0x2c, 0x8c, 0x65, 0x43, 0x20, 0x55, 0x51, 0x59, 0xd8, 0x9c, 0xe4, 0x28, 0x58, 0x38,
	0x3e, 0x96, 0x6d, 0x2a, 0xb8, 0x03, 0x46, 0x46, 0xc6, 0x5e, 0x08, 0xec, 0x1d, 0x3b, 0x32, 0x21,
	0xd4, 0x2e, 0x5c, 0x06, 0xca, 0x4f, 0x25, 0x86, 0xde, 0x00, 0x9b, 0x8f, 0x9f, 0xd7, 0x8f, 0x8f,
	0xf4, 0xfa, 0x14, 0x6d, 0x81, 0x56, 0x58, 0x96, 0xf3, 0xa2, 0x60, 0xb3, 0x51, 0x02, 0x8e, 0x8f,
	0x58, 0x8e, 0xb3, 0x48, 0x1b, 0x74, 0x18, 0x74, 0x1a, 0x1e, 0x95, 0x3c, 0x6a, 0x78, 0xaf, 0x93,
	0x63, 0x8e, 0x55, 0x80, 0x95, 0xa7, 0x3a, 0xdb, 0x1b, 0x6c, 0x77, 0x81, 0x82, 0x52, 0x50, 0x65,
	0x06, 0xef, 0xc4, 0x0f, 0x63, 0xd0, 0x92, 0xa7, 0x70, 0x2d, 0x72, 0xc3, 0x9d, 0x40, 0x15, 0x43,
	0x8a, 0x26, 0xb3, 0x13, 0x83, 0x1a, 0x2d, 0x97, 0x41, 0xc7, 0xdf, 0x71, 0xc2, 0x49, 0xe8, 0x92,
	0x3e, 0x19, 0xee, 0xc7, 0xf5, 0x10, 0xf4, 0xfd, 0x76, 0x06, 0x36, 0x35, 0x42, 0x97, 0x6f, 0xba,
	0xff, 0x2a, 0xf6, 0xfb, 0x2a, 0x98, 0xfa, 0x7b, 0xa6, 0x56, 0x75, 0x5b, 0xfd, 0xd6, 0xb0, 0x7d,
	0x7c, 0x1a, 0x6d, 0xdb, 0x3e, 0x1a, 0x73, 0xc9, 0x55, 0x0a, 0x66, 0x8a, 0x17, 0xa8, 0x52, 0x50,
	0xce, 0x70, 0x07, 0xd9, 0x04, 0x51, 0x5e, 0x09, 0x75, 0x3f, 0xfe, 0xbf, 0xf8, 0x0c, 0xbd, 0x78,
	0xa3, 0x3a, 0x3f, 0x78, 0x9e, 0x87, 0xde, 0xeb, 0x3c, 0xf4, 0xbe, 0xe7, 0x21, 0x19, 0xbc, 0x11,
	0x9f, 0xde, 0xe8, 0x8c, 0xbb, 0x3f, 0xb9, 0xfe, 0xf8, 0x72, 0xb1, 0xa2, 0x64, 0xb9, 0xa2, 0xe4,
	0x6b, 0x45, 0xc9, 0xcb, 0x9a, 0x7a, 0xcb, 0x35, 0xf5, 0x3e, 0xd6, 0xd4, 0xbb, 0x3d, 0xca, 0x85,
	0xbb, 0x7b, 0x48, 0xa2, 0x14, 0x0b, 0xd6, 0x7c, 0x7b, 0x28, 0x79, 0x62, 0x37, 0x03, 0x9b, 0x8d,
	0xce, 0xd8, 0x63, 0x5d, 0xad, 0x7b, 0xd2, 0x60, 0x93, 0xdd, 0xaa, 0xd1, 0x93, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x07, 0xd6, 0x58, 0xf0, 0x43, 0x02, 0x00, 0x00,
}

func (this *ReplaceMigrationRecordsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplaceMigrationRecordsProposal)
	if !ok {
		that2, ok := that.(ReplaceMigrationRecordsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Records) != len(that1.Records) {
		return false
	}
	for i := range this.Records {
		if !this.Records[i].Equal(&that1.Records[i]) {
			return false
		}
	}
	return true
}
func (this *UpdateMigrationRecordsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateMigrationRecordsProposal)
	if !ok {
		that2, ok := that.(UpdateMigrationRecordsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Records) != len(that1.Records) {
		return false
	}
	for i := range this.Records {
		if !this.Records[i].Equal(&that1.Records[i]) {
			return false
		}
	}
	return true
}
func (m *ReplaceMigrationRecordsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceMigrationRecordsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplaceMigrationRecordsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateMigrationRecordsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateMigrationRecordsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateMigrationRecordsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReplaceMigrationRecordsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *UpdateMigrationRecordsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplaceMigrationRecordsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: ReplaceMigrationRecordsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceMigrationRecordsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, BalancerToConcentratedPoolLink{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *UpdateMigrationRecordsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateMigrationRecordsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateMigrationRecordsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, BalancerToConcentratedPoolLink{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
