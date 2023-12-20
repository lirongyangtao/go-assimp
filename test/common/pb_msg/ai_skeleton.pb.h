// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: common/pb_msg/ai_skeleton.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_common_2fpb_5fmsg_2fai_5fskeleton_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_common_2fpb_5fmsg_2fai_5fskeleton_2eproto_2epb_2eh

#include <limits>
#include <string>
#include <type_traits>

#include "google/protobuf/port_def.inc"
#if PROTOBUF_VERSION < 4024000
#error "This file was generated by a newer version of protoc which is"
#error "incompatible with your Protocol Buffer headers. Please update"
#error "your headers."
#endif  // PROTOBUF_VERSION

#if 4024000 < PROTOBUF_MIN_PROTOC_VERSION
#error "This file was generated by an older version of protoc which is"
#error "incompatible with your Protocol Buffer headers. Please"
#error "regenerate this file with a newer version of protoc."
#endif  // PROTOBUF_MIN_PROTOC_VERSION
#include "google/protobuf/port_undef.inc"
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/arena.h"
#include "google/protobuf/arenastring.h"
#include "google/protobuf/generated_message_tctable_decl.h"
#include "google/protobuf/generated_message_util.h"
#include "google/protobuf/metadata_lite.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/message.h"
#include "google/protobuf/repeated_field.h"  // IWYU pragma: export
#include "google/protobuf/extension_set.h"  // IWYU pragma: export
#include "google/protobuf/unknown_field_set.h"
#include "common/pb_msg/common.pb.h"
#include "common/pb_msg/ai_node.pb.h"
#include "common/pb_msg/ai_vertex_weight.pb.h"
#include "common/pb_msg/ai_mesh.pb.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"

#define PROTOBUF_INTERNAL_EXPORT_common_2fpb_5fmsg_2fai_5fskeleton_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_common_2fpb_5fmsg_2fai_5fskeleton_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_common_2fpb_5fmsg_2fai_5fskeleton_2eproto;
namespace pb_msg {
class AiSkeleton;
struct AiSkeletonDefaultTypeInternal;
extern AiSkeletonDefaultTypeInternal _AiSkeleton_default_instance_;
class AiSkeletonBone;
struct AiSkeletonBoneDefaultTypeInternal;
extern AiSkeletonBoneDefaultTypeInternal _AiSkeletonBone_default_instance_;
}  // namespace pb_msg
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace pb_msg {

// ===================================================================


// -------------------------------------------------------------------

class AiSkeleton final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb_msg.AiSkeleton) */ {
 public:
  inline AiSkeleton() : AiSkeleton(nullptr) {}
  ~AiSkeleton() override;
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR AiSkeleton(::google::protobuf::internal::ConstantInitialized);

  AiSkeleton(const AiSkeleton& from);
  AiSkeleton(AiSkeleton&& from) noexcept
    : AiSkeleton() {
    *this = ::std::move(from);
  }

