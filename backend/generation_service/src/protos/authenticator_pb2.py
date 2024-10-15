# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: protos/authenticator.proto
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
    'protos/authenticator.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1aprotos/authenticator.proto\x12\rauthenticator\"7\n\x11\x43reateUserRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x14\n\x12\x43reateUserResponse\"0\n\x1a\x41uthenticateSessionRequest\x12\x12\n\nsession_id\x18\x01 \x01(\t\"O\n\x1b\x41uthenticateSessionResponse\x12\x30\n\x0cuser_session\x18\x01 \x01(\x0b\x32\x1a.authenticator.UserSession\":\n\x14\x43reateSessionRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"+\n\x15\x43reateSessionResponse\x12\x12\n\nsession_id\x18\x01 \x01(\t\"*\n\x14\x44\x65leteSessionRequest\x12\x12\n\nsession_id\x18\x01 \x01(\t\"\x17\n\x15\x44\x65leteSessionResponse\"G\n\x06\x44\x42User\x12\x0f\n\x07user_id\x18\x01 \x01(\x04\x12\x10\n\x08username\x18\x02 \x01(\t\x12\x1a\n\x12hash_salt_password\x18\x03 \x01(\t\"0\n\x0bUserSession\x12\x0f\n\x07user_id\x18\x01 \x01(\x04\x12\x10\n\x08username\x18\x02 \x01(\t2\x86\x03\n\x0bUserService\x12Q\n\nCreateUser\x12 .authenticator.CreateUserRequest\x1a!.authenticator.CreateUserResponse\x12l\n\x13\x41uthenticateSession\x12).authenticator.AuthenticateSessionRequest\x1a*.authenticator.AuthenticateSessionResponse\x12Z\n\rCreateSession\x12#.authenticator.CreateSessionRequest\x1a$.authenticator.CreateSessionResponse\x12Z\n\rDeleteSession\x12#.authenticator.DeleteSessionRequest\x1a$.authenticator.DeleteSessionResponseB(Z&%replace%/internal/proto/authenticatorb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.authenticator_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z&%replace%/internal/proto/authenticator'
  _globals['_CREATEUSERREQUEST']._serialized_start=45
  _globals['_CREATEUSERREQUEST']._serialized_end=100
  _globals['_CREATEUSERRESPONSE']._serialized_start=102
  _globals['_CREATEUSERRESPONSE']._serialized_end=122
  _globals['_AUTHENTICATESESSIONREQUEST']._serialized_start=124
  _globals['_AUTHENTICATESESSIONREQUEST']._serialized_end=172
  _globals['_AUTHENTICATESESSIONRESPONSE']._serialized_start=174
  _globals['_AUTHENTICATESESSIONRESPONSE']._serialized_end=253
  _globals['_CREATESESSIONREQUEST']._serialized_start=255
  _globals['_CREATESESSIONREQUEST']._serialized_end=313
  _globals['_CREATESESSIONRESPONSE']._serialized_start=315
  _globals['_CREATESESSIONRESPONSE']._serialized_end=358
  _globals['_DELETESESSIONREQUEST']._serialized_start=360
  _globals['_DELETESESSIONREQUEST']._serialized_end=402
  _globals['_DELETESESSIONRESPONSE']._serialized_start=404
  _globals['_DELETESESSIONRESPONSE']._serialized_end=427
  _globals['_DBUSER']._serialized_start=429
  _globals['_DBUSER']._serialized_end=500
  _globals['_USERSESSION']._serialized_start=502
  _globals['_USERSESSION']._serialized_end=550
  _globals['_USERSERVICE']._serialized_start=553
  _globals['_USERSERVICE']._serialized_end=943
# @@protoc_insertion_point(module_scope)
