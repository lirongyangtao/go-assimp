// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: common/pb_msg/ai_animation.proto

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

//   - An animation consists of key-frame data for a number of nodes. For
//     each node affected by the animation a separate series of data is given.
type AiAnimation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//   - The name of the animation. If the modeling package this data was
	//     exported from does support only a single animation channel, this
	//     name is usually empty (length is zero).
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// * Duration of the animation in ticks.
	Duration float64 `protobuf:"fixed64,2,opt,name=Duration,proto3" json:"Duration,omitempty"`
	// * Ticks per second. 0 if not specified in the imported file
	TicksPerSecond float64 `protobuf:"fixed64,3,opt,name=TicksPerSecond,proto3" json:"TicksPerSecond,omitempty"`
	//   - The node animation channels. Each channel affects a single node.
	//     The array is mNumChannels in size.
	Channels []*AiNodeAnim `protobuf:"bytes,4,rep,name=Channels,proto3" json:"Channels,omitempty"`
	//   - The mesh animation channels. Each channel affects a single mesh.
	//     The array is mNumMeshChannels in size.
	MeshChannels []*AiMeshAnim `protobuf:"bytes,5,rep,name=MeshChannels,proto3" json:"MeshChannels,omitempty"`
	//   - The morph mesh animation channels. Each channel affects a single mesh.
	//     The array is mNumMorphMeshChannels in size.
	MorphMeshChannels []*AiMeshMorphAnim `protobuf:"bytes,6,rep,name=MorphMeshChannels,proto3" json:"MorphMeshChannels,omitempty"`
}

func (x *AiAnimation) Reset() {
	*x = AiAnimation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiAnimation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiAnimation) ProtoMessage() {}

func (x *AiAnimation) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiAnimation.ProtoReflect.Descriptor instead.
func (*AiAnimation) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{0}
}

func (x *AiAnimation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AiAnimation) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *AiAnimation) GetTicksPerSecond() float64 {
	if x != nil {
		return x.TicksPerSecond
	}
	return 0
}

func (x *AiAnimation) GetChannels() []*AiNodeAnim {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *AiAnimation) GetMeshChannels() []*AiMeshAnim {
	if x != nil {
		return x.MeshChannels
	}
	return nil
}

func (x *AiAnimation) GetMorphMeshChannels() []*AiMeshMorphAnim {
	if x != nil {
		return x.MorphMeshChannels
	}
	return nil
}

// * Describes a morphing animation of a given mesh.
type AiMeshMorphAnim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//   - Name of the mesh to be animated. An empty string is not allowed,
	//     animated meshes need to be named (not necessarily uniquely,
	//     the name can basically serve as wildcard to select a group
	//     of meshes with similar animation setup)
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// * Key frames of the animation. May not be nullptr.
	Keys []*AiMeshMorphKey `protobuf:"bytes,2,rep,name=Keys,proto3" json:"Keys,omitempty"`
}

func (x *AiMeshMorphAnim) Reset() {
	*x = AiMeshMorphAnim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiMeshMorphAnim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiMeshMorphAnim) ProtoMessage() {}

func (x *AiMeshMorphAnim) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiMeshMorphAnim.ProtoReflect.Descriptor instead.
func (*AiMeshMorphAnim) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{1}
}

func (x *AiMeshMorphAnim) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AiMeshMorphAnim) GetKeys() []*AiMeshMorphKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

