// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/network/kafka_broker/v3/kafka_broker.proto

package envoy_extensions_filters_network_kafka_broker_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type KafkaBroker struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_kafka_broker_stats>`.
	StatPrefix           string   `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaBroker) Reset()         { *m = KafkaBroker{} }
func (m *KafkaBroker) String() string { return proto.CompactTextString(m) }
func (*KafkaBroker) ProtoMessage()    {}
func (*KafkaBroker) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0a5d572866d713, []int{0}
}
func (m *KafkaBroker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KafkaBroker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KafkaBroker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KafkaBroker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaBroker.Merge(m, src)
}
func (m *KafkaBroker) XXX_Size() int {
	return m.Size()
}
func (m *KafkaBroker) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaBroker.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaBroker proto.InternalMessageInfo

func (m *KafkaBroker) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*KafkaBroker)(nil), "envoy.extensions.filters.network.kafka_broker.v3.KafkaBroker")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/kafka_broker/v3/kafka_broker.proto", fileDescriptor_4c0a5d572866d713)
}

var fileDescriptor_4c0a5d572866d713 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0xcf, 0x4e, 0x4c,
	0xcb, 0x4e, 0x8c, 0x4f, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x2f, 0x33, 0x46, 0xe1, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x19, 0x80, 0x0d, 0xd1, 0x43, 0x18, 0xa2, 0x07, 0x35, 0x44, 0x0f,
	0x6a, 0x88, 0x1e, 0x8a, 0xa6, 0x32, 0x63, 0x29, 0xd9, 0xd2, 0x94, 0x82, 0x44, 0xfd, 0xc4, 0xbc,
	0xbc, 0xfc, 0x92, 0xc4, 0x12, 0xb0, 0xb5, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x10, 0x03, 0xa5,
	0x14, 0x31, 0xa4, 0xcb, 0x52, 0x8b, 0x40, 0x26, 0x67, 0xe6, 0xa5, 0x43, 0x95, 0x88, 0x97, 0x25,
	0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x09, 0xa5, 0x5a, 0x2e, 0x6e, 0x6f,
	0x90, 0x6d, 0x4e, 0x60, 0xcb, 0x84, 0x34, 0xb8, 0xb8, 0x41, 0x46, 0xc7, 0x17, 0x14, 0xa5, 0xa6,
	0x65, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xb1, 0xff, 0x72, 0x62, 0x29, 0x62, 0x52,
	0x60, 0x0c, 0xe2, 0x02, 0xc9, 0x05, 0x80, 0xa5, 0xac, 0x5c, 0x66, 0x1d, 0xed, 0x90, 0xb3, 0xe7,
	0xb2, 0x85, 0x78, 0x26, 0x39, 0x3f, 0x2f, 0x2d, 0x33, 0x1d, 0xea, 0x11, 0x1c, 0xfe, 0x30, 0x4a,
	0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x43, 0xb2, 0xcf, 0x29, 0xe3, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0xdc, 0xd5, 0x70, 0xe2, 0x22, 0x1b, 0x93, 0x00, 0x23,
	0x97, 0x5d, 0x66, 0xbe, 0x1e, 0xd8, 0xdc, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0x52, 0xc3, 0xcb,
	0x49, 0x00, 0xc9, 0x8a, 0x00, 0x90, 0x37, 0x03, 0x18, 0x93, 0xd8, 0xc0, 0xfe, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x19, 0xb7, 0x76, 0xed, 0xc3, 0x01, 0x00, 0x00,
}

func (m *KafkaBroker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KafkaBroker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KafkaBroker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintKafkaBroker(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKafkaBroker(dAtA []byte, offset int, v uint64) int {
	offset -= sovKafkaBroker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KafkaBroker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovKafkaBroker(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovKafkaBroker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKafkaBroker(x uint64) (n int) {
	return sovKafkaBroker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KafkaBroker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKafkaBroker
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
			return fmt.Errorf("proto: KafkaBroker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KafkaBroker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafkaBroker
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
				return ErrInvalidLengthKafkaBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKafkaBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKafkaBroker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKafkaBroker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKafkaBroker
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
func skipKafkaBroker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKafkaBroker
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
					return 0, ErrIntOverflowKafkaBroker
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
					return 0, ErrIntOverflowKafkaBroker
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
				return 0, ErrInvalidLengthKafkaBroker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKafkaBroker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKafkaBroker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKafkaBroker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKafkaBroker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKafkaBroker = fmt.Errorf("proto: unexpected end of group")
)
