// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/policies.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common1 "github.com/inklabsfoundation/inkchain/protos/msp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Policy_PolicyType int32

const (
	Policy_UNKNOWN       Policy_PolicyType = 0
	Policy_SIGNATURE     Policy_PolicyType = 1
	Policy_MSP           Policy_PolicyType = 2
	Policy_IMPLICIT_META Policy_PolicyType = 3
)

var Policy_PolicyType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SIGNATURE",
	2: "MSP",
	3: "IMPLICIT_META",
}
var Policy_PolicyType_value = map[string]int32{
	"UNKNOWN":       0,
	"SIGNATURE":     1,
	"MSP":           2,
	"IMPLICIT_META": 3,
}

func (x Policy_PolicyType) String() string {
	return proto.EnumName(Policy_PolicyType_name, int32(x))
}
func (Policy_PolicyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

type ImplicitMetaPolicy_Rule int32

const (
	ImplicitMetaPolicy_ANY      ImplicitMetaPolicy_Rule = 0
	ImplicitMetaPolicy_ALL      ImplicitMetaPolicy_Rule = 1
	ImplicitMetaPolicy_MAJORITY ImplicitMetaPolicy_Rule = 2
)

var ImplicitMetaPolicy_Rule_name = map[int32]string{
	0: "ANY",
	1: "ALL",
	2: "MAJORITY",
}
var ImplicitMetaPolicy_Rule_value = map[string]int32{
	"ANY":      0,
	"ALL":      1,
	"MAJORITY": 2,
}

func (x ImplicitMetaPolicy_Rule) String() string {
	return proto.EnumName(ImplicitMetaPolicy_Rule_name, int32(x))
}
func (ImplicitMetaPolicy_Rule) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{3, 0} }

// Policy expresses a policy which the orderer can evaluate, because there has been some desire expressed to support
// multiple policy engines, this is typed as a oneof for now
type Policy struct {
	Type  int32  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Policy) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Policy) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// SignaturePolicyEnvelope wraps a SignaturePolicy and includes a version for future enhancements
type SignaturePolicyEnvelope struct {
	Version    int32                   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Rule       *SignaturePolicy        `protobuf:"bytes,2,opt,name=rule" json:"rule,omitempty"`
	Identities []*common1.MSPPrincipal `protobuf:"bytes,3,rep,name=identities" json:"identities,omitempty"`
}

func (m *SignaturePolicyEnvelope) Reset()                    { *m = SignaturePolicyEnvelope{} }
func (m *SignaturePolicyEnvelope) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicyEnvelope) ProtoMessage()               {}
func (*SignaturePolicyEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *SignaturePolicyEnvelope) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SignaturePolicyEnvelope) GetRule() *SignaturePolicy {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *SignaturePolicyEnvelope) GetIdentities() []*common1.MSPPrincipal {
	if m != nil {
		return m.Identities
	}
	return nil
}

// SignaturePolicy is a recursive message structure which defines a featherweight DSL for describing
// policies which are more complicated than 'exactly this signature'.  The NOutOf operator is sufficent
// to express AND as well as OR, as well as of course N out of the following M policies
// SignedBy implies that the signature is from a valid certificate which is signed by the trusted
// authority specified in the bytes.  This will be the certificate itself for a self-signed certificate
// and will be the CA for more traditional certificates
type SignaturePolicy struct {
	// Types that are valid to be assigned to Type:
	//	*SignaturePolicy_SignedBy
	//	*SignaturePolicy_NOutOf_
	Type isSignaturePolicy_Type `protobuf_oneof:"Type"`
}

func (m *SignaturePolicy) Reset()                    { *m = SignaturePolicy{} }
func (m *SignaturePolicy) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy) ProtoMessage()               {}
func (*SignaturePolicy) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type isSignaturePolicy_Type interface {
	isSignaturePolicy_Type()
}

