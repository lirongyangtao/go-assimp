// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: common/pb_msg/ai_meta_data.proto

#include "common/pb_msg/ai_meta_data.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
#include "google/protobuf/generated_message_tctable_impl.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::google::protobuf;
namespace _pbi = ::google::protobuf::internal;
namespace _fl = ::google::protobuf::internal::field_layout;
namespace pb_msg {
        template <typename>
PROTOBUF_CONSTEXPR AiMetadata::AiMetadata(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.keys_)*/ {},
      /*decltype(_impl_.values_)*/ {},
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct AiMetadataDefaultTypeInternal {
  PROTOBUF_CONSTEXPR AiMetadataDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~AiMetadataDefaultTypeInternal() {}
  union {
    AiMetadata _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 AiMetadataDefaultTypeInternal _AiMetadata_default_instance_;
        template <typename>
PROTOBUF_CONSTEXPR AiMetadataEntry::AiMetadataEntry(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.data_)*/ {
          &::_pbi::fixed_address_empty_string,
          ::_pbi::ConstantInitialized{},
      },
      /*decltype(_impl_.type_)*/ 0,
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct AiMetadataEntryDefaultTypeInternal {
  PROTOBUF_CONSTEXPR AiMetadataEntryDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~AiMetadataEntryDefaultTypeInternal() {}
  union {
    AiMetadataEntry _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 AiMetadataEntryDefaultTypeInternal _AiMetadataEntry_default_instance_;
}  // namespace pb_msg
static ::_pb::Metadata file_level_metadata_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto[2];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto = nullptr;
const ::uint32_t TableStruct_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::pb_msg::AiMetadata, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::pb_msg::AiMetadata, _impl_.keys_),
    PROTOBUF_FIELD_OFFSET(::pb_msg::AiMetadata, _impl_.values_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::pb_msg::AiMetadataEntry, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::pb_msg::AiMetadataEntry, _impl_.type_),
    PROTOBUF_FIELD_OFFSET(::pb_msg::AiMetadataEntry, _impl_.data_),
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        {0, -1, -1, sizeof(::pb_msg::AiMetadata)},
        {10, -1, -1, sizeof(::pb_msg::AiMetadataEntry)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::pb_msg::_AiMetadata_default_instance_._instance,
    &::pb_msg::_AiMetadataEntry_default_instance_._instance,
};
const char descriptor_table_protodef_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n common/pb_msg/ai_meta_data.proto\022\006pb_m"
    "sg\"C\n\nAiMetadata\022\014\n\004Keys\030\001 \003(\t\022\'\n\006Values"
    "\030\002 \003(\0132\027.pb_msg.AiMetadataEntry\"-\n\017AiMet"
    "adataEntry\022\014\n\004Type\030\001 \001(\005\022\014\n\004Data\030\002 \001(\014B\017"
    "Z\rcommon/pb_msgb\006proto3"
};
static ::absl::once_flag descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto = {
    false,
    false,
    183,
    descriptor_table_protodef_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto,
    "common/pb_msg/ai_meta_data.proto",
    &descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto_once,
    nullptr,
    0,
    2,
    schemas,
    file_default_instances,
    TableStruct_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto::offsets,
    file_level_metadata_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto,
    file_level_enum_descriptors_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto,
    file_level_service_descriptors_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto_getter() {
  return &descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto(&descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto);
namespace pb_msg {
// ===================================================================

class AiMetadata::_Internal {
 public:
};

AiMetadata::AiMetadata(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:pb_msg.AiMetadata)
}
AiMetadata::AiMetadata(const AiMetadata& from) : ::google::protobuf::Message() {
  AiMetadata* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.keys_){from._impl_.keys_},
      decltype(_impl_.values_){from._impl_.values_},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);

  // @@protoc_insertion_point(copy_constructor:pb_msg.AiMetadata)
}
inline void AiMetadata::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.keys_){arena},
      decltype(_impl_.values_){arena},
      /*decltype(_impl_._cached_size_)*/ {},
  };
}
AiMetadata::~AiMetadata() {
  // @@protoc_insertion_point(destructor:pb_msg.AiMetadata)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void AiMetadata::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _internal_mutable_keys()->~RepeatedPtrField();
  _impl_.values_.~RepeatedPtrField();
}
void AiMetadata::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

PROTOBUF_NOINLINE void AiMetadata::Clear() {
// @@protoc_insertion_point(message_clear_start:pb_msg.AiMetadata)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _internal_mutable_keys()->Clear();
  _internal_mutable_values()->Clear();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* AiMetadata::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<1, 2, 1, 30, 2> AiMetadata::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    2, 8,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967292,  // skipmap
    offsetof(decltype(_table_), field_entries),
    2,  // num_field_entries
    1,  // num_aux_entries
    offsetof(decltype(_table_), aux_entries),
    &_AiMetadata_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    // repeated .pb_msg.AiMetadataEntry Values = 2;
    {::_pbi::TcParser::FastMtR1,
     {18, 63, 0, PROTOBUF_FIELD_OFFSET(AiMetadata, _impl_.values_)}},
    // repeated string Keys = 1;
    {::_pbi::TcParser::FastUR1,
     {10, 63, 0, PROTOBUF_FIELD_OFFSET(AiMetadata, _impl_.keys_)}},
  }}, {{
    65535, 65535
  }}, {{
    // repeated string Keys = 1;
    {PROTOBUF_FIELD_OFFSET(AiMetadata, _impl_.keys_), 0, 0,
    (0 | ::_fl::kFcRepeated | ::_fl::kUtf8String | ::_fl::kRepSString)},
    // repeated .pb_msg.AiMetadataEntry Values = 2;
    {PROTOBUF_FIELD_OFFSET(AiMetadata, _impl_.values_), 0, 0,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
  }}, {{
    {::_pbi::TcParser::GetTable<::pb_msg::AiMetadataEntry>()},
  }}, {{
    "\21\4\0\0\0\0\0\0"
    "pb_msg.AiMetadata"
    "Keys"
  }},
};

::uint8_t* AiMetadata::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:pb_msg.AiMetadata)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // repeated string Keys = 1;
  for (int i = 0, n = this->_internal_keys_size(); i < n; ++i) {
    const auto& s = this->_internal_keys().Get(i);
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        s.data(), static_cast<int>(s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "pb_msg.AiMetadata.Keys");
    target = stream->WriteString(1, s, target);
  }

  // repeated .pb_msg.AiMetadataEntry Values = 2;
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_values_size()); i < n; i++) {
    const auto& repfield = this->_internal_values().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(2, repfield, repfield.GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb_msg.AiMetadata)
  return target;
}

::size_t AiMetadata::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb_msg.AiMetadata)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated string Keys = 1;
  total_size += 1 * ::google::protobuf::internal::FromIntSize(_internal_keys().size());
  for (int i = 0, n = _internal_keys().size(); i < n; ++i) {
    total_size += ::google::protobuf::internal::WireFormatLite::StringSize(
        _internal_keys().Get(i));
  }
  // repeated .pb_msg.AiMetadataEntry Values = 2;
  total_size += 1UL * this->_internal_values_size();
  for (const auto& msg : this->_internal_values()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData AiMetadata::_class_data_ = {
    ::google::protobuf::Message::CopyWithSourceCheck,
    AiMetadata::MergeImpl
};
const ::google::protobuf::Message::ClassData*AiMetadata::GetClassData() const { return &_class_data_; }


void AiMetadata::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<AiMetadata*>(&to_msg);
  auto& from = static_cast<const AiMetadata&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:pb_msg.AiMetadata)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_internal_mutable_keys()->MergeFrom(from._internal_keys());
  _this->_internal_mutable_values()->MergeFrom(from._internal_values());
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void AiMetadata::CopyFrom(const AiMetadata& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb_msg.AiMetadata)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool AiMetadata::IsInitialized() const {
  return true;
}

void AiMetadata::InternalSwap(AiMetadata* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.keys_.InternalSwap(&other->_impl_.keys_);
  _impl_.values_.InternalSwap(&other->_impl_.values_);
}

::google::protobuf::Metadata AiMetadata::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto_getter, &descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto_once,
      file_level_metadata_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto[0]);
}
// ===================================================================

class AiMetadataEntry::_Internal {
 public:
};

AiMetadataEntry::AiMetadataEntry(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:pb_msg.AiMetadataEntry)
}
AiMetadataEntry::AiMetadataEntry(const AiMetadataEntry& from) : ::google::protobuf::Message() {
  AiMetadataEntry* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.data_){},
      decltype(_impl_.type_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  _impl_.data_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.data_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_data().empty()) {
    _this->_impl_.data_.Set(from._internal_data(), _this->GetArenaForAllocation());
  }
  _this->_impl_.type_ = from._impl_.type_;

  // @@protoc_insertion_point(copy_constructor:pb_msg.AiMetadataEntry)
}
inline void AiMetadataEntry::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.data_){},
      decltype(_impl_.type_){0},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _impl_.data_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.data_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}
AiMetadataEntry::~AiMetadataEntry() {
  // @@protoc_insertion_point(destructor:pb_msg.AiMetadataEntry)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void AiMetadataEntry::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.data_.Destroy();
}
void AiMetadataEntry::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

PROTOBUF_NOINLINE void AiMetadataEntry::Clear() {
// @@protoc_insertion_point(message_clear_start:pb_msg.AiMetadataEntry)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.data_.ClearToEmpty();
  _impl_.type_ = 0;
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* AiMetadataEntry::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<1, 2, 0, 0, 2> AiMetadataEntry::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    2, 8,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967292,  // skipmap
    offsetof(decltype(_table_), field_entries),
    2,  // num_field_entries
    0,  // num_aux_entries
    offsetof(decltype(_table_), field_names),  // no aux_entries
    &_AiMetadataEntry_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    // bytes Data = 2;
    {::_pbi::TcParser::FastBS1,
     {18, 63, 0, PROTOBUF_FIELD_OFFSET(AiMetadataEntry, _impl_.data_)}},
    // int32 Type = 1;
    {::_pbi::TcParser::SingularVarintNoZag1<::uint32_t, offsetof(AiMetadataEntry, _impl_.type_), 63>(),
     {8, 63, 0, PROTOBUF_FIELD_OFFSET(AiMetadataEntry, _impl_.type_)}},
  }}, {{
    65535, 65535
  }}, {{
    // int32 Type = 1;
    {PROTOBUF_FIELD_OFFSET(AiMetadataEntry, _impl_.type_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kInt32)},
    // bytes Data = 2;
    {PROTOBUF_FIELD_OFFSET(AiMetadataEntry, _impl_.data_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kBytes | ::_fl::kRepAString)},
  }},
  // no aux_entries
  {{
  }},
};

::uint8_t* AiMetadataEntry::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:pb_msg.AiMetadataEntry)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // int32 Type = 1;
  if (this->_internal_type() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::
        WriteInt32ToArrayWithField<1>(
            stream, this->_internal_type(), target);
  }

  // bytes Data = 2;
  if (!this->_internal_data().empty()) {
    const std::string& _s = this->_internal_data();
    target = stream->WriteBytesMaybeAliased(2, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb_msg.AiMetadataEntry)
  return target;
}

::size_t AiMetadataEntry::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb_msg.AiMetadataEntry)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // bytes Data = 2;
  if (!this->_internal_data().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::BytesSize(
                                    this->_internal_data());
  }

  // int32 Type = 1;
  if (this->_internal_type() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(
        this->_internal_type());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData AiMetadataEntry::_class_data_ = {
    ::google::protobuf::Message::CopyWithSourceCheck,
    AiMetadataEntry::MergeImpl
};
const ::google::protobuf::Message::ClassData*AiMetadataEntry::GetClassData() const { return &_class_data_; }


void AiMetadataEntry::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<AiMetadataEntry*>(&to_msg);
  auto& from = static_cast<const AiMetadataEntry&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:pb_msg.AiMetadataEntry)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_data().empty()) {
    _this->_internal_set_data(from._internal_data());
  }
  if (from._internal_type() != 0) {
    _this->_internal_set_type(from._internal_type());
  }
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void AiMetadataEntry::CopyFrom(const AiMetadataEntry& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb_msg.AiMetadataEntry)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool AiMetadataEntry::IsInitialized() const {
  return true;
}

void AiMetadataEntry::InternalSwap(AiMetadataEntry* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.data_, lhs_arena,
                                       &other->_impl_.data_, rhs_arena);
        swap(_impl_.type_, other->_impl_.type_);
}

::google::protobuf::Metadata AiMetadataEntry::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto_getter, &descriptor_table_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto_once,
      file_level_metadata_common_2fpb_5fmsg_2fai_5fmeta_5fdata_2eproto[1]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace pb_msg
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"