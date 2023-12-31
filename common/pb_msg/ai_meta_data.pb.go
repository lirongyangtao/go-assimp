// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: common/pb_msg/ai_meta_data.proto

package pb_msg

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

type AiMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * Arrays of keys, may not be NULL. Entries in this array may not be NULL as well.
	Keys []string `protobuf:"bytes,1,rep,name=Keys,proto3" json:"Keys,omitempty"`
	// * Arrays of values, may not be NULL. Entries in this array may be NULL if the
	// corresponding property key has no assigned value.
	Values []*AiMetadataEntry `protobuf:"bytes,2,rep,name=Values,proto3" json:"Values,omitempty"`
}

func (x *AiMetadata) Reset() {
	*x = AiMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_meta_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiMetadata) ProtoMessage() {}

func (x *AiMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_meta_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiMetadata.ProtoReflect.Descriptor instead.
func (*AiMetadata) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_meta_data_proto_rawDescGZIP(), []int{0}
}

func (x *AiMetadata) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *AiMetadata) GetValues() []*AiMetadataEntry {
	if x != nil {
		return x.Values
	}
	return nil
}

type AiMetadataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32  `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *AiMetadataEntry) Reset() {
	*x = AiMetadataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_meta_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiMetadataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiMetadataEntry) ProtoMessage() {}

func (x *AiMetadataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_meta_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiMetadataEntry.ProtoReflect.Descriptor instead.
func (*AiMetadataEntry) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_meta_data_proto_rawDescGZIP(), []int{1}
}

func (x *AiMetadataEntry) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AiMetadataEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_common_pb_msg_ai_meta_data_proto protoreflect.FileDescriptor

var file_common_pb_msg_ai_meta_data_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2f,
	0x61, 0x69, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x51, 0x0a, 0x0a, 0x41, 0x69,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2f, 0x0a, 0x06,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x39, 0x0a,
	0x0f, 0x41, 0x69, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x5a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_pb_msg_ai_meta_data_proto_rawDescOnce sync.Once
	file_common_pb_msg_ai_meta_data_proto_rawDescData = file_common_pb_msg_ai_meta_data_proto_rawDesc
)

func file_common_pb_msg_ai_meta_data_proto_rawDescGZIP() []byte {
	file_common_pb_msg_ai_meta_data_proto_rawDescOnce.Do(func() {
		file_common_pb_msg_ai_meta_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pb_msg_ai_meta_data_proto_rawDescData)
	})
	return file_common_pb_msg_ai_meta_data_proto_rawDescData
}

var file_common_pb_msg_ai_meta_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_pb_msg_ai_meta_data_proto_goTypes = []interface{}{
	(*AiMetadata)(nil),      // 0: pb_msg.AiMetadata
	(*AiMetadataEntry)(nil), // 1: pb_msg.AiMetadataEntry
}
var file_common_pb_msg_ai_meta_data_proto_depIdxs = []int32{
	1, // 0: pb_msg.AiMetadata.Values:type_name -> pb_msg.AiMetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_pb_msg_ai_meta_data_proto_init() }
func file_common_pb_msg_ai_meta_data_proto_init() {
	if File_common_pb_msg_ai_meta_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_pb_msg_ai_meta_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiMetadata); i {
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
		file_common_pb_msg_ai_meta_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiMetadataEntry); i {
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
			RawDescriptor: file_common_pb_msg_ai_meta_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pb_msg_ai_meta_data_proto_goTypes,
		DependencyIndexes: file_common_pb_msg_ai_meta_data_proto_depIdxs,
		MessageInfos:      file_common_pb_msg_ai_meta_data_proto_msgTypes,
	}.Build()
	File_common_pb_msg_ai_meta_data_proto = out.File
	file_common_pb_msg_ai_meta_data_proto_rawDesc = nil
	file_common_pb_msg_ai_meta_data_proto_goTypes = nil
	file_common_pb_msg_ai_meta_data_proto_depIdxs = nil
}
