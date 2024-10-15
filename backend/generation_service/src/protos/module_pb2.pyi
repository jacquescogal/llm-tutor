from src.protos import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateModuleRequest(_message.Message):
    __slots__ = ("user_id", "module_name", "module_description", "is_public")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_NAME_FIELD_NUMBER: _ClassVar[int]
    MODULE_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_PUBLIC_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_name: str
    module_description: str
    is_public: bool
    def __init__(self, user_id: _Optional[int] = ..., module_name: _Optional[str] = ..., module_description: _Optional[str] = ..., is_public: bool = ...) -> None: ...

class CreateModuleResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetPublicModulesRequest(_message.Message):
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

class GetPublicModulesResponse(_message.Message):
    __slots__ = ("modules",)
    MODULES_FIELD_NUMBER: _ClassVar[int]
    modules: _containers.RepeatedCompositeFieldContainer[FullModule]
    def __init__(self, modules: _Optional[_Iterable[_Union[FullModule, _Mapping]]] = ...) -> None: ...

class GetPrivateModulesByUserIdRequest(_message.Message):
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

class GetPrivateModulesByUserIdResponse(_message.Message):
    __slots__ = ("modules",)
    MODULES_FIELD_NUMBER: _ClassVar[int]
    modules: _containers.RepeatedCompositeFieldContainer[FullModule]
    def __init__(self, modules: _Optional[_Iterable[_Union[FullModule, _Mapping]]] = ...) -> None: ...

class GetFavouriteModulesByUserIdRequest(_message.Message):
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

class GetFavouriteModulesByUserIdResponse(_message.Message):
    __slots__ = ("modules",)
    MODULES_FIELD_NUMBER: _ClassVar[int]
    modules: _containers.RepeatedCompositeFieldContainer[FullModule]
    def __init__(self, modules: _Optional[_Iterable[_Union[FullModule, _Mapping]]] = ...) -> None: ...

class GetModuleByIdRequest(_message.Message):
    __slots__ = ("user_id", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ...) -> None: ...

class GetModuleByIdResponse(_message.Message):
    __slots__ = ("module",)
    MODULE_FIELD_NUMBER: _ClassVar[int]
    module: FullModule
    def __init__(self, module: _Optional[_Union[FullModule, _Mapping]] = ...) -> None: ...

class GetModulesBySubjectIdRequest(_message.Message):
    __slots__ = ("user_id", "subject_id", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    subject_id: int
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., subject_id: _Optional[int] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetModulesBySubjectIdResponse(_message.Message):
    __slots__ = ("modules",)
    MODULES_FIELD_NUMBER: _ClassVar[int]
    modules: _containers.RepeatedCompositeFieldContainer[FullModule]
    def __init__(self, modules: _Optional[_Iterable[_Union[FullModule, _Mapping]]] = ...) -> None: ...

class GetModulesByNameSearchRequest(_message.Message):
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

class GetModulesByNameSearchResponse(_message.Message):
    __slots__ = ("modules",)
    MODULES_FIELD_NUMBER: _ClassVar[int]
    modules: _containers.RepeatedCompositeFieldContainer[FullModule]
    def __init__(self, modules: _Optional[_Iterable[_Union[FullModule, _Mapping]]] = ...) -> None: ...

class UpdateModuleRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "module_name", "module_description", "is_public")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_NAME_FIELD_NUMBER: _ClassVar[int]
    MODULE_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_PUBLIC_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    module_name: str
    module_description: str
    is_public: bool
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., module_name: _Optional[str] = ..., module_description: _Optional[str] = ..., is_public: bool = ...) -> None: ...

class UpdateModuleResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteModuleRequest(_message.Message):
    __slots__ = ("user_id", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ...) -> None: ...

class DeleteModuleResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SetUserModuleFavouriteRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "is_favourite")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    IS_FAVOURITE_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    is_favourite: bool
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., is_favourite: bool = ...) -> None: ...

class SetUserModuleFavouriteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SetUserModuleRoleRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "user_module_role")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    USER_MODULE_ROLE_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    user_module_role: _common_pb2.UserModuleRole
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., user_module_role: _Optional[_Union[_common_pb2.UserModuleRole, str]] = ...) -> None: ...

class SetUserModuleRoleResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class FullModule(_message.Message):
    __slots__ = ("module", "user_module_role", "is_favourite")
    MODULE_FIELD_NUMBER: _ClassVar[int]
    USER_MODULE_ROLE_FIELD_NUMBER: _ClassVar[int]
    IS_FAVOURITE_FIELD_NUMBER: _ClassVar[int]
    module: DBModule
    user_module_role: _common_pb2.UserModuleRole
    is_favourite: bool
    def __init__(self, module: _Optional[_Union[DBModule, _Mapping]] = ..., user_module_role: _Optional[_Union[_common_pb2.UserModuleRole, str]] = ..., is_favourite: bool = ...) -> None: ...

class DBModule(_message.Message):
    __slots__ = ("module_id", "module_name", "module_description", "is_public", "created_time", "updated_time")
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_NAME_FIELD_NUMBER: _ClassVar[int]
    MODULE_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_PUBLIC_FIELD_NUMBER: _ClassVar[int]
    CREATED_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATED_TIME_FIELD_NUMBER: _ClassVar[int]
    module_id: int
    module_name: str
    module_description: str
    is_public: bool
    created_time: int
    updated_time: int
    def __init__(self, module_id: _Optional[int] = ..., module_name: _Optional[str] = ..., module_description: _Optional[str] = ..., is_public: bool = ..., created_time: _Optional[int] = ..., updated_time: _Optional[int] = ...) -> None: ...

class DBUserModuleMap(_message.Message):
    __slots__ = ("user_id", "module_id", "user_module_role", "is_favourite")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    USER_MODULE_ROLE_FIELD_NUMBER: _ClassVar[int]
    IS_FAVOURITE_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    user_module_role: _common_pb2.UserModuleRole
    is_favourite: bool
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., user_module_role: _Optional[_Union[_common_pb2.UserModuleRole, str]] = ..., is_favourite: bool = ...) -> None: ...
