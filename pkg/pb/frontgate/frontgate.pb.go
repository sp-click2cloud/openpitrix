// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/frontgate/frontgate.proto

package pbfrontgate // import "openpitrix.io/openpitrix/pkg/pb/frontgate"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "openpitrix.io/openpitrix/pkg/pb/types"

import "bufio"
import "crypto/tls"
import "errors"
import "io"
import "log"
import "net"
import "net/http"
import "net/rpc"
import "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Id                   string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ListenPort           int32                    `protobuf:"varint,2,opt,name=listen_port,json=listenPort" json:"listen_port,omitempty"`
	PilotHost            string                   `protobuf:"bytes,3,opt,name=pilot_host,json=pilotHost" json:"pilot_host,omitempty"`
	PilotPort            int32                    `protobuf:"varint,4,opt,name=pilot_port,json=pilotPort" json:"pilot_port,omitempty"`
	NodeList             *types.FrontgateEndpoint `protobuf:"bytes,5,opt,name=node_list,json=nodeList" json:"node_list,omitempty"`
	EtcdConfig           *types.EtcdConfig        `protobuf:"bytes,6,opt,name=etcd_config,json=etcdConfig" json:"etcd_config,omitempty"`
	ConfdConfig          *types.ConfdConfig       `protobuf:"bytes,7,opt,name=confd_config,json=confdConfig" json:"confd_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_ddd615a9b5015cbb, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Config) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *Config) GetPilotHost() string {
	if m != nil {
		return m.PilotHost
	}
	return ""
}

func (m *Config) GetPilotPort() int32 {
	if m != nil {
		return m.PilotPort
	}
	return 0
}

func (m *Config) GetNodeList() *types.FrontgateEndpoint {
	if m != nil {
		return m.NodeList
	}
	return nil
}

func (m *Config) GetEtcdConfig() *types.EtcdConfig {
	if m != nil {
		return m.EtcdConfig
	}
	return nil
}

func (m *Config) GetConfdConfig() *types.ConfdConfig {
	if m != nil {
		return m.ConfdConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "openpitrix.frontgate.Config")
}

type FrontgateService interface {
	GetPilotConfig(in *types.Empty, out *types.PilotConfig) error
	GetFrontgateConfig(in *types.Empty, out *types.FrontgateConfig) error
	SetFrontgateConfig(in *types.FrontgateConfig, out *types.Empty) error
	SetFrontgateNodeConfig(in *types.FrontgateConfig, out *types.Empty) error
	GetDroneList(in *types.Empty, out *types.DroneIdList) error
	GetDroneConfig(in *types.DroneEndpoint, out *types.DroneConfig) error
	SetDroneConfig(in *types.SetDroneConfigRequest, out *types.Empty) error
	GetConfdConfig(in *types.ConfdEndpoint, out *types.ConfdConfig) error
	IsConfdRunning(in *types.ConfdEndpoint, out *types.Bool) error
	StartConfd(in *types.ConfdEndpoint, out *types.Empty) error
	StopConfd(in *types.ConfdEndpoint, out *types.Empty) error
	RegisterMetadata(in *types.SubTask_RegisterMetadata, out *types.Empty) error
	DeregisterMetadata(in *types.SubTask_DeregisterMetadata, out *types.Empty) error
	RegisterCmd(in *types.SubTask_RegisterCmd, out *types.Empty) error
	DeregisterCmd(in *types.SubTask_DeregisterCmd, out *types.Empty) error
	ReportSubTaskStatus(in *types.SubTaskStatus, out *types.Empty) error
	GetEtcdValuesByPrefix(in *types.String, out *types.StringMap) error
	GetEtcdValues(in *types.StringList, out *types.StringMap) error
	SetEtcdValues(in *types.StringMap, out *types.Empty) error
	PingPilot(in *types.Empty, out *types.Empty) error
	PingFrontgate(in *types.Empty, out *types.Empty) error
	PingFrontgateNode(in *types.Empty, out *types.Empty) error
	PingDrone(in *types.DroneEndpoint, out *types.Empty) error
	RunCommand(in *types.String, out *types.String) error
	RunCommandOnDrone(in *types.DroneEndpoint, out *types.String) error
	HeartBeat(in *types.Empty, out *types.Empty) error
}

// AcceptFrontgateServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func AcceptFrontgateServiceClient(lis net.Listener, x FrontgateService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("openpitrix.frontgate.FrontgateService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// RegisterFrontgateService publish the given FrontgateService implementation on the server.
func RegisterFrontgateService(srv *rpc.Server, x FrontgateService) error {
	if err := srv.RegisterName("openpitrix.frontgate.FrontgateService", x); err != nil {
		return err
	}
	return nil
}

// NewFrontgateServiceServer returns a new FrontgateService Server.
func NewFrontgateServiceServer(x FrontgateService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("openpitrix.frontgate.FrontgateService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// ListenAndServeFrontgateService listen announces on the local network address laddr
// and serves the given FrontgateService implementation.
func ListenAndServeFrontgateService(network, addr string, x FrontgateService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("openpitrix.frontgate.FrontgateService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// ServeFrontgateService serves the given FrontgateService implementation.
func ServeFrontgateService(conn io.ReadWriteCloser, x FrontgateService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("openpitrix.frontgate.FrontgateService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeConn(conn)
}

type FrontgateServiceClient struct {
	*rpc.Client
}

// NewFrontgateServiceClient returns a FrontgateService stub to handle
// requests to the set of FrontgateService at the other end of the connection.
func NewFrontgateServiceClient(conn io.ReadWriteCloser) *FrontgateServiceClient {
	c := rpc.NewClient(conn)
	return &FrontgateServiceClient{c}
}

func (c *FrontgateServiceClient) GetPilotConfig(in *types.Empty) (out *types.PilotConfig, err error) {
	if in == nil {
		in = new(types.Empty)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.PilotConfig)
	if err = c.Call("openpitrix.frontgate.FrontgateService.GetPilotConfig", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetPilotConfig(in *types.Empty, out *types.PilotConfig, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.Empty)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.GetPilotConfig",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetFrontgateConfig(in *types.Empty) (out *types.FrontgateConfig, err error) {
	if in == nil {
		in = new(types.Empty)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.FrontgateConfig)
	if err = c.Call("openpitrix.frontgate.FrontgateService.GetFrontgateConfig", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetFrontgateConfig(in *types.Empty, out *types.FrontgateConfig, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.Empty)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.GetFrontgateConfig",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) SetFrontgateConfig(in *types.FrontgateConfig) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.FrontgateConfig)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.SetFrontgateConfig", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncSetFrontgateConfig(in *types.FrontgateConfig, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.FrontgateConfig)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.SetFrontgateConfig",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) SetFrontgateNodeConfig(in *types.FrontgateConfig) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.FrontgateConfig)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.SetFrontgateNodeConfig", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncSetFrontgateNodeConfig(in *types.FrontgateConfig, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.FrontgateConfig)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.SetFrontgateNodeConfig",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetDroneList(in *types.Empty) (out *types.DroneIdList, err error) {
	if in == nil {
		in = new(types.Empty)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.DroneIdList)
	if err = c.Call("openpitrix.frontgate.FrontgateService.GetDroneList", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetDroneList(in *types.Empty, out *types.DroneIdList, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.Empty)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.GetDroneList",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetDroneConfig(in *types.DroneEndpoint) (out *types.DroneConfig, err error) {
	if in == nil {
		in = new(types.DroneEndpoint)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.DroneConfig)
	if err = c.Call("openpitrix.frontgate.FrontgateService.GetDroneConfig", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetDroneConfig(in *types.DroneEndpoint, out *types.DroneConfig, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.DroneEndpoint)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.GetDroneConfig",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) SetDroneConfig(in *types.SetDroneConfigRequest) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.SetDroneConfigRequest)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.SetDroneConfig", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncSetDroneConfig(in *types.SetDroneConfigRequest, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.SetDroneConfigRequest)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.SetDroneConfig",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetConfdConfig(in *types.ConfdEndpoint) (out *types.ConfdConfig, err error) {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.ConfdConfig)
	if err = c.Call("openpitrix.frontgate.FrontgateService.GetConfdConfig", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetConfdConfig(in *types.ConfdEndpoint, out *types.ConfdConfig, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.GetConfdConfig",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) IsConfdRunning(in *types.ConfdEndpoint) (out *types.Bool, err error) {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Bool)
	if err = c.Call("openpitrix.frontgate.FrontgateService.IsConfdRunning", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncIsConfdRunning(in *types.ConfdEndpoint, out *types.Bool, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.IsConfdRunning",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) StartConfd(in *types.ConfdEndpoint) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.StartConfd", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncStartConfd(in *types.ConfdEndpoint, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.StartConfd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) StopConfd(in *types.ConfdEndpoint) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.StopConfd", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncStopConfd(in *types.ConfdEndpoint, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.ConfdEndpoint)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.StopConfd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RegisterMetadata(in *types.SubTask_RegisterMetadata) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.SubTask_RegisterMetadata)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.RegisterMetadata", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncRegisterMetadata(in *types.SubTask_RegisterMetadata, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.SubTask_RegisterMetadata)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.RegisterMetadata",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) DeregisterMetadata(in *types.SubTask_DeregisterMetadata) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.SubTask_DeregisterMetadata)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.DeregisterMetadata", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncDeregisterMetadata(in *types.SubTask_DeregisterMetadata, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.SubTask_DeregisterMetadata)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.DeregisterMetadata",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RegisterCmd(in *types.SubTask_RegisterCmd) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.SubTask_RegisterCmd)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.RegisterCmd", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncRegisterCmd(in *types.SubTask_RegisterCmd, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.SubTask_RegisterCmd)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.RegisterCmd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) DeregisterCmd(in *types.SubTask_DeregisterCmd) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.SubTask_DeregisterCmd)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.DeregisterCmd", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncDeregisterCmd(in *types.SubTask_DeregisterCmd, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.SubTask_DeregisterCmd)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.DeregisterCmd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) ReportSubTaskStatus(in *types.SubTaskStatus) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.SubTaskStatus)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.ReportSubTaskStatus", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncReportSubTaskStatus(in *types.SubTaskStatus, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.SubTaskStatus)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.ReportSubTaskStatus",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetEtcdValuesByPrefix(in *types.String) (out *types.StringMap, err error) {
	if in == nil {
		in = new(types.String)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.StringMap)
	if err = c.Call("openpitrix.frontgate.FrontgateService.GetEtcdValuesByPrefix", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetEtcdValuesByPrefix(in *types.String, out *types.StringMap, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.String)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.GetEtcdValuesByPrefix",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetEtcdValues(in *types.StringList) (out *types.StringMap, err error) {
	if in == nil {
		in = new(types.StringList)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.StringMap)
	if err = c.Call("openpitrix.frontgate.FrontgateService.GetEtcdValues", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetEtcdValues(in *types.StringList, out *types.StringMap, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.StringList)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.GetEtcdValues",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) SetEtcdValues(in *types.StringMap) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.StringMap)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.SetEtcdValues", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncSetEtcdValues(in *types.StringMap, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.StringMap)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.SetEtcdValues",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) PingPilot(in *types.Empty) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.Empty)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.PingPilot", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncPingPilot(in *types.Empty, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.Empty)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.PingPilot",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) PingFrontgate(in *types.Empty) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.Empty)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.PingFrontgate", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncPingFrontgate(in *types.Empty, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.Empty)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.PingFrontgate",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) PingFrontgateNode(in *types.Empty) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.Empty)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.PingFrontgateNode", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncPingFrontgateNode(in *types.Empty, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.Empty)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.PingFrontgateNode",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) PingDrone(in *types.DroneEndpoint) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.DroneEndpoint)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.PingDrone", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncPingDrone(in *types.DroneEndpoint, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.DroneEndpoint)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.PingDrone",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RunCommand(in *types.String) (out *types.String, err error) {
	if in == nil {
		in = new(types.String)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.String)
	if err = c.Call("openpitrix.frontgate.FrontgateService.RunCommand", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncRunCommand(in *types.String, out *types.String, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.String)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.RunCommand",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RunCommandOnDrone(in *types.DroneEndpoint) (out *types.String, err error) {
	if in == nil {
		in = new(types.DroneEndpoint)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.String)
	if err = c.Call("openpitrix.frontgate.FrontgateService.RunCommandOnDrone", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncRunCommandOnDrone(in *types.DroneEndpoint, out *types.String, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.DroneEndpoint)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.RunCommandOnDrone",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) HeartBeat(in *types.Empty) (out *types.Empty, err error) {
	if in == nil {
		in = new(types.Empty)
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new(types.Empty)
	if err = c.Call("openpitrix.frontgate.FrontgateService.HeartBeat", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *FrontgateServiceClient) AsyncHeartBeat(in *types.Empty, out *types.Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(types.Empty)
	}
	return c.Go(
		"openpitrix.frontgate.FrontgateService.HeartBeat",
		in, out,
		done,
	)
}

// DialFrontgateService connects to an FrontgateService at the specified network address.
func DialFrontgateService(network, addr string) (*FrontgateServiceClient, error) {
	c, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &FrontgateServiceClient{c}, nil
}

// DialFrontgateServiceTimeout connects to an FrontgateService at the specified network address.
func DialFrontgateServiceTimeout(network, addr string, timeout time.Duration) (*FrontgateServiceClient, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &FrontgateServiceClient{rpc.NewClient(conn)}, nil
}

// DialFrontgateServiceHTTP connects to an HTTP RPC server at the specified network address
// listening on the default HTTP RPC path.
func DialFrontgateServiceHTTP(network, address string) (*FrontgateServiceClient, error) {
	return DialFrontgateServiceHTTPPath(network, address, rpc.DefaultRPCPath)
}

// DialFrontgateServiceHTTPPath connects to an HTTP RPC server
// at the specified network address and path.
func DialFrontgateServiceHTTPPath(network, address, path string) (*FrontgateServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return dialFrontgateServicePath(network, address, path, conn)
}

// DialFrontgateServiceHTTPS connects to an HTTPS RPC server at the specified network address
// listening on the default HTTP RPC path.
func DialFrontgateServiceHTTPS(network, address string, tlsConfig *tls.Config) (*FrontgateServiceClient, error) {
	return DialFrontgateServiceHTTPSPath(network, address, rpc.DefaultRPCPath, tlsConfig)
}

// DialFrontgateServiceHTTPSPath connects to an HTTPS RPC server
// at the specified network address and path.
func DialFrontgateServiceHTTPSPath(network, address, path string, tlsConfig *tls.Config) (*FrontgateServiceClient, error) {
	conn, err := tls.Dial(network, address, tlsConfig)
	if err != nil {
		return nil, err
	}
	return dialFrontgateServicePath(network, address, path, conn)
}

func dialFrontgateServicePath(network, address, path string, conn net.Conn) (*FrontgateServiceClient, error) {
	const net_rpc_connected = "200 Connected to Go RPC"

	io.WriteString(conn, "CONNECT "+path+" HTTP/1.0\n\n")

	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "CONNECT"})
	if err == nil && resp.Status == net_rpc_connected {
		return &FrontgateServiceClient{rpc.NewClient(conn)}, nil
	}
	if err == nil {
		err = errors.New("unexpected HTTP response: " + resp.Status)
	}
	conn.Close()
	return nil, &net.OpError{
		Op:   "dial-http",
		Net:  network + " " + address,
		Addr: nil,
		Err:  err,
	}
}

func init() {
	proto.RegisterFile("openpitrix/frontgate/frontgate.proto", fileDescriptor_frontgate_ddd615a9b5015cbb)
}

var fileDescriptor_frontgate_ddd615a9b5015cbb = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x6f, 0x4f, 0x1a, 0x4b,
	0x14, 0xc6, 0x03, 0xf7, 0xea, 0xbd, 0x1c, 0x84, 0xab, 0x73, 0x5b, 0x4b, 0x50, 0x23, 0xfd, 0x17,
	0x4d, 0xd3, 0x68, 0x62, 0xd3, 0xf4, 0x45, 0x63, 0x62, 0x41, 0x44, 0x9b, 0xaa, 0x74, 0xd7, 0xb4,
	0x49, 0xfb, 0x82, 0x2c, 0xec, 0x40, 0x27, 0xc2, 0xcc, 0x74, 0xf6, 0xd0, 0xe8, 0x57, 0xeb, 0x77,
	0xe9, 0x77, 0x69, 0x66, 0x16, 0xd8, 0xc5, 0xdd, 0x01, 0x6b, 0xfb, 0x86, 0x2c, 0x73, 0xce, 0xf3,
	0x9b, 0xe7, 0xcc, 0xd9, 0x3d, 0x19, 0x78, 0x22, 0x24, 0xe5, 0x92, 0xa1, 0x62, 0x57, 0xbb, 0x5d,
	0x25, 0x38, 0xf6, 0x3c, 0xa4, 0xd1, 0xd3, 0x8e, 0x54, 0x02, 0x05, 0xb9, 0x17, 0x65, 0xed, 0x4c,
	0x62, 0xe5, 0xf5, 0x98, 0x16, 0xaf, 0x25, 0x0d, 0xc2, 0xdf, 0x50, 0x53, 0x5e, 0x4b, 0x44, 0x29,
	0x76, 0xfc, 0x51, 0x30, 0x29, 0xed, 0x08, 0xde, 0xb5, 0x47, 0x7d, 0x25, 0xf8, 0xc8, 0x4c, 0xb9,
	0x92, 0x88, 0xde, 0xb0, 0x9b, 0xa2, 0x97, 0xac, 0x2f, 0xd0, 0x6a, 0x0c, 0xbd, 0xe0, 0x32, 0x0c,
	0x3e, 0xfa, 0x9e, 0x85, 0xc5, 0x9a, 0xe0, 0x5d, 0xd6, 0x23, 0x45, 0xc8, 0x32, 0xbf, 0x94, 0xa9,
	0x64, 0xb6, 0x73, 0x4e, 0x96, 0xf9, 0x64, 0x13, 0xf2, 0x7d, 0x16, 0x20, 0xe5, 0x2d, 0x29, 0x14,
	0x96, 0xb2, 0x95, 0xcc, 0xf6, 0x82, 0x03, 0xe1, 0x52, 0x53, 0x28, 0x24, 0x1b, 0x00, 0x66, 0x9f,
	0xd6, 0x17, 0x11, 0x60, 0xe9, 0x2f, 0x23, 0xcc, 0x99, 0x95, 0x63, 0x11, 0xc4, 0xc2, 0x46, 0xfe,
	0xb7, 0x91, 0x87, 0x61, 0xa3, 0x3e, 0x80, 0x1c, 0x17, 0x3e, 0x6d, 0x69, 0x60, 0x69, 0xa1, 0x92,
	0xd9, 0xce, 0xef, 0x3d, 0xde, 0x89, 0x9d, 0x7b, 0x78, 0xb6, 0x47, 0xe3, 0x52, 0xeb, 0xdc, 0x97,
	0x82, 0x71, 0x74, 0xfe, 0xd5, 0xaa, 0x77, 0x2c, 0x40, 0xb2, 0x0f, 0x79, 0x7d, 0xc4, 0xad, 0x8e,
	0xf1, 0x5f, 0x5a, 0x34, 0x8c, 0xf5, 0x24, 0xa3, 0x8e, 0x1d, 0x3f, 0xac, 0xd1, 0x01, 0x3a, 0x79,
	0x26, 0x07, 0xb0, 0x64, 0x9a, 0x30, 0xd6, 0xff, 0x63, 0xf4, 0x1b, 0x49, 0xbd, 0xce, 0x1f, 0x03,
	0xf2, 0x9d, 0xe8, 0xcf, 0xde, 0x8f, 0xff, 0x60, 0x79, 0x62, 0xd0, 0xa5, 0xea, 0x1b, 0xeb, 0x50,
	0x72, 0x0c, 0xc5, 0x06, 0xc5, 0xa6, 0xae, 0x73, 0xb4, 0xd1, 0x83, 0x14, 0x4b, 0x03, 0x89, 0xd7,
	0xe5, 0x94, 0xbd, 0xe2, 0xba, 0x26, 0x90, 0x06, 0xc5, 0xc9, 0x06, 0xf3, 0x68, 0x0f, 0x67, 0x9c,
	0x5e, 0x44, 0x74, 0x93, 0xc4, 0xf9, 0xc2, 0xb2, 0x6d, 0x53, 0x72, 0x01, 0xab, 0x71, 0xe2, 0x99,
	0xf0, 0xff, 0x04, 0xf5, 0x08, 0x96, 0x1a, 0x14, 0x0f, 0xf5, 0x47, 0x60, 0x3a, 0xfd, 0x2b, 0x27,
	0x68, 0x54, 0x27, 0xbe, 0xd1, 0x35, 0x4d, 0x2f, 0xcc, 0xca, 0xc8, 0xd5, 0xa6, 0x45, 0x30, 0x7e,
	0xbd, 0xac, 0xc4, 0x91, 0xde, 0x81, 0xa2, 0x3b, 0x4d, 0xdc, 0x4a, 0x0a, 0xa6, 0x33, 0x1c, 0xfa,
	0x75, 0x48, 0x03, 0xb4, 0x57, 0x1b, 0xba, 0x8c, 0xbd, 0x65, 0x69, 0x2e, 0x4d, 0x78, 0x96, 0xcb,
	0xb8, 0xfe, 0x04, 0x8a, 0x27, 0x81, 0x59, 0x70, 0x86, 0x9c, 0x33, 0x7e, 0x0b, 0xe2, 0x6a, 0x32,
	0xa1, 0x2a, 0x44, 0x9f, 0x1c, 0x01, 0xb8, 0xe8, 0xa9, 0xd0, 0xde, 0x7c, 0x8c, 0xb5, 0xc8, 0x3a,
	0xe4, 0x5c, 0x14, 0xf2, 0x77, 0x31, 0x1f, 0x61, 0xd9, 0xa1, 0x3d, 0x3d, 0x83, 0xd4, 0x29, 0x45,
	0xcf, 0xf7, 0xd0, 0x23, 0xcf, 0x52, 0x3a, 0x30, 0x6c, 0x5f, 0x78, 0xc1, 0x65, 0xeb, 0x66, 0xae,
	0x1d, 0xfc, 0x19, 0xc8, 0x21, 0x55, 0x37, 0xd1, 0xcf, 0xed, 0xe8, 0x64, 0xb6, 0x1d, 0x7e, 0x0a,
	0xf9, 0xb1, 0x93, 0xda, 0xc0, 0x27, 0x4f, 0xe7, 0x1b, 0xae, 0x0d, 0x7c, 0x3b, 0xee, 0x3d, 0x14,
	0xa2, 0xdd, 0x35, 0x70, 0xeb, 0x36, 0x36, 0x67, 0x22, 0xcf, 0xe1, 0x7f, 0x87, 0xea, 0x41, 0x3d,
	0xd2, 0xb9, 0xe8, 0xe1, 0x30, 0x48, 0x6b, 0xd4, 0x54, 0x82, 0x1d, 0x78, 0x06, 0xf7, 0x1b, 0x14,
	0xf5, 0xe8, 0xfd, 0xe0, 0xf5, 0x87, 0x34, 0xa8, 0x5e, 0x37, 0x15, 0xed, 0xb2, 0x2b, 0x52, 0x4a,
	0x41, 0xa2, 0x62, 0xbc, 0x57, 0x5e, 0xb3, 0x45, 0x4e, 0x3d, 0x49, 0xde, 0x42, 0x61, 0x8a, 0x47,
	0xd6, 0x6d, 0xd9, 0xfa, 0xcb, 0x9f, 0xcd, 0xaa, 0x43, 0xc1, 0x9d, 0x62, 0xcd, 0xca, 0xb6, 0x97,
	0xb8, 0x0f, 0xb9, 0x26, 0xe3, 0x3d, 0x33, 0xb2, 0xed, 0x23, 0xca, 0x2a, 0x7f, 0x03, 0x05, 0x2d,
	0x9f, 0x0c, 0xc5, 0x3b, 0x20, 0xea, 0xb0, 0x32, 0x85, 0xd0, 0xe3, 0xf7, 0x4e, 0x18, 0x53, 0x88,
	0x99, 0x59, 0xf3, 0x27, 0xa4, 0x15, 0x73, 0x00, 0xe0, 0x0c, 0x79, 0x4d, 0x0c, 0x06, 0x1e, 0xf7,
	0x67, 0xf4, 0xd9, 0x1a, 0x21, 0x67, 0xb0, 0x12, 0x11, 0xce, 0xf9, 0x2d, 0x0d, 0xd9, 0x79, 0xfb,
	0x90, 0x3b, 0xa6, 0x9e, 0xc2, 0x2a, 0xf5, 0xee, 0xd0, 0xa1, 0xea, 0xab, 0x4f, 0x2f, 0x63, 0x11,
	0x26, 0x76, 0x63, 0x37, 0x29, 0x79, 0xd9, 0xdb, 0x95, 0xed, 0xe8, 0x2a, 0xf6, 0x5a, 0xb6, 0x27,
	0xcf, 0xed, 0x45, 0x73, 0xb9, 0x7a, 0xf1, 0x33, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xd0, 0xbe, 0x9a,
	0x6e, 0x0a, 0x00, 0x00,
}
