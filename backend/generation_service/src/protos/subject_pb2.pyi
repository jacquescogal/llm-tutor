from src.protos import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateSubjectRequest(_message.Message):
    __slots__ = ("user_id", "subject_name", "subject_description", "is_public")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_NAME_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_PUBLIC_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_name: str
    subject_description: str
    is_public: bool
    def __init__(self, user_id: _Optional[int] = ..., subject_name: _Optional[str] = ..., subject_description: _Optional[str] = ..., is_public: bool = ...) -> None: ...

class CreateSubjectResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetPublicSubjectsRequest(_message.Message):
    __slots__ = ("user_id", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetPublicSubjectsResponse(_message.Message):
    __slots__ = ("subjects",)
    SUBJECTS_FIELD_NUMBER: _ClassVar[int]
    subjects: _containers.RepeatedCompositeFieldContainer[FullSubject]
    def __init__(self, subjects: _Optional[_Iterable[_Union[FullSubject, _Mapping]]] = ...) -> None: ...

class GetPrivateSubjectsByUserIdRequest(_message.Message):
    __slots__ = ("user_id", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetPrivateSubjectsByUserIdResponse(_message.Message):
    __slots__ = ("subjects",)
    SUBJECTS_FIELD_NUMBER: _ClassVar[int]
    subjects: _containers.RepeatedCompositeFieldContainer[FullSubject]
    def __init__(self, subjects: _Optional[_Iterable[_Union[FullSubject, _Mapping]]] = ...) -> None: ...

class GetFavouriteSubjectsByUserIdRequest(_message.Message):
    __slots__ = ("user_id", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetFavouriteSubjectsByUserIdResponse(_message.Message):
    __slots__ = ("subjects",)
    SUBJECTS_FIELD_NUMBER: _ClassVar[int]
    subjects: _containers.RepeatedCompositeFieldContainer[FullSubject]
    def __init__(self, subjects: _Optional[_Iterable[_Union[FullSubject, _Mapping]]] = ...) -> None: ...

class GetSubjectByIdRequest(_message.Message):
    __slots__ = ("user_id", "subject_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ...) -> None: ...

class GetSubjectByIdResponse(_message.Message):
    __slots__ = ("subject",)
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    subject: FullSubject
    def __init__(self, subject: _Optional[_Union[FullSubject, _Mapping]] = ...) -> None: ...

class GetSubjectsByUserIdRequest(_message.Message):
    __slots__ = ("user_id", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetSubjectsByUserIdResponse(_message.Message):
    __slots__ = ("subjects",)
    SUBJECTS_FIELD_NUMBER: _ClassVar[int]
    subjects: _containers.RepeatedCompositeFieldContainer[FullSubject]
    def __init__(self, subjects: _Optional[_Iterable[_Union[FullSubject, _Mapping]]] = ...) -> None: ...

class GetSubjectsByNameSearchRequest(_message.Message):
    __slots__ = ("user_id", "search_query", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SEARCH_QUERY_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    search_query: str
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., search_query: _Optional[str] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetSubjectsByNameSearchResponse(_message.Message):
    __slots__ = ("subjects",)
    SUBJECTS_FIELD_NUMBER: _ClassVar[int]
    subjects: _containers.RepeatedCompositeFieldContainer[FullSubject]
    def __init__(self, subjects: _Optional[_Iterable[_Union[FullSubject, _Mapping]]] = ...) -> None: ...

class UpdateSubjectRequest(_message.Message):
    __slots__ = ("user_id", "subject_id", "subject_name", "subject_description", "is_public")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_NAME_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_PUBLIC_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    subject_name: str
    subject_description: str
    is_public: bool
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ..., subject_name: _Optional[str] = ..., subject_description: _Optional[str] = ..., is_public: bool = ...) -> None: ...

class UpdateSubjectResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteSubjectRequest(_message.Message):
    __slots__ = ("user_id", "subject_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ...) -> None: ...

class DeleteSubjectResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SetUserSubjectFavouriteRequest(_message.Message):
    __slots__ = ("user_id", "subject_id", "is_favourite")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    IS_FAVOURITE_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    is_favourite: bool
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ..., is_favourite: bool = ...) -> None: ...

class SetUserSubjectFavouriteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SetUserSubjectRoleRequest(_message.Message):
    __slots__ = ("user_id", "subject_id", "user_subject_role")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    USER_SUBJECT_ROLE_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    user_subject_role: _common_pb2.UserSubjectRole
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ..., user_subject_role: _Optional[_Union[_common_pb2.UserSubjectRole, str]] = ...) -> None: ...

class SetUserSubjectRoleResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SetSubjectModuleMappingRequest(_message.Message):
    __slots__ = ("user_id", "subject_id", "module_ids")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_IDS_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    module_ids: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ..., module_ids: _Optional[_Iterable[int]] = ...) -> None: ...

class SetSubjectModuleMappingResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class FullSubject(_message.Message):
    __slots__ = ("subject", "user_subject_role", "is_favourite")
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    USER_SUBJECT_ROLE_FIELD_NUMBER: _ClassVar[int]
    IS_FAVOURITE_FIELD_NUMBER: _ClassVar[int]
    subject: DBSubject
    user_subject_role: _common_pb2.UserSubjectRole
    is_favourite: bool
    def __init__(self, subject: _Optional[_Union[DBSubject, _Mapping]] = ..., user_subject_role: _Optional[_Union[_common_pb2.UserSubjectRole, str]] = ..., is_favourite: bool = ...) -> None: ...

class DBUserSubjectMap(_message.Message):
    __slots__ = ("user_id", "subject_id", "user_subject_role", "is_favourite")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    USER_SUBJECT_ROLE_FIELD_NUMBER: _ClassVar[int]
    IS_FAVOURITE_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    user_subject_role: _common_pb2.UserSubjectRole
    is_favourite: bool
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ..., user_subject_role: _Optional[_Union[_common_pb2.UserSubjectRole, str]] = ..., is_favourite: bool = ...) -> None: ...

class DBSubject(_message.Message):
    __slots__ = ("subject_id", "subject_name", "subject_description", "is_public", "created_time", "updated_time")
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_NAME_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_PUBLIC_FIELD_NUMBER: _ClassVar[int]
    CREATED_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATED_TIME_FIELD_NUMBER: _ClassVar[int]
    subject_id: int
    subject_name: str
    subject_description: str
    is_public: bool
    created_time: int
    updated_time: int
    def __init__(self, subject_id: _Optional[int] = ..., subject_name: _Optional[str] = ..., subject_description: _Optional[str] = ..., is_public: bool = ..., created_time: _Optional[int] = ..., updated_time: _Optional[int] = ...) -> None: ...

class DBSubjectModuleMap(_message.Message):
    __slots__ = ("subject_id", "module_id")
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    subject_id: int
    module_id: int
    def __init__(self, subject_id: _Optional[int] = ..., module_id: _Optional[int] = ...) -> None: ...
