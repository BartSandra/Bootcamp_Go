// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: heart_beat.proto

package heart

import (
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

type PulseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port     int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	IsClient bool  `protobuf:"varint,2,opt,name=is_client,json=isClient,proto3" json:"is_client,omitempty"`
}

func (x *PulseRequest) Reset() {
	*x = PulseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heart_beat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulseRequest) ProtoMessage() {}

func (x *PulseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heart_beat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulseRequest.ProtoReflect.Descriptor instead.
func (*PulseRequest) Descriptor() ([]byte, []int) {
	return file_heart_beat_proto_rawDescGZIP(), []int{0}
}

func (x *PulseRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PulseRequest) GetIsClient() bool {
	if x != nil {
		return x.IsClient
	}
	return false
}

type PulseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port                   int32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	ClusterPorts           []int32 `protobuf:"varint,2,rep,packed,name=cluster_ports,json=clusterPorts,proto3" json:"cluster_ports,omitempty"`
	ReplicationCoefficient int32   `protobuf:"varint,3,opt,name=replication_coefficient,json=replicationCoefficient,proto3" json:"replication_coefficient,omitempty"`
}

func (x *PulseResponse) Reset() {
	*x = PulseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heart_beat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulseResponse) ProtoMessage() {}

func (x *PulseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heart_beat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulseResponse.ProtoReflect.Descriptor instead.
func (*PulseResponse) Descriptor() ([]byte, []int) {
	return file_heart_beat_proto_rawDescGZIP(), []int{1}
}

func (x *PulseResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PulseResponse) GetClusterPorts() []int32 {
	if x != nil {
		return x.ClusterPorts
	}
	return nil
}

func (x *PulseResponse) GetReplicationCoefficient() int32 {
	if x != nil {
		return x.ReplicationCoefficient
	}
	return 0
}

var File_heart_beat_proto protoreflect.FileDescriptor

var file_heart_beat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x61, 0x72, 0x74, 0x22, 0x3f, 0x0a, 0x0c, 0x50, 0x75, 0x6c,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x50,
	0x75, 0x6c, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x17, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x48,
	0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x2e, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x2e, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heart_beat_proto_rawDescOnce sync.Once
	file_heart_beat_proto_rawDescData = file_heart_beat_proto_rawDesc
)

func file_heart_beat_proto_rawDescGZIP() []byte {
	file_heart_beat_proto_rawDescOnce.Do(func() {
		file_heart_beat_proto_rawDescData = protoimpl.X.CompressGZIP(file_heart_beat_proto_rawDescData)
	})
	return file_heart_beat_proto_rawDescData
}

var file_heart_beat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_heart_beat_proto_goTypes = []interface{}{
	(*PulseRequest)(nil),  // 0: heart.PulseRequest
	(*PulseResponse)(nil), // 1: heart.PulseResponse
}
var file_heart_beat_proto_depIdxs = []int32{
	0, // 0: heart.HeartBeatService.Pulse:input_type -> heart.PulseRequest
	1, // 1: heart.HeartBeatService.Pulse:output_type -> heart.PulseResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heart_beat_proto_init() }
func file_heart_beat_proto_init() {
	if File_heart_beat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heart_beat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulseRequest); i {
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
		file_heart_beat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulseResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_heart_beat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heart_beat_proto_goTypes,
		DependencyIndexes: file_heart_beat_proto_depIdxs,
		MessageInfos:      file_heart_beat_proto_msgTypes,
	}.Build()
	File_heart_beat_proto = out.File
	file_heart_beat_proto_rawDesc = nil
	file_heart_beat_proto_goTypes = nil
	file_heart_beat_proto_depIdxs = nil
}