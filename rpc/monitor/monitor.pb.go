// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: monitor/monitor.proto

package monitor

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MonitorConfig_TargetType int32

const (
	MonitorConfig_SERIAL MonitorConfig_TargetType = 0
)

// Enum value maps for MonitorConfig_TargetType.
var (
	MonitorConfig_TargetType_name = map[int32]string{
		0: "SERIAL",
	}
	MonitorConfig_TargetType_value = map[string]int32{
		"SERIAL": 0,
	}
)

func (x MonitorConfig_TargetType) Enum() *MonitorConfig_TargetType {
	p := new(MonitorConfig_TargetType)
	*p = x
	return p
}

func (x MonitorConfig_TargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MonitorConfig_TargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_monitor_monitor_proto_enumTypes[0].Descriptor()
}

func (MonitorConfig_TargetType) Type() protoreflect.EnumType {
	return &file_monitor_monitor_proto_enumTypes[0]
}

func (x MonitorConfig_TargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MonitorConfig_TargetType.Descriptor instead.
func (MonitorConfig_TargetType) EnumDescriptor() ([]byte, []int) {
	return file_monitor_monitor_proto_rawDescGZIP(), []int{1, 0}
}

// The top-level message sent by the client for the `StreamingOpen` method.
// Multiple `StreamingOpenReq` messages can be sent but the first message
// must contain a `monitor_config` message to initialize the monitor target.
// All subsequent messages must contain bytes to be sent to the target
// and must not contain a `monitor_config` message.
type StreamingOpenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Content must be either a monitor config or data to be sent.
	//
	// Types that are assignable to Content:
	//	*StreamingOpenReq_MonitorConfig
	//	*StreamingOpenReq_Data
	Content isStreamingOpenReq_Content `protobuf_oneof:"content"`
}

func (x *StreamingOpenReq) Reset() {
	*x = StreamingOpenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingOpenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingOpenReq) ProtoMessage() {}

func (x *StreamingOpenReq) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingOpenReq.ProtoReflect.Descriptor instead.
func (*StreamingOpenReq) Descriptor() ([]byte, []int) {
	return file_monitor_monitor_proto_rawDescGZIP(), []int{0}
}

func (m *StreamingOpenReq) GetContent() isStreamingOpenReq_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *StreamingOpenReq) GetMonitorConfig() *MonitorConfig {
	if x, ok := x.GetContent().(*StreamingOpenReq_MonitorConfig); ok {
		return x.MonitorConfig
	}
	return nil
}

func (x *StreamingOpenReq) GetData() []byte {
	if x, ok := x.GetContent().(*StreamingOpenReq_Data); ok {
		return x.Data
	}
	return nil
}

type isStreamingOpenReq_Content interface {
	isStreamingOpenReq_Content()
}

type StreamingOpenReq_MonitorConfig struct {
	// Provides information to the monitor that specifies which is the target.
	// The first `StreamingOpenReq` message must contain a `monitor_config`
	// message.
	MonitorConfig *MonitorConfig `protobuf:"bytes,1,opt,name=monitorConfig,proto3,oneof"`
}

type StreamingOpenReq_Data struct {
	// The data to be sent to the target being monitored.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*StreamingOpenReq_MonitorConfig) isStreamingOpenReq_Content() {}

func (*StreamingOpenReq_Data) isStreamingOpenReq_Content() {}

// Tells the monitor which target to open and provides additional parameters
// that might be needed to configure the target or the monitor itself.
type MonitorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The target name.
	Target string                   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Type   MonitorConfig_TargetType `protobuf:"varint,2,opt,name=type,proto3,enum=cc.arduino.cli.monitor.MonitorConfig_TargetType" json:"type,omitempty"`
	// Additional parameters that might be needed to configure the target or the
	// monitor itself.
	AdditionalConfig *_struct.Struct `protobuf:"bytes,3,opt,name=additionalConfig,proto3" json:"additionalConfig,omitempty"`
}

func (x *MonitorConfig) Reset() {
	*x = MonitorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorConfig) ProtoMessage() {}

func (x *MonitorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorConfig.ProtoReflect.Descriptor instead.
func (*MonitorConfig) Descriptor() ([]byte, []int) {
	return file_monitor_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *MonitorConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *MonitorConfig) GetType() MonitorConfig_TargetType {
	if x != nil {
		return x.Type
	}
	return MonitorConfig_SERIAL
}

func (x *MonitorConfig) GetAdditionalConfig() *_struct.Struct {
	if x != nil {
		return x.AdditionalConfig
	}
	return nil
}

//
type StreamingOpenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The data received from the target.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamingOpenResp) Reset() {
	*x = StreamingOpenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_monitor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingOpenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingOpenResp) ProtoMessage() {}