// * Binds a morph anim mesh to a specific point in time.
type AiMeshMorphKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * The time of this key
	Time float64 `protobuf:"fixed64,1,opt,name=Time,proto3" json:"Time,omitempty"`
	// * The values and weights at the time of this key
	//   - mValues: index of attachment mesh to apply weight at the same position in mWeights
	//   - mWeights: weight to apply to the blend shape index at the same position in mValues
	Values  []uint32  `protobuf:"varint,2,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	Weights []float64 `protobuf:"fixed64,3,rep,packed,name=Weights,proto3" json:"Weights,omitempty"`
}

func (x *AiMeshMorphKey) Reset() {
	*x = AiMeshMorphKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiMeshMorphKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiMeshMorphKey) ProtoMessage() {}

func (x *AiMeshMorphKey) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiMeshMorphKey.ProtoReflect.Descriptor instead.
func (*AiMeshMorphKey) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{2}
}

func (x *AiMeshMorphKey) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AiMeshMorphKey) GetValues() []uint32 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *AiMeshMorphKey) GetWeights() []float64 {
	if x != nil {
		return x.Weights
	}
	return nil
}

//   - Describes the animation of a single node. The name specifies the
//     bone/node which is affected by this animation channel. The keyframes
//     are given in three separate series of values, one each for position,
//     rotation and scaling. The transformation matrix computed from these
//     values replaces the node's original transformation matrix at a
//     specific time.
//     This means all keys are absolute and not relative to the bone default pose.
//     The order in which the transformations are applied is
//
//   - as usual - scaling, rotation, translation.
//
//     @note All keys are returned in their correct, chronological order.
//     Duplicate keys don't pass the validation step. Most likely there
//     will be no negative time values, but they are not forbidden also ( so
//     implementations need to cope with them! )
type AiNodeAnim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//   - The name of the node affected by this animation. The node
	//     must exist and it must be unique.
	NodeName string `protobuf:"bytes,1,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	// * The position keys of this animation channel. Positions are
	// specified as 3D vector. The array is mNumPositionKeys in size.
	//
	// If there are position keys, there will also be at least one
	// scaling and one rotation key.
	PositionKeys []*AiVectorKey `protobuf:"bytes,2,rep,name=PositionKeys,proto3" json:"PositionKeys,omitempty"`
	//   - The rotation keys of this animation channel. Rotations are
	//     given as quaternions,  which are 4D vectors. The array is
	//     mNumRotationKeys in size.
	//
	// If there are rotation keys, there will also be at least one
	// scaling and one position key.
	RotationKeys []*AiQuatKey `protobuf:"bytes,3,rep,name=RotationKeys,proto3" json:"RotationKeys,omitempty"`
	//   - The scaling keys of this animation channel. Scalings are
	//     specified as 3D vector. The array is mNumScalingKeys in size.
	//
	// If there are scaling keys, there will also be at least one
	// position and one rotation key.
	ScalingKeys []*AiVectorKey `protobuf:"bytes,4,rep,name=ScalingKeys,proto3" json:"ScalingKeys,omitempty"`
	//   - Defines how the animation behaves before the first
	//     key is encountered.
	//
	//     The default value is aiAnimBehaviour_DEFAULT (the original
	//     transformation matrix of the affected node is used).
	PreState int32 `protobuf:"varint,5,opt,name=PreState,proto3" json:"PreState,omitempty"`
	//   - Defines how the animation behaves after the last
	//     key was processed.
	//
	//     The default value is aiAnimBehaviour_DEFAULT (the original
	//     transformation matrix of the affected node is taken).
	PostState int32 `protobuf:"varint,6,opt,name=PostState,proto3" json:"PostState,omitempty"`
}

func (x *AiNodeAnim) Reset() {
	*x = AiNodeAnim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiNodeAnim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiNodeAnim) ProtoMessage() {}

func (x *AiNodeAnim) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiNodeAnim.ProtoReflect.Descriptor instead.
func (*AiNodeAnim) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{3}
}

func (x *AiNodeAnim) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AiNodeAnim) GetPositionKeys() []*AiVectorKey {
	if x != nil {
		return x.PositionKeys
	}
	return nil
}

func (x *AiNodeAnim) GetRotationKeys() []*AiQuatKey {
	if x != nil {
		return x.RotationKeys
	}
	return nil
}

func (x *AiNodeAnim) GetScalingKeys() []*AiVectorKey {
	if x != nil {
		return x.ScalingKeys
	}
	return nil
}

func (x *AiNodeAnim) GetPreState() int32 {
	if x != nil {
		return x.PreState
	}
	return 0
}

func (x *AiNodeAnim) GetPostState() int32 {
	if x != nil {
		return x.PostState
	}
	return 0
}

//   - Describes vertex-based animations for a single mesh or a group of
//     meshes. Meshes carry the animation data for each frame in their
//     aiMesh::mAnimMeshes array. The purpose of aiMeshAnim is to
//     define keyframes linking each mesh attachment to a particular
//     point in time.
type AiMeshAnim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//   - Name of the mesh to be animated. An empty string is not allowed,
	//     animated meshes need to be named (not necessarily uniquely,
	//     the name can basically serve as wild-card to select a group
	//     of meshes with similar animation setup)
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// * Key frames of the animation. May not be nullptr.
	Keys []*AiMeshKey `protobuf:"bytes,2,rep,name=Keys,proto3" json:"Keys,omitempty"`
}

func (x *AiMeshAnim) Reset() {
	*x = AiMeshAnim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiMeshAnim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiMeshAnim) ProtoMessage() {}

func (x *AiMeshAnim) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiMeshAnim.ProtoReflect.Descriptor instead.
func (*AiMeshAnim) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{4}
}

