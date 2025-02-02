// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go-internal/handshake.proto

package handshake

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

type BoxEnvelope struct {
	Box                  []byte   `protobuf:"bytes,1,opt,name=box,proto3" json:"box,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoxEnvelope) Reset()         { *m = BoxEnvelope{} }
func (m *BoxEnvelope) String() string { return proto.CompactTextString(m) }
func (*BoxEnvelope) ProtoMessage()    {}
func (*BoxEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc780342ca42053, []int{0}
}
func (m *BoxEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BoxEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BoxEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BoxEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoxEnvelope.Merge(m, src)
}
func (m *BoxEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *BoxEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_BoxEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_BoxEnvelope proto.InternalMessageInfo

func (m *BoxEnvelope) GetBox() []byte {
	if m != nil {
		return m.Box
	}
	return nil
}

type HelloPayload struct {
	EphemeralPubKey      []byte   `protobuf:"bytes,1,opt,name=ephemeral_pub_key,json=ephemeralPubKey,proto3" json:"ephemeral_pub_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloPayload) Reset()         { *m = HelloPayload{} }
func (m *HelloPayload) String() string { return proto.CompactTextString(m) }
func (*HelloPayload) ProtoMessage()    {}
func (*HelloPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc780342ca42053, []int{1}
}
func (m *HelloPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloPayload.Merge(m, src)
}
func (m *HelloPayload) XXX_Size() int {
	return m.Size()
}
func (m *HelloPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloPayload.DiscardUnknown(m)
}

var xxx_messageInfo_HelloPayload proto.InternalMessageInfo

func (m *HelloPayload) GetEphemeralPubKey() []byte {
	if m != nil {
		return m.EphemeralPubKey
	}
	return nil
}

type RequesterAuthenticatePayload struct {
	RequesterAccountId   []byte   `protobuf:"bytes,1,opt,name=requester_account_id,json=requesterAccountId,proto3" json:"requester_account_id,omitempty"`
	RequesterAccountSig  []byte   `protobuf:"bytes,2,opt,name=requester_account_sig,json=requesterAccountSig,proto3" json:"requester_account_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequesterAuthenticatePayload) Reset()         { *m = RequesterAuthenticatePayload{} }
func (m *RequesterAuthenticatePayload) String() string { return proto.CompactTextString(m) }
func (*RequesterAuthenticatePayload) ProtoMessage()    {}
func (*RequesterAuthenticatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc780342ca42053, []int{2}
}
func (m *RequesterAuthenticatePayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequesterAuthenticatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequesterAuthenticatePayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequesterAuthenticatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequesterAuthenticatePayload.Merge(m, src)
}
func (m *RequesterAuthenticatePayload) XXX_Size() int {
	return m.Size()
}
func (m *RequesterAuthenticatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_RequesterAuthenticatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_RequesterAuthenticatePayload proto.InternalMessageInfo

func (m *RequesterAuthenticatePayload) GetRequesterAccountId() []byte {
	if m != nil {
		return m.RequesterAccountId
	}
	return nil
}

func (m *RequesterAuthenticatePayload) GetRequesterAccountSig() []byte {
	if m != nil {
		return m.RequesterAccountSig
	}
	return nil
}

type ResponderAcceptPayload struct {
	ResponderAccountSig  []byte   `protobuf:"bytes,1,opt,name=responder_account_sig,json=responderAccountSig,proto3" json:"responder_account_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponderAcceptPayload) Reset()         { *m = ResponderAcceptPayload{} }
func (m *ResponderAcceptPayload) String() string { return proto.CompactTextString(m) }
func (*ResponderAcceptPayload) ProtoMessage()    {}
func (*ResponderAcceptPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc780342ca42053, []int{3}
}
func (m *ResponderAcceptPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponderAcceptPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponderAcceptPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponderAcceptPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponderAcceptPayload.Merge(m, src)
}
func (m *ResponderAcceptPayload) XXX_Size() int {
	return m.Size()
}
func (m *ResponderAcceptPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponderAcceptPayload.DiscardUnknown(m)
}

var xxx_messageInfo_ResponderAcceptPayload proto.InternalMessageInfo

func (m *ResponderAcceptPayload) GetResponderAccountSig() []byte {
	if m != nil {
		return m.ResponderAccountSig
	}
	return nil
}

type RequesterAcknowledgePayload struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequesterAcknowledgePayload) Reset()         { *m = RequesterAcknowledgePayload{} }
func (m *RequesterAcknowledgePayload) String() string { return proto.CompactTextString(m) }
func (*RequesterAcknowledgePayload) ProtoMessage()    {}
func (*RequesterAcknowledgePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc780342ca42053, []int{4}
}
func (m *RequesterAcknowledgePayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequesterAcknowledgePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequesterAcknowledgePayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequesterAcknowledgePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequesterAcknowledgePayload.Merge(m, src)
}
func (m *RequesterAcknowledgePayload) XXX_Size() int {
	return m.Size()
}
func (m *RequesterAcknowledgePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_RequesterAcknowledgePayload.DiscardUnknown(m)
}

var xxx_messageInfo_RequesterAcknowledgePayload proto.InternalMessageInfo

func (m *RequesterAcknowledgePayload) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*BoxEnvelope)(nil), "handshake.BoxEnvelope")
	proto.RegisterType((*HelloPayload)(nil), "handshake.HelloPayload")
	proto.RegisterType((*RequesterAuthenticatePayload)(nil), "handshake.RequesterAuthenticatePayload")
	proto.RegisterType((*ResponderAcceptPayload)(nil), "handshake.ResponderAcceptPayload")
	proto.RegisterType((*RequesterAcknowledgePayload)(nil), "handshake.RequesterAcknowledgePayload")
}