func (x *StreamingOpenResp) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_monitor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingOpenResp.ProtoReflect.Descriptor instead.
func (*StreamingOpenResp) Descriptor() ([]byte, []int) {
	return file_monitor_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *StreamingOpenResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_monitor_monitor_proto protoreflect.FileDescriptor

var file_monitor_monitor_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75,
	0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01,
	0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x4d, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x63, 0x2e, 0x61,
	0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x00, 0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0d, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x44, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x63, 0x2e,
	0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x18, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x10,
	0x00, 0x22, 0x27, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x70,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x75, 0x0a, 0x07, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x6a, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75,
	0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x29, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c,
	0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2d,
	0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_monitor_proto_rawDescOnce sync.Once
	file_monitor_monitor_proto_rawDescData = file_monitor_monitor_proto_rawDesc
)

func file_monitor_monitor_proto_rawDescGZIP() []byte {
	file_monitor_monitor_proto_rawDescOnce.Do(func() {
		file_monitor_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_monitor_proto_rawDescData)
	})
	return file_monitor_monitor_proto_rawDescData
}

var file_monitor_monitor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_monitor_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_monitor_monitor_proto_goTypes = []interface{}{
	(MonitorConfig_TargetType)(0), // 0: cc.arduino.cli.monitor.MonitorConfig.TargetType
	(*StreamingOpenReq)(nil),      // 1: cc.arduino.cli.monitor.StreamingOpenReq
	(*MonitorConfig)(nil),         // 2: cc.arduino.cli.monitor.MonitorConfig
	(*StreamingOpenResp)(nil),     // 3: cc.arduino.cli.monitor.StreamingOpenResp
	(*_struct.Struct)(nil),        // 4: google.protobuf.Struct
}
var file_monitor_monitor_proto_depIdxs = []int32{
	2, // 0: cc.arduino.cli.monitor.StreamingOpenReq.monitorConfig:type_name -> cc.arduino.cli.monitor.MonitorConfig
	0, // 1: cc.arduino.cli.monitor.MonitorConfig.type:type_name -> cc.arduino.cli.monitor.MonitorConfig.TargetType
	4, // 2: cc.arduino.cli.monitor.MonitorConfig.additionalConfig:type_name -> google.protobuf.Struct
	1, // 3: cc.arduino.cli.monitor.Monitor.StreamingOpen:input_type -> cc.arduino.cli.monitor.StreamingOpenReq
	3, // 4: cc.arduino.cli.monitor.Monitor.StreamingOpen:output_type -> cc.arduino.cli.monitor.StreamingOpenResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_monitor_monitor_proto_init() }
func file_monitor_monitor_proto_init() {
	if File_monitor_monitor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingOpenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_monitor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingOpenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_monitor_monitor_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StreamingOpenReq_MonitorConfig)(nil),
		(*StreamingOpenReq_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitor_monitor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitor_monitor_proto_goTypes,
		DependencyIndexes: file_monitor_monitor_proto_depIdxs,
		EnumInfos:         file_monitor_monitor_proto_enumTypes,
		MessageInfos:      file_monitor_monitor_proto_msgTypes,
	}.Build()
	File_monitor_monitor_proto = out.File
	file_monitor_monitor_proto_rawDesc = nil
	file_monitor_monitor_proto_goTypes = nil
	file_monitor_monitor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorClient interface {
	// Open a bidirectional monitor stream. This can be used to implement
	// something similar to the Arduino IDE's Serial Monitor.
	StreamingOpen(ctx context.Context, opts ...grpc.CallOption) (Monitor_StreamingOpenClient, error)
}

type monitorClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorClient(cc grpc.ClientConnInterface) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) StreamingOpen(ctx context.Context, opts ...grpc.CallOption) (Monitor_StreamingOpenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Monitor_serviceDesc.Streams[0], "/cc.arduino.cli.monitor.Monitor/StreamingOpen", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorStreamingOpenClient{stream}
	return x, nil
}

type Monitor_StreamingOpenClient interface {
	Send(*StreamingOpenReq) error
	Recv() (*StreamingOpenResp, error)
	grpc.ClientStream
}

type monitorStreamingOpenClient struct {
	grpc.ClientStream
}

func (x *monitorStreamingOpenClient) Send(m *StreamingOpenReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *monitorStreamingOpenClient) Recv() (*StreamingOpenResp, error) {
	m := new(StreamingOpenResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorServer is the server API for Monitor service.
type MonitorServer interface {
	// Open a bidirectional monitor stream. This can be used to implement
	// something similar to the Arduino IDE's Serial Monitor.
	StreamingOpen(Monitor_StreamingOpenServer) error
}

// UnimplementedMonitorServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorServer struct {
}

func (*UnimplementedMonitorServer) StreamingOpen(Monitor_StreamingOpenServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingOpen not implemented")
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_StreamingOpen_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MonitorServer).StreamingOpen(&monitorStreamingOpenServer{stream})
}

type Monitor_StreamingOpenServer interface {
	Send(*StreamingOpenResp) error
	Recv() (*StreamingOpenReq, error)
	grpc.ServerStream
}

type monitorStreamingOpenServer struct {
	grpc.ServerStream
}

func (x *monitorStreamingOpenServer) Send(m *StreamingOpenResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *monitorStreamingOpenServer) Recv() (*StreamingOpenReq, error) {
	m := new(StreamingOpenReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cc.arduino.cli.monitor.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOpen",
			Handler:       _Monitor_StreamingOpen_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "monitor/monitor.proto",
}