func (x *AiMeshAnim) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AiMeshAnim) GetKeys() []*AiMeshKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

// * A time-value pair specifying a certain 3D vector for the given time.
type AiVectorKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * The time of this key
	Time float64 `protobuf:"fixed64,1,opt,name=Time,proto3" json:"Time,omitempty"`
	// * The value of this key
	Value *AiVector3D `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *AiVectorKey) Reset() {
	*x = AiVectorKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiVectorKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiVectorKey) ProtoMessage() {}

func (x *AiVectorKey) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiVectorKey.ProtoReflect.Descriptor instead.
func (*AiVectorKey) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{5}
}

func (x *AiVectorKey) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AiVectorKey) GetValue() *AiVector3D {
	if x != nil {
		return x.Value
	}
	return nil
}

//   - A time-value pair specifying a rotation for the given time.
//     Rotations are expressed with quaternions.
type AiQuatKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * The time of this key
	Time float64 `protobuf:"fixed64,1,opt,name=Time,proto3" json:"Time,omitempty"`
	// * The value of this key
	Value *AiQuaternion `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *AiQuatKey) Reset() {
	*x = AiQuatKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiQuatKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiQuatKey) ProtoMessage() {}

func (x *AiQuatKey) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiQuatKey.ProtoReflect.Descriptor instead.
func (*AiQuatKey) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{6}
}

func (x *AiQuatKey) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AiQuatKey) GetValue() *AiQuaternion {
	if x != nil {
		return x.Value
	}
	return nil
}

