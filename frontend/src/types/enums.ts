// from proto
export enum UserSubjectRole {
    USER_SUBJECT_ROLE_UNDEFINED = 0,
    USER_SUBJECT_ROLE_OWNER = 1, // can give access to other users + editor rights
    USER_SUBJECT_ROLE_ADMIN = 2, // can give access to other users + editor rights
    USER_SUBJECT_ROLE_EDITOR = 3, // CRUD + viewer rights
    USER_SUBJECT_ROLE_VIEWER = 4 // allows reads when subject is not public
}

export enum UserModuleRole {
    USER_MODULE_ROLE_UNDEFINED = 0,
    USER_MODULE_ROLE_OWNER = 1, // can give access to other users + editor rights
    USER_MODULE_ROLE_ADMIN = 2, // can give access to other users + editor rights
    USER_MODULE_ROLE_EDITOR = 3, // CRUD + viewer rights
    USER_MODULE_ROLE_VIEWER = 4 // allows reads when doc is not public
}


export enum UploadStatus {
    UPLOAD_STATUS_NOT_STARTED = 0,
    UPLOAD_STATUS_QUEUEING = 1,
    UPLOAD_STATUS_PROCESSING = 2,
    UPLOAD_STATUS_SUCCESS = 3,
    UPLOAD_STATUS_FAILED = 4,
}

export enum QuestionType {
    QUESTION_TYPE_UNDEFINED = 0,
    QUESTION_TYPE_MCQ = 1,
    QUESTION_TYPE_MULTI_ANSWER_MCQ = 2,
    QUESTION_TYPE_OPEN_ENDED = 3,
}

export enum OrderByDirection {
    ORDER_BY_DIRECTION_UNDEFINED = 0,
    ORDER_BY_DIRECTION_ASC = 1,
    ORDER_BY_DIRECTION_DESC = 2
}

export enum OrderByField {
    ORDER_BY_FIELD_UNDEFINED = 0,
    ORDER_BY_FIELD_ID = 1,
    ORDER_BY_FIELD_TITLE = 2,
    ORDER_BY_FIELD_CREATED_TIME = 3,
    ORDER_BY_FIELD_UPDATED_TIME = 4
}
