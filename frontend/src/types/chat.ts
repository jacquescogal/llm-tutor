export enum ID_TYPE {
    UNSPECIFIED = 0,
    MODULE = 1,
    SUBJECT = 2,
    DOCUMENT = 3
}
export interface ChatHistory {
    chat_history: Chat[];
}
export interface Chat {
    entity_type: ENTITY_TYPE;
    content: string;
}

export enum ENTITY_TYPE {
    UNDEFINED = 0,
    SYSTEM  = 1,
    USER = 2,
    BOT = 3
}