  inline AiSkeleton& operator=(const AiSkeleton& from) {
    CopyFrom(from);
    return *this;
  }
  inline AiSkeleton& operator=(AiSkeleton&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const AiSkeleton& default_instance() {
    return *internal_default_instance();
  }
  static inline const AiSkeleton* internal_default_instance() {
    return reinterpret_cast<const AiSkeleton*>(
               &_AiSkeleton_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(AiSkeleton& a, AiSkeleton& b) {
    a.Swap(&b);
  }
  inline void Swap(AiSkeleton* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(AiSkeleton* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  AiSkeleton* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<AiSkeleton>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const AiSkeleton& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom( const AiSkeleton& from) {
    AiSkeleton::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  ::size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::google::protobuf::internal::ParseContext* ctx) final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target, ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(AiSkeleton* other);

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "pb_msg.AiSkeleton";
  }
  protected:
  explicit AiSkeleton(::google::protobuf::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::google::protobuf::Message::ClassData*GetClassData() const final;

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kBonesFieldNumber = 2,
    kNameFieldNumber = 1,
  };
  // repeated .pb_msg.AiSkeletonBone Bones = 2;
  int bones_size() const;
  private:
  int _internal_bones_size() const;

  public:
  void clear_bones() ;
  ::pb_msg::AiSkeletonBone* mutable_bones(int index);
  ::google::protobuf::RepeatedPtrField< ::pb_msg::AiSkeletonBone >*
      mutable_bones();
  private:
  const ::google::protobuf::RepeatedPtrField<::pb_msg::AiSkeletonBone>& _internal_bones() const;
  ::google::protobuf::RepeatedPtrField<::pb_msg::AiSkeletonBone>* _internal_mutable_bones();
  public:
  const ::pb_msg::AiSkeletonBone& bones(int index) const;
  ::pb_msg::AiSkeletonBone* add_bones();
  const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiSkeletonBone >&
      bones() const;
  // string Name = 1;
  void clear_name() ;
  const std::string& name() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_name(Arg_&& arg, Args_... args);
  std::string* mutable_name();
  PROTOBUF_NODISCARD std::string* release_name();
  void set_allocated_name(std::string* ptr);

  private:
  const std::string& _internal_name() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_name(
      const std::string& value);
  std::string* _internal_mutable_name();

  public:
  // @@protoc_insertion_point(class_scope:pb_msg.AiSkeleton)
 private:
  class _Internal;

  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<1, 2, 1, 30, 2> _table_;
  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::google::protobuf::RepeatedPtrField< ::pb_msg::AiSkeletonBone > bones_;
    ::google::protobuf::internal::ArenaStringPtr name_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_common_2fpb_5fmsg_2fai_5fskeleton_2eproto;
};// -------------------------------------------------------------------

class AiSkeletonBone final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb_msg.AiSkeletonBone) */ {
 public:
  inline AiSkeletonBone() : AiSkeletonBone(nullptr) {}
  ~AiSkeletonBone() override;
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR AiSkeletonBone(::google::protobuf::internal::ConstantInitialized);

  AiSkeletonBone(const AiSkeletonBone& from);
  AiSkeletonBone(AiSkeletonBone&& from) noexcept
    : AiSkeletonBone() {
    *this = ::std::move(from);
  }

  inline AiSkeletonBone& operator=(const AiSkeletonBone& from) {
    CopyFrom(from);
    return *this;
  }
  inline AiSkeletonBone& operator=(AiSkeletonBone&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const AiSkeletonBone& default_instance() {
    return *internal_default_instance();
  }
  static inline const AiSkeletonBone* internal_default_instance() {
    return reinterpret_cast<const AiSkeletonBone*>(
               &_AiSkeletonBone_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(AiSkeletonBone& a, AiSkeletonBone& b) {
    a.Swap(&b);
  }
  inline void Swap(AiSkeletonBone* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(AiSkeletonBone* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  AiSkeletonBone* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<AiSkeletonBone>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const AiSkeletonBone& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom( const AiSkeletonBone& from) {
    AiSkeletonBone::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  ::size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::google::protobuf::internal::ParseContext* ctx) final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target, ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(AiSkeletonBone* other);

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "pb_msg.AiSkeletonBone";
  }
  protected:
  explicit AiSkeletonBone(::google::protobuf::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::google::protobuf::Message::ClassData*GetClassData() const final;

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kArmatureFieldNumber = 2,
    kNodeFieldNumber = 3,
    kMeshIdFieldNumber = 5,
    kWeightsFieldNumber = 6,
    kOffsetMatrixFieldNumber = 7,
    kLocalMatrixFieldNumber = 8,
    kParentFieldNumber = 1,
  };
  // repeated .pb_msg.AiNode Armature = 2;
  int armature_size() const;
  private:
  int _internal_armature_size() const;

  public:
  void clear_armature() ;
  ::pb_msg::AiNode* mutable_armature(int index);
  ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >*
      mutable_armature();
  private:
  const ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>& _internal_armature() const;
  ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>* _internal_mutable_armature();
  public:
  const ::pb_msg::AiNode& armature(int index) const;
  ::pb_msg::AiNode* add_armature();
  const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >&
      armature() const;
  // repeated .pb_msg.AiNode Node = 3;
  int node_size() const;
  private:
  int _internal_node_size() const;

  public:
  void clear_node() ;
  ::pb_msg::AiNode* mutable_node(int index);
  ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >*
      mutable_node();
  private:
  const ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>& _internal_node() const;
  ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>* _internal_mutable_node();
  public:
  const ::pb_msg::AiNode& node(int index) const;
  ::pb_msg::AiNode* add_node();
  const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >&
      node() const;
  // repeated .pb_msg.AiMesh MeshId = 5;
  int meshid_size() const;
  private:
  int _internal_meshid_size() const;

  public:
  void clear_meshid() ;
  ::pb_msg::AiMesh* mutable_meshid(int index);
  ::google::protobuf::RepeatedPtrField< ::pb_msg::AiMesh >*
      mutable_meshid();
  private:
  const ::google::protobuf::RepeatedPtrField<::pb_msg::AiMesh>& _internal_meshid() const;
  ::google::protobuf::RepeatedPtrField<::pb_msg::AiMesh>* _internal_mutable_meshid();
  public:
  const ::pb_msg::AiMesh& meshid(int index) const;
  ::pb_msg::AiMesh* add_meshid();
  const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiMesh >&
      meshid() const;
  // repeated .pb_msg.AiVertexWeight Weights = 6;
  int weights_size() const;
  private:
  int _internal_weights_size() const;

  public:
  void clear_weights() ;
  ::pb_msg::AiVertexWeight* mutable_weights(int index);
  ::google::protobuf::RepeatedPtrField< ::pb_msg::AiVertexWeight >*
      mutable_weights();
  private:
  const ::google::protobuf::RepeatedPtrField<::pb_msg::AiVertexWeight>& _internal_weights() const;
  ::google::protobuf::RepeatedPtrField<::pb_msg::AiVertexWeight>* _internal_mutable_weights();
  public:
  const ::pb_msg::AiVertexWeight& weights(int index) const;
  ::pb_msg::AiVertexWeight* add_weights();
  const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiVertexWeight >&
      weights() const;
  // .pb_msg.AiMatrix4x4 OffsetMatrix = 7;
  bool has_offsetmatrix() const;
  void clear_offsetmatrix() ;
  const ::pb_msg::AiMatrix4x4& offsetmatrix() const;
  PROTOBUF_NODISCARD ::pb_msg::AiMatrix4x4* release_offsetmatrix();
  ::pb_msg::AiMatrix4x4* mutable_offsetmatrix();
  void set_allocated_offsetmatrix(::pb_msg::AiMatrix4x4* value);
  void unsafe_arena_set_allocated_offsetmatrix(::pb_msg::AiMatrix4x4* value);
  ::pb_msg::AiMatrix4x4* unsafe_arena_release_offsetmatrix();

  private:
  const ::pb_msg::AiMatrix4x4& _internal_offsetmatrix() const;
  ::pb_msg::AiMatrix4x4* _internal_mutable_offsetmatrix();

  public:
  // .pb_msg.AiMatrix4x4 LocalMatrix = 8;
  bool has_localmatrix() const;
  void clear_localmatrix() ;
  const ::pb_msg::AiMatrix4x4& localmatrix() const;
  PROTOBUF_NODISCARD ::pb_msg::AiMatrix4x4* release_localmatrix();
  ::pb_msg::AiMatrix4x4* mutable_localmatrix();
  void set_allocated_localmatrix(::pb_msg::AiMatrix4x4* value);
  void unsafe_arena_set_allocated_localmatrix(::pb_msg::AiMatrix4x4* value);
  ::pb_msg::AiMatrix4x4* unsafe_arena_release_localmatrix();

  private:
  const ::pb_msg::AiMatrix4x4& _internal_localmatrix() const;
  ::pb_msg::AiMatrix4x4* _internal_mutable_localmatrix();

  public:
  // int32 Parent = 1;
  void clear_parent() ;
  ::int32_t parent() const;
  void set_parent(::int32_t value);

  private:
  ::int32_t _internal_parent() const;
  void _internal_set_parent(::int32_t value);

  public:
  // @@protoc_insertion_point(class_scope:pb_msg.AiSkeletonBone)
 private:
  class _Internal;

  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<3, 7, 6, 0, 2> _table_;
  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::google::protobuf::internal::HasBits<1> _has_bits_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode > armature_;
    ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode > node_;
    ::google::protobuf::RepeatedPtrField< ::pb_msg::AiMesh > meshid_;
    ::google::protobuf::RepeatedPtrField< ::pb_msg::AiVertexWeight > weights_;
    ::pb_msg::AiMatrix4x4* offsetmatrix_;
    ::pb_msg::AiMatrix4x4* localmatrix_;
    ::int32_t parent_;
    PROTOBUF_TSAN_DECLARE_MEMBER;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_common_2fpb_5fmsg_2fai_5fskeleton_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// AiSkeleton

// string Name = 1;
inline void AiSkeleton::clear_name() {
  _impl_.name_.ClearToEmpty();
}
inline const std::string& AiSkeleton::name() const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeleton.Name)
  return _internal_name();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void AiSkeleton::set_name(Arg_&& arg,
                                                     Args_... args) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.name_.Set(static_cast<Arg_&&>(arg), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:pb_msg.AiSkeleton.Name)
}
inline std::string* AiSkeleton::mutable_name() {
  std::string* _s = _internal_mutable_name();
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeleton.Name)
  return _s;
}
inline const std::string& AiSkeleton::_internal_name() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.name_.Get();
}
inline void AiSkeleton::_internal_set_name(const std::string& value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.name_.Set(value, GetArenaForAllocation());
}
inline std::string* AiSkeleton::_internal_mutable_name() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  return _impl_.name_.Mutable( GetArenaForAllocation());
}
inline std::string* AiSkeleton::release_name() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:pb_msg.AiSkeleton.Name)
  return _impl_.name_.Release();
}
inline void AiSkeleton::set_allocated_name(std::string* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_.name_.SetAllocated(value, GetArenaForAllocation());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.name_.IsDefault()) {
          _impl_.name_.Set("", GetArenaForAllocation());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:pb_msg.AiSkeleton.Name)
}

