// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbs/configuration/configuration.proto

/*
Package configuration is a generated protocol buffer package.

It is generated from these files:
	pbs/configuration/configuration.proto

It has these top-level messages:
	Auth
	RequestByKey
	RequestByKeys
	Client
	Key
	Config
	Configs
	ConfigMeta
	ConfigMetas
	Empty
*/
package configuration

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_PULL   MessageType = 0
	MessageType_CHECK  MessageType = 1
	MessageType_ENSURE MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "PULL",
	1: "CHECK",
	2: "ENSURE",
}
var MessageType_value = map[string]int32{
	"PULL":   0,
	"CHECK":  1,
	"ENSURE": 2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Environment int32

const (
	Environment_LOCAL       Environment = 0
	Environment_DEVELOPMENT Environment = 1
	Environment_STAGING     Environment = 2
	Environment_PRODUCTION  Environment = 3
)

var Environment_name = map[int32]string{
	0: "LOCAL",
	1: "DEVELOPMENT",
	2: "STAGING",
	3: "PRODUCTION",
}
var Environment_value = map[string]int32{
	"LOCAL":       0,
	"DEVELOPMENT": 1,
	"STAGING":     2,
	"PRODUCTION":  3,
}

func (x Environment) String() string {
	return proto.EnumName(Environment_name, int32(x))
}
func (Environment) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Auth struct {
	Appid     string `protobuf:"bytes,1,opt,name=appid" json:"appid,omitempty"`
	Appsecret string `protobuf:"bytes,2,opt,name=appsecret" json:"appsecret,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Auth) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *Auth) GetAppsecret() string {
	if m != nil {
		return m.Appsecret
	}
	return ""
}

type RequestByKey struct {
	Key *Key        `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Env Environment `protobuf:"varint,2,opt,name=env,enum=configuration.Environment" json:"env,omitempty"`
}

func (m *RequestByKey) Reset()                    { *m = RequestByKey{} }
func (m *RequestByKey) String() string            { return proto.CompactTextString(m) }
func (*RequestByKey) ProtoMessage()               {}
func (*RequestByKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RequestByKey) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RequestByKey) GetEnv() Environment {
	if m != nil {
		return m.Env
	}
	return Environment_LOCAL
}

type RequestByKeys struct {
	Keys   []*Key      `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Type   MessageType `protobuf:"varint,2,opt,name=type,enum=configuration.MessageType" json:"type,omitempty"`
	Env    Environment `protobuf:"varint,3,opt,name=env,enum=configuration.Environment" json:"env,omitempty"`
	Client *Client     `protobuf:"bytes,4,opt,name=client" json:"client,omitempty"`
}

func (m *RequestByKeys) Reset()                    { *m = RequestByKeys{} }
func (m *RequestByKeys) String() string            { return proto.CompactTextString(m) }
func (*RequestByKeys) ProtoMessage()               {}
func (*RequestByKeys) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestByKeys) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *RequestByKeys) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_PULL
}

func (m *RequestByKeys) GetEnv() Environment {
	if m != nil {
		return m.Env
	}
	return Environment_LOCAL
}

func (m *RequestByKeys) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

type Client struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Main  string `protobuf:"bytes,2,opt,name=main" json:"main,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Client) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Client) GetMain() string {
	if m != nil {
		return m.Main
	}
	return ""
}

