// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/types/frontgate.proto

package pbtypes // import "openpitrix.io/openpitrix/pkg/pb/types"

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

type FrontgateId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload" json:"payload"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateId) Reset()         { *m = FrontgateId{} }
func (m *FrontgateId) String() string { return proto.CompactTextString(m) }
func (*FrontgateId) ProtoMessage()    {}
func (*FrontgateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_eee64f93028b97e1, []int{0}
}
func (m *FrontgateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateId.Unmarshal(m, b)
}
func (m *FrontgateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateId.Marshal(b, m, deterministic)
}
func (dst *FrontgateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateId.Merge(dst, src)
}
func (m *FrontgateId) XXX_Size() int {
	return xxx_messageInfo_FrontgateId.Size(m)
}
func (m *FrontgateId) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateId.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateId proto.InternalMessageInfo

func (m *FrontgateId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrontgateId) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type FrontgateNodeId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id"`
	NodeId               string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload" json:"payload"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateNodeId) Reset()         { *m = FrontgateNodeId{} }
func (m *FrontgateNodeId) String() string { return proto.CompactTextString(m) }
func (*FrontgateNodeId) ProtoMessage()    {}
func (*FrontgateNodeId) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_eee64f93028b97e1, []int{1}
}
func (m *FrontgateNodeId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateNodeId.Unmarshal(m, b)
}
func (m *FrontgateNodeId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateNodeId.Marshal(b, m, deterministic)
}
func (dst *FrontgateNodeId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateNodeId.Merge(dst, src)
}
func (m *FrontgateNodeId) XXX_Size() int {
	return xxx_messageInfo_FrontgateNodeId.Size(m)
}
func (m *FrontgateNodeId) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateNodeId.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateNodeId proto.InternalMessageInfo