type SignaturePolicy_SignedBy struct {
	SignedBy int32 `protobuf:"varint,1,opt,name=signed_by,json=signedBy,oneof"`
}
type SignaturePolicy_NOutOf_ struct {
	NOutOf *SignaturePolicy_NOutOf `protobuf:"bytes,2,opt,name=n_out_of,json=nOutOf,oneof"`
}

func (*SignaturePolicy_SignedBy) isSignaturePolicy_Type() {}
func (*SignaturePolicy_NOutOf_) isSignaturePolicy_Type()  {}

func (m *SignaturePolicy) GetType() isSignaturePolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SignaturePolicy) GetSignedBy() int32 {
	if x, ok := m.GetType().(*SignaturePolicy_SignedBy); ok {
		return x.SignedBy
	}
	return 0
}

func (m *SignaturePolicy) GetNOutOf() *SignaturePolicy_NOutOf {
	if x, ok := m.GetType().(*SignaturePolicy_NOutOf_); ok {
		return x.NOutOf
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SignaturePolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SignaturePolicy_OneofMarshaler, _SignaturePolicy_OneofUnmarshaler, _SignaturePolicy_OneofSizer, []interface{}{
		(*SignaturePolicy_SignedBy)(nil),
		(*SignaturePolicy_NOutOf_)(nil),
	}
}

func _SignaturePolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_NOutOf_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NOutOf); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SignaturePolicy.Type has unexpected type %T", x)
	}
	return nil
}

