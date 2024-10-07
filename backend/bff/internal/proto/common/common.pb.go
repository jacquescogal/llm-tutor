// common.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: common.proto

package common

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

type UserSubjectRole int32

const (
	UserSubjectRole_USER_SUBJECT_ROLE_UNDEFINED UserSubjectRole = 0
	UserSubjectRole_USER_SUBJECT_ROLE_OWNER     UserSubjectRole = 1 // can give access to other users + editor rights
	UserSubjectRole_USER_SUBJECT_ROLE_ADMIN     UserSubjectRole = 2 // can give access to other users + editor rights
	UserSubjectRole_USER_SUBJECT_ROLE_EDITOR    UserSubjectRole = 3 // CRUD + viewer rights
	UserSubjectRole_USER_SUBJECT_ROLE_VIEWER    UserSubjectRole = 4 // allows reads when subject is not public
)

// Enum value maps for UserSubjectRole.
var (
	UserSubjectRole_name = map[int32]string{
		0: "USER_SUBJECT_ROLE_UNDEFINED",
		1: "USER_SUBJECT_ROLE_OWNER",
		2: "USER_SUBJECT_ROLE_ADMIN",
		3: "USER_SUBJECT_ROLE_EDITOR",
		4: "USER_SUBJECT_ROLE_VIEWER",
	}
	UserSubjectRole_value = map[string]int32{
		"USER_SUBJECT_ROLE_UNDEFINED": 0,
		"USER_SUBJECT_ROLE_OWNER":     1,
		"USER_SUBJECT_ROLE_ADMIN":     2,
		"USER_SUBJECT_ROLE_EDITOR":    3,
		"USER_SUBJECT_ROLE_VIEWER":    4,
	}
)

func (x UserSubjectRole) Enum() *UserSubjectRole {
	p := new(UserSubjectRole)
	*p = x
	return p
}

func (x UserSubjectRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSubjectRole) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (UserSubjectRole) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x UserSubjectRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSubjectRole.Descriptor instead.
func (UserSubjectRole) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type UserModuleRole int32

const (
	UserModuleRole_USER_MODULE_ROLE_UNDEFINED UserModuleRole = 0
	UserModuleRole_USER_MODULE_ROLE_OWNER     UserModuleRole = 1 // can give access to other users + editor rights
	UserModuleRole_USER_MODULE_ROLE_ADMIN     UserModuleRole = 2 // can give access to other users + editor rights
	UserModuleRole_USER_MODULE_ROLE_EDITOR    UserModuleRole = 3 // CRUD + viewer rights
	UserModuleRole_USER_MODULE_ROLE_VIEWER    UserModuleRole = 4 // allows reads when doc is not public
)

// Enum value maps for UserModuleRole.
var (
	UserModuleRole_name = map[int32]string{
		0: "USER_MODULE_ROLE_UNDEFINED",
		1: "USER_MODULE_ROLE_OWNER",
		2: "USER_MODULE_ROLE_ADMIN",
		3: "USER_MODULE_ROLE_EDITOR",
		4: "USER_MODULE_ROLE_VIEWER",
	}
	UserModuleRole_value = map[string]int32{
		"USER_MODULE_ROLE_UNDEFINED": 0,
		"USER_MODULE_ROLE_OWNER":     1,
		"USER_MODULE_ROLE_ADMIN":     2,
		"USER_MODULE_ROLE_EDITOR":    3,
		"USER_MODULE_ROLE_VIEWER":    4,
	}
)

func (x UserModuleRole) Enum() *UserModuleRole {
	p := new(UserModuleRole)
	*p = x
	return p
}

func (x UserModuleRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserModuleRole) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (UserModuleRole) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x UserModuleRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserModuleRole.Descriptor instead.
func (UserModuleRole) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type UploadStatus int32

const (
	UploadStatus_UPLOAD_STATUS_NOT_STARTED      UploadStatus = 0
	UploadStatus_UPLOAD_STATUS_SUCCESS          UploadStatus = 1
	UploadStatus_UPLOAD_STATUS_FAILED           UploadStatus = 2
	UploadStatus_UPLOAD_STATUS_UPLOADING        UploadStatus = 3
	UploadStatus_UPLOAD_STATUS_PENDING_APPROVAL UploadStatus = 4
	UploadStatus_UPLOAD_STATUS_INSERTING        UploadStatus = 5
)

