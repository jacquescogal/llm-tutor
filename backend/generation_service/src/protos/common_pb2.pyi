from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class UserSubjectRole(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    USER_SUBJECT_ROLE_UNDEFINED: _ClassVar[UserSubjectRole]
    USER_SUBJECT_ROLE_OWNER: _ClassVar[UserSubjectRole]
    USER_SUBJECT_ROLE_ADMIN: _ClassVar[UserSubjectRole]
    USER_SUBJECT_ROLE_EDITOR: _ClassVar[UserSubjectRole]
    USER_SUBJECT_ROLE_VIEWER: _ClassVar[UserSubjectRole]

class UserModuleRole(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    USER_MODULE_ROLE_UNDEFINED: _ClassVar[UserModuleRole]
    USER_MODULE_ROLE_OWNER: _ClassVar[UserModuleRole]
    USER_MODULE_ROLE_ADMIN: _ClassVar[UserModuleRole]
    USER_MODULE_ROLE_EDITOR: _ClassVar[UserModuleRole]
    USER_MODULE_ROLE_VIEWER: _ClassVar[UserModuleRole]

class UploadStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UPLOAD_STATUS_NOT_STARTED: _ClassVar[UploadStatus]
    UPLOAD_STATUS_QUEUEING: _ClassVar[UploadStatus]
    UPLOAD_STATUS_PROCESSING: _ClassVar[UploadStatus]
    UPLOAD_STATUS_SUCCESS: _ClassVar[UploadStatus]
    UPLOAD_STATUS_FAILED: _ClassVar[UploadStatus]

class QuestionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    QUESTION_TYPE_UNDEFINED: _ClassVar[QuestionType]
    QUESTION_TYPE_MCQ: _ClassVar[QuestionType]
    QUESTION_TYPE_MULTI_ANSWER_MCQ: _ClassVar[QuestionType]
    QUESTION_TYPE_OPEN_ANSWER: _ClassVar[QuestionType]

class ORDER_BY_DIRECTION(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ORDER_BY_DIRECTION_UNDEFINED: _ClassVar[ORDER_BY_DIRECTION]
    ORDER_BY_DIRECTION_ASC: _ClassVar[ORDER_BY_DIRECTION]
    ORDER_BY_DIRECTION_DESC: _ClassVar[ORDER_BY_DIRECTION]

class ORDER_BY_FIELD(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ORDER_BY_FIELD_UNDEFINED: _ClassVar[ORDER_BY_FIELD]
    ORDER_BY_FIELD_ID: _ClassVar[ORDER_BY_FIELD]
    ORDER_BY_FIELD_TITLE: _ClassVar[ORDER_BY_FIELD]
    ORDER_BY_FIELD_CREATED_TIME: _ClassVar[ORDER_BY_FIELD]
    ORDER_BY_FIELD_UPDATED_TIME: _ClassVar[ORDER_BY_FIELD]

class IDType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ID_UNSPECIFIED: _ClassVar[IDType]
    ID_MODULE: _ClassVar[IDType]
    ID_SUBJECT: _ClassVar[IDType]
    ID_DOCUMENT: _ClassVar[IDType]
USER_SUBJECT_ROLE_UNDEFINED: UserSubjectRole
USER_SUBJECT_ROLE_OWNER: UserSubjectRole
USER_SUBJECT_ROLE_ADMIN: UserSubjectRole
USER_SUBJECT_ROLE_EDITOR: UserSubjectRole
USER_SUBJECT_ROLE_VIEWER: UserSubjectRole
USER_MODULE_ROLE_UNDEFINED: UserModuleRole
USER_MODULE_ROLE_OWNER: UserModuleRole
USER_MODULE_ROLE_ADMIN: UserModuleRole
USER_MODULE_ROLE_EDITOR: UserModuleRole
USER_MODULE_ROLE_VIEWER: UserModuleRole
UPLOAD_STATUS_NOT_STARTED: UploadStatus
UPLOAD_STATUS_QUEUEING: UploadStatus
UPLOAD_STATUS_PROCESSING: UploadStatus
UPLOAD_STATUS_SUCCESS: UploadStatus
UPLOAD_STATUS_FAILED: UploadStatus
QUESTION_TYPE_UNDEFINED: QuestionType
QUESTION_TYPE_MCQ: QuestionType
QUESTION_TYPE_MULTI_ANSWER_MCQ: QuestionType
QUESTION_TYPE_OPEN_ANSWER: QuestionType
ORDER_BY_DIRECTION_UNDEFINED: ORDER_BY_DIRECTION
ORDER_BY_DIRECTION_ASC: ORDER_BY_DIRECTION
ORDER_BY_DIRECTION_DESC: ORDER_BY_DIRECTION
ORDER_BY_FIELD_UNDEFINED: ORDER_BY_FIELD
ORDER_BY_FIELD_ID: ORDER_BY_FIELD
ORDER_BY_FIELD_TITLE: ORDER_BY_FIELD
ORDER_BY_FIELD_CREATED_TIME: ORDER_BY_FIELD
ORDER_BY_FIELD_UPDATED_TIME: ORDER_BY_FIELD
ID_UNSPECIFIED: IDType
ID_MODULE: IDType
ID_SUBJECT: IDType
ID_DOCUMENT: IDType
