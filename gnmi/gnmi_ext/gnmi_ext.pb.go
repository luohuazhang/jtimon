// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gnmi_ext.proto

/*
Package gnmi_ext is a generated protocol buffer package.

Package gnmi_ext defines a set of extensions messages which can be optionally
included with the request and response messages of gNMI RPCs. A set of
well-known extensions are defined within this file, along with a registry for
extensions defined outside of this package.

It is generated from these files:
	gnmi_ext.proto

It has these top-level messages:
	Extension
	RegisteredExtension
	MasterArbitration
	Uint128
	Role
*/
package gnmi_ext

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RegisteredExtension is an enumeration acting as a registry for extensions
// defined by external sources.
type ExtensionID int32

const (
	ExtensionID_EID_UNSET ExtensionID = 0
	// Juniper Telemetry header
	ExtensionID_EID_JUNIPER_TELEMETRY_HEADER ExtensionID = 1
	// An experimental extension that may be used during prototyping of a new
	// extension.
	ExtensionID_EID_EXPERIMENTAL ExtensionID = 999
)

var ExtensionID_name = map[int32]string{
	0:   "EID_UNSET",
	1:   "EID_JUNIPER_TELEMETRY_HEADER",
	999: "EID_EXPERIMENTAL",
}
var ExtensionID_value = map[string]int32{
	"EID_UNSET":                    0,
	"EID_JUNIPER_TELEMETRY_HEADER": 1,
	"EID_EXPERIMENTAL":             999,
}

func (x ExtensionID) String() string {
	return proto.EnumName(ExtensionID_name, int32(x))
}
func (ExtensionID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The Extension message contains a single gNMI extension.
type Extension struct {
	// Types that are valid to be assigned to Ext:
	//	*Extension_RegisteredExt
	//	*Extension_MasterArbitration
	Ext isExtension_Ext `protobuf_oneof:"ext"`
}

func (m *Extension) Reset()                    { *m = Extension{} }
func (m *Extension) String() string            { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()               {}
func (*Extension) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isExtension_Ext interface{ isExtension_Ext() }

type Extension_RegisteredExt struct {
	RegisteredExt *RegisteredExtension `protobuf:"bytes,1,opt,name=registered_ext,json=registeredExt,oneof"`
}
type Extension_MasterArbitration struct {
	MasterArbitration *MasterArbitration `protobuf:"bytes,2,opt,name=master_arbitration,json=masterArbitration,oneof"`
}

func (*Extension_RegisteredExt) isExtension_Ext()     {}
func (*Extension_MasterArbitration) isExtension_Ext() {}

func (m *Extension) GetExt() isExtension_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (m *Extension) GetRegisteredExt() *RegisteredExtension {
	if x, ok := m.GetExt().(*Extension_RegisteredExt); ok {
		return x.RegisteredExt
	}
	return nil
}

func (m *Extension) GetMasterArbitration() *MasterArbitration {
	if x, ok := m.GetExt().(*Extension_MasterArbitration); ok {
		return x.MasterArbitration
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Extension) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Extension_OneofMarshaler, _Extension_OneofUnmarshaler, _Extension_OneofSizer, []interface{}{
		(*Extension_RegisteredExt)(nil),
		(*Extension_MasterArbitration)(nil),
	}
}

func _Extension_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Extension)
	// ext
	switch x := m.Ext.(type) {
	case *Extension_RegisteredExt:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RegisteredExt); err != nil {
			return err
		}
	case *Extension_MasterArbitration:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MasterArbitration); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Extension.Ext has unexpected type %T", x)
	}
	return nil
}

