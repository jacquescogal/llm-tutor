from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class DocumentProcessingJob(_message.Message):
    __slots__ = ("user_id", "module_id", "doc_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    doc_id: int
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_id: _Optional[int] = ...) -> None: ...
