from src.protos import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateQuestionRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "question_title", "question_blob", "question_type", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TITLE_FIELD_NUMBER: _ClassVar[int]
    QUESTION_BLOB_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TYPE_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    question_title: str
    question_blob: bytes
    question_type: _common_pb2.QuestionType
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., question_title: _Optional[str] = ..., question_blob: _Optional[bytes] = ..., question_type: _Optional[_Union[_common_pb2.QuestionType, str]] = ..., module_id: _Optional[int] = ...) -> None: ...

class CreateQuestionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetQuestionByIdRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "question_id", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    question_id: int
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., question_id: _Optional[int] = ..., module_id: _Optional[int] = ...) -> None: ...

class GetQuestionByIdResponse(_message.Message):
    __slots__ = ("question",)
    QUESTION_FIELD_NUMBER: _ClassVar[int]
    question: DBQuestion
    def __init__(self, question: _Optional[_Union[DBQuestion, _Mapping]] = ...) -> None: ...

class GetQuestionsByDocIdRequest(_message.Message):
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

class GetQuestionsByDocIdResponse(_message.Message):
    __slots__ = ("questions",)
    QUESTIONS_FIELD_NUMBER: _ClassVar[int]
    questions: _containers.RepeatedCompositeFieldContainer[DBQuestion]
    def __init__(self, questions: _Optional[_Iterable[_Union[DBQuestion, _Mapping]]] = ...) -> None: ...

class GetQuestionsByQuestionTitleSearchRequest(_message.Message):
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

class GetQuestionsByQuestionTitleSearchResponse(_message.Message):
    __slots__ = ("questions",)
    QUESTIONS_FIELD_NUMBER: _ClassVar[int]
    questions: _containers.RepeatedCompositeFieldContainer[DBQuestion]
    def __init__(self, questions: _Optional[_Iterable[_Union[DBQuestion, _Mapping]]] = ...) -> None: ...

class UpdateQuestionRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "question_id", "question_title", "question_blob", "question_type", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TITLE_FIELD_NUMBER: _ClassVar[int]
    QUESTION_BLOB_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TYPE_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    question_id: int
    question_title: str
    question_blob: bytes
    question_type: _common_pb2.QuestionType
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., question_id: _Optional[int] = ..., question_title: _Optional[str] = ..., question_blob: _Optional[bytes] = ..., question_type: _Optional[_Union[_common_pb2.QuestionType, str]] = ..., module_id: _Optional[int] = ...) -> None: ...

class UpdateQuestionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteQuestionRequest(_message.Message):
    __slots__ = ("user_id", "doc_id", "question_id", "module_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_ID_FIELD_NUMBER: _ClassVar[int]
    MODULE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    doc_id: int
    question_id: int
    module_id: int
    def __init__(self, user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., question_id: _Optional[int] = ..., module_id: _Optional[int] = ...) -> None: ...

class DeleteQuestionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DBQuestion(_message.Message):
    __slots__ = ("question_id", "user_id", "doc_id", "question_title", "question_blob", "question_type", "created_time", "updated_time")
    QUESTION_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DOC_ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TITLE_FIELD_NUMBER: _ClassVar[int]
    QUESTION_BLOB_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TYPE_FIELD_NUMBER: _ClassVar[int]
    CREATED_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATED_TIME_FIELD_NUMBER: _ClassVar[int]
    question_id: int
    user_id: int
    doc_id: int
    question_title: str
    question_blob: bytes
    question_type: _common_pb2.QuestionType
    created_time: int
    updated_time: int
    def __init__(self, question_id: _Optional[int] = ..., user_id: _Optional[int] = ..., doc_id: _Optional[int] = ..., question_title: _Optional[str] = ..., question_blob: _Optional[bytes] = ..., question_type: _Optional[_Union[_common_pb2.QuestionType, str]] = ..., created_time: _Optional[int] = ..., updated_time: _Optional[int] = ...) -> None: ...

class MCQQuestion(_message.Message):
    __slots__ = ("choices",)
    CHOICES_FIELD_NUMBER: _ClassVar[int]
    choices: _containers.RepeatedCompositeFieldContainer[MCQChoice]
    def __init__(self, choices: _Optional[_Iterable[_Union[MCQChoice, _Mapping]]] = ...) -> None: ...

class MCQChoice(_message.Message):
    __slots__ = ("choice", "is_correct")
    CHOICE_FIELD_NUMBER: _ClassVar[int]
    IS_CORRECT_FIELD_NUMBER: _ClassVar[int]
    choice: str
    is_correct: bool
    def __init__(self, choice: _Optional[str] = ..., is_correct: bool = ...) -> None: ...

class TextInputQuestion(_message.Message):
    __slots__ = ("answer",)
    ANSWER_FIELD_NUMBER: _ClassVar[int]
    answer: str
    def __init__(self, answer: _Optional[str] = ...) -> None: ...
