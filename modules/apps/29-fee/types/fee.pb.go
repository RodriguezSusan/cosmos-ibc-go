// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fee/v1/fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/ibc-go/modules/core/04-channel/types"
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

// Fee interface
// See Fee Payment Middleware spec:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-029-fee-payment#fee-middleware-contract
type Fee struct {
	ReceiveFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=receive_fee,json=receiveFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"receive_fee" yaml:"receive_fee"`
	AckFee     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=ack_fee,json=ackFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"ack_fee" yaml:"ack_fee"`
	TimeoutFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=timeout_fee,json=timeoutFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"timeout_fee" yaml:"timeout_fee"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{0}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetReceiveFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ReceiveFee
	}
	return nil
}

func (m *Fee) GetAckFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.AckFee
	}
	return nil
}

func (m *Fee) GetTimeoutFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TimeoutFee
	}
	return nil
}

// Fee associated with a packet_id
type IdentifiedPacketFee struct {
	PacketId *types1.PacketId `protobuf:"bytes,1,opt,name=packet_id,json=packetId,proto3" json:"packet_id,omitempty" yaml:"packet_id"`
	Fee      *Fee             `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Relayers []string         `protobuf:"bytes,3,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *IdentifiedPacketFee) Reset()         { *m = IdentifiedPacketFee{} }
func (m *IdentifiedPacketFee) String() string { return proto.CompactTextString(m) }
func (*IdentifiedPacketFee) ProtoMessage()    {}
func (*IdentifiedPacketFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{1}
}
func (m *IdentifiedPacketFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedPacketFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedPacketFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedPacketFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedPacketFee.Merge(m, src)
}
func (m *IdentifiedPacketFee) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedPacketFee) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedPacketFee.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedPacketFee proto.InternalMessageInfo

func (m *IdentifiedPacketFee) GetPacketId() *types1.PacketId {
	if m != nil {
		return m.PacketId
	}
	return nil
}

func (m *IdentifiedPacketFee) GetFee() *Fee {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *IdentifiedPacketFee) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

func init() {
	proto.RegisterType((*Fee)(nil), "ibc.applications.fee.v1.Fee")
	proto.RegisterType((*IdentifiedPacketFee)(nil), "ibc.applications.fee.v1.IdentifiedPacketFee")
}

func init() { proto.RegisterFile("ibc/applications/fee/v1/fee.proto", fileDescriptor_cb3319f1af2a53e5) }

var fileDescriptor_cb3319f1af2a53e5 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbd, 0x8a, 0x1b, 0x31,
	0x10, 0xc7, 0xad, 0x2c, 0x5c, 0xee, 0x64, 0x08, 0x61, 0x73, 0x90, 0x8b, 0x49, 0xd6, 0x97, 0xad,
	0xdc, 0x58, 0xc2, 0x4e, 0x95, 0x94, 0x0e, 0x2c, 0x1c, 0xa4, 0x38, 0xb6, 0x4c, 0x73, 0x68, 0xb5,
	0xe3, 0x3d, 0xb1, 0x1f, 0x5a, 0x56, 0xf2, 0x82, 0xdb, 0x34, 0x69, 0xf3, 0x1c, 0x69, 0xf3, 0x12,
	0x57, 0x5e, 0x99, 0xca, 0x09, 0xf6, 0x1b, 0xdc, 0x13, 0x04, 0x7d, 0x9c, 0x31, 0x84, 0x10, 0x5c,
	0x69, 0x24, 0xcd, 0x5f, 0xbf, 0x99, 0xd1, 0x0c, 0x7e, 0x2b, 0x32, 0x4e, 0x59, 0xdb, 0x56, 0x82,
	0x33, 0x2d, 0x64, 0xa3, 0xe8, 0x12, 0x80, 0xf6, 0x33, 0xb3, 0x90, 0xb6, 0x93, 0x5a, 0x86, 0x2f,
	0x45, 0xc6, 0xc9, 0xa1, 0x0b, 0x31, 0x77, 0xfd, 0x6c, 0x14, 0x71, 0xa9, 0x6a, 0xa9, 0x68, 0xc6,
	0x94, 0x91, 0x64, 0xa0, 0xd9, 0x8c, 0x72, 0x29, 0x1a, 0x27, 0x1c, 0x9d, 0x17, 0xb2, 0x90, 0xd6,
	0xa4, 0xc6, 0xf2, 0xa7, 0x96, 0xc8, 0x65, 0x07, 0x94, 0xdf, 0xb2, 0xa6, 0x81, 0xca, 0xd0, 0xbc,
	0xe9, 0x5c, 0xe2, 0xaf, 0x01, 0x0e, 0x12, 0x80, 0xf0, 0x0b, 0xc2, 0xc3, 0x0e, 0x38, 0x88, 0x1e,
	0x6e, 0x96, 0x00, 0x17, 0xe8, 0x32, 0x98, 0x0c, 0xe7, 0xaf, 0x88, 0xe3, 0x12, 0xc3, 0x25, 0x9e,
	0x4b, 0x3e, 0x4a, 0xd1, 0x2c, 0x92, 0xbb, 0xcd, 0x78, 0xf0, 0xb0, 0x19, 0x87, 0x6b, 0x56, 0x57,
	0x1f, 0xe2, 0x03, 0x6d, 0xfc, 0xfd, 0xd7, 0x78, 0x52, 0x08, 0x7d, 0xbb, 0xca, 0x08, 0x97, 0x35,
	0xf5, 0xa1, 0xbb, 0x65, 0xaa, 0xf2, 0x92, 0xea, 0x75, 0x0b, 0xca, 0x3e, 0xa3, 0x52, 0xec, 0x95,
	0x26, 0x88, 0x1e, 0x3f, 0x65, 0xbc, 0xb4, 0xfc, 0x27, 0xff, 0xe3, 0x2f, 0x3c, 0xff, 0x99, 0xe3,
	0x7b, 0xdd, 0x71, 0xec, 0x13, 0xc6, 0xcb, 0xc7, 0xe4, 0xb5, 0xa8, 0x41, 0xae, 0xb4, 0x85, 0x07,
	0x47, 0x26, 0x7f, 0xa0, 0x3d, 0x32, 0x79, 0xaf, 0x4c, 0x00, 0xe2, 0x1f, 0x08, 0xbf, 0xb8, 0xca,
	0xa1, 0xd1, 0x62, 0x29, 0x20, 0xbf, 0x66, 0xbc, 0x04, 0x73, 0x1e, 0x5e, 0xe3, 0xb3, 0xd6, 0x6e,
	0x6e, 0x44, 0x7e, 0x81, 0x2e, 0xd1, 0x64, 0x38, 0x7f, 0x43, 0x4c, 0x9f, 0x98, 0x8f, 0x25, 0x8f,
	0xbf, 0xd9, 0xcf, 0x88, 0x93, 0x5c, 0xe5, 0x8b, 0xf3, 0x87, 0xcd, 0xf8, 0xb9, 0x8b, 0x6c, 0xaf,
	0x8c, 0xd3, 0xd3, 0xd6, 0xdf, 0x87, 0x04, 0x07, 0xae, 0xc4, 0xe6, 0xad, 0xd7, 0xe4, 0x1f, 0x3d,
	0x47, 0x12, 0x80, 0xd4, 0x38, 0x86, 0x23, 0x7c, 0xda, 0x41, 0xc5, 0xd6, 0xd0, 0x29, 0x5b, 0x9a,
	0xb3, 0x74, 0xbf, 0x5f, 0x7c, 0xba, 0xdb, 0x46, 0xe8, 0x7e, 0x1b, 0xa1, 0xdf, 0xdb, 0x08, 0x7d,
	0xdb, 0x45, 0x83, 0xfb, 0x5d, 0x34, 0xf8, 0xb9, 0x8b, 0x06, 0x9f, 0xe7, 0x7f, 0x57, 0x41, 0x64,
	0x7c, 0x5a, 0x48, 0x5a, 0xcb, 0x7c, 0x55, 0x81, 0x32, 0xb3, 0xa0, 0xe8, 0xfc, 0xfd, 0xd4, 0x8c,
	0x81, 0xad, 0x4a, 0x76, 0x62, 0x9b, 0xf2, 0xdd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x74,
	0xda, 0xb7, 0x2b, 0x03, 0x00, 0x00,
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TimeoutFee) > 0 {
		for iNdEx := len(m.TimeoutFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeoutFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AckFee) > 0 {
		for iNdEx := len(m.AckFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AckFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ReceiveFee) > 0 {
		for iNdEx := len(m.ReceiveFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReceiveFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IdentifiedPacketFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedPacketFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedPacketFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintFee(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.PacketId != nil {
		{
			size, err := m.PacketId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ReceiveFee) > 0 {
		for _, e := range m.ReceiveFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.AckFee) > 0 {
		for _, e := range m.AckFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.TimeoutFee) > 0 {
		for _, e := range m.TimeoutFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *IdentifiedPacketFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketId != nil {
		l = m.PacketId.Size()
		n += 1 + l + sovFee(uint64(l))
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovFee(uint64(l))
	}
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiveFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiveFee = append(m.ReceiveFee, types.Coin{})
			if err := m.ReceiveFee[len(m.ReceiveFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckFee = append(m.AckFee, types.Coin{})
			if err := m.AckFee[len(m.AckFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutFee = append(m.TimeoutFee, types.Coin{})
			if err := m.TimeoutFee[len(m.TimeoutFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func (m *IdentifiedPacketFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: IdentifiedPacketFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedPacketFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PacketId == nil {
				m.PacketId = &types1.PacketId{}
			}
			if err := m.PacketId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &Fee{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)