type Key struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Config struct {
	Config  string `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Config) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *Config) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type Configs struct {
	Configs []*Config `protobuf:"bytes,1,rep,name=configs" json:"configs,omitempty"`
}

func (m *Configs) Reset()                    { *m = Configs{} }
func (m *Configs) String() string            { return proto.CompactTextString(m) }
func (*Configs) ProtoMessage()               {}
func (*Configs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Configs) GetConfigs() []*Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

type ConfigMeta struct {
	Key     string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *ConfigMeta) Reset()                    { *m = ConfigMeta{} }
func (m *ConfigMeta) String() string            { return proto.CompactTextString(m) }
func (*ConfigMeta) ProtoMessage()               {}
func (*ConfigMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ConfigMeta) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigMeta) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ConfigMetas struct {
	Metas []*ConfigMeta `protobuf:"bytes,1,rep,name=metas" json:"metas,omitempty"`
}

func (m *ConfigMetas) Reset()                    { *m = ConfigMetas{} }
func (m *ConfigMetas) String() string            { return proto.CompactTextString(m) }
func (*ConfigMetas) ProtoMessage()               {}
func (*ConfigMetas) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ConfigMetas) GetMetas() []*ConfigMeta {
	if m != nil {
		return m.Metas
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*Auth)(nil), "configuration.Auth")
	proto.RegisterType((*RequestByKey)(nil), "configuration.RequestByKey")
	proto.RegisterType((*RequestByKeys)(nil), "configuration.RequestByKeys")
	proto.RegisterType((*Client)(nil), "configuration.Client")
	proto.RegisterType((*Key)(nil), "configuration.Key")
	proto.RegisterType((*Config)(nil), "configuration.Config")
	proto.RegisterType((*Configs)(nil), "configuration.Configs")
	proto.RegisterType((*ConfigMeta)(nil), "configuration.ConfigMeta")
	proto.RegisterType((*ConfigMetas)(nil), "configuration.ConfigMetas")
	proto.RegisterType((*Empty)(nil), "configuration.Empty")
	proto.RegisterEnum("configuration.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("configuration.Environment", Environment_name, Environment_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Configuration service

type ConfigurationClient interface {
	Signin(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Client, error)
	Get(ctx context.Context, in *RequestByKey, opts ...grpc.CallOption) (*Config, error)
	Pull(ctx context.Context, in *RequestByKeys, opts ...grpc.CallOption) (*Configs, error)
	Check(ctx context.Context, in *RequestByKeys, opts ...grpc.CallOption) (*ConfigMetas, error)
	Ensure(ctx context.Context, in *RequestByKeys, opts ...grpc.CallOption) (*Empty, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) Signin(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := grpc.Invoke(ctx, "/configuration.Configuration/Signin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Get(ctx context.Context, in *RequestByKey, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := grpc.Invoke(ctx, "/configuration.Configuration/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Pull(ctx context.Context, in *RequestByKeys, opts ...grpc.CallOption) (*Configs, error) {
	out := new(Configs)
	err := grpc.Invoke(ctx, "/configuration.Configuration/Pull", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Check(ctx context.Context, in *RequestByKeys, opts ...grpc.CallOption) (*ConfigMetas, error) {
	out := new(ConfigMetas)
	err := grpc.Invoke(ctx, "/configuration.Configuration/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Ensure(ctx context.Context, in *RequestByKeys, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/configuration.Configuration/Ensure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Configuration service

type ConfigurationServer interface {
	Signin(context.Context, *Auth) (*Client, error)
	Get(context.Context, *RequestByKey) (*Config, error)
	Pull(context.Context, *RequestByKeys) (*Configs, error)
	Check(context.Context, *RequestByKeys) (*ConfigMetas, error)
	Ensure(context.Context, *RequestByKeys) (*Empty, error)
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configuration.Configuration/Signin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Signin(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configuration.Configuration/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Get(ctx, req.(*RequestByKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configuration.Configuration/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Pull(ctx, req.(*RequestByKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configuration.Configuration/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Check(ctx, req.(*RequestByKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Ensure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Ensure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configuration.Configuration/Ensure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Ensure(ctx, req.(*RequestByKeys))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "configuration.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signin",
			Handler:    _Configuration_Signin_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Configuration_Get_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _Configuration_Pull_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Configuration_Check_Handler,
		},
		{
			MethodName: "Ensure",
			Handler:    _Configuration_Ensure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbs/configuration/configuration.proto",
}

func init() { proto.RegisterFile("pbs/configuration/configuration.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0xc5, 0xd8, 0x98, 0x8f, 0xf1, 0x47, 0x6a, 0x4d, 0xd3, 0x94, 0xd2, 0x5c, 0x44, 0xab, 0xb6,
	0xaa, 0xa2, 0x96, 0x48, 0x54, 0xaa, 0x2a, 0x2a, 0x45, 0xa2, 0x8e, 0x4b, 0x23, 0x7e, 0x65, 0xa0,
	0xf7, 0x86, 0x6e, 0x89, 0x05, 0xac, 0x5d, 0xef, 0x82, 0xe4, 0xf7, 0xeb, 0x4d, 0xdf, 0xaa, 0xda,
	0x35, 0x88, 0x9f, 0x40, 0xd2, 0x3b, 0xcf, 0xcc, 0x99, 0x39, 0x67, 0xcf, 0xec, 0x1a, 0x5e, 0x47,
	0x23, 0x7e, 0x35, 0x0e, 0xd9, 0xcf, 0x60, 0xb2, 0x88, 0x7d, 0x11, 0x84, 0x6c, 0x37, 0xaa, 0x44,
	0x71, 0x28, 0x42, 0x2c, 0xee, 0x24, 0x49, 0x0d, 0x8c, 0xfa, 0x42, 0xdc, 0xe1, 0x29, 0xe4, 0xfc,
	0x28, 0x0a, 0x7e, 0x94, 0xb4, 0x0b, 0xed, 0x6d, 0xc1, 0x4b, 0x03, 0x3c, 0x87, 0x82, 0x1f, 0x45,
	0x9c, 0x8e, 0x63, 0x2a, 0x4a, 0x59, 0x55, 0xd9, 0x24, 0xc8, 0x08, 0xfe, 0xf7, 0xe8, 0xaf, 0x05,
	0xe5, 0xe2, 0x4b, 0xd2, 0xa4, 0x09, 0xbe, 0x02, 0x7d, 0x4a, 0x13, 0x35, 0xc1, 0xaa, 0x62, 0x65,
	0x97, 0xbd, 0x49, 0x13, 0x4f, 0x96, 0xf1, 0x1d, 0xe8, 0x94, 0x2d, 0xd5, 0xb4, 0x93, 0x6a, 0x79,
	0x0f, 0xe5, 0xb2, 0x65, 0x10, 0x87, 0x6c, 0x4e, 0x99, 0xf0, 0x24, 0x8c, 0xfc, 0xd6, 0xa0, 0xb8,
	0x4d, 0xc2, 0xf1, 0x0d, 0x18, 0x53, 0x9a, 0xf0, 0x92, 0x76, 0xa1, 0x1f, 0xa1, 0x51, 0x75, 0xac,
	0x80, 0x21, 0x92, 0x88, 0x1e, 0x21, 0x6a, 0x53, 0xce, 0xfd, 0x09, 0x1d, 0x24, 0x11, 0xf5, 0x14,
	0x6e, 0xad, 0x4b, 0xff, 0x27, 0x5d, 0xf8, 0x1e, 0xcc, 0xf1, 0x2c, 0xa0, 0x4c, 0x94, 0x0c, 0x75,
	0xdc, 0x67, 0x7b, 0x0d, 0x8e, 0x2a, 0x7a, 0x2b, 0x10, 0xa9, 0x82, 0x99, 0x66, 0xa4, 0xd1, 0x22,
	0x9c, 0x52, 0xb6, 0x36, 0x5a, 0x05, 0x88, 0x60, 0xcc, 0xfd, 0x80, 0xad, 0x3c, 0x56, 0xdf, 0xe4,
	0x39, 0xe8, 0xd2, 0x55, 0x7b, 0xe3, 0x6a, 0x41, 0x39, 0x48, 0x6a, 0x60, 0x3a, 0x8a, 0x0c, 0xcf,
	0xc0, 0x4c, 0x69, 0x57, 0xe5, 0x55, 0x84, 0x25, 0xc8, 0x2f, 0x69, 0xcc, 0x83, 0x30, 0x9d, 0xa8,
	0x7b, 0xeb, 0x90, 0xd4, 0x20, 0x9f, 0xf6, 0x72, 0xbc, 0x82, 0x7c, 0x0a, 0x5f, 0x7b, 0x79, 0xef,
	0x0c, 0x2a, 0xf2, 0xd6, 0x28, 0xf2, 0x09, 0x20, 0x4d, 0xb5, 0xa9, 0xf0, 0xef, 0xeb, 0x7a, 0x80,
	0xf5, 0x1a, 0xac, 0x4d, 0xa7, 0x64, 0xce, 0xcd, 0xe5, 0xc7, 0x8a, 0xf7, 0xc5, 0x41, 0x5e, 0x09,
	0xf5, 0x52, 0x1c, 0xc9, 0x43, 0xce, 0x9d, 0x47, 0x22, 0xb9, 0xac, 0x80, 0xb5, 0xb5, 0x39, 0xfc,
	0x0f, 0x8c, 0xde, 0xb0, 0xd5, 0xb2, 0x33, 0x58, 0x80, 0x9c, 0xf3, 0xcd, 0x75, 0x9a, 0xb6, 0x86,
	0x00, 0xa6, 0xdb, 0xe9, 0x0f, 0x3d, 0xd7, 0xce, 0x5e, 0x7e, 0x05, 0x6b, 0x6b, 0x75, 0x12, 0xd5,
	0xea, 0x3a, 0x75, 0xd9, 0xf0, 0x04, 0xac, 0x1b, 0xf7, 0xbb, 0xdb, 0xea, 0xf6, 0xda, 0x6e, 0x67,
	0x60, 0x6b, 0x68, 0x41, 0xbe, 0x3f, 0xa8, 0x37, 0x6e, 0x3b, 0x0d, 0x3b, 0x8b, 0x27, 0x00, 0x3d,
	0xaf, 0x7b, 0x33, 0x74, 0x06, 0xb7, 0xdd, 0x8e, 0xad, 0x57, 0xff, 0x64, 0xa1, 0xe8, 0x6c, 0x8b,
	0xc4, 0x8f, 0x60, 0xf6, 0x83, 0x09, 0x0b, 0x18, 0x3e, 0xdd, 0x93, 0x2f, 0xdf, 0x53, 0xf9, 0xf0,
	0x7d, 0x20, 0x19, 0xfc, 0x0c, 0x7a, 0x83, 0x0a, 0x7c, 0xb9, 0x57, 0xdf, 0xbe, 0xe3, 0xe5, 0xc3,
	0x8b, 0x20, 0x19, 0xbc, 0x06, 0xa3, 0xb7, 0x98, 0xcd, 0xf0, 0xfc, 0x81, 0x6e, 0x5e, 0x3e, 0x3b,
	0xd8, 0xce, 0x49, 0x06, 0x1d, 0xc8, 0x39, 0x77, 0x74, 0x3c, 0x7d, 0x64, 0x40, 0xf9, 0xe8, 0x42,
	0xb8, 0x12, 0x61, 0xba, 0x8c, 0x2f, 0x62, 0xfa, 0xc8, 0x94, 0xd3, 0xfd, 0x37, 0x24, 0x37, 0x48,
	0x32, 0x23, 0x53, 0xfd, 0x88, 0x3e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xe0, 0x48, 0x6a,
	0xb1, 0x04, 0x00, 0x00,
}
