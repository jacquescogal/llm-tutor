from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateUserRequest(_message.Message):
    __slots__ = ("username", "password")
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    username: str
    password: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class CreateUserResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class AuthenticateSessionRequest(_message.Message):
    __slots__ = ("session_id",)
    SESSION_ID_FIELD_NUMBER: _ClassVar[int]
    session_id: str
    def __init__(self, session_id: _Optional[str] = ...) -> None: ...

class AuthenticateSessionResponse(_message.Message):
    __slots__ = ("user_session",)
    USER_SESSION_FIELD_NUMBER: _ClassVar[int]
    user_session: UserSession
    def __init__(self, user_session: _Optional[_Union[UserSession, _Mapping]] = ...) -> None: ...

class CreateSessionRequest(_message.Message):
    __slots__ = ("username", "password")
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    username: str
    password: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class CreateSessionResponse(_message.Message):
    __slots__ = ("session_id",)
    SESSION_ID_FIELD_NUMBER: _ClassVar[int]
    session_id: str
    def __init__(self, session_id: _Optional[str] = ...) -> None: ...

class DeleteSessionRequest(_message.Message):
    __slots__ = ("session_id",)
    SESSION_ID_FIELD_NUMBER: _ClassVar[int]
    session_id: str
    def __init__(self, session_id: _Optional[str] = ...) -> None: ...

class DeleteSessionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DBUser(_message.Message):
    __slots__ = ("user_id", "username", "hash_salt_password")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    HASH_SALT_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    username: str
    hash_salt_password: str
    def __init__(self, user_id: _Optional[int] = ..., username: _Optional[str] = ..., hash_salt_password: _Optional[str] = ...) -> None: ...

class UserSession(_message.Message):
    __slots__ = ("user_id", "username")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    username: str
    def __init__(self, user_id: _Optional[int] = ..., username: _Optional[str] = ...) -> None: ...
