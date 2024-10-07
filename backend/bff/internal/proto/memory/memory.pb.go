// memory.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: memory.proto

package memory

import (
	common "bff/internal/proto/common"
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

// MemoryService
type CreateMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DocId         uint64 `protobuf:"varint,2,opt,name=doc_id,json=docId,proto3" json:"doc_id,omitempty"`
	MemoryTitle   string `protobuf:"bytes,3,opt,name=memory_title,json=memoryTitle,proto3" json:"memory_title,omitempty"`
	MemoryContent string `protobuf:"bytes,4,opt,name=memory_content,json=memoryContent,proto3" json:"memory_content,omitempty"`
	IsPublic      bool   `protobuf:"varint,5,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *CreateMemoryRequest) Reset() {
	*x = CreateMemoryRequest{}
	mi := &file_memory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemoryRequest) ProtoMessage() {}

func (x *CreateMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemoryRequest.ProtoReflect.Descriptor instead.
func (*CreateMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMemoryRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateMemoryRequest) GetDocId() uint64 {
	if x != nil {
		return x.DocId
	}
	return 0
}

func (x *CreateMemoryRequest) GetMemoryTitle() string {
	if x != nil {
		return x.MemoryTitle
	}
	return ""
}

func (x *CreateMemoryRequest) GetMemoryContent() string {
	if x != nil {
		return x.MemoryContent
	}
	return ""
}

func (x *CreateMemoryRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type CreateMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateMemoryResponse) Reset() {
	*x = CreateMemoryResponse{}
	mi := &file_memory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemoryResponse) ProtoMessage() {}

func (x *CreateMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemoryResponse.ProtoReflect.Descriptor instead.
func (*CreateMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{1}
}

type GetMemoryByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MemoryId uint64 `protobuf:"varint,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
}

func (x *GetMemoryByIdRequest) Reset() {
	*x = GetMemoryByIdRequest{}
	mi := &file_memory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemoryByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoryByIdRequest) ProtoMessage() {}

func (x *GetMemoryByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoryByIdRequest.ProtoReflect.Descriptor instead.
func (*GetMemoryByIdRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{2}
}

func (x *GetMemoryByIdRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetMemoryByIdRequest) GetMemoryId() uint64 {
	if x != nil {
		return x.MemoryId
	}
	return 0
}

type GetMemoryByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memory *DBMemory `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *GetMemoryByIdResponse) Reset() {
	*x = GetMemoryByIdResponse{}
	mi := &file_memory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemoryByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoryByIdResponse) ProtoMessage() {}

func (x *GetMemoryByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoryByIdResponse.ProtoReflect.Descriptor instead.
func (*GetMemoryByIdResponse) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{3}
}

func (x *GetMemoryByIdResponse) GetMemory() *DBMemory {
	if x != nil {
		return x.Memory
	}
	return nil
}

type GetMemoriesByDocIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           uint64                    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DocId            uint64                    `protobuf:"varint,2,opt,name=doc_id,json=docId,proto3" json:"doc_id,omitempty"`
	PageNumber       uint32                    `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize         uint32                    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	OrderByField     common.ORDER_BY_FIELD     `protobuf:"varint,5,opt,name=order_by_field,json=orderByField,proto3,enum=common.ORDER_BY_FIELD" json:"order_by_field,omitempty"`
	OrderByDirection common.ORDER_BY_DIRECTION `protobuf:"varint,6,opt,name=order_by_direction,json=orderByDirection,proto3,enum=common.ORDER_BY_DIRECTION" json:"order_by_direction,omitempty"`
}

func (x *GetMemoriesByDocIdRequest) Reset() {
	*x = GetMemoriesByDocIdRequest{}
	mi := &file_memory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemoriesByDocIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoriesByDocIdRequest) ProtoMessage() {}

func (x *GetMemoriesByDocIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoriesByDocIdRequest.ProtoReflect.Descriptor instead.
func (*GetMemoriesByDocIdRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{4}
}

func (x *GetMemoriesByDocIdRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetMemoriesByDocIdRequest) GetDocId() uint64 {
	if x != nil {
		return x.DocId
	}
	return 0
}

func (x *GetMemoriesByDocIdRequest) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetMemoriesByDocIdRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetMemoriesByDocIdRequest) GetOrderByField() common.ORDER_BY_FIELD {
	if x != nil {
		return x.OrderByField
	}
	return common.ORDER_BY_FIELD(0)
}

func (x *GetMemoriesByDocIdRequest) GetOrderByDirection() common.ORDER_BY_DIRECTION {
	if x != nil {
		return x.OrderByDirection
	}
	return common.ORDER_BY_DIRECTION(0)
}

type GetMemoriesByDocIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memories []*DBMemory `protobuf:"bytes,1,rep,name=memories,proto3" json:"memories,omitempty"`
}

func (x *GetMemoriesByDocIdResponse) Reset() {
	*x = GetMemoriesByDocIdResponse{}
	mi := &file_memory_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemoriesByDocIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoriesByDocIdResponse) ProtoMessage() {}

func (x *GetMemoriesByDocIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoriesByDocIdResponse.ProtoReflect.Descriptor instead.
func (*GetMemoriesByDocIdResponse) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{5}
}

func (x *GetMemoriesByDocIdResponse) GetMemories() []*DBMemory {
	if x != nil {
		return x.Memories
	}
	return nil
}

type GetMemoriesByMemoryTitleSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           uint64                    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SearchQuery      string                    `protobuf:"bytes,2,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty"`
	PageNumber       uint32                    `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize         uint32                    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	OrderByField     common.ORDER_BY_FIELD     `protobuf:"varint,5,opt,name=order_by_field,json=orderByField,proto3,enum=common.ORDER_BY_FIELD" json:"order_by_field,omitempty"`
	OrderByDirection common.ORDER_BY_DIRECTION `protobuf:"varint,6,opt,name=order_by_direction,json=orderByDirection,proto3,enum=common.ORDER_BY_DIRECTION" json:"order_by_direction,omitempty"`
}

func (x *GetMemoriesByMemoryTitleSearchRequest) Reset() {
	*x = GetMemoriesByMemoryTitleSearchRequest{}
	mi := &file_memory_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemoriesByMemoryTitleSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoriesByMemoryTitleSearchRequest) ProtoMessage() {}

func (x *GetMemoriesByMemoryTitleSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoriesByMemoryTitleSearchRequest.ProtoReflect.Descriptor instead.
func (*GetMemoriesByMemoryTitleSearchRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{6}
}

func (x *GetMemoriesByMemoryTitleSearchRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetMemoriesByMemoryTitleSearchRequest) GetSearchQuery() string {
	if x != nil {
		return x.SearchQuery
	}
	return ""
}

func (x *GetMemoriesByMemoryTitleSearchRequest) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetMemoriesByMemoryTitleSearchRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetMemoriesByMemoryTitleSearchRequest) GetOrderByField() common.ORDER_BY_FIELD {
	if x != nil {
		return x.OrderByField
	}
	return common.ORDER_BY_FIELD(0)
}

func (x *GetMemoriesByMemoryTitleSearchRequest) GetOrderByDirection() common.ORDER_BY_DIRECTION {
	if x != nil {
		return x.OrderByDirection
	}
	return common.ORDER_BY_DIRECTION(0)
}

type GetMemoriesByMemoryTitleSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memories []*DBMemory `protobuf:"bytes,1,rep,name=memories,proto3" json:"memories,omitempty"`
}

func (x *GetMemoriesByMemoryTitleSearchResponse) Reset() {
	*x = GetMemoriesByMemoryTitleSearchResponse{}
	mi := &file_memory_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemoriesByMemoryTitleSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoriesByMemoryTitleSearchResponse) ProtoMessage() {}

func (x *GetMemoriesByMemoryTitleSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoriesByMemoryTitleSearchResponse.ProtoReflect.Descriptor instead.
func (*GetMemoriesByMemoryTitleSearchResponse) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{7}
}

func (x *GetMemoriesByMemoryTitleSearchResponse) GetMemories() []*DBMemory {
	if x != nil {
		return x.Memories
	}
	return nil
}

type UpdateMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MemoryId      uint64 `protobuf:"varint,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	MemoryTitle   string `protobuf:"bytes,3,opt,name=memory_title,json=memoryTitle,proto3" json:"memory_title,omitempty"`
	MemoryContent string `protobuf:"bytes,4,opt,name=memory_content,json=memoryContent,proto3" json:"memory_content,omitempty"`
}

func (x *UpdateMemoryRequest) Reset() {
	*x = UpdateMemoryRequest{}
	mi := &file_memory_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemoryRequest) ProtoMessage() {}

func (x *UpdateMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateMemoryRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateMemoryRequest) GetMemoryId() uint64 {
	if x != nil {
		return x.MemoryId
	}
	return 0
}

func (x *UpdateMemoryRequest) GetMemoryTitle() string {
	if x != nil {
		return x.MemoryTitle
	}
	return ""
}

func (x *UpdateMemoryRequest) GetMemoryContent() string {
	if x != nil {
		return x.MemoryContent
	}
	return ""
}

type UpdateMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMemoryResponse) Reset() {
	*x = UpdateMemoryResponse{}
	mi := &file_memory_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemoryResponse) ProtoMessage() {}

func (x *UpdateMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{9}
}

type DeleteMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MemoryId uint64 `protobuf:"varint,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
}

func (x *DeleteMemoryRequest) Reset() {
	*x = DeleteMemoryRequest{}
	mi := &file_memory_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoryRequest) ProtoMessage() {}

func (x *DeleteMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteMemoryRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteMemoryRequest) GetMemoryId() uint64 {
	if x != nil {
		return x.MemoryId
	}
	return 0
}

type DeleteMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMemoryResponse) Reset() {
	*x = DeleteMemoryResponse{}
	mi := &file_memory_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoryResponse) ProtoMessage() {}

func (x *DeleteMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoryResponse.ProtoReflect.Descriptor instead.
func (*DeleteMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{11}
}

type DBMemory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId      uint64 `protobuf:"varint,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	DocId         uint64 `protobuf:"varint,2,opt,name=doc_id,json=docId,proto3" json:"doc_id,omitempty"`
	MemoryTitle   string `protobuf:"bytes,3,opt,name=memory_title,json=memoryTitle,proto3" json:"memory_title,omitempty"`
	MemoryContent string `protobuf:"bytes,4,opt,name=memory_content,json=memoryContent,proto3" json:"memory_content,omitempty"`
	CreatedTime   uint64 `protobuf:"varint,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime   uint64 `protobuf:"varint,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
}

func (x *DBMemory) Reset() {
	*x = DBMemory{}
	mi := &file_memory_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DBMemory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBMemory) ProtoMessage() {}

func (x *DBMemory) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBMemory.ProtoReflect.Descriptor instead.
func (*DBMemory) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{12}
}

func (x *DBMemory) GetMemoryId() uint64 {
	if x != nil {
		return x.MemoryId
	}
	return 0
}

func (x *DBMemory) GetDocId() uint64 {
	if x != nil {
		return x.DocId
	}
	return 0
}

func (x *DBMemory) GetMemoryTitle() string {
	if x != nil {
		return x.MemoryTitle
	}
	return ""
}

func (x *DBMemory) GetMemoryContent() string {
	if x != nil {
		return x.MemoryContent
	}
	return ""
}

func (x *DBMemory) GetCreatedTime() uint64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *DBMemory) GetUpdatedTime() uint64 {
	if x != nil {
		return x.UpdatedTime
	}
	return 0
}

var File_memory_proto protoreflect.FileDescriptor

var file_memory_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x42, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x91, 0x02, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x44, 0x6f,
	0x63, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x42, 0x59, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x48, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x52, 0x10,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x4a, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42,
	0x79, 0x44, 0x6f, 0x63, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x42, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0xa9, 0x02, 0x0a,
	0x25, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x3c, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x48,
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x52, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x42,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x95, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x16, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x44, 0x42, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x9c, 0x04, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x42, 0x79, 0x44, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x44, 0x6f,
	0x63, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42,
	0x79, 0x44, 0x6f, 0x63, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x2d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x62, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memory_proto_rawDescOnce sync.Once
	file_memory_proto_rawDescData = file_memory_proto_rawDesc
)

func file_memory_proto_rawDescGZIP() []byte {
	file_memory_proto_rawDescOnce.Do(func() {
		file_memory_proto_rawDescData = protoimpl.X.CompressGZIP(file_memory_proto_rawDescData)
	})
	return file_memory_proto_rawDescData
}