func (m *FrontgateNodeId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrontgateNodeId) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *FrontgateNodeId) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type FrontgateIdList struct {
	IdList               []string `protobuf:"bytes,1,rep,name=id_list,json=idList" json:"id_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateIdList) Reset()         { *m = FrontgateIdList{} }
func (m *FrontgateIdList) String() string { return proto.CompactTextString(m) }
func (*FrontgateIdList) ProtoMessage()    {}
func (*FrontgateIdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_eee64f93028b97e1, []int{2}
}
func (m *FrontgateIdList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateIdList.Unmarshal(m, b)
}
func (m *FrontgateIdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateIdList.Marshal(b, m, deterministic)
}
func (dst *FrontgateIdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateIdList.Merge(dst, src)
}
func (m *FrontgateIdList) XXX_Size() int {
	return xxx_messageInfo_FrontgateIdList.Size(m)
}
func (m *FrontgateIdList) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateIdList.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateIdList proto.InternalMessageInfo

func (m *FrontgateIdList) GetIdList() []string {
	if m != nil {
		return m.IdList
	}
	return nil
}

type FrontgateConfig struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id" json:"id"`
	NodeId               string               `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id"`
	Host                 string               `protobuf:"bytes,3,opt,name=host" json:"host"`
	ListenPort           int32                `protobuf:"varint,4,opt,name=listen_port,json=listenPort" json:"listen_port"`
	PilotHost            string               `protobuf:"bytes,5,opt,name=pilot_host,json=pilotHost" json:"pilot_host"`
	PilotPort            int32                `protobuf:"varint,6,opt,name=pilot_port,json=pilotPort" json:"pilot_port"`
	NodeList             []*FrontgateEndpoint `protobuf:"bytes,7,rep,name=node_list,json=nodeList" json:"node_list"`
	EtcdConfig           *EtcdConfig          `protobuf:"bytes,8,opt,name=etcd_config,json=etcdConfig" json:"etcd_config"`
	ConfdConfig          *ConfdConfig         `protobuf:"bytes,9,opt,name=confd_config,json=confdConfig" json:"confd_config"`
	LogLevel             string               `protobuf:"bytes,10,opt,name=log_level,json=logLevel" json:"log_level"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrontgateConfig) Reset()         { *m = FrontgateConfig{} }
func (m *FrontgateConfig) String() string { return proto.CompactTextString(m) }
func (*FrontgateConfig) ProtoMessage()    {}
func (*FrontgateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_eee64f93028b97e1, []int{3}
}
func (m *FrontgateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateConfig.Unmarshal(m, b)
}
func (m *FrontgateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateConfig.Marshal(b, m, deterministic)
}
func (dst *FrontgateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateConfig.Merge(dst, src)
}
func (m *FrontgateConfig) XXX_Size() int {
	return xxx_messageInfo_FrontgateConfig.Size(m)
}
func (m *FrontgateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateConfig proto.InternalMessageInfo

func (m *FrontgateConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrontgateConfig) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *FrontgateConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *FrontgateConfig) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *FrontgateConfig) GetPilotHost() string {
	if m != nil {
		return m.PilotHost
	}
	return ""
}

func (m *FrontgateConfig) GetPilotPort() int32 {
	if m != nil {
		return m.PilotPort
	}
	return 0
}

func (m *FrontgateConfig) GetNodeList() []*FrontgateEndpoint {
	if m != nil {
		return m.NodeList
	}
	return nil
}

func (m *FrontgateConfig) GetEtcdConfig() *EtcdConfig {
	if m != nil {
		return m.EtcdConfig
	}
	return nil
}

func (m *FrontgateConfig) GetConfdConfig() *ConfdConfig {
	if m != nil {
		return m.ConfdConfig
	}
	return nil
}

func (m *FrontgateConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type FrontgateEndpoint struct {
	FrontgateId          string   `protobuf:"bytes,1,opt,name=frontgate_id,json=frontgateId" json:"frontgate_id"`
	NodeIp               string   `protobuf:"bytes,2,opt,name=node_ip,json=nodeIp" json:"node_ip"`
	NodePort             int32    `protobuf:"varint,3,opt,name=node_port,json=nodePort" json:"node_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateEndpoint) Reset()         { *m = FrontgateEndpoint{} }
func (m *FrontgateEndpoint) String() string { return proto.CompactTextString(m) }
func (*FrontgateEndpoint) ProtoMessage()    {}
func (*FrontgateEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_eee64f93028b97e1, []int{4}
}
func (m *FrontgateEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateEndpoint.Unmarshal(m, b)
}
func (m *FrontgateEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateEndpoint.Marshal(b, m, deterministic)
}
func (dst *FrontgateEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateEndpoint.Merge(dst, src)
}
func (m *FrontgateEndpoint) XXX_Size() int {
	return xxx_messageInfo_FrontgateEndpoint.Size(m)
}
func (m *FrontgateEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateEndpoint proto.InternalMessageInfo

func (m *FrontgateEndpoint) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *FrontgateEndpoint) GetNodeIp() string {
	if m != nil {
		return m.NodeIp
	}
	return ""
}

func (m *FrontgateEndpoint) GetNodePort() int32 {
	if m != nil {
		return m.NodePort
	}
	return 0
}

func init() {
	proto.RegisterType((*FrontgateId)(nil), "openpitrix.types.FrontgateId")
	proto.RegisterType((*FrontgateNodeId)(nil), "openpitrix.types.FrontgateNodeId")
	proto.RegisterType((*FrontgateIdList)(nil), "openpitrix.types.FrontgateIdList")
	proto.RegisterType((*FrontgateConfig)(nil), "openpitrix.types.FrontgateConfig")
	proto.RegisterType((*FrontgateEndpoint)(nil), "openpitrix.types.FrontgateEndpoint")
}

func init() {
	proto.RegisterFile("openpitrix/types/frontgate.proto", fileDescriptor_frontgate_eee64f93028b97e1)
}

var fileDescriptor_frontgate_eee64f93028b97e1 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x86, 0xc9, 0x64, 0x77, 0x66, 0x52, 0x59, 0xfc, 0xe8, 0x8b, 0x61, 0x3f, 0x30, 0xc6, 0x4b,
	0x10, 0x4c, 0x60, 0x3d, 0x78, 0x10, 0x61, 0x51, 0x56, 0x5c, 0x58, 0x44, 0x82, 0x27, 0x2f, 0x61,
	0x26, 0xdd, 0x89, 0x8d, 0x4d, 0x57, 0x93, 0x34, 0xe2, 0xfc, 0x08, 0xff, 0xb3, 0x74, 0x65, 0x92,
	0xc9, 0x18, 0x0f, 0x7b, 0x4a, 0xba, 0xab, 0x9e, 0xb7, 0xab, 0xea, 0x2d, 0x88, 0xd1, 0x08, 0x6d,
	0xa4, 0x6d, 0xe5, 0xef, 0xdc, 0xee, 0x8c, 0xe8, 0xf2, 0xba, 0x45, 0x6d, 0x9b, 0x8d, 0x15, 0x99,
	0x69, 0xd1, 0x22, 0x7b, 0x72, 0xc8, 0xc8, 0x28, 0xe3, 0xfc, 0x62, 0xc6, 0x08, 0x5b, 0xf1, 0x3e,
	0xfd, 0xfc, 0x72, 0x16, 0xac, 0x50, 0xd7, 0xfb, 0x68, 0xf2, 0x16, 0xc2, 0x4f, 0x83, 0xfe, 0x1d,
	0x67, 0x8f, 0x60, 0x21, 0x79, 0xe4, 0xc5, 0x5e, 0x1a, 0x14, 0x0b, 0xc9, 0x59, 0x04, 0x2b, 0xb3,
	0xd9, 0x29, 0xdc, 0xf0, 0x68, 0x41, 0x97, 0xc3, 0x31, 0xf9, 0x06, 0x8f, 0x47, 0xf0, 0x0b, 0xf2,
	0xff, 0xc1, 0xcf, 0x60, 0xa5, 0x91, 0x8b, 0x52, 0x0e, 0xf0, 0x52, 0xf7, 0x89, 0x13, 0x55, 0xff,
	0x58, 0xf5, 0xd5, 0x44, 0xf5, 0x8e, 0xdf, 0xcb, 0xce, 0x3a, 0x15, 0xc9, 0x4b, 0x25, 0x3b, 0x1b,
	0x79, 0xb1, 0xef, 0x54, 0x24, 0x05, 0x92, 0x3f, 0xfe, 0x24, 0xf9, 0x23, 0xea, 0x5a, 0x36, 0x0f,
	0x2f, 0x81, 0xc1, 0xc9, 0x0f, 0xec, 0xec, 0xfe, 0x7d, 0xfa, 0x67, 0xcf, 0x21, 0x74, 0xcf, 0x08,
	0x5d, 0x1a, 0x6c, 0x6d, 0x74, 0x12, 0x7b, 0xe9, 0x69, 0x01, 0xfd, 0xd5, 0x57, 0x6c, 0x2d, 0xbb,
	0x02, 0x30, 0x52, 0xa1, 0x2d, 0x09, 0x3d, 0x25, 0x34, 0xa0, 0x9b, 0xcf, 0x8e, 0x1f, 0xc3, 0x84,
	0x2f, 0x09, 0xef, 0xc3, 0x44, 0xdf, 0x40, 0x40, 0xb5, 0x50, 0x2b, 0xab, 0xd8, 0x4f, 0xc3, 0xeb,
	0x97, 0xd9, 0xbf, 0x5e, 0x66, 0x63, 0x47, 0xb7, 0x9a, 0x1b, 0x94, 0xda, 0x16, 0x6b, 0x47, 0xd1,
	0x28, 0xde, 0x43, 0xe8, 0x8c, 0x2d, 0x2b, 0x6a, 0x36, 0x5a, 0xc7, 0x5e, 0x1a, 0x5e, 0x5f, 0xce,
	0x35, 0x6e, 0x6d, 0xc5, 0xfb, 0x81, 0x14, 0x20, 0xc6, 0x7f, 0x76, 0x03, 0x67, 0x64, 0xfd, 0xc0,
	0x07, 0xc4, 0x5f, 0xcd, 0x79, 0x97, 0x3f, 0x08, 0x84, 0xd5, 0xe1, 0xc0, 0x2e, 0x20, 0x50, 0xd8,
	0x94, 0x4a, 0xfc, 0x12, 0x2a, 0x02, 0xea, 0x7f, 0xad, 0xb0, 0xb9, 0x77, 0xe7, 0x44, 0xc1, 0xd3,
	0x59, 0xf1, 0xec, 0x05, 0x9c, 0x8d, 0xfb, 0x5b, 0x8e, 0xd6, 0x84, 0xf5, 0x64, 0xe7, 0x46, 0x8f,
	0xcc, 0x91, 0x47, 0xc6, 0xbd, 0x46, 0x01, 0x1a, 0xa7, 0x4f, 0xe3, 0xa4, 0x59, 0xb8, 0x69, 0x7e,
	0xc8, 0xbf, 0xbf, 0x9e, 0xd4, 0x2d, 0x31, 0x9f, 0xac, 0xb9, 0xf9, 0xd9, 0xe4, 0x66, 0xdb, 0x6f,
	0xfb, 0x3b, 0xb3, 0xa5, 0xef, 0x76, 0x49, 0x0b, 0xff, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7c, 0x85, 0xf0, 0xa8, 0x61, 0x03, 0x00, 0x00,
}
