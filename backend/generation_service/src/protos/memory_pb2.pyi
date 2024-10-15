from src.protos import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateMemoryRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "memory_title", "memory_content", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_TITLE_FIELD_NUMBER: _ClassVar[int]
    MEMORY_CONTENT_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    memory_title: str
    memory_content: str
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., memory_title: _Optional[str] = ..., memory_content: _Optional[str] = ..., module_id: _Optional[int] = ...) -> None: ...

class CreateMemoryResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetMemoryByIdRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "memory_id", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    memory_id: int
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., memory_id: _Optional[int] = ..., module_id: _Optional[int] = ...) -> None: ...

class GetMemoryByIdResponse(_message.Message):
    __slots__ = ("memory",)
    MEMORY_FIELD_NUMBER: _ClassVar[int]
    memory: DBMemory
    def __init__(self, memory: _Optional[_Union[DBMemory, _Mapping]] = ...) -> None: ...

class GetMemoriesByDocIdRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "page_number", "page_size", "order_by_field", "order_by_direction", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ..., module_id: _Optional[int] = ...) -> None: ...

class GetMemoriesByDocIdResponse(_message.Message):
    __slots__ = ("memories",)
    MEMORIES_FIELD_NUMBER: _ClassVar[int]
    memories: _containers.RepeatedCompositeFieldContainer[DBMemory]
    def __init__(self, memories: _Optional[_Iterable[_Union[DBMemory, _Mapping]]] = ...) -> None: ...

class GetMemoriesByMemoryTitleSearchRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "search_query", "page_number", "page_size", "order_by_field", "order_by_direction", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    SEARCH_QUERY_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    search_query: str
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., search_query: _Optional[str] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ..., module_id: _Optional[int] = ...) -> None: ...

class GetMemoriesByMemoryTitleSearchResponse(_message.Message):
    __slots__ = ("memories",)
    MEMORIES_FIELD_NUMBER: _ClassVar[int]
    memories: _containers.RepeatedCompositeFieldContainer[DBMemory]
    def __init__(self, memories: _Optional[_Iterable[_Union[DBMemory, _Mapping]]] = ...) -> None: ...

class UpdateMemoryRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "memory_id", "memory_title", "memory_content", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_TITLE_FIELD_NUMBER: _ClassVar[int]
    MEMORY_CONTENT_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    memory_id: int
    memory_title: str
    memory_content: str
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., memory_id: _Optional[int] = ..., memory_title: _Optional[str] = ..., memory_content: _Optional[str] = ..., module_id: _Optional[int] = ...) -> None: ...

class UpdateMemoryResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteMemoryRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "memory_id", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    memory_id: int
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., memory_id: _Optional[int] = ..., module_id: _Optional[int] = ...) -> None: ...

class DeleteMemoryResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DBMemory(_message.Message):
    __slots__ = ("memory_id", "user_id", "doc_id", "memory_title", "memory_content", "created_time", "updated_time", "vector_uuid")
    MEMORY_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_TITLE_FIELD_NUMBER: _ClassVar[int]
    MEMORY_CONTENT_FIELD_NUMBER: _ClassVar[int]
    CREATED_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATED_TIME_FIELD_NUMBER: _ClassVar[int]
    VECTOR_UUID_FIELD_NUMBER: _ClassVar[int]
    memory_id: int
    user_id: int
    doc_id: int
    memory_title: str
    memory_content: str
    created_time: int
    updated_time: int
    vector_uuid: str
    def __init__(self, memory_id: _Optional[int] = ..., user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., memory_title: _Optional[str] = ..., memory_content: _Optional[str] = ..., created_time: _Optional[int] = ..., updated_time: _Optional[int] = ..., vector_uuid: _Optional[str] = ...) -> None: ...
