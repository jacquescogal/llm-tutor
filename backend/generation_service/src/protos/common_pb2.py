# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: protos/common.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'protos/common.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13protos/common.proto\x12\x06\x63ommon*\xa8\x01\n\x0fUserSubjectRole\x12\x1f\n\x1bUSER_SUBJECT_ROLE_UNDEFINED\x10\x00\x12\x1b\n\x17USER_SUBJECT_ROLE_OWNER\x10\x01\x12\x1b\n\x17USER_SUBJECT_ROLE_ADMIN\x10\x02\x12\x1c\n\x18USER_SUBJECT_ROLE_EDITOR\x10\x03\x12\x1c\n\x18USER_SUBJECT_ROLE_VIEWER\x10\x04*\xa2\x01\n\x0eUserModuleRole\x12\x1e\n\x1aUSER_MODULE_ROLE_UNDEFINED\x10\x00\x12\x1a\n\x16USER_MODULE_ROLE_OWNER\x10\x01\x12\x1a\n\x16USER_MODULE_ROLE_ADMIN\x10\x02\x12\x1b\n\x17USER_MODULE_ROLE_EDITOR\x10\x03\x12\x1b\n\x17USER_MODULE_ROLE_VIEWER\x10\x04*\x9c\x01\n\x0cUploadStatus\x12\x1d\n\x19UPLOAD_STATUS_NOT_STARTED\x10\x00\x12\x1a\n\x16UPLOAD_STATUS_QUEUEING\x10\x01\x12\x1c\n\x18UPLOAD_STATUS_PROCESSING\x10\x02\x12\x19\n\x15UPLOAD_STATUS_SUCCESS\x10\x03\x12\x18\n\x14UPLOAD_STATUS_FAILED\x10\x04*\x85\x01\n\x0cQuestionType\x12\x1b\n\x17QUESTION_TYPE_UNDEFINED\x10\x00\x12\x15\n\x11QUESTION_TYPE_MCQ\x10\x01\x12\"\n\x1eQUESTION_TYPE_MULTI_ANSWER_MCQ\x10\x02\x12\x1d\n\x19QUESTION_TYPE_OPEN_ANSWER\x10\x03*o\n\x12ORDER_BY_DIRECTION\x12 \n\x1cORDER_BY_DIRECTION_UNDEFINED\x10\x00\x12\x1a\n\x16ORDER_BY_DIRECTION_ASC\x10\x01\x12\x1b\n\x17ORDER_BY_DIRECTION_DESC\x10\x02*\xa1\x01\n\x0eORDER_BY_FIELD\x12\x1c\n\x18ORDER_BY_FIELD_UNDEFINED\x10\x00\x12\x15\n\x11ORDER_BY_FIELD_ID\x10\x01\x12\x18\n\x14ORDER_BY_FIELD_TITLE\x10\x02\x12\x1f\n\x1bORDER_BY_FIELD_CREATED_TIME\x10\x03\x12\x1f\n\x1bORDER_BY_FIELD_UPDATED_TIME\x10\x04*L\n\x06IDType\x12\x12\n\x0eID_UNSPECIFIED\x10\x00\x12\r\n\tID_MODULE\x10\x01\x12\x0e\n\nID_SUBJECT\x10\x02\x12\x0f\n\x0bID_DOCUMENT\x10\x03\x42!Z\x1f%replace%/internal/proto/commonb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.common_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\037%replace%/internal/proto/common'
  _globals['_USERSUBJECTROLE']._serialized_start=32
  _globals['_USERSUBJECTROLE']._serialized_end=200
  _globals['_USERMODULEROLE']._serialized_start=203
  _globals['_USERMODULEROLE']._serialized_end=365
  _globals['_UPLOADSTATUS']._serialized_start=368
  _globals['_UPLOADSTATUS']._serialized_end=524
  _globals['_QUESTIONTYPE']._serialized_start=527
  _globals['_QUESTIONTYPE']._serialized_end=660
  _globals['_ORDER_BY_DIRECTION']._serialized_start=662
  _globals['_ORDER_BY_DIRECTION']._serialized_end=773
  _globals['_ORDER_BY_FIELD']._serialized_start=776
  _globals['_ORDER_BY_FIELD']._serialized_end=937
  _globals['_IDTYPE']._serialized_start=939
  _globals['_IDTYPE']._serialized_end=1015
# @@protoc_insertion_point(module_scope)