// repeated .pb_msg.AiSkeletonBone Bones = 2;
inline int AiSkeleton::_internal_bones_size() const {
  return _internal_bones().size();
}
inline int AiSkeleton::bones_size() const {
  return _internal_bones_size();
}
inline void AiSkeleton::clear_bones() {
  _internal_mutable_bones()->Clear();
}
inline ::pb_msg::AiSkeletonBone* AiSkeleton::mutable_bones(int index) {
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeleton.Bones)
  return _internal_mutable_bones()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::pb_msg::AiSkeletonBone >*
AiSkeleton::mutable_bones() {
  // @@protoc_insertion_point(field_mutable_list:pb_msg.AiSkeleton.Bones)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_bones();
}
inline const ::pb_msg::AiSkeletonBone& AiSkeleton::bones(int index) const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeleton.Bones)
    return _internal_bones().Get(index);
}
inline ::pb_msg::AiSkeletonBone* AiSkeleton::add_bones() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::pb_msg::AiSkeletonBone* _add = _internal_mutable_bones()->Add();
  // @@protoc_insertion_point(field_add:pb_msg.AiSkeleton.Bones)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiSkeletonBone >&
AiSkeleton::bones() const {
  // @@protoc_insertion_point(field_list:pb_msg.AiSkeleton.Bones)
  return _internal_bones();
}
inline const ::google::protobuf::RepeatedPtrField<::pb_msg::AiSkeletonBone>&
AiSkeleton::_internal_bones() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.bones_;
}
inline ::google::protobuf::RepeatedPtrField<::pb_msg::AiSkeletonBone>*
AiSkeleton::_internal_mutable_bones() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.bones_;
}