func _Extension_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Extension)
	switch tag {
	case 1: // ext.registered_ext
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegisteredExtension)
		err := b.DecodeMessage(msg)
		m.Ext = &Extension_RegisteredExt{msg}
		return true, err
	case 2: // ext.master_arbitration
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MasterArbitration)
		err := b.DecodeMessage(msg)
		m.Ext = &Extension_MasterArbitration{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Extension_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Extension)
	// ext
	switch x := m.Ext.(type) {
	case *Extension_RegisteredExt:
		s := proto.Size(x.RegisteredExt)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Extension_MasterArbitration:
		s := proto.Size(x.MasterArbitration)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The RegisteredExtension message defines an extension which is defined outside
// of this file.
type RegisteredExtension struct {
	Id  ExtensionID `protobuf:"varint,1,opt,name=id,enum=gnmi_ext.ExtensionID" json:"id,omitempty"`
	Msg []byte      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *RegisteredExtension) Reset()                    { *m = RegisteredExtension{} }
func (m *RegisteredExtension) String() string            { return proto.CompactTextString(m) }
func (*RegisteredExtension) ProtoMessage()               {}
func (*RegisteredExtension) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisteredExtension) GetId() ExtensionID {
	if m != nil {
		return m.Id
	}
	return ExtensionID_EID_UNSET
}