// * Binds a anim-mesh to a specific point in time.
type AiMeshKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * The time of this key
	Time float64 `protobuf:"fixed64,1,opt,name=Time,proto3" json:"Time,omitempty"`
	//   - Index into the aiMesh::mAnimMeshes array of the
	//     mesh corresponding to the #aiMeshAnim hosting this
	//     key frame. The referenced anim mesh is evaluated
	//     according to the rules defined in the docs for #aiAnimMesh.
	Value uint32 `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *AiMeshKey) Reset() {
	*x = AiMeshKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_msg_ai_animation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiMeshKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiMeshKey) ProtoMessage() {}

func (x *AiMeshKey) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_msg_ai_animation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiMeshKey.ProtoReflect.Descriptor instead.
func (*AiMeshKey) Descriptor() ([]byte, []int) {
	return file_common_pb_msg_ai_animation_proto_rawDescGZIP(), []int{7}
}

func (x *AiMeshKey) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AiMeshKey) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_common_pb_msg_ai_animation_proto protoreflect.FileDescriptor

var file_common_pb_msg_ai_animation_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2f,
	0x61, 0x69, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x0b, 0x41, 0x69, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x50,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e,
	0x54, 0x69, 0x63, 0x6b, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x2e,
	0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x4e, 0x6f, 0x64, 0x65,
	0x41, 0x6e, 0x69, 0x6d, 0x52, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x36,
	0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69,
	0x4d, 0x65, 0x73, 0x68, 0x41, 0x6e, 0x69, 0x6d, 0x52, 0x0c, 0x4d, 0x65, 0x73, 0x68, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x4d,
	0x65, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x4d, 0x65, 0x73,
	0x68, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x41, 0x6e, 0x69, 0x6d, 0x52, 0x11, 0x4d, 0x6f, 0x72, 0x70,
	0x68, 0x4d, 0x65, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22, 0x51, 0x0a,
	0x0f, 0x41, 0x69, 0x4d, 0x65, 0x73, 0x68, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x41, 0x6e, 0x69, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x4d, 0x65,
	0x73, 0x68, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x4b, 0x65, 0x79, 0x73,
	0x22, 0x56, 0x0a, 0x0e, 0x41, 0x69, 0x4d, 0x65, 0x73, 0x68, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x0a, 0x41, 0x69, 0x4e,
	0x6f, 0x64, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x6d,
	0x73, 0x67, 0x2e, 0x41, 0x69, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x0c,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x35, 0x0a, 0x0c,
	0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x51, 0x75,
	0x61, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x0c, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4b, 0x65,
	0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73,
	0x67, 0x2e, 0x41, 0x69, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x53,
	0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x72,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x47, 0x0a, 0x0a, 0x41, 0x69, 0x4d, 0x65, 0x73, 0x68, 0x41, 0x6e,
	0x69, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69,
	0x4d, 0x65, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x4b, 0x0a,
	0x0b, 0x41, 0x69, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x33, 0x44, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4b, 0x0a, 0x09, 0x41, 0x69,
	0x51, 0x75, 0x61, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x5f,
	0x6d, 0x73, 0x67, 0x2e, 0x41, 0x69, 0x51, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x09, 0x41, 0x69, 0x4d, 0x65, 0x73,
	0x68, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0f,
	0x5a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_pb_msg_ai_animation_proto_rawDescOnce sync.Once
	file_common_pb_msg_ai_animation_proto_rawDescData = file_common_pb_msg_ai_animation_proto_rawDesc
)

func file_common_pb_msg_ai_animation_proto_rawDescGZIP() []byte {
	file_common_pb_msg_ai_animation_proto_rawDescOnce.Do(func() {
		file_common_pb_msg_ai_animation_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pb_msg_ai_animation_proto_rawDescData)
	})
	return file_common_pb_msg_ai_animation_proto_rawDescData
}

var file_common_pb_msg_ai_animation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_pb_msg_ai_animation_proto_goTypes = []interface{}{
	(*AiAnimation)(nil),     // 0: pb_msg.AiAnimation
	(*AiMeshMorphAnim)(nil), // 1: pb_msg.AiMeshMorphAnim
	(*AiMeshMorphKey)(nil),  // 2: pb_msg.AiMeshMorphKey
	(*AiNodeAnim)(nil),      // 3: pb_msg.AiNodeAnim
	(*AiMeshAnim)(nil),      // 4: pb_msg.AiMeshAnim
	(*AiVectorKey)(nil),     // 5: pb_msg.AiVectorKey
	(*AiQuatKey)(nil),       // 6: pb_msg.AiQuatKey
	(*AiMeshKey)(nil),       // 7: pb_msg.AiMeshKey
	(*AiVector3D)(nil),      // 8: pb_msg.AiVector3D
	(*AiQuaternion)(nil),    // 9: pb_msg.AiQuaternion
}
var file_common_pb_msg_ai_animation_proto_depIdxs = []int32{
	3,  // 0: pb_msg.AiAnimation.Channels:type_name -> pb_msg.AiNodeAnim
	4,  // 1: pb_msg.AiAnimation.MeshChannels:type_name -> pb_msg.AiMeshAnim
	1,  // 2: pb_msg.AiAnimation.MorphMeshChannels:type_name -> pb_msg.AiMeshMorphAnim
	2,  // 3: pb_msg.AiMeshMorphAnim.Keys:type_name -> pb_msg.AiMeshMorphKey
	5,  // 4: pb_msg.AiNodeAnim.PositionKeys:type_name -> pb_msg.AiVectorKey
	6,  // 5: pb_msg.AiNodeAnim.RotationKeys:type_name -> pb_msg.AiQuatKey
	5,  // 6: pb_msg.AiNodeAnim.ScalingKeys:type_name -> pb_msg.AiVectorKey
	7,  // 7: pb_msg.AiMeshAnim.Keys:type_name -> pb_msg.AiMeshKey
	8,  // 8: pb_msg.AiVectorKey.Value:type_name -> pb_msg.AiVector3D
	9,  // 9: pb_msg.AiQuatKey.Value:type_name -> pb_msg.AiQuaternion
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_common_pb_msg_ai_animation_proto_init() }
func file_common_pb_msg_ai_animation_proto_init() {
	if File_common_pb_msg_ai_animation_proto != nil {
		return
	}
	file_common_pb_msg_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_pb_msg_ai_animation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiAnimation); i {
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
		file_common_pb_msg_ai_animation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiMeshMorphAnim); i {
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
		file_common_pb_msg_ai_animation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiMeshMorphKey); i {
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
		file_common_pb_msg_ai_animation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiNodeAnim); i {
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
		file_common_pb_msg_ai_animation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiMeshAnim); i {
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
		file_common_pb_msg_ai_animation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiVectorKey); i {
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
		file_common_pb_msg_ai_animation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiQuatKey); i {
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
		file_common_pb_msg_ai_animation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiMeshKey); i {
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
			RawDescriptor: file_common_pb_msg_ai_animation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pb_msg_ai_animation_proto_goTypes,
		DependencyIndexes: file_common_pb_msg_ai_animation_proto_depIdxs,
		MessageInfos:      file_common_pb_msg_ai_animation_proto_msgTypes,
	}.Build()
	File_common_pb_msg_ai_animation_proto = out.File
	file_common_pb_msg_ai_animation_proto_rawDesc = nil
	file_common_pb_msg_ai_animation_proto_goTypes = nil
	file_common_pb_msg_ai_animation_proto_depIdxs = nil
}