// -------------------------------------------------------------------

// AiSkeletonBone

// int32 Parent = 1;
inline void AiSkeletonBone::clear_parent() {
  _impl_.parent_ = 0;
}
inline ::int32_t AiSkeletonBone::parent() const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeletonBone.Parent)
  return _internal_parent();
}
inline void AiSkeletonBone::set_parent(::int32_t value) {
  _internal_set_parent(value);
  // @@protoc_insertion_point(field_set:pb_msg.AiSkeletonBone.Parent)
}
inline ::int32_t AiSkeletonBone::_internal_parent() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.parent_;
}
inline void AiSkeletonBone::_internal_set_parent(::int32_t value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.parent_ = value;
}

// repeated .pb_msg.AiNode Armature = 2;
inline int AiSkeletonBone::_internal_armature_size() const {
  return _internal_armature().size();
}
inline int AiSkeletonBone::armature_size() const {
  return _internal_armature_size();
}
inline ::pb_msg::AiNode* AiSkeletonBone::mutable_armature(int index) {
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeletonBone.Armature)
  return _internal_mutable_armature()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >*
AiSkeletonBone::mutable_armature() {
  // @@protoc_insertion_point(field_mutable_list:pb_msg.AiSkeletonBone.Armature)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_armature();
}
inline const ::pb_msg::AiNode& AiSkeletonBone::armature(int index) const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeletonBone.Armature)
    return _internal_armature().Get(index);
}
inline ::pb_msg::AiNode* AiSkeletonBone::add_armature() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::pb_msg::AiNode* _add = _internal_mutable_armature()->Add();
  // @@protoc_insertion_point(field_add:pb_msg.AiSkeletonBone.Armature)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >&
