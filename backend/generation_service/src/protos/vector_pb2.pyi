from src.protos import common_pb2 as _common_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateMemoryVectorRequest(_message.Message):
    __slots__ = ("module_id", "memory_id", "memory_title", "memory_content")
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_ID_FIELD_NUMBER: _ClassVar[int]
    MEMORY_TITLE_FIELD_NUMBER: _ClassVar[int]
    MEMORY_CONTENT_FIELD_NUMBER: _ClassVar[int]
    module_id: int
    memory_id: int
    memory_title: str
    memory_content: str
    def __init__(self, module_id: _Optional[int] = ..., memory_id: _Optional[int] = ..., memory_title: _Optional[str] = ..., memory_content: _Optional[str] = ...) -> None: ...

class CreateMemoryVectorResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SearchMemoryVectorRequest(_message.Message):
    __slots__ = ("search_query", "limit", "id", "id_type")
    SEARCH_QUERY_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    ID_TYPE_FIELD_NUMBER: _ClassVar[int]
    search_query: str
    limit: int
    id: int
    id_type: _common_pb2.IDType
    def __init__(self, search_query: _Optional[str] = ..., limit: _Optional[int] = ..., id: _Optional[int] = ..., id_type: _Optional[_Union[_common_pb2.IDType, str]] = ...) -> None: ...

class SearchMemoryVectorResponse(_message.Message):
    __slots__ = ("json_response",)
    JSON_RESPONSE_FIELD_NUMBER: _ClassVar[int]
    json_response: bytes
    def __init__(self, json_response: _Optional[bytes] = ...) -> None: ...