func (m *RegisteredExtension) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// MasterArbitration is used to select the master among multiple gNMI clients
// with the same Roles. The client with the largest election_id is honored as
// the master.
// The document about gNMI master arbitration can be found at
// https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-master-arbitration.md
type MasterArbitration struct {
	Role       *Role    `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	ElectionId *Uint128 `protobuf:"bytes,2,opt,name=election_id,json=electionId" json:"election_id,omitempty"`
}

func (m *MasterArbitration) Reset()                    { *m = MasterArbitration{} }
func (m *MasterArbitration) String() string            { return proto.CompactTextString(m) }
func (*MasterArbitration) ProtoMessage()               {}
func (*MasterArbitration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MasterArbitration) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *MasterArbitration) GetElectionId() *Uint128 {
	if m != nil {
		return m.ElectionId
	}
	return nil
}

// Representation of unsigned 128-bit integer.
type Uint128 struct {
	High uint64 `protobuf:"varint,1,opt,name=high" json:"high,omitempty"`
	Low  uint64 `protobuf:"varint,2,opt,name=low" json:"low,omitempty"`
}

func (m *Uint128) Reset()                    { *m = Uint128{} }
func (m *Uint128) String() string            { return proto.CompactTextString(m) }
func (*Uint128) ProtoMessage()               {}
func (*Uint128) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Uint128) GetHigh() uint64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Uint128) GetLow() uint64 {
	if m != nil {
		return m.Low
	}
	return 0
}

// There can be one master for each role. The role is identified by its id.
type Role struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Extension)(nil), "gnmi_ext.Extension")
	proto.RegisterType((*RegisteredExtension)(nil), "gnmi_ext.RegisteredExtension")
	proto.RegisterType((*MasterArbitration)(nil), "gnmi_ext.MasterArbitration")
	proto.RegisterType((*Uint128)(nil), "gnmi_ext.Uint128")
	proto.RegisterType((*Role)(nil), "gnmi_ext.Role")
	proto.RegisterEnum("gnmi_ext.ExtensionID", ExtensionID_name, ExtensionID_value)
}

func init() { proto.RegisterFile("gnmi_ext.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6a, 0xe2, 0x40,
	0x14, 0xc6, 0x8d, 0x66, 0xd7, 0xf5, 0xb8, 0x86, 0x38, 0x8b, 0x8b, 0xd0, 0x16, 0x24, 0x50, 0xe8,
	0x1f, 0x30, 0x34, 0xbd, 0xe9, 0xad, 0xe2, 0x14, 0x53, 0x34, 0xc8, 0x34, 0x81, 0xb6, 0x37, 0x41,
	0xcd, 0x34, 0x0e, 0x4d, 0x32, 0x12, 0xa7, 0xd4, 0x47, 0xea, 0x9b, 0xf5, 0x35, 0xca, 0x0c, 0x9a,
	0x48, 0xdb, 0xbb, 0xf3, 0xe7, 0x9b, 0xef, 0x7c, 0x3f, 0x18, 0x30, 0xe2, 0x2c, 0x65, 0x21, 0xdd,
	0x8a, 0xfe, 0x3a, 0xe7, 0x82, 0xa3, 0x3f, 0xfb, 0xde, 0x7a, 0xd7, 0xa0, 0x81, 0xb7, 0x82, 0x66,
	0x1b, 0xc6, 0x33, 0x74, 0x0b, 0x46, 0x4e, 0x63, 0xb6, 0x11, 0x34, 0xa7, 0x91, 0xdc, 0x77, 0xb5,
	0x9e, 0x76, 0xd6, 0x74, 0x4e, 0xfa, 0x85, 0x01, 0x29, 0xf6, 0xc5, 0xb3, 0x71, 0x85, 0xb4, 0xf2,
	0xc3, 0x31, 0x9a, 0x00, 0x4a, 0xe7, 0xb2, 0x0d, 0xe7, 0xf9, 0x82, 0x89, 0x7c, 0x2e, 0x18, 0xcf,
	0xba, 0x55, 0xe5, 0x75, 0x54, 0x7a, 0x4d, 0x95, 0x66, 0x50, 0x4a, 0xc6, 0x15, 0xd2, 0x4e, 0xbf,
	0x0e, 0x87, 0xbf, 0xa0, 0x26, 0xa3, 0x7a, 0xf0, 0xef, 0x87, 0xe3, 0xe8, 0x14, 0xaa, 0x2c, 0x52,
	0x39, 0x0d, 0xa7, 0x53, 0x7a, 0x17, 0x02, 0x77, 0x44, 0xaa, 0x2c, 0x42, 0x26, 0xd4, 0xd2, 0x4d,
	0xac, 0x32, 0xfc, 0x25, 0xb2, 0xb4, 0x5e, 0xa0, 0xfd, 0x2d, 0x00, 0xb2, 0x40, 0xcf, 0x79, 0x42,
	0x77, 0xdc, 0xc6, 0x01, 0x37, 0x4f, 0x28, 0x51, 0x3b, 0xe4, 0x40, 0x93, 0x26, 0x74, 0x29, 0xf5,
	0x21, 0x8b, 0x76, 0x58, 0xed, 0x52, 0x1a, 0xb0, 0x4c, 0x5c, 0x39, 0x37, 0x04, 0xf6, 0x2a, 0x37,
	0xb2, 0x6c, 0xa8, 0xef, 0xc6, 0x08, 0x81, 0xbe, 0x62, 0xf1, 0x4a, 0x9d, 0xd0, 0x89, 0xaa, 0x65,
	0xba, 0x84, 0xbf, 0x29, 0x2b, 0x9d, 0xc8, 0xd2, 0xfa, 0x0f, 0xba, 0x3c, 0x89, 0x8c, 0x02, 0xaf,
	0x21, 0x39, 0x2e, 0x02, 0x68, 0x1e, 0xa0, 0xa1, 0x16, 0x34, 0xb0, 0x3b, 0x0a, 0x03, 0xef, 0x1e,
	0xfb, 0x66, 0x05, 0xf5, 0xe0, 0x58, 0xb6, 0x77, 0x81, 0xe7, 0xce, 0x30, 0x09, 0x7d, 0x3c, 0xc1,
	0x53, 0xec, 0x93, 0xc7, 0x70, 0x8c, 0x07, 0x23, 0x4c, 0x4c, 0x0d, 0x75, 0xc0, 0x94, 0x0a, 0xfc,
	0x30, 0xc3, 0xc4, 0x9d, 0x62, 0xcf, 0x1f, 0x4c, 0xcc, 0x8f, 0xfa, 0xf0, 0xf2, 0xe9, 0x3c, 0x66,
	0x62, 0xf5, 0xba, 0xe8, 0x2f, 0x79, 0x6a, 0xf3, 0x35, 0xcd, 0x96, 0x3c, 0x7b, 0x66, 0xb1, 0x2d,
	0xa9, 0x6c, 0xf5, 0x6b, 0xec, 0x3d, 0xe0, 0xe2, 0xb7, 0xea, 0xaf, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x34, 0x11, 0x32, 0x57, 0x02, 0x00, 0x00,
}
