// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: common/pb_msg/ai_node.proto

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

// -------------------------------------------------------------------------------
type AiNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * The name of the node.
	//
	// The name might be empty (length of zero) but all nodes which
	// need to be referenced by either bones or animations are named.
	// Multiple nodes may have the same name, except for nodes which are referenced
	// by bones (see #aiBone and #aiMesh::mBones). Their names *must* be unique.
	//
	// Cameras and lights reference a specific node by name - if there
	// are multiple nodes with this name, they are assigned to each of them.
	// <br>
	// There are no limitations with regard to the characters contained in
	// the name string as it is usually taken directly from the source file.
	//
	// Implementations should be able to handle tokens such as whitespace, tabs,
	// line feeds, quotation marks, ampersands etc.
	//
	// Sometimes assimp introduces new nodes not present in the source file
	// into the hierarchy (usually out of necessity because sometimes the
	// source hierarchy format is simply not compatible). Their names are
	// surrounded by @verbatim <> @endverbatim e.g.
	//
	//	@verbatim<DummyRootNode> @endverbatim.
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// * The transformation relative to the node's parent.
	Transformation *AiMatrix4X4 `protobuf:"bytes,2,opt,name=Transformation,proto3" json:"Transformation,omitempty"`
	// * The child nodes of this node. nullptr if mNumChildren is 0.
	Children []*AiNode `protobuf:"bytes,3,rep,name=Children,proto3" json:"Children,omitempty"`
	// * The meshes of this node. Each entry is an index into the
	// mesh list of the #aiScene.
	Meshes []int32 `protobuf:"varint,4,rep,packed,name=Meshes,proto3" json:"Meshes,omitempty"`
	//   - Metadata associated with this node or nullptr if there is no metadata.
	//     Whether any metadata is generated depends on the source file format. See the
	//
	// @link importer_notes @endlink page for more information on every source file
	// format. Importers that don't document any metadata don't write any.
	MetaData *AiMetadata `protobuf:"bytes,5,opt,name=MetaData,proto3" json:"MetaData,omitempty"`
}

func (x *AiNode) Reset() {
	*x = AiNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiNode) ProtoMessage() {}

func (x *AiNode) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiNode.ProtoReflect.Descriptor instead.
func (*AiNode) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_node_proto_rawDescGZIP(), []int{0}
}

func (x *AiNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AiNode) GetTransformation() *AiMatrix4X4 {
	if x != nil {
		return x.Transformation
	}
	return nil
}

func (x *AiNode) GetChildren() []*AiNode {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *AiNode) GetMeshes() []int32 {
	if x != nil {
		return x.Meshes
	}
	return nil
}

func (x *AiNode) GetMetaData() *AiMetadata {
	if x != nil {
		return x.MetaData
	}
	return nil
}

var File_common_pb_msg_ai_node_proto protoreflect.FileDescriptor

var file_common_pb_msg_ai_node_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2f,
	0x61, 0x69, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x62, 0x5f, 0x6d, 0x73, 0x67, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x5f, 0x6d, 0x73, 0x67, 0x2f, 0x61, 0x69, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x06, 0x41, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x5f,
	0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x34, 0x78, 0x34, 0x52,
	0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4d,
	0x65, 0x73, 0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x4d, 0x65, 0x73,
	0x68, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41,
	0x69, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x0f, 0x5a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x5f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_pb_msg_ai_node_proto_rawDescOnce sync.Once
	file_common_pb_msg_ai_node_proto_rawDescData = file_common_pb_msg_ai_node_proto_rawDesc
)

func file_common_pb_msg_ai_node_proto_rawDescGZIP() []byte {
	file_common_pb_msg_ai_node_proto_rawDescOnce.Do(func() {
		file_common_pb_msg_ai_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pb_msg_ai_node_proto_rawDescData)
	})
	return file_common_pb_msg_ai_node_proto_rawDescData
}

var file_common_pb_msg_ai_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_pb_msg_ai_node_proto_goTypes = []interface{}{
	(*AiNode)(nil),      // 0: pb_msg.AiNode
	(*AiMatrix4X4)(nil), // 1: pb_msg.AiMatrix4x4
	(*AiMetadata)(nil),  // 2: pb_msg.AiMetadata
}
var file_common_pb_msg_ai_node_proto_depIdxs = []int32{
	1, // 0: pb_msg.AiNode.Transformation:type_name -> pb_msg.AiMatrix4x4
	0, // 1: pb_msg.AiNode.Children:type_name -> pb_msg.AiNode
	2, // 2: pb_msg.AiNode.MetaData:type_name -> pb_msg.AiMetadata
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_pb_msg_ai_node_proto_init() }
func file_common_pb_msg_ai_node_proto_init() {
	if File_common_pb_msg_ai_node_proto != nil {
		return
	}
	file_common_pb_msg_ai_meta_data_proto_init()
	file_common_pb_msg_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_pb_msg_ai_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiNode); i {
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
			RawDescriptor: file_common_pb_msg_ai_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pb_msg_ai_node_proto_goTypes,
		DependencyIndexes: file_common_pb_msg_ai_node_proto_depIdxs,
		MessageInfos:      file_common_pb_msg_ai_node_proto_msgTypes,
	}.Build()
	File_common_pb_msg_ai_node_proto = out.File
	file_common_pb_msg_ai_node_proto_rawDesc = nil
	file_common_pb_msg_ai_node_proto_goTypes = nil
	file_common_pb_msg_ai_node_proto_depIdxs = nil
}
