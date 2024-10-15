from src.protos import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateDocRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "doc_name", "doc_description", "doc_summary", "upload_status", "s3_object_key")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_NAME_FIELD_NUMBER: _ClassVar[int]
    DOC_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DOC_SUMMARY_FIELD_NUMBER: _ClassVar[int]
    UPLOAD_STATUS_FIELD_NUMBER: _ClassVar[int]
    S3_OBJECT_KEY_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    doc_name: str
    doc_description: str
    doc_summary: str
    upload_status: _common_pb2.UploadStatus
    s3_object_key: str
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_name: _Optional[str] = ..., doc_description: _Optional[str] = ..., doc_summary: _Optional[str] = ..., upload_status: _Optional[_Union[_common_pb2.UploadStatus, str]] = ..., s3_object_key: _Optional[str] = ...) -> None: ...

class CreateDocResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetDocByIdRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "doc_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    doc_id: int
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_id: _Optional[int] = ...) -> None: ...

class GetDocByIdResponse(_message.Message):
    __slots__ = ("doc",)
    DOC_FIELD_NUMBER: _ClassVar[int]
    doc: DBDoc
    def __init__(self, doc: _Optional[_Union[DBDoc, _Mapping]] = ...) -> None: ...

class GetDocsByModuleIdRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetDocsByModuleIdResponse(_message.Message):
    __slots__ = ("docs",)
    DOCS_FIELD_NUMBER: _ClassVar[int]
    docs: _containers.RepeatedCompositeFieldContainer[DBDoc]
    def __init__(self, docs: _Optional[_Iterable[_Union[DBDoc, _Mapping]]] = ...) -> None: ...

class GetDocsByNameSearchRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "search_query", "page_number", "page_size", "order_by_field", "order_by_direction")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    SEARCH_QUERY_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    search_query: str
    page_number: int
    page_size: int
    order_by_field: _common_pb2.ORDER_BY_FIELD
    order_by_direction: _common_pb2.ORDER_BY_DIRECTION
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., search_query: _Optional[str] = ..., page_number: _Optional[int] = ..., page_size: _Optional[int] = ..., order_by_field: _Optional[_Union[_common_pb2.ORDER_BY_FIELD, str]] = ..., order_by_direction: _Optional[_Union[_common_pb2.ORDER_BY_DIRECTION, str]] = ...) -> None: ...

class GetDocsByNameSearchResponse(_message.Message):
    __slots__ = ("docs",)
    DOCS_FIELD_NUMBER: _ClassVar[int]
    docs: _containers.RepeatedCompositeFieldContainer[DBDoc]
    def __init__(self, docs: _Optional[_Iterable[_Union[DBDoc, _Mapping]]] = ...) -> None: ...

class UpdateDocRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "doc_id", "doc_name", "doc_description", "doc_summary", "upload_status", "s3_object_key")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_NAME_FIELD_NUMBER: _ClassVar[int]
    DOC_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DOC_SUMMARY_FIELD_NUMBER: _ClassVar[int]
    UPLOAD_STATUS_FIELD_NUMBER: _ClassVar[int]
    S3_OBJECT_KEY_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    doc_id: int
    doc_name: str
    doc_description: str
    doc_summary: str
    upload_status: _common_pb2.UploadStatus
    s3_object_key: str
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., doc_name: _Optional[str] = ..., doc_description: _Optional[str] = ..., doc_summary: _Optional[str] = ..., upload_status: _Optional[_Union[_common_pb2.UploadStatus, str]] = ..., s3_object_key: _Optional[str] = ...) -> None: ...

class UpdateDocResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteDocRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "doc_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    doc_id: int
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_id: _Optional[int] = ...) -> None: ...

class DeleteDocResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UpdateSummaryRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "doc_id", "doc_summary", "upload_status")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_SUMMARY_FIELD_NUMBER: _ClassVar[int]
    UPLOAD_STATUS_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    doc_id: int
    doc_summary: str
    upload_status: _common_pb2.UploadStatus
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., doc_summary: _Optional[str] = ..., upload_status: _Optional[_Union[_common_pb2.UploadStatus, str]] = ...) -> None: ...

class UpdateSummaryResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UpdateUploadStatusRequest(_message.Message):
    __slots__ = ("user_id", "module_id", "doc_id", "upload_status")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    UPLOAD_STATUS_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    module_id: int
    doc_id: int
    upload_status: _common_pb2.UploadStatus
    def __init__(self, user_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., upload_status: _Optional[_Union[_common_pb2.UploadStatus, str]] = ...) -> None: ...

class UpdateUploadStatusResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DBDoc(_message.Message):
    __slots__ = ("doc_id", "module_id", "doc_name", "doc_description", "doc_summary", "upload_status", "s3_object_key", "created_time", "updated_time")
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_NAME_FIELD_NUMBER: _ClassVar[int]
    DOC_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DOC_SUMMARY_FIELD_NUMBER: _ClassVar[int]
    UPLOAD_STATUS_FIELD_NUMBER: _ClassVar[int]
    S3_OBJECT_KEY_FIELD_NUMBER: _ClassVar[int]
    CREATED_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATED_TIME_FIELD_NUMBER: _ClassVar[int]
    doc_id: int
    module_id: int
    doc_name: str
    doc_description: str
    doc_summary: str
    upload_status: _common_pb2.UploadStatus
    s3_object_key: str
    created_time: int
    updated_time: int
    def __init__(self, doc_id: _Optional[int] = ..., module_id: _Optional[int] = ..., doc_name: _Optional[str] = ..., doc_description: _Optional[str] = ..., doc_summary: _Optional[str] = ..., upload_status: _Optional[_Union[_common_pb2.UploadStatus, str]] = ..., s3_object_key: _Optional[str] = ..., created_time: _Optional[int] = ..., updated_time: _Optional[int] = ...) -> None: ...
