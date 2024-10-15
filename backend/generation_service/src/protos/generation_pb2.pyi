from src.protos import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EntityType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ENTITY_TYPE_UNSPECIFIED: _ClassVar[EntityType]
    ENTITY_TYPE_SYSTEM: _ClassVar[EntityType]
    ENTITY_TYPE_USER: _ClassVar[EntityType]
    ENTITY_TYPE_BOT: _ClassVar[EntityType]
ENTITY_TYPE_UNSPECIFIED: EntityType
ENTITY_TYPE_SYSTEM: EntityType
ENTITY_TYPE_USER: EntityType
ENTITY_TYPE_BOT: EntityType

class CreateGenerationRequest(_message.Message):
    __slots__ = ("user_id", "id", "id_type", "chat_history")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    ID_TYPE_FIELD_NUMBER: _ClassVar[int]
    CHAT_HISTORY_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    id: int
    id_type: _common_pb2.IDType
    chat_history: _containers.RepeatedCompositeFieldContainer[ChatHistory]
    def __init__(self, user_id: _Optional[int] = ..., id: _Optional[int] = ..., id_type: _Optional[_Union[_common_pb2.IDType, str]] = ..., chat_history: _Optional[_Iterable[_Union[ChatHistory, _Mapping]]] = ...) -> None: ...

class CreateGenerationResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: str
    def __init__(self, response: _Optional[str] = ...) -> None: ...

class ChatHistory(_message.Message):
    __slots__ = ("entity_type", "content")
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    entity_type: EntityType
    content: str
    def __init__(self, entity_type: _Optional[_Union[EntityType, str]] = ..., content: _Optional[str] = ...) -> None: ...