AiSkeletonBone::armature() const {
  // @@protoc_insertion_point(field_list:pb_msg.AiSkeletonBone.Armature)
  return _internal_armature();
}
inline const ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>&
AiSkeletonBone::_internal_armature() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.armature_;
}
inline ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>*
AiSkeletonBone::_internal_mutable_armature() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.armature_;
}

// repeated .pb_msg.AiNode Node = 3;
inline int AiSkeletonBone::_internal_node_size() const {
  return _internal_node().size();
}
inline int AiSkeletonBone::node_size() const {
  return _internal_node_size();
}
inline ::pb_msg::AiNode* AiSkeletonBone::mutable_node(int index) {
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeletonBone.Node)
  return _internal_mutable_node()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >*
AiSkeletonBone::mutable_node() {
  // @@protoc_insertion_point(field_mutable_list:pb_msg.AiSkeletonBone.Node)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_node();
}
inline const ::pb_msg::AiNode& AiSkeletonBone::node(int index) const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeletonBone.Node)
    return _internal_node().Get(index);
}
inline ::pb_msg::AiNode* AiSkeletonBone::add_node() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::pb_msg::AiNode* _add = _internal_mutable_node()->Add();
  // @@protoc_insertion_point(field_add:pb_msg.AiSkeletonBone.Node)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiNode >&
AiSkeletonBone::node() const {
  // @@protoc_insertion_point(field_list:pb_msg.AiSkeletonBone.Node)
  return _internal_node();
}
inline const ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>&
AiSkeletonBone::_internal_node() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.node_;
}
inline ::google::protobuf::RepeatedPtrField<::pb_msg::AiNode>*
AiSkeletonBone::_internal_mutable_node() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.node_;
}

// repeated .pb_msg.AiMesh MeshId = 5;
inline int AiSkeletonBone::_internal_meshid_size() const {
  return _internal_meshid().size();
}
inline int AiSkeletonBone::meshid_size() const {
  return _internal_meshid_size();
}
inline ::pb_msg::AiMesh* AiSkeletonBone::mutable_meshid(int index) {
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeletonBone.MeshId)
  return _internal_mutable_meshid()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::pb_msg::AiMesh >*
AiSkeletonBone::mutable_meshid() {
  // @@protoc_insertion_point(field_mutable_list:pb_msg.AiSkeletonBone.MeshId)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_meshid();
}
inline const ::pb_msg::AiMesh& AiSkeletonBone::meshid(int index) const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeletonBone.MeshId)
    return _internal_meshid().Get(index);
}
inline ::pb_msg::AiMesh* AiSkeletonBone::add_meshid() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::pb_msg::AiMesh* _add = _internal_mutable_meshid()->Add();
  // @@protoc_insertion_point(field_add:pb_msg.AiSkeletonBone.MeshId)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiMesh >&
AiSkeletonBone::meshid() const {
  // @@protoc_insertion_point(field_list:pb_msg.AiSkeletonBone.MeshId)
  return _internal_meshid();
}
inline const ::google::protobuf::RepeatedPtrField<::pb_msg::AiMesh>&
AiSkeletonBone::_internal_meshid() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.meshid_;
}
inline ::google::protobuf::RepeatedPtrField<::pb_msg::AiMesh>*
AiSkeletonBone::_internal_mutable_meshid() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.meshid_;
}

// repeated .pb_msg.AiVertexWeight Weights = 6;
inline int AiSkeletonBone::_internal_weights_size() const {
  return _internal_weights().size();
}
inline int AiSkeletonBone::weights_size() const {
  return _internal_weights_size();
}
inline ::pb_msg::AiVertexWeight* AiSkeletonBone::mutable_weights(int index) {
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeletonBone.Weights)
  return _internal_mutable_weights()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::pb_msg::AiVertexWeight >*
