// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/router/v3alpha/router.proto

package envoy_config_filter_http_router_v3alpha

import (
	fmt "fmt"
	v3alpha "github.com/datawire/ambassador/pkg/api/envoy/config/filter/accesslog/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// [#next-free-field: 7]
type Router struct {
	// Whether the router generates dynamic cluster statistics. Defaults to
	// true. Can be disabled in high performance scenarios.
	DynamicStats *types.BoolValue `protobuf:"bytes,1,opt,name=dynamic_stats,json=dynamicStats,proto3" json:"dynamic_stats,omitempty"`
	// Whether to start a child span for egress routed calls. This can be
	// useful in scenarios where other filters (auth, ratelimit, etc.) make
	// outbound calls and have child spans rooted at the same ingress
	// parent. Defaults to false.
	StartChildSpan bool `protobuf:"varint,2,opt,name=start_child_span,json=startChildSpan,proto3" json:"start_child_span,omitempty"`
	// Configuration for HTTP upstream logs emitted by the router. Upstream logs
	// are configured in the same way as access logs, but each log entry represents
	// an upstream request. Presuming retries are configured, multiple upstream
	// requests may be made for each downstream (inbound) request.
	UpstreamLog []*v3alpha.AccessLog `protobuf:"bytes,3,rep,name=upstream_log,json=upstreamLog,proto3" json:"upstream_log,omitempty"`
	// Do not add any additional *x-envoy-* headers to requests or responses. This
	// only affects the :ref:`router filter generated *x-envoy-* headers
	// <config_http_filters_router_headers_set>`, other Envoy filters and the HTTP
	// connection manager may continue to set *x-envoy-* headers.
	SuppressEnvoyHeaders bool `protobuf:"varint,4,opt,name=suppress_envoy_headers,json=suppressEnvoyHeaders,proto3" json:"suppress_envoy_headers,omitempty"`
	// Specifies a list of HTTP headers to strictly validate. Envoy will reject a
	// request and respond with HTTP status 400 if the request contains an invalid
	// value for any of the headers listed in this field. Strict header checking
	// is only supported for the following headers:
	//
	// Value must be a ','-delimited list (i.e. no spaces) of supported retry
	// policy values:
	//
	// * :ref:`config_http_filters_router_x-envoy-retry-grpc-on`
	// * :ref:`config_http_filters_router_x-envoy-retry-on`
	//
	// Value must be an integer:
	//
	// * :ref:`config_http_filters_router_x-envoy-max-retries`
	// * :ref:`config_http_filters_router_x-envoy-upstream-rq-timeout-ms`
	// * :ref:`config_http_filters_router_x-envoy-upstream-rq-per-try-timeout-ms`
	StrictCheckHeaders []string `protobuf:"bytes,5,rep,name=strict_check_headers,json=strictCheckHeaders,proto3" json:"strict_check_headers,omitempty"`
	// If not set, ingress Envoy will ignore
	// :ref:`config_http_filters_router_x-envoy-expected-rq-timeout-ms` header, populated by egress
	// Envoy, when deriving timeout for upstream cluster.
	RespectExpectedRqTimeout bool     `protobuf:"varint,6,opt,name=respect_expected_rq_timeout,json=respectExpectedRqTimeout,proto3" json:"respect_expected_rq_timeout,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdead4d5191ff12b, []int{0}
}
func (m *Router) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Router.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(m, src)
}
func (m *Router) XXX_Size() int {
	return m.Size()
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetDynamicStats() *types.BoolValue {
	if m != nil {
		return m.DynamicStats
	}
	return nil
}

func (m *Router) GetStartChildSpan() bool {
	if m != nil {
		return m.StartChildSpan
	}
	return false
}

func (m *Router) GetUpstreamLog() []*v3alpha.AccessLog {
	if m != nil {
		return m.UpstreamLog
	}
	return nil
}

func (m *Router) GetSuppressEnvoyHeaders() bool {
	if m != nil {
		return m.SuppressEnvoyHeaders
	}
	return false
}

func (m *Router) GetStrictCheckHeaders() []string {
	if m != nil {
		return m.StrictCheckHeaders
	}
	return nil
}

func (m *Router) GetRespectExpectedRqTimeout() bool {
	if m != nil {
		return m.RespectExpectedRqTimeout
	}
	return false
}

func init() {
	proto.RegisterType((*Router)(nil), "envoy.config.filter.http.router.v3alpha.Router")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/router/v3alpha/router.proto", fileDescriptor_fdead4d5191ff12b)
}

var fileDescriptor_fdead4d5191ff12b = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0x8f, 0xd3, 0x3e,
	0x18, 0x97, 0xff, 0xbd, 0x7f, 0x05, 0xe9, 0x81, 0x4e, 0xe1, 0x80, 0xa8, 0x48, 0x55, 0x75, 0x03,
	0x74, 0x89, 0x83, 0xee, 0x65, 0x44, 0x88, 0x9e, 0x4e, 0x62, 0xb8, 0xe1, 0xe4, 0x22, 0xd6, 0xc8,
	0x97, 0x3c, 0x4d, 0x23, 0x92, 0xd8, 0x7d, 0xec, 0x94, 0xe6, 0x0b, 0x20, 0x9d, 0xc4, 0x80, 0x18,
	0xf8, 0x30, 0x4c, 0x8c, 0xb0, 0xf1, 0x11, 0x50, 0x37, 0xbe, 0x05, 0xb2, 0x9d, 0x14, 0x90, 0x3a,
	0x30, 0xc5, 0xfe, 0xbd, 0xf9, 0x67, 0xe7, 0xf1, 0x4e, 0xa1, 0x5a, 0x89, 0x26, 0x4a, 0x44, 0x35,
	0xcf, 0xb3, 0x68, 0x9e, 0x17, 0x1a, 0x30, 0x5a, 0x68, 0x2d, 0x23, 0x14, 0xb5, 0x59, 0xaf, 0x4e,
	0x78, 0x21, 0x17, 0xbc, 0xdd, 0x52, 0x89, 0x42, 0x0b, 0xff, 0x89, 0x75, 0x51, 0xe7, 0xa2, 0xce,
	0x45, 0x8d, 0x8b, 0xb6, 0xb2, 0xd6, 0x35, 0x3c, 0xdb, 0x15, 0xcf, 0x93, 0x04, 0x94, 0x2a, 0x44,
	0xb6, 0x0d, 0xdf, 0x22, 0x2e, 0x7f, 0x38, 0xca, 0x84, 0xc8, 0x0a, 0x88, 0xec, 0xee, 0xba, 0x9e,
	0x47, 0x6f, 0x91, 0x4b, 0x09, 0xa8, 0x5a, 0xfe, 0xe1, 0x8a, 0x17, 0x79, 0xca, 0x35, 0x44, 0xdd,
	0xc2, 0x11, 0x47, 0x37, 0x7b, 0x5e, 0x9f, 0xd9, 0x0a, 0xfe, 0x73, 0xef, 0x4e, 0xda, 0x54, 0xbc,
	0xcc, 0x93, 0x58, 0x69, 0xae, 0x55, 0x40, 0xc6, 0x64, 0x32, 0x38, 0x1e, 0x52, 0x97, 0x4d, 0xbb,
	0x6c, 0x3a, 0x15, 0xa2, 0x78, 0xcd, 0x8b, 0x1a, 0xd8, 0x7e, 0x6b, 0x98, 0x19, 0xbd, 0x3f, 0xf1,
	0x0e, 0x94, 0xe6, 0xa8, 0xe3, 0x64, 0x91, 0x17, 0x69, 0xac, 0x24, 0xaf, 0x82, 0xff, 0xc6, 0x64,
	0x72, 0x8b, 0xdd, 0xb5, 0xf8, 0xb9, 0x81, 0x67, 0x92, 0x57, 0xfe, 0xcc, 0xdb, 0xaf, 0xa5, 0xd2,
	0x08, 0xbc, 0x8c, 0x0b, 0x91, 0x05, 0xbd, 0x71, 0x6f, 0x32, 0x38, 0x7e, 0x4a, 0x77, 0xbd, 0xd2,
	0xef, 0xab, 0xb6, 0x97, 0xa7, 0x2f, 0x2c, 0x72, 0x29, 0x32, 0x36, 0xe8, 0x52, 0x2e, 0x45, 0xe6,
	0x9f, 0x7a, 0x0f, 0x54, 0x2d, 0x25, 0x82, 0x52, 0xb1, 0x0d, 0x8a, 0x17, 0xc0, 0x53, 0x40, 0x15,
	0xec, 0xd9, 0x12, 0x87, 0x1d, 0x7b, 0x61, 0xc8, 0x97, 0x8e, 0xf3, 0xbf, 0x11, 0xef, 0x50, 0x69,
	0xcc, 0x13, 0x53, 0x1b, 0x92, 0x37, 0x5b, 0xd3, 0xff, 0xe3, 0xde, 0xe4, 0xf6, 0xf4, 0x13, 0xf9,
	0xfc, 0xf3, 0x4b, 0xef, 0x03, 0xf9, 0x48, 0xde, 0x93, 0xa3, 0x1b, 0x82, 0xef, 0x08, 0x1b, 0xad,
	0x43, 0x1b, 0x1f, 0x76, 0x87, 0x87, 0xb8, 0x0c, 0x75, 0x5e, 0x82, 0xa8, 0x75, 0x58, 0x2a, 0xf6,
	0x78, 0x17, 0x2f, 0x01, 0x43, 0x8d, 0xcd, 0x9f, 0xba, 0x7b, 0x9d, 0xae, 0xe4, 0xeb, 0x10, 0x41,
	0x63, 0x0e, 0x8a, 0xdd, 0xef, 0x40, 0x03, 0x34, 0x61, 0x86, 0x32, 0x09, 0x45, 0xc5, 0x0e, 0xfe,
	0x86, 0x45, 0xc5, 0x7c, 0x57, 0xfa, 0xdc, 0x74, 0xee, 0xee, 0xf2, 0xcc, 0x7b, 0x84, 0xa0, 0x24,
	0x24, 0x3a, 0x86, 0xb5, 0xf9, 0x40, 0x1a, 0xe3, 0x32, 0x6e, 0xcf, 0x0c, 0xfa, 0xf6, 0x19, 0x82,
	0x56, 0x72, 0xd1, 0x2a, 0xd8, 0xf2, 0x95, 0xe3, 0xa7, 0xec, 0xeb, 0x66, 0x44, 0xbe, 0x6f, 0x46,
	0xe4, 0xc7, 0x66, 0x44, 0xbc, 0xb3, 0x5c, 0xb8, 0xff, 0x21, 0x51, 0xac, 0x1b, 0xfa, 0x8f, 0x03,
	0x3c, 0x1d, 0xb8, 0x69, 0xba, 0x32, 0xc3, 0x72, 0x45, 0xae, 0xfb, 0x76, 0x6a, 0x4e, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xdc, 0xf2, 0xd8, 0x70, 0x37, 0x03, 0x00, 0x00,
}

func (m *Router) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Router) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Router) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RespectExpectedRqTimeout {
		i--
		if m.RespectExpectedRqTimeout {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.StrictCheckHeaders) > 0 {
		for iNdEx := len(m.StrictCheckHeaders) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StrictCheckHeaders[iNdEx])
			copy(dAtA[i:], m.StrictCheckHeaders[iNdEx])
			i = encodeVarintRouter(dAtA, i, uint64(len(m.StrictCheckHeaders[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.SuppressEnvoyHeaders {
		i--
		if m.SuppressEnvoyHeaders {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.UpstreamLog) > 0 {
		for iNdEx := len(m.UpstreamLog) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UpstreamLog[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRouter(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.StartChildSpan {
		i--
		if m.StartChildSpan {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.DynamicStats != nil {
		{
			size, err := m.DynamicStats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRouter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRouter(dAtA []byte, offset int, v uint64) int {
	offset -= sovRouter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Router) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DynamicStats != nil {
		l = m.DynamicStats.Size()
		n += 1 + l + sovRouter(uint64(l))
	}
	if m.StartChildSpan {
		n += 2
	}
	if len(m.UpstreamLog) > 0 {
		for _, e := range m.UpstreamLog {
			l = e.Size()
			n += 1 + l + sovRouter(uint64(l))
		}
	}
	if m.SuppressEnvoyHeaders {
		n += 2
	}
	if len(m.StrictCheckHeaders) > 0 {
		for _, s := range m.StrictCheckHeaders {
			l = len(s)
			n += 1 + l + sovRouter(uint64(l))
		}
	}
	if m.RespectExpectedRqTimeout {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRouter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRouter(x uint64) (n int) {
	return sovRouter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Router) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRouter
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
			return fmt.Errorf("proto: Router: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Router: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DynamicStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DynamicStats == nil {
				m.DynamicStats = &types.BoolValue{}
			}
			if err := m.DynamicStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartChildSpan", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
			m.StartChildSpan = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpstreamLog", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpstreamLog = append(m.UpstreamLog, &v3alpha.AccessLog{})
			if err := m.UpstreamLog[len(m.UpstreamLog)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuppressEnvoyHeaders", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
			m.SuppressEnvoyHeaders = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrictCheckHeaders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StrictCheckHeaders = append(m.StrictCheckHeaders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RespectExpectedRqTimeout", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
			m.RespectExpectedRqTimeout = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRouter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRouter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRouter
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
func skipRouter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRouter
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
					return 0, ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRouter
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
				return 0, ErrInvalidLengthRouter
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRouter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRouter
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRouter(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRouter
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRouter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRouter   = fmt.Errorf("proto: integer overflow")
)
