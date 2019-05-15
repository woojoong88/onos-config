// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/northbound/proto/diags.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ChangesRequest is a message for specifying GetChanges query parameters.
type ChangesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangesRequest) Reset()         { *m = ChangesRequest{} }
func (m *ChangesRequest) String() string { return proto.CompactTextString(m) }
func (*ChangesRequest) ProtoMessage()    {}
func (*ChangesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{0}
}

func (m *ChangesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangesRequest.Unmarshal(m, b)
}
func (m *ChangesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangesRequest.Marshal(b, m, deterministic)
}
func (m *ChangesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangesRequest.Merge(m, src)
}
func (m *ChangesRequest) XXX_Size() int {
	return xxx_messageInfo_ChangesRequest.Size(m)
}
func (m *ChangesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangesRequest proto.InternalMessageInfo

// ConfigRequest is a message for specifying GetConfigurations query parameters.
type ConfigRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{1}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

// ChangeValue is an individual Path/Value combination in a Change
type ChangeValue struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Removed              bool     `protobuf:"varint,3,opt,name=removed,proto3" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeValue) Reset()         { *m = ChangeValue{} }
func (m *ChangeValue) String() string { return proto.CompactTextString(m) }
func (*ChangeValue) ProtoMessage()    {}
func (*ChangeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{2}
}

func (m *ChangeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeValue.Unmarshal(m, b)
}
func (m *ChangeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeValue.Marshal(b, m, deterministic)
}
func (m *ChangeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeValue.Merge(m, src)
}
func (m *ChangeValue) XXX_Size() int {
	return xxx_messageInfo_ChangeValue.Size(m)
}
func (m *ChangeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeValue.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeValue proto.InternalMessageInfo

func (m *ChangeValue) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChangeValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ChangeValue) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

// Change is a descriptor of a submitted configuration change targeted at a single device.
type Change struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Desc                 string               `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Changevalues         []*ChangeValue       `protobuf:"bytes,4,rep,name=changevalues,proto3" json:"changevalues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{3}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Change) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Change) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Change) GetChangevalues() []*ChangeValue {
	if m != nil {
		return m.Changevalues
	}
	return nil
}

// Change is a descriptor of a submitted configuration change targeted at a single device.
type Configuration struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Deviceid             string               `protobuf:"bytes,2,opt,name=deviceid,proto3" json:"deviceid,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	User                 string               `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Desc                 string               `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
	ChangeIDs            []string             `protobuf:"bytes,7,rep,name=changeIDs,proto3" json:"changeIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{4}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Configuration) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

func (m *Configuration) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Configuration) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *Configuration) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Configuration) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Configuration) GetChangeIDs() []string {
	if m != nil {
		return m.ChangeIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangesRequest)(nil), "proto.ChangesRequest")
	proto.RegisterType((*ConfigRequest)(nil), "proto.ConfigRequest")
	proto.RegisterType((*ChangeValue)(nil), "proto.ChangeValue")
	proto.RegisterType((*Change)(nil), "proto.Change")
	proto.RegisterType((*Configuration)(nil), "proto.Configuration")
}

func init() { proto.RegisterFile("pkg/northbound/proto/diags.proto", fileDescriptor_c7617167392ea9d7) }

var fileDescriptor_c7617167392ea9d7 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xce, 0x9b, 0x30,
	0x10, 0x94, 0x13, 0xf2, 0xc3, 0xd2, 0xa4, 0xad, 0x95, 0x4a, 0x08, 0x55, 0x2a, 0xe2, 0xc4, 0x09,
	0xa2, 0x34, 0xea, 0xbd, 0x4a, 0xa4, 0xaa, 0xc7, 0x5a, 0x55, 0xef, 0x0e, 0x76, 0x08, 0x6a, 0x82,
	0x29, 0x36, 0x79, 0x84, 0x3e, 0x41, 0x5f, 0xb4, 0x6f, 0xf0, 0xc9, 0x36, 0xce, 0x07, 0x97, 0xef,
	0x3b, 0xb1, 0x3b, 0xcc, 0x8e, 0x67, 0x56, 0x0b, 0x71, 0xf3, 0xbb, 0xcc, 0x6b, 0xd1, 0xaa, 0xcb,
	0x49, 0x74, 0x35, 0xcb, 0x9b, 0x56, 0x28, 0x91, 0xb3, 0x8a, 0x96, 0x32, 0x33, 0x35, 0x9e, 0x99,
	0x4f, 0xf4, 0xa9, 0x14, 0xa2, 0xbc, 0x72, 0x4b, 0x38, 0x75, 0xe7, 0x5c, 0x55, 0x37, 0x2e, 0x15,
	0xbd, 0x35, 0x96, 0x97, 0xbc, 0x83, 0xf5, 0xe1, 0x42, 0xeb, 0x92, 0x4b, 0xc2, 0xff, 0x74, 0x5c,
	0xaa, 0xe4, 0x2d, 0xac, 0x0e, 0xa2, 0x3e, 0x57, 0xa5, 0x03, 0x7e, 0x40, 0x60, 0x29, 0xbf, 0xe8,
	0xb5, 0xe3, 0x18, 0x83, 0xd7, 0x50, 0x75, 0x09, 0x51, 0x8c, 0x52, 0x9f, 0x98, 0x1a, 0x6f, 0x60,
	0x76, 0xd7, 0x3f, 0xc3, 0x89, 0x01, 0x6d, 0x83, 0x43, 0x58, 0xb4, 0xfc, 0x26, 0xee, 0x9c, 0x85,
	0xd3, 0x18, 0xa5, 0x4b, 0xe2, 0xda, 0xe4, 0x1f, 0x82, 0xb9, 0xd5, 0xc4, 0x19, 0x78, 0xda, 0x93,
	0x91, 0x0b, 0x76, 0x51, 0x66, 0x0d, 0x67, 0xce, 0x70, 0xf6, 0xd3, 0x19, 0x26, 0x86, 0x87, 0xd7,
	0x30, 0xa9, 0x58, 0xff, 0xce, 0xa4, 0x62, 0xda, 0x0e, 0xe3, 0xb2, 0x30, 0x2f, 0xf8, 0xc4, 0xd4,
	0xf8, 0x0b, 0xbc, 0x29, 0x8c, 0xba, 0xf1, 0x21, 0x43, 0x2f, 0x9e, 0xa6, 0xc1, 0x0e, 0x5b, 0xd1,
	0x6c, 0x10, 0x86, 0x8c, 0x78, 0xc9, 0x7f, 0xe4, 0xb2, 0x77, 0x2d, 0x55, 0x95, 0xa8, 0xb5, 0x7a,
	0x4d, 0x7b, 0x77, 0x3e, 0x31, 0x35, 0x8e, 0x60, 0xc9, 0xf8, 0xbd, 0x2a, 0xf8, 0xc3, 0xc7, 0xa3,
	0xc7, 0x7b, 0x58, 0x14, 0x2d, 0xa7, 0xaa, 0x8f, 0xfc, 0x72, 0x20, 0x47, 0xd5, 0x53, 0x5d, 0xc3,
	0xcc, 0x94, 0xf7, 0xfa, 0x54, 0x4f, 0xd5, 0xde, 0x3a, 0xc9, 0xdb, 0x70, 0x66, 0xbd, 0xe9, 0xfa,
	0xb1, 0x8d, 0xf9, 0x60, 0x1b, 0x1f, 0xc1, 0xb7, 0x29, 0xbf, 0x1f, 0x65, 0xb8, 0x88, 0xa7, 0xa9,
	0x4f, 0x9e, 0x81, 0xdd, 0x5f, 0x04, 0x81, 0xcd, 0x7c, 0xd4, 0xe7, 0x83, 0xf7, 0x00, 0xdf, 0xb8,
	0xea, 0x6f, 0x02, 0x7f, 0x18, 0xed, 0xcc, 0xdd, 0x48, 0xb4, 0x1a, 0xc1, 0x5b, 0x84, 0xbf, 0xc2,
	0x7b, 0x3d, 0x35, 0xdc, 0x9d, 0xc4, 0x1b, 0xc7, 0x1a, 0x9e, 0x53, 0x34, 0x46, 0x7b, 0xf2, 0x16,
	0x9d, 0xe6, 0x06, 0xfe, 0xfc, 0x14, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xbd, 0x74, 0x07, 0xdc, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigDiagsClient is the client API for ConfigDiags service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigDiagsClient interface {
	// GetChanges returns a stream of submitted changes objects.
	GetChanges(ctx context.Context, in *ChangesRequest, opts ...grpc.CallOption) (ConfigDiags_GetChangesClient, error)
	// GetConfigurations returns a stream of submitted configurations aimed at individual devices.
	GetConfigurations(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ConfigDiags_GetConfigurationsClient, error)
}

type configDiagsClient struct {
	cc *grpc.ClientConn
}

func NewConfigDiagsClient(cc *grpc.ClientConn) ConfigDiagsClient {
	return &configDiagsClient{cc}
}

func (c *configDiagsClient) GetChanges(ctx context.Context, in *ChangesRequest, opts ...grpc.CallOption) (ConfigDiags_GetChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConfigDiags_serviceDesc.Streams[0], "/proto.ConfigDiags/GetChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &configDiagsGetChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfigDiags_GetChangesClient interface {
	Recv() (*Change, error)
	grpc.ClientStream
}

type configDiagsGetChangesClient struct {
	grpc.ClientStream
}

func (x *configDiagsGetChangesClient) Recv() (*Change, error) {
	m := new(Change)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *configDiagsClient) GetConfigurations(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ConfigDiags_GetConfigurationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConfigDiags_serviceDesc.Streams[1], "/proto.ConfigDiags/GetConfigurations", opts...)
	if err != nil {
		return nil, err
	}
	x := &configDiagsGetConfigurationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfigDiags_GetConfigurationsClient interface {
	Recv() (*Configuration, error)
	grpc.ClientStream
}

type configDiagsGetConfigurationsClient struct {
	grpc.ClientStream
}

func (x *configDiagsGetConfigurationsClient) Recv() (*Configuration, error) {
	m := new(Configuration)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfigDiagsServer is the server API for ConfigDiags service.
type ConfigDiagsServer interface {
	// GetChanges returns a stream of submitted changes objects.
	GetChanges(*ChangesRequest, ConfigDiags_GetChangesServer) error
	// GetConfigurations returns a stream of submitted configurations aimed at individual devices.
	GetConfigurations(*ConfigRequest, ConfigDiags_GetConfigurationsServer) error
}

// UnimplementedConfigDiagsServer can be embedded to have forward compatible implementations.
type UnimplementedConfigDiagsServer struct {
}

func (*UnimplementedConfigDiagsServer) GetChanges(req *ChangesRequest, srv ConfigDiags_GetChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChanges not implemented")
}
func (*UnimplementedConfigDiagsServer) GetConfigurations(req *ConfigRequest, srv ConfigDiags_GetConfigurationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfigurations not implemented")
}

func RegisterConfigDiagsServer(s *grpc.Server, srv ConfigDiagsServer) {
	s.RegisterService(&_ConfigDiags_serviceDesc, srv)
}

func _ConfigDiags_GetChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigDiagsServer).GetChanges(m, &configDiagsGetChangesServer{stream})
}

type ConfigDiags_GetChangesServer interface {
	Send(*Change) error
	grpc.ServerStream
}

type configDiagsGetChangesServer struct {
	grpc.ServerStream
}

func (x *configDiagsGetChangesServer) Send(m *Change) error {
	return x.ServerStream.SendMsg(m)
}

func _ConfigDiags_GetConfigurations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfigRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigDiagsServer).GetConfigurations(m, &configDiagsGetConfigurationsServer{stream})
}

type ConfigDiags_GetConfigurationsServer interface {
	Send(*Configuration) error
	grpc.ServerStream
}

type configDiagsGetConfigurationsServer struct {
	grpc.ServerStream
}

func (x *configDiagsGetConfigurationsServer) Send(m *Configuration) error {
	return x.ServerStream.SendMsg(m)
}

var _ConfigDiags_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ConfigDiags",
	HandlerType: (*ConfigDiagsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChanges",
			Handler:       _ConfigDiags_GetChanges_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetConfigurations",
			Handler:       _ConfigDiags_GetConfigurations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/northbound/proto/diags.proto",
}