AiSkeletonBone::mutable_weights() {
  // @@protoc_insertion_point(field_mutable_list:pb_msg.AiSkeletonBone.Weights)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_weights();
}
inline const ::pb_msg::AiVertexWeight& AiSkeletonBone::weights(int index) const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeletonBone.Weights)
    return _internal_weights().Get(index);
}
inline ::pb_msg::AiVertexWeight* AiSkeletonBone::add_weights() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::pb_msg::AiVertexWeight* _add = _internal_mutable_weights()->Add();
  // @@protoc_insertion_point(field_add:pb_msg.AiSkeletonBone.Weights)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::pb_msg::AiVertexWeight >&
AiSkeletonBone::weights() const {
  // @@protoc_insertion_point(field_list:pb_msg.AiSkeletonBone.Weights)
  return _internal_weights();
}
inline const ::google::protobuf::RepeatedPtrField<::pb_msg::AiVertexWeight>&
AiSkeletonBone::_internal_weights() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.weights_;
}
inline ::google::protobuf::RepeatedPtrField<::pb_msg::AiVertexWeight>*
AiSkeletonBone::_internal_mutable_weights() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.weights_;
}

// .pb_msg.AiMatrix4x4 OffsetMatrix = 7;
inline bool AiSkeletonBone::has_offsetmatrix() const {
  bool value = (_impl_._has_bits_[0] & 0x00000001u) != 0;
  PROTOBUF_ASSUME(!value || _impl_.offsetmatrix_ != nullptr);
  return value;
}
inline const ::pb_msg::AiMatrix4x4& AiSkeletonBone::_internal_offsetmatrix() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  const ::pb_msg::AiMatrix4x4* p = _impl_.offsetmatrix_;
  return p != nullptr ? *p : reinterpret_cast<const ::pb_msg::AiMatrix4x4&>(::pb_msg::_AiMatrix4x4_default_instance_);
}
inline const ::pb_msg::AiMatrix4x4& AiSkeletonBone::offsetmatrix() const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeletonBone.OffsetMatrix)
  return _internal_offsetmatrix();
}
inline void AiSkeletonBone::unsafe_arena_set_allocated_offsetmatrix(::pb_msg::AiMatrix4x4* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.offsetmatrix_);
  }
  _impl_.offsetmatrix_ = reinterpret_cast<::pb_msg::AiMatrix4x4*>(value);
  if (value != nullptr) {
    _impl_._has_bits_[0] |= 0x00000001u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000001u;
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:pb_msg.AiSkeletonBone.OffsetMatrix)
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::release_offsetmatrix() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);

  _impl_._has_bits_[0] &= ~0x00000001u;
  ::pb_msg::AiMatrix4x4* released = _impl_.offsetmatrix_;
  _impl_.offsetmatrix_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old = reinterpret_cast<::google::protobuf::MessageLite*>(released);
  released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  if (GetArenaForAllocation() == nullptr) {
    delete old;
  }
#else   // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return released;
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::unsafe_arena_release_offsetmatrix() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:pb_msg.AiSkeletonBone.OffsetMatrix)

  _impl_._has_bits_[0] &= ~0x00000001u;
  ::pb_msg::AiMatrix4x4* temp = _impl_.offsetmatrix_;
  _impl_.offsetmatrix_ = nullptr;
  return temp;
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::_internal_mutable_offsetmatrix() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_._has_bits_[0] |= 0x00000001u;
  if (_impl_.offsetmatrix_ == nullptr) {
    auto* p = CreateMaybeMessage<::pb_msg::AiMatrix4x4>(GetArenaForAllocation());
    _impl_.offsetmatrix_ = reinterpret_cast<::pb_msg::AiMatrix4x4*>(p);
  }
  return _impl_.offsetmatrix_;
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::mutable_offsetmatrix() {
  ::pb_msg::AiMatrix4x4* _msg = _internal_mutable_offsetmatrix();
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeletonBone.OffsetMatrix)
  return _msg;
}
inline void AiSkeletonBone::set_allocated_offsetmatrix(::pb_msg::AiMatrix4x4* value) {
  ::google::protobuf::Arena* message_arena = GetArenaForAllocation();
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (message_arena == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.offsetmatrix_);
  }

  if (value != nullptr) {
    ::google::protobuf::Arena* submessage_arena =
        ::google::protobuf::Arena::InternalGetOwningArena(reinterpret_cast<::google::protobuf::MessageLite*>(value));
    if (message_arena != submessage_arena) {
      value = ::google::protobuf::internal::GetOwnedMessage(message_arena, value, submessage_arena);
    }
    _impl_._has_bits_[0] |= 0x00000001u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000001u;
  }

  _impl_.offsetmatrix_ = reinterpret_cast<::pb_msg::AiMatrix4x4*>(value);
  // @@protoc_insertion_point(field_set_allocated:pb_msg.AiSkeletonBone.OffsetMatrix)
}

