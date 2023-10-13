// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: sf/substreams/sink/subgraph/v1/services.proto

package pbsubgraph

import (
	_ "github.com/streamingfast/substreams/pb/sf/substreams"
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

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Containing both create table statements and index creation statements
	Schema                       string         `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	SubgraphYaml                 string         `protobuf:"bytes,2,opt,name=subgraph_yaml,json=subgraphYaml,proto3" json:"subgraph_yaml,omitempty"`
	PostgresDirectProtocolAccess bool           `protobuf:"varint,3,opt,name=postgres_direct_protocol_access,json=postgresDirectProtocolAccess,proto3" json:"postgres_direct_protocol_access,omitempty"`
	PgwebFrontend                *PGWebFrontend `protobuf:"bytes,4,opt,name=pgweb_frontend,json=pgwebFrontend,proto3" json:"pgweb_frontend,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_sf_substreams_sink_subgraph_v1_services_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Service) GetSubgraphYaml() string {
	if x != nil {
		return x.SubgraphYaml
	}
	return ""
}

func (x *Service) GetPostgresDirectProtocolAccess() bool {
	if x != nil {
		return x.PostgresDirectProtocolAccess
	}
	return false
}

func (x *Service) GetPgwebFrontend() *PGWebFrontend {
	if x != nil {
		return x.PgwebFrontend
	}
	return nil
}

type PGWebFrontend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *PGWebFrontend) Reset() {
	*x = PGWebFrontend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PGWebFrontend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PGWebFrontend) ProtoMessage() {}

func (x *PGWebFrontend) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PGWebFrontend.ProtoReflect.Descriptor instead.
func (*PGWebFrontend) Descriptor() ([]byte, []int) {
	return file_sf_substreams_sink_subgraph_v1_services_proto_rawDescGZIP(), []int{1}
}

func (x *PGWebFrontend) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_sf_substreams_sink_subgraph_v1_services_proto protoreflect.FileDescriptor

var file_sf_substreams_sink_subgraph_v1_services_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x73, 0x69, 0x6e, 0x6b, 0x2f, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73,
	0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0x89, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2b, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x5f, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xc2, 0x89, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x59, 0x61, 0x6d, 0x6c, 0x12, 0x45, 0x0a, 0x1f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c,
	0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x54, 0x0a, 0x0e,
	0x70, 0x67, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x47, 0x57, 0x65, 0x62, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x52, 0x0d, 0x70, 0x67, 0x77, 0x65, 0x62, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x50, 0x47, 0x57, 0x65, 0x62, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x90, 0x02,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2d, 0x73, 0x75, 0x62, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2d, 0x73, 0x71, 0x6c, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x73, 0x75, 0x62,
	0x67, 0x72, 0x61, 0x70, 0x68, 0xa2, 0x02, 0x04, 0x53, 0x53, 0x53, 0x53, 0xaa, 0x02, 0x1e, 0x53,
	0x66, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x53, 0x69, 0x6e,
	0x6b, 0x2e, 0x53, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e,
	0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x5c, 0x53, 0x69,
	0x6e, 0x6b, 0x5c, 0x53, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x2a, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x5c, 0x53,
	0x69, 0x6e, 0x6b, 0x5c, 0x53, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x53, 0x66,
	0x3a, 0x3a, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x3a, 0x3a, 0x53, 0x69,
	0x6e, 0x6b, 0x3a, 0x3a, 0x53, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_substreams_sink_subgraph_v1_services_proto_rawDescOnce sync.Once
	file_sf_substreams_sink_subgraph_v1_services_proto_rawDescData = file_sf_substreams_sink_subgraph_v1_services_proto_rawDesc
)

func file_sf_substreams_sink_subgraph_v1_services_proto_rawDescGZIP() []byte {
	file_sf_substreams_sink_subgraph_v1_services_proto_rawDescOnce.Do(func() {
		file_sf_substreams_sink_subgraph_v1_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_sink_subgraph_v1_services_proto_rawDescData)
	})
	return file_sf_substreams_sink_subgraph_v1_services_proto_rawDescData
}

var file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sf_substreams_sink_subgraph_v1_services_proto_goTypes = []interface{}{
	(*Service)(nil),       // 0: sf.substreams.sink.subgraph.v1.Service
	(*PGWebFrontend)(nil), // 1: sf.substreams.sink.subgraph.v1.PGWebFrontend
}
var file_sf_substreams_sink_subgraph_v1_services_proto_depIdxs = []int32{
	1, // 0: sf.substreams.sink.subgraph.v1.Service.pgweb_frontend:type_name -> sf.substreams.sink.subgraph.v1.PGWebFrontend
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sf_substreams_sink_subgraph_v1_services_proto_init() }
func file_sf_substreams_sink_subgraph_v1_services_proto_init() {
	if File_sf_substreams_sink_subgraph_v1_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PGWebFrontend); i {
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
			RawDescriptor: file_sf_substreams_sink_subgraph_v1_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_sink_subgraph_v1_services_proto_goTypes,
		DependencyIndexes: file_sf_substreams_sink_subgraph_v1_services_proto_depIdxs,
		MessageInfos:      file_sf_substreams_sink_subgraph_v1_services_proto_msgTypes,
	}.Build()
	File_sf_substreams_sink_subgraph_v1_services_proto = out.File
	file_sf_substreams_sink_subgraph_v1_services_proto_rawDesc = nil
	file_sf_substreams_sink_subgraph_v1_services_proto_goTypes = nil
	file_sf_substreams_sink_subgraph_v1_services_proto_depIdxs = nil
}