func _SignaturePolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SignaturePolicy)
	switch tag {
	case 1: // Type.signed_by
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &SignaturePolicy_SignedBy{int32(x)}
		return true, err
	case 2: // Type.n_out_of
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicy_NOutOf)
		err := b.DecodeMessage(msg)
		m.Type = &SignaturePolicy_NOutOf_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SignaturePolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_NOutOf_:
		s := proto.Size(x.NOutOf)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SignaturePolicy_NOutOf struct {
	N     int32              `protobuf:"varint,1,opt,name=n" json:"n,omitempty"`
	Rules []*SignaturePolicy `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *SignaturePolicy_NOutOf) Reset()                    { *m = SignaturePolicy_NOutOf{} }
func (m *SignaturePolicy_NOutOf) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy_NOutOf) ProtoMessage()               {}
func (*SignaturePolicy_NOutOf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 0} }

func (m *SignaturePolicy_NOutOf) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *SignaturePolicy_NOutOf) GetRules() []*SignaturePolicy {
	if m != nil {
		return m.Rules
	}
	return nil
}

// ImplicitMetaPolicy is a policy type which depends on the hierarchical nature of the configuration
// It is implicit because the rule is generate implicitly based on the number of sub policies
// It is meta because it depends only on the result of other policies
// When evaluated, this policy iterates over all immediate child sub-groups, retrieves the policy
// of name sub_policy, evaluates the collection and applies the rule.
// For example, with 4 sub-groups, and a policy name of "foo", ImplicitMetaPolicy retrieves
// each sub-group, retrieves policy "foo" for each subgroup, evaluates it, and, in the case of ANY
// 1 satisfied is sufficient, ALL would require 4 signatures, and MAJORITY would require 3 signatures.
type ImplicitMetaPolicy struct {
	SubPolicy string                  `protobuf:"bytes,1,opt,name=sub_policy,json=subPolicy" json:"sub_policy,omitempty"`
	Rule      ImplicitMetaPolicy_Rule `protobuf:"varint,2,opt,name=rule,enum=common.ImplicitMetaPolicy_Rule" json:"rule,omitempty"`
}

func (m *ImplicitMetaPolicy) Reset()                    { *m = ImplicitMetaPolicy{} }
func (m *ImplicitMetaPolicy) String() string            { return proto.CompactTextString(m) }
func (*ImplicitMetaPolicy) ProtoMessage()               {}
func (*ImplicitMetaPolicy) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ImplicitMetaPolicy) GetSubPolicy() string {
	if m != nil {
		return m.SubPolicy
	}
	return ""
}

func (m *ImplicitMetaPolicy) GetRule() ImplicitMetaPolicy_Rule {
	if m != nil {
		return m.Rule
	}
	return ImplicitMetaPolicy_ANY
}

func init() {
	proto.RegisterType((*Policy)(nil), "common.Policy")
	proto.RegisterType((*SignaturePolicyEnvelope)(nil), "common.SignaturePolicyEnvelope")
	proto.RegisterType((*SignaturePolicy)(nil), "common.SignaturePolicy")
	proto.RegisterType((*SignaturePolicy_NOutOf)(nil), "common.SignaturePolicy.NOutOf")
	proto.RegisterType((*ImplicitMetaPolicy)(nil), "common.ImplicitMetaPolicy")
	proto.RegisterEnum("common.Policy_PolicyType", Policy_PolicyType_name, Policy_PolicyType_value)
	proto.RegisterEnum("common.ImplicitMetaPolicy_Rule", ImplicitMetaPolicy_Rule_name, ImplicitMetaPolicy_Rule_value)
}

func init() { proto.RegisterFile("common/policies.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5f, 0x8f, 0xd2, 0x4e,
	0x14, 0xa5, 0xfc, 0x29, 0x70, 0x61, 0x7f, 0xbf, 0x3a, 0x59, 0x03, 0xd9, 0x44, 0x25, 0x7d, 0x30,
	0x24, 0x6a, 0x9b, 0x80, 0x4f, 0xbe, 0x81, 0x12, 0xb7, 0x4a, 0x0b, 0x19, 0xd8, 0x98, 0xf5, 0xa5,
	0x69, 0x61, 0x60, 0x27, 0x96, 0x99, 0xa6, 0x33, 0x43, 0xe4, 0x5b, 0xf8, 0xe4, 0x97, 0xf1, 0xcb,
	0x99, 0x76, 0xa8, 0xd9, 0xec, 0x66, 0xdf, 0xee, 0xb9, 0x3d, 0xe7, 0xf6, 0x9c, 0x3b, 0x17, 0x9e,
	0x6f, 0xf8, 0xe1, 0xc0, 0x99, 0x9b, 0xf2, 0x84, 0x6e, 0x28, 0x11, 0x4e, 0x9a, 0x71, 0xc9, 0x91,
	0xa9, 0xdb, 0x57, 0xbd, 0x83, 0x48, 0xdd, 0x83, 0x48, 0xc3, 0x34, 0xa3, 0x6c, 0x43, 0xd3, 0x28,
	0xd1, 0x04, 0xfb, 0x27, 0x98, 0xcb, 0x5c, 0x72, 0x42, 0x08, 0xea, 0xf2, 0x94, 0x92, 0xbe, 0x31,
	0x30, 0x86, 0x0d, 0x5c, 0xd4, 0xe8, 0x12, 0x1a, 0xc7, 0x28, 0x51, 0xa4, 0x5f, 0x1d, 0x18, 0xc3,
	0x2e, 0xd6, 0xc0, 0xfe, 0x04, 0xa0, 0x35, 0xeb, 0x9c, 0xd3, 0x81, 0xe6, 0x4d, 0xf0, 0x35, 0x58,
	0x7c, 0x0b, 0xac, 0x0a, 0xba, 0x80, 0xf6, 0xca, 0xfb, 0x1c, 0x4c, 0xd6, 0x37, 0x78, 0x66, 0x19,
	0xa8, 0x09, 0x35, 0x7f, 0xb5, 0xb4, 0xaa, 0xe8, 0x19, 0x5c, 0x78, 0xfe, 0x72, 0xee, 0x7d, 0xf4,
	0xd6, 0xa1, 0x3f, 0x5b, 0x4f, 0xac, 0x9a, 0xfd, 0xdb, 0x80, 0xde, 0x8a, 0xee, 0x59, 0x24, 0x55,
	0x46, 0xf4, 0xbc, 0x19, 0x3b, 0x92, 0x84, 0xa7, 0x04, 0xf5, 0xa1, 0x79, 0x24, 0x99, 0xa0, 0x9c,
	0x9d, 0xed, 0x94, 0x10, 0xbd, 0x81, 0x7a, 0xa6, 0x12, 0x6d, 0xa8, 0x33, 0xea, 0x39, 0x3a, 0x9f,
	0xf3, 0x60, 0x10, 0x2e, 0x48, 0xe8, 0x3d, 0x00, 0xdd, 0x12, 0x26, 0xa9, 0xa4, 0x44, 0xf4, 0x6b,
	0x83, 0xda, 0xb0, 0x33, 0xba, 0x2c, 0x25, 0xfe, 0x6a, 0xb9, 0x2c, 0x97, 0x81, 0xef, 0xf1, 0xec,
	0x3f, 0x06, 0xfc, 0xff, 0x60, 0x1e, 0x7a, 0x01, 0x6d, 0x41, 0xf7, 0x8c, 0x6c, 0xc3, 0xf8, 0xa4,
	0x2d, 0x5d, 0x57, 0x70, 0x4b, 0xb7, 0xa6, 0x27, 0xf4, 0x01, 0x5a, 0x2c, 0xe4, 0x4a, 0x86, 0x7c,
	0x77, 0x76, 0xf6, 0xf2, 0x09, 0x67, 0x4e, 0xb0, 0x50, 0x72, 0xb1, 0xbb, 0xae, 0x60, 0x93, 0x15,
	0xd5, 0xd5, 0x0c, 0x4c, 0xdd, 0x43, 0x5d, 0x30, 0xca, 0xbc, 0x06, 0x43, 0xef, 0xa0, 0x91, 0x87,
	0x10, 0xfd, 0x6a, 0xe1, 0xfb, 0xc9, 0xa8, 0x9a, 0x35, 0x35, 0xa1, 0x9e, 0x3f, 0x87, 0xfd, 0xcb,
	0x00, 0xe4, 0x1d, 0xd2, 0xfc, 0x0a, 0xa4, 0x4f, 0x64, 0xf4, 0x2f, 0x00, 0x08, 0x15, 0x87, 0xc5,
	0x79, 0xe8, 0x04, 0x6d, 0xdc, 0x16, 0x2a, 0x3e, 0x7f, 0x1e, 0xdf, 0x5b, 0xeb, 0x7f, 0xa3, 0x57,
	0xe5, 0xbf, 0x1e, 0x0f, 0x72, 0xb0, 0x4a, 0x88, 0x5e, 0xaf, 0xfd, 0x1a, 0xea, 0x39, 0xca, 0x5f,
	0x79, 0x12, 0xdc, 0x5a, 0x95, 0xa2, 0x98, 0xcf, 0x2d, 0x03, 0x75, 0xa1, 0xe5, 0x4f, 0xbe, 0x2c,
	0xb0, 0xb7, 0xbe, 0xb5, 0xaa, 0xd3, 0x0d, 0xbc, 0xe5, 0xd9, 0xde, 0xa1, 0xec, 0x47, 0x12, 0xc5,
	0x62, 0xc7, 0x15, 0xdb, 0x46, 0x92, 0x72, 0x96, 0x77, 0x36, 0x77, 0x11, 0x65, 0xfa, 0x16, 0xc5,
	0xf9, 0xaf, 0xdf, 0xc7, 0x7b, 0x2a, 0xef, 0x54, 0x9c, 0x43, 0xf7, 0x91, 0xc8, 0x2d, 0x45, 0xae,
	0x16, 0xb9, 0x5a, 0x14, 0x9b, 0x05, 0x1c, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x91, 0x94, 0x54,
	0x55, 0x09, 0x03, 0x00, 0x00,
}