// .pb_msg.AiMatrix4x4 LocalMatrix = 8;
inline bool AiSkeletonBone::has_localmatrix() const {
  bool value = (_impl_._has_bits_[0] & 0x00000002u) != 0;
  PROTOBUF_ASSUME(!value || _impl_.localmatrix_ != nullptr);
  return value;
}
inline const ::pb_msg::AiMatrix4x4& AiSkeletonBone::_internal_localmatrix() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  const ::pb_msg::AiMatrix4x4* p = _impl_.localmatrix_;
  return p != nullptr ? *p : reinterpret_cast<const ::pb_msg::AiMatrix4x4&>(::pb_msg::_AiMatrix4x4_default_instance_);
}
inline const ::pb_msg::AiMatrix4x4& AiSkeletonBone::localmatrix() const {
  // @@protoc_insertion_point(field_get:pb_msg.AiSkeletonBone.LocalMatrix)
  return _internal_localmatrix();
}
inline void AiSkeletonBone::unsafe_arena_set_allocated_localmatrix(::pb_msg::AiMatrix4x4* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.localmatrix_);
  }
  _impl_.localmatrix_ = reinterpret_cast<::pb_msg::AiMatrix4x4*>(value);
  if (value != nullptr) {
    _impl_._has_bits_[0] |= 0x00000002u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000002u;
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:pb_msg.AiSkeletonBone.LocalMatrix)
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::release_localmatrix() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);

  _impl_._has_bits_[0] &= ~0x00000002u;
  ::pb_msg::AiMatrix4x4* released = _impl_.localmatrix_;
  _impl_.localmatrix_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old = reinterpret_cast<::google::protobuf::MessageLite*>(released);
  released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  if (GetArenaForAllocation() == nullptr) {
    delete old;
  }
#else   // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return released;
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::unsafe_arena_release_localmatrix() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:pb_msg.AiSkeletonBone.LocalMatrix)

  _impl_._has_bits_[0] &= ~0x00000002u;
  ::pb_msg::AiMatrix4x4* temp = _impl_.localmatrix_;
  _impl_.localmatrix_ = nullptr;
  return temp;
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::_internal_mutable_localmatrix() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_._has_bits_[0] |= 0x00000002u;
  if (_impl_.localmatrix_ == nullptr) {
    auto* p = CreateMaybeMessage<::pb_msg::AiMatrix4x4>(GetArenaForAllocation());
    _impl_.localmatrix_ = reinterpret_cast<::pb_msg::AiMatrix4x4*>(p);
  }
  return _impl_.localmatrix_;
}
inline ::pb_msg::AiMatrix4x4* AiSkeletonBone::mutable_localmatrix() {
  ::pb_msg::AiMatrix4x4* _msg = _internal_mutable_localmatrix();
  // @@protoc_insertion_point(field_mutable:pb_msg.AiSkeletonBone.LocalMatrix)
  return _msg;
}
inline void AiSkeletonBone::set_allocated_localmatrix(::pb_msg::AiMatrix4x4* value) {
  ::google::protobuf::Arena* message_arena = GetArenaForAllocation();
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (message_arena == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.localmatrix_);
  }

  if (value != nullptr) {
    ::google::protobuf::Arena* submessage_arena =
        ::google::protobuf::Arena::InternalGetOwningArena(reinterpret_cast<::google::protobuf::MessageLite*>(value));
    if (message_arena != submessage_arena) {
      value = ::google::protobuf::internal::GetOwnedMessage(message_arena, value, submessage_arena);
    }
    _impl_._has_bits_[0] |= 0x00000002u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000002u;
  }

  _impl_.localmatrix_ = reinterpret_cast<::pb_msg::AiMatrix4x4*>(value);
  // @@protoc_insertion_point(field_set_allocated:pb_msg.AiSkeletonBone.LocalMatrix)
}

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace pb_msg


// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_common_2fpb_5fmsg_2fai_5fskeleton_2eproto_2epb_2eh