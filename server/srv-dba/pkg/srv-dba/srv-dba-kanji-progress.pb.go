// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: kioku/srv-dba/v1/srv-dba-kanji-progress.proto

package srv_dba

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KanjiProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kanji      *Kanji                 `protobuf:"bytes,2,opt,name=kanji,proto3" json:"kanji,omitempty"`
	SrsLevel   uint32                 `protobuf:"varint,3,opt,name=srs_level,json=srsLevel,proto3" json:"srs_level,omitempty"`
	UnlockDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=unlock_date,json=unlockDate,proto3" json:"unlock_date,omitempty"`
	NextDate   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=next_date,json=nextDate,proto3" json:"next_date,omitempty"`
	BurnDate   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=burn_date,json=burnDate,proto3" json:"burn_date,omitempty"`
}

func (x *KanjiProgress) Reset() {
	*x = KanjiProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KanjiProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KanjiProgress) ProtoMessage() {}

func (x *KanjiProgress) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KanjiProgress.ProtoReflect.Descriptor instead.
func (*KanjiProgress) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{0}
}

func (x *KanjiProgress) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KanjiProgress) GetKanji() *Kanji {
	if x != nil {
		return x.Kanji
	}
	return nil
}

func (x *KanjiProgress) GetSrsLevel() uint32 {
	if x != nil {
		return x.SrsLevel
	}
	return 0
}

func (x *KanjiProgress) GetUnlockDate() *timestamppb.Timestamp {
	if x != nil {
		return x.UnlockDate
	}
	return nil
}

func (x *KanjiProgress) GetNextDate() *timestamppb.Timestamp {
	if x != nil {
		return x.NextDate
	}
	return nil
}

func (x *KanjiProgress) GetBurnDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BurnDate
	}
	return nil
}

type GetKanjiProgressByIdV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	KanjiId uint64 `protobuf:"varint,2,opt,name=kanji_id,json=kanjiId,proto3" json:"kanji_id,omitempty"`
}

func (x *GetKanjiProgressByIdV1Request) Reset() {
	*x = GetKanjiProgressByIdV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKanjiProgressByIdV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKanjiProgressByIdV1Request) ProtoMessage() {}

func (x *GetKanjiProgressByIdV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKanjiProgressByIdV1Request.ProtoReflect.Descriptor instead.
func (*GetKanjiProgressByIdV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{1}
}

func (x *GetKanjiProgressByIdV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetKanjiProgressByIdV1Request) GetKanjiId() uint64 {
	if x != nil {
		return x.KanjiId
	}
	return 0
}

type GetKanjiProgressByIdV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KanjiProgress *KanjiProgress `protobuf:"bytes,1,opt,name=kanji_progress,json=kanjiProgress,proto3" json:"kanji_progress,omitempty"`
}

func (x *GetKanjiProgressByIdV1Response) Reset() {
	*x = GetKanjiProgressByIdV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKanjiProgressByIdV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKanjiProgressByIdV1Response) ProtoMessage() {}

func (x *GetKanjiProgressByIdV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKanjiProgressByIdV1Response.ProtoReflect.Descriptor instead.
func (*GetKanjiProgressByIdV1Response) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{2}
}

func (x *GetKanjiProgressByIdV1Response) GetKanjiProgress() *KanjiProgress {
	if x != nil {
		return x.KanjiProgress
	}
	return nil
}

type ListKanjiProgressByIdsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	KanjiId []uint64 `protobuf:"varint,2,rep,packed,name=kanji_id,json=kanjiId,proto3" json:"kanji_id,omitempty"`
}

func (x *ListKanjiProgressByIdsV1Request) Reset() {
	*x = ListKanjiProgressByIdsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKanjiProgressByIdsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKanjiProgressByIdsV1Request) ProtoMessage() {}

func (x *ListKanjiProgressByIdsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKanjiProgressByIdsV1Request.ProtoReflect.Descriptor instead.
func (*ListKanjiProgressByIdsV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{3}
}

func (x *ListKanjiProgressByIdsV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListKanjiProgressByIdsV1Request) GetKanjiId() []uint64 {
	if x != nil {
		return x.KanjiId
	}
	return nil
}

type ListKanjiProgressByTimeV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListKanjiProgressByTimeV1Request) Reset() {
	*x = ListKanjiProgressByTimeV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKanjiProgressByTimeV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKanjiProgressByTimeV1Request) ProtoMessage() {}

func (x *ListKanjiProgressByTimeV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKanjiProgressByTimeV1Request.ProtoReflect.Descriptor instead.
func (*ListKanjiProgressByTimeV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{4}
}

func (x *ListKanjiProgressByTimeV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListKanjiProgressByTimeV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListKanjiProgressByTimeV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListKanjiProgressBySrsLevelV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SrsLevel uint32 `protobuf:"varint,2,opt,name=srs_level,json=srsLevel,proto3" json:"srs_level,omitempty"`
	Limit    uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   uint64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListKanjiProgressBySrsLevelV1Request) Reset() {
	*x = ListKanjiProgressBySrsLevelV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKanjiProgressBySrsLevelV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKanjiProgressBySrsLevelV1Request) ProtoMessage() {}

func (x *ListKanjiProgressBySrsLevelV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKanjiProgressBySrsLevelV1Request.ProtoReflect.Descriptor instead.
func (*ListKanjiProgressBySrsLevelV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{5}
}

func (x *ListKanjiProgressBySrsLevelV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListKanjiProgressBySrsLevelV1Request) GetSrsLevel() uint32 {
	if x != nil {
		return x.SrsLevel
	}
	return 0
}

func (x *ListKanjiProgressBySrsLevelV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListKanjiProgressBySrsLevelV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListKanjiProgressV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KanjiProgress []*KanjiProgress `protobuf:"bytes,1,rep,name=kanji_progress,json=kanjiProgress,proto3" json:"kanji_progress,omitempty"`
}

func (x *ListKanjiProgressV1Response) Reset() {
	*x = ListKanjiProgressV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKanjiProgressV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKanjiProgressV1Response) ProtoMessage() {}

func (x *ListKanjiProgressV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKanjiProgressV1Response.ProtoReflect.Descriptor instead.
func (*ListKanjiProgressV1Response) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{6}
}

func (x *ListKanjiProgressV1Response) GetKanjiProgress() []*KanjiProgress {
	if x != nil {
		return x.KanjiProgress
	}
	return nil
}

type AddKanjiProgressV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	KanjiId []uint64 `protobuf:"varint,2,rep,packed,name=kanji_id,json=kanjiId,proto3" json:"kanji_id,omitempty"`
}

func (x *AddKanjiProgressV1Request) Reset() {
	*x = AddKanjiProgressV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKanjiProgressV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKanjiProgressV1Request) ProtoMessage() {}

func (x *AddKanjiProgressV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKanjiProgressV1Request.ProtoReflect.Descriptor instead.
func (*AddKanjiProgressV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{7}
}

func (x *AddKanjiProgressV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddKanjiProgressV1Request) GetKanjiId() []uint64 {
	if x != nil {
		return x.KanjiId
	}
	return nil
}

type UpdateKanjiProgressV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProgressId uint64                 `protobuf:"varint,1,opt,name=progress_id,json=progressId,proto3" json:"progress_id,omitempty"`
	SrsLevel   uint32                 `protobuf:"varint,2,opt,name=srs_level,json=srsLevel,proto3" json:"srs_level,omitempty"`
	NextDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=next_date,json=nextDate,proto3" json:"next_date,omitempty"`
	BurnDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=burn_date,json=burnDate,proto3" json:"burn_date,omitempty"`
}

func (x *UpdateKanjiProgressV1Request) Reset() {
	*x = UpdateKanjiProgressV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKanjiProgressV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKanjiProgressV1Request) ProtoMessage() {}

func (x *UpdateKanjiProgressV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKanjiProgressV1Request.ProtoReflect.Descriptor instead.
func (*UpdateKanjiProgressV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateKanjiProgressV1Request) GetProgressId() uint64 {
	if x != nil {
		return x.ProgressId
	}
	return 0
}

func (x *UpdateKanjiProgressV1Request) GetSrsLevel() uint32 {
	if x != nil {
		return x.SrsLevel
	}
	return 0
}

func (x *UpdateKanjiProgressV1Request) GetNextDate() *timestamppb.Timestamp {
	if x != nil {
		return x.NextDate
	}
	return nil
}

func (x *UpdateKanjiProgressV1Request) GetBurnDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BurnDate
	}
	return nil
}

type DefaultKanjiProgressV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DefaultKanjiProgressV1Response) Reset() {
	*x = DefaultKanjiProgressV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultKanjiProgressV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultKanjiProgressV1Response) ProtoMessage() {}

func (x *DefaultKanjiProgressV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultKanjiProgressV1Response.ProtoReflect.Descriptor instead.
func (*DefaultKanjiProgressV1Response) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP(), []int{9}
}

func (x *DefaultKanjiProgressV1Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_kioku_srv_dba_v1_srv_dba_kanji_progress_proto protoreflect.FileDescriptor

var file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2d, 0x6b, 0x61, 0x6e, 0x6a, 0x69,
	0x2d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x72,
	0x76, 0x5f, 0x64, 0x62, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2d, 0x6b, 0x61, 0x6e, 0x6a, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0d, 0x4b, 0x61, 0x6e, 0x6a, 0x69,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x05, 0x6b, 0x61, 0x6e, 0x6a,
	0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x64, 0x62, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x52, 0x05, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x72, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x72, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x62, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x62, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x65, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79,
	0x49, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x08, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x49,
	0x64, 0x22, 0x6f, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x69,
	0x6f, 0x6b, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x64,
	0x62, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0d, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x6c, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x6b, 0x61, 0x6e, 0x6a, 0x69,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x92, 0x01,
	0x06, 0x22, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x49, 0x64,
	0x22, 0x7c, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x32, 0x03, 0x10, 0xe8, 0x07, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xa8,
	0x01, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x53, 0x72, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x73, 0x72, 0x73,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x2a, 0x04, 0x18, 0x09, 0x20, 0x00, 0x52, 0x08, 0x73, 0x72, 0x73, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x32, 0x03, 0x10, 0xe8, 0x07, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x6c, 0x0a, 0x1b, 0x4c, 0x69, 0x73,
	0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x6b, 0x61, 0x6e, 0x6a,
	0x69, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x73, 0x72, 0x76, 0x5f, 0x64, 0x62, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x6e, 0x6a, 0x69,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x66, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x4b, 0x61,
	0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06,
	0x22, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x49, 0x64, 0x22,
	0xe2, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x73, 0x72,
	0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x2a, 0x04, 0x18, 0x09, 0x20, 0x00, 0x52, 0x08, 0x73, 0x72, 0x73, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x62,
	0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x75, 0x72, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x0a, 0x1e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4b,
	0x61, 0x6e, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6f, 0x6d, 0x61, 0x7a, 0x69, 0x73, 0x2f, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x3b, 0x73, 0x72, 0x76, 0x5f, 0x64, 0x62, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescOnce sync.Once
	file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescData = file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDesc
)

func file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescGZIP() []byte {
	file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescOnce.Do(func() {
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescData = protoimpl.X.CompressGZIP(file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescData)
	})
	return file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDescData
}

var file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_goTypes = []interface{}{
	(*KanjiProgress)(nil),                        // 0: kioku.server.srv_dba.v1.KanjiProgress
	(*GetKanjiProgressByIdV1Request)(nil),        // 1: kioku.server.srv_dba.v1.GetKanjiProgressByIdV1Request
	(*GetKanjiProgressByIdV1Response)(nil),       // 2: kioku.server.srv_dba.v1.GetKanjiProgressByIdV1Response
	(*ListKanjiProgressByIdsV1Request)(nil),      // 3: kioku.server.srv_dba.v1.ListKanjiProgressByIdsV1Request
	(*ListKanjiProgressByTimeV1Request)(nil),     // 4: kioku.server.srv_dba.v1.ListKanjiProgressByTimeV1Request
	(*ListKanjiProgressBySrsLevelV1Request)(nil), // 5: kioku.server.srv_dba.v1.ListKanjiProgressBySrsLevelV1Request
	(*ListKanjiProgressV1Response)(nil),          // 6: kioku.server.srv_dba.v1.ListKanjiProgressV1Response
	(*AddKanjiProgressV1Request)(nil),            // 7: kioku.server.srv_dba.v1.AddKanjiProgressV1Request
	(*UpdateKanjiProgressV1Request)(nil),         // 8: kioku.server.srv_dba.v1.UpdateKanjiProgressV1Request
	(*DefaultKanjiProgressV1Response)(nil),       // 9: kioku.server.srv_dba.v1.DefaultKanjiProgressV1Response
	(*Kanji)(nil),                                // 10: kioku.server.srv_dba.v1.Kanji
	(*timestamppb.Timestamp)(nil),                // 11: google.protobuf.Timestamp
}
var file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_depIdxs = []int32{
	10, // 0: kioku.server.srv_dba.v1.KanjiProgress.kanji:type_name -> kioku.server.srv_dba.v1.Kanji
	11, // 1: kioku.server.srv_dba.v1.KanjiProgress.unlock_date:type_name -> google.protobuf.Timestamp
	11, // 2: kioku.server.srv_dba.v1.KanjiProgress.next_date:type_name -> google.protobuf.Timestamp
	11, // 3: kioku.server.srv_dba.v1.KanjiProgress.burn_date:type_name -> google.protobuf.Timestamp
	0,  // 4: kioku.server.srv_dba.v1.GetKanjiProgressByIdV1Response.kanji_progress:type_name -> kioku.server.srv_dba.v1.KanjiProgress
	0,  // 5: kioku.server.srv_dba.v1.ListKanjiProgressV1Response.kanji_progress:type_name -> kioku.server.srv_dba.v1.KanjiProgress
	11, // 6: kioku.server.srv_dba.v1.UpdateKanjiProgressV1Request.next_date:type_name -> google.protobuf.Timestamp
	11, // 7: kioku.server.srv_dba.v1.UpdateKanjiProgressV1Request.burn_date:type_name -> google.protobuf.Timestamp
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_init() }
func file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_init() {
	if File_kioku_srv_dba_v1_srv_dba_kanji_progress_proto != nil {
		return
	}
	file_kioku_srv_dba_v1_srv_dba_kanji_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KanjiProgress); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKanjiProgressByIdV1Request); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKanjiProgressByIdV1Response); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKanjiProgressByIdsV1Request); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKanjiProgressByTimeV1Request); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKanjiProgressBySrsLevelV1Request); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKanjiProgressV1Response); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKanjiProgressV1Request); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKanjiProgressV1Request); i {
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
		file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultKanjiProgressV1Response); i {
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
			RawDescriptor: file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_goTypes,
		DependencyIndexes: file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_depIdxs,
		MessageInfos:      file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_msgTypes,
	}.Build()
	File_kioku_srv_dba_v1_srv_dba_kanji_progress_proto = out.File
	file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_rawDesc = nil
	file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_goTypes = nil
	file_kioku_srv_dba_v1_srv_dba_kanji_progress_proto_depIdxs = nil
}
