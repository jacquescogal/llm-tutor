// document_job.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: document_job.proto

package document_job

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

// use this job to get from the document service the full document and then process it
type DocumentProcessingJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ModuleId uint64 `protobuf:"varint,2,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	DocId    uint64 `protobuf:"varint,3,opt,name=doc_id,json=docId,proto3" json:"doc_id,omitempty"`
}

func (x *DocumentProcessingJob) Reset() {
	*x = DocumentProcessingJob{}
	mi := &file_document_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DocumentProcessingJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentProcessingJob) ProtoMessage() {}

func (x *DocumentProcessingJob) ProtoReflect() protoreflect.Message {
	mi := &file_document_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentProcessingJob.ProtoReflect.Descriptor instead.
func (*DocumentProcessingJob) Descriptor() ([]byte, []int) {
	return file_document_job_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentProcessingJob) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DocumentProcessingJob) GetModuleId() uint64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *DocumentProcessingJob) GetDocId() uint64 {
	if x != nil {
		return x.DocId
	}
	return 0
}

var File_document_job_proto protoreflect.FileDescriptor

var file_document_job_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6a,
	0x6f, 0x62, 0x22, 0x64, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x42, 0x21, 0x5a, 0x1f, 0x62, 0x66, 0x66, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_document_job_proto_rawDescOnce sync.Once
	file_document_job_proto_rawDescData = file_document_job_proto_rawDesc
)

func file_document_job_proto_rawDescGZIP() []byte {
	file_document_job_proto_rawDescOnce.Do(func() {
		file_document_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_document_job_proto_rawDescData)
	})
	return file_document_job_proto_rawDescData
}

var file_document_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_document_job_proto_goTypes = []any{
	(*DocumentProcessingJob)(nil), // 0: document_job.DocumentProcessingJob
}
var file_document_job_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_document_job_proto_init() }
func file_document_job_proto_init() {
	if File_document_job_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_document_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_document_job_proto_goTypes,
		DependencyIndexes: file_document_job_proto_depIdxs,
		MessageInfos:      file_document_job_proto_msgTypes,
	}.Build()
	File_document_job_proto = out.File
	file_document_job_proto_rawDesc = nil
	file_document_job_proto_goTypes = nil
	file_document_job_proto_depIdxs = nil
}