// Enum value maps for UploadStatus.
var (
	UploadStatus_name = map[int32]string{
		0: "UPLOAD_STATUS_NOT_STARTED",
		1: "UPLOAD_STATUS_SUCCESS",
		2: "UPLOAD_STATUS_FAILED",
		3: "UPLOAD_STATUS_UPLOADING",
		4: "UPLOAD_STATUS_PENDING_APPROVAL",
		5: "UPLOAD_STATUS_INSERTING",
	}
	UploadStatus_value = map[string]int32{
		"UPLOAD_STATUS_NOT_STARTED":      0,
		"UPLOAD_STATUS_SUCCESS":          1,
		"UPLOAD_STATUS_FAILED":           2,
		"UPLOAD_STATUS_UPLOADING":        3,
		"UPLOAD_STATUS_PENDING_APPROVAL": 4,
		"UPLOAD_STATUS_INSERTING":        5,
	}
)

func (x UploadStatus) Enum() *UploadStatus {
	p := new(UploadStatus)
	*p = x
	return p
}

func (x UploadStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (UploadStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x UploadStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatus.Descriptor instead.
func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

type QuestionType int32

const (
	QuestionType_QUESTION_TYPE_UNDEFINED        QuestionType = 0
	QuestionType_QUESTION_TYPE_MCQ              QuestionType = 1
	QuestionType_QUESTION_TYPE_MULTI_ANSWER_MCQ QuestionType = 2
	QuestionType_QUESTION_TYPE_SHORT_ANSWER     QuestionType = 3
	QuestionType_QUESTION_TYPE_LONG_ANSWER      QuestionType = 4
)

// Enum value maps for QuestionType.
var (
	QuestionType_name = map[int32]string{
		0: "QUESTION_TYPE_UNDEFINED",
		1: "QUESTION_TYPE_MCQ",
		2: "QUESTION_TYPE_MULTI_ANSWER_MCQ",
		3: "QUESTION_TYPE_SHORT_ANSWER",
		4: "QUESTION_TYPE_LONG_ANSWER",
	}
	QuestionType_value = map[string]int32{
		"QUESTION_TYPE_UNDEFINED":        0,
		"QUESTION_TYPE_MCQ":              1,
		"QUESTION_TYPE_MULTI_ANSWER_MCQ": 2,
		"QUESTION_TYPE_SHORT_ANSWER":     3,
		"QUESTION_TYPE_LONG_ANSWER":      4,
	}
)

func (x QuestionType) Enum() *QuestionType {
	p := new(QuestionType)
	*p = x
	return p
}

func (x QuestionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestionType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[3].Descriptor()
}

func (QuestionType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[3]
}

func (x QuestionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestionType.Descriptor instead.
func (QuestionType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

type ORDER_BY_DIRECTION int32

const (
	ORDER_BY_DIRECTION_ORDER_BY_DIRECTION_UNDEFINED ORDER_BY_DIRECTION = 0
	ORDER_BY_DIRECTION_ORDER_BY_DIRECTION_ASC       ORDER_BY_DIRECTION = 1
	ORDER_BY_DIRECTION_ORDER_BY_DIRECTION_DESC      ORDER_BY_DIRECTION = 2
)

// Enum value maps for ORDER_BY_DIRECTION.
var (
	ORDER_BY_DIRECTION_name = map[int32]string{
		0: "ORDER_BY_DIRECTION_UNDEFINED",
		1: "ORDER_BY_DIRECTION_ASC",
		2: "ORDER_BY_DIRECTION_DESC",
	}
	ORDER_BY_DIRECTION_value = map[string]int32{
		"ORDER_BY_DIRECTION_UNDEFINED": 0,
		"ORDER_BY_DIRECTION_ASC":       1,
		"ORDER_BY_DIRECTION_DESC":      2,
	}
)

func (x ORDER_BY_DIRECTION) Enum() *ORDER_BY_DIRECTION {
	p := new(ORDER_BY_DIRECTION)
	*p = x
	return p
}

func (x ORDER_BY_DIRECTION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ORDER_BY_DIRECTION) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[4].Descriptor()
}

func (ORDER_BY_DIRECTION) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[4]
}

func (x ORDER_BY_DIRECTION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ORDER_BY_DIRECTION.Descriptor instead.
func (ORDER_BY_DIRECTION) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

type ORDER_BY_FIELD int32

const (
	ORDER_BY_FIELD_ORDER_BY_FIELD_UNDEFINED    ORDER_BY_FIELD = 0
	ORDER_BY_FIELD_ORDER_BY_FIELD_ID           ORDER_BY_FIELD = 1
	ORDER_BY_FIELD_ORDER_BY_FIELD_TITLE        ORDER_BY_FIELD = 2
	ORDER_BY_FIELD_ORDER_BY_FIELD_CREATED_TIME ORDER_BY_FIELD = 3
	ORDER_BY_FIELD_ORDER_BY_FIELD_UPDATED_TIME ORDER_BY_FIELD = 4
)

// Enum value maps for ORDER_BY_FIELD.
var (
	ORDER_BY_FIELD_name = map[int32]string{
		0: "ORDER_BY_FIELD_UNDEFINED",
		1: "ORDER_BY_FIELD_ID",
		2: "ORDER_BY_FIELD_TITLE",
		3: "ORDER_BY_FIELD_CREATED_TIME",
		4: "ORDER_BY_FIELD_UPDATED_TIME",
	}
	ORDER_BY_FIELD_value = map[string]int32{
		"ORDER_BY_FIELD_UNDEFINED":    0,
		"ORDER_BY_FIELD_ID":           1,
		"ORDER_BY_FIELD_TITLE":        2,
		"ORDER_BY_FIELD_CREATED_TIME": 3,
		"ORDER_BY_FIELD_UPDATED_TIME": 4,
	}
)

func (x ORDER_BY_FIELD) Enum() *ORDER_BY_FIELD {
	p := new(ORDER_BY_FIELD)
	*p = x
	return p
}

func (x ORDER_BY_FIELD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ORDER_BY_FIELD) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[5].Descriptor()
}

func (ORDER_BY_FIELD) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[5]
}

func (x ORDER_BY_FIELD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ORDER_BY_FIELD.Descriptor instead.
func (ORDER_BY_FIELD) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0xa8, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x55,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f,
	0x52, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10,
	0x04, 0x2a, 0xa2, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x56, 0x49,
	0x45, 0x57, 0x45, 0x52, 0x10, 0x04, 0x2a, 0xc0, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x50, 0x4c, 0x4f, 0x41,
	0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x55,
	0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x50, 0x4c,
	0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x50, 0x4c, 0x4f,
	0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17,
	0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x53, 0x45, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x2a, 0xa5, 0x01, 0x0a, 0x0c, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x43, 0x51, 0x10, 0x01, 0x12, 0x22,
	0x0a, 0x1e, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x5f, 0x4d, 0x43, 0x51,
	0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52,
	0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10,
	0x04, 0x2a, 0x6f, 0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42,
	0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x53, 0x43,
	0x10, 0x02, 0x2a, 0xa1, 0x01, 0x0a, 0x0e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42,
	0x59, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x49, 0x54,
	0x4c, 0x45, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42,
	0x59, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x10, 0x04, 0x42, 0x1b, 0x5a, 0x19, 0x62, 0x66, 0x66, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_common_proto_goTypes = []any{
	(UserSubjectRole)(0),    // 0: common.UserSubjectRole
	(UserModuleRole)(0),     // 1: common.UserModuleRole
	(UploadStatus)(0),       // 2: common.UploadStatus
	(QuestionType)(0),       // 3: common.QuestionType
	(ORDER_BY_DIRECTION)(0), // 4: common.ORDER_BY_DIRECTION
	(ORDER_BY_FIELD)(0),     // 5: common.ORDER_BY_FIELD
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