var file_memory_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_memory_proto_goTypes = []any{
	(*CreateMemoryRequest)(nil),                    // 0: memory.CreateMemoryRequest
	(*CreateMemoryResponse)(nil),                   // 1: memory.CreateMemoryResponse
	(*GetMemoryByIdRequest)(nil),                   // 2: memory.GetMemoryByIdRequest
	(*GetMemoryByIdResponse)(nil),                  // 3: memory.GetMemoryByIdResponse
	(*GetMemoriesByDocIdRequest)(nil),              // 4: memory.GetMemoriesByDocIdRequest
	(*GetMemoriesByDocIdResponse)(nil),             // 5: memory.GetMemoriesByDocIdResponse
	(*GetMemoriesByMemoryTitleSearchRequest)(nil),  // 6: memory.GetMemoriesByMemoryTitleSearchRequest
	(*GetMemoriesByMemoryTitleSearchResponse)(nil), // 7: memory.GetMemoriesByMemoryTitleSearchResponse
	(*UpdateMemoryRequest)(nil),                    // 8: memory.UpdateMemoryRequest
	(*UpdateMemoryResponse)(nil),                   // 9: memory.UpdateMemoryResponse
	(*DeleteMemoryRequest)(nil),                    // 10: memory.DeleteMemoryRequest
	(*DeleteMemoryResponse)(nil),                   // 11: memory.DeleteMemoryResponse
	(*DBMemory)(nil),                               // 12: memory.DBMemory
	(common.ORDER_BY_FIELD)(0),                     // 13: common.ORDER_BY_FIELD
	(common.ORDER_BY_DIRECTION)(0),                 // 14: common.ORDER_BY_DIRECTION
}
var file_memory_proto_depIdxs = []int32{
	12, // 0: memory.GetMemoryByIdResponse.memory:type_name -> memory.DBMemory
	13, // 1: memory.GetMemoriesByDocIdRequest.order_by_field:type_name -> common.ORDER_BY_FIELD
	14, // 2: memory.GetMemoriesByDocIdRequest.order_by_direction:type_name -> common.ORDER_BY_DIRECTION
	12, // 3: memory.GetMemoriesByDocIdResponse.memories:type_name -> memory.DBMemory
	13, // 4: memory.GetMemoriesByMemoryTitleSearchRequest.order_by_field:type_name -> common.ORDER_BY_FIELD
	14, // 5: memory.GetMemoriesByMemoryTitleSearchRequest.order_by_direction:type_name -> common.ORDER_BY_DIRECTION
	12, // 6: memory.GetMemoriesByMemoryTitleSearchResponse.memories:type_name -> memory.DBMemory
	0,  // 7: memory.MemoryService.CreateMemory:input_type -> memory.CreateMemoryRequest
	2,  // 8: memory.MemoryService.GetMemoryById:input_type -> memory.GetMemoryByIdRequest
	4,  // 9: memory.MemoryService.GetMemoriesByDocId:input_type -> memory.GetMemoriesByDocIdRequest
	6,  // 10: memory.MemoryService.GetMemoriesByMemoryTitleSearch:input_type -> memory.GetMemoriesByMemoryTitleSearchRequest
	8,  // 11: memory.MemoryService.UpdateMemory:input_type -> memory.UpdateMemoryRequest
	10, // 12: memory.MemoryService.DeleteMemory:input_type -> memory.DeleteMemoryRequest
	1,  // 13: memory.MemoryService.CreateMemory:output_type -> memory.CreateMemoryResponse
	3,  // 14: memory.MemoryService.GetMemoryById:output_type -> memory.GetMemoryByIdResponse
	5,  // 15: memory.MemoryService.GetMemoriesByDocId:output_type -> memory.GetMemoriesByDocIdResponse
	7,  // 16: memory.MemoryService.GetMemoriesByMemoryTitleSearch:output_type -> memory.GetMemoriesByMemoryTitleSearchResponse
	9,  // 17: memory.MemoryService.UpdateMemory:output_type -> memory.UpdateMemoryResponse
	11, // 18: memory.MemoryService.DeleteMemory:output_type -> memory.DeleteMemoryResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_memory_proto_init() }
func file_memory_proto_init() {
	if File_memory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_memory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memory_proto_goTypes,
		DependencyIndexes: file_memory_proto_depIdxs,
		MessageInfos:      file_memory_proto_msgTypes,
	}.Build()
	File_memory_proto = out.File
	file_memory_proto_rawDesc = nil
	file_memory_proto_goTypes = nil
	file_memory_proto_depIdxs = nil
}