func init() { proto.RegisterFile("go-internal/handshake.proto", fileDescriptor_7dc780342ca42053) }

var fileDescriptor_7dc780342ca42053 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0xef, 0x03, 0xff, 0x8c, 0x05, 0x35, 0x56, 0x29, 0x56, 0xaa, 0x64, 0x21, 0x22,
	0xd8, 0x88, 0x0a, 0x82, 0xbb, 0x16, 0x04, 0x45, 0x17, 0x25, 0xee, 0xdc, 0x84, 0xc9, 0xe4, 0x32,
	0x09, 0x1d, 0xe7, 0xc6, 0xcc, 0x44, 0x9b, 0xbd, 0x0f, 0xe7, 0xd2, 0x47, 0x90, 0x3e, 0x89, 0x74,
	0x9a, 0x49, 0xd1, 0xee, 0xce, 0x9c, 0x73, 0xcf, 0x6f, 0xb8, 0x5c, 0xda, 0x15, 0x78, 0x96, 0x29,
	0x03, 0x85, 0x62, 0x32, 0x48, 0x99, 0x4a, 0x74, 0xca, 0xc6, 0xd0, 0xcf, 0x0b, 0x34, 0xe8, 0xad,
	0x37, 0xc6, 0x7e, 0x5b, 0xa0, 0x40, 0xeb, 0x06, 0x33, 0x35, 0x1f, 0xf0, 0x0f, 0xe9, 0xc6, 0x10,
	0x27, 0xb7, 0xea, 0x0d, 0x24, 0xe6, 0xe0, 0x6d, 0xd1, 0xff, 0x31, 0x4e, 0x3a, 0xe4, 0x88, 0x9c,
	0xb4, 0xc2, 0x99, 0xf4, 0x6f, 0x68, 0xeb, 0x0e, 0xa4, 0xc4, 0x11, 0xab, 0x24, 0xb2, 0xc4, 0x3b,
	0xa5, 0xdb, 0x90, 0xa7, 0xf0, 0x02, 0x05, 0x93, 0x51, 0x5e, 0xc6, 0xd1, 0x18, 0xaa, 0x7a, 0x7e,
	0xb3, 0x09, 0x46, 0x65, 0xfc, 0x00, 0x95, 0xff, 0x41, 0xe8, 0x41, 0x08, 0xaf, 0x25, 0x68, 0x03,
	0xc5, 0xa0, 0x34, 0x29, 0x28, 0x93, 0x71, 0x66, 0xc0, 0xc1, 0xce, 0x69, 0xbb, 0x70, 0x79, 0xc4,
	0x38, 0xc7, 0x52, 0x99, 0x28, 0x4b, 0x6a, 0x9e, 0xd7, 0x64, 0x83, 0x79, 0x74, 0x9f, 0x78, 0x17,
	0x74, 0x77, 0xb9, 0xa1, 0x33, 0xd1, 0xf9, 0x67, 0x2b, 0x3b, 0x7f, 0x2b, 0x4f, 0x99, 0xf0, 0x1f,
	0xe9, 0x5e, 0x08, 0x3a, 0x47, 0x95, 0x58, 0x1b, 0x72, 0xe3, 0xfe, 0xb7, 0xb4, 0x3a, 0xf9, 0x45,
	0x23, 0x8e, 0xb6, 0xa8, 0x39, 0xda, 0x35, 0xed, 0x2e, 0x76, 0xe2, 0x63, 0x85, 0xef, 0x12, 0x12,
	0xd1, 0xac, 0xd4, 0xa1, 0xab, 0xba, 0xe4, 0x1c, 0xb4, 0xb6, 0x90, 0xb5, 0xd0, 0x3d, 0x87, 0x57,
	0x9f, 0xd3, 0x1e, 0xf9, 0x9a, 0xf6, 0xc8, 0xf7, 0xb4, 0x47, 0x9e, 0x8f, 0x63, 0x28, 0x4c, 0xd5,
	0x37, 0xc0, 0xd3, 0xc0, 0xca, 0x40, 0x60, 0xb0, 0x7c, 0xc7, 0x78, 0xc5, 0xde, 0xe9, 0xf2, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0xb4, 0x63, 0x43, 0xe7, 0x01, 0x00, 0x00,
}

func (m *BoxEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BoxEnvelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BoxEnvelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Box) > 0 {
		i -= len(m.Box)
		copy(dAtA[i:], m.Box)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.Box)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelloPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.EphemeralPubKey) > 0 {
		i -= len(m.EphemeralPubKey)
		copy(dAtA[i:], m.EphemeralPubKey)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.EphemeralPubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequesterAuthenticatePayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequesterAuthenticatePayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequesterAuthenticatePayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RequesterAccountSig) > 0 {
		i -= len(m.RequesterAccountSig)
		copy(dAtA[i:], m.RequesterAccountSig)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.RequesterAccountSig)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RequesterAccountId) > 0 {
		i -= len(m.RequesterAccountId)
		copy(dAtA[i:], m.RequesterAccountId)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.RequesterAccountId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponderAcceptPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponderAcceptPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponderAcceptPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ResponderAccountSig) > 0 {
		i -= len(m.ResponderAccountSig)
		copy(dAtA[i:], m.ResponderAccountSig)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.ResponderAccountSig)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequesterAcknowledgePayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequesterAcknowledgePayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequesterAcknowledgePayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHandshake(dAtA []byte, offset int, v uint64) int {
	offset -= sovHandshake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BoxEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Box)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HelloPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EphemeralPubKey)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequesterAuthenticatePayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequesterAccountId)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.RequesterAccountSig)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponderAcceptPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResponderAccountSig)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequesterAcknowledgePayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHandshake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHandshake(x uint64) (n int) {
	return sovHandshake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BoxEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: BoxEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BoxEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Box", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Box = append(m.Box[:0], dAtA[iNdEx:postIndex]...)
			if m.Box == nil {
				m.Box = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: HelloPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EphemeralPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EphemeralPubKey = append(m.EphemeralPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EphemeralPubKey == nil {
				m.EphemeralPubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequesterAuthenticatePayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: RequesterAuthenticatePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequesterAuthenticatePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequesterAccountId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequesterAccountId = append(m.RequesterAccountId[:0], dAtA[iNdEx:postIndex]...)
			if m.RequesterAccountId == nil {
				m.RequesterAccountId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequesterAccountSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequesterAccountSig = append(m.RequesterAccountSig[:0], dAtA[iNdEx:postIndex]...)
			if m.RequesterAccountSig == nil {
				m.RequesterAccountSig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponderAcceptPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: ResponderAcceptPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponderAcceptPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponderAccountSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponderAccountSig = append(m.ResponderAccountSig[:0], dAtA[iNdEx:postIndex]...)
			if m.ResponderAccountSig == nil {
				m.ResponderAccountSig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequesterAcknowledgePayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: RequesterAcknowledgePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequesterAcknowledgePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHandshake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandshake
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
					return 0, ErrIntOverflowHandshake
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
					return 0, ErrIntOverflowHandshake
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
				return 0, ErrInvalidLengthHandshake
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHandshake
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHandshake
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHandshake        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandshake          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHandshake = fmt.Errorf("proto: unexpected end of group")
)
