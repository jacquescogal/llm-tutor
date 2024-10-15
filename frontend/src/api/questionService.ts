import { OrderByDirection, OrderByField, QuestionType } from "../types/enums";
import { apiClient } from "./client";

// Payload and Response Interfaces

export interface CreateQuestionPayload {
  user_id: number;
  doc_id: number;
  question_title: string;
  question_blob: Uint8Array;
  question_type: string;
  module_id: number;
}

export interface CreateQuestionResponse {
  message: string;
}

export const createQuestion = async (
  module_id: number,
  document_id: number,
  payload: CreateQuestionPayload
): Promise<CreateQuestionResponse> => {
  const response = await apiClient.post(`/module/${module_id}/document/${document_id}/question`, payload);
  return response.data;
};

export interface GetQuestionByIdPayload {
  user_id: number;
}

export interface GetQuestionByIdResponse {
  question: DBQuestion;
}

export const getQuestionById = async (
  module_id: number,
  document_id: number,
  question_id: number,
  payload: GetQuestionByIdPayload
): Promise<GetQuestionByIdResponse> => {
  const response = await apiClient.get(`/module/${module_id}/document/${document_id}/question/${question_id}`, {
    params: { user_id: payload.user_id },
  });
  return response.data;
};

export interface GetQuestionsByDocIdPayload {
  page_number: number;
  
  sort_by: OrderByField;
  order: OrderByDirection;
}


export const getQuestionsByDocId = async (
  module_id: number,
  document_id: number,
  payload: GetQuestionsByDocIdPayload
): Promise<QuestionReturn[]> => {
  const response = await apiClient.get(`/module/${module_id}/document/${document_id}/question`, {
    params: payload,
  });
  return response.data;
};

export interface GetQuestionsByQuestionTitleSearchPayload {
  user_id: number;
  search_query: string;
  page_number: number;
  
  sort_by: OrderByField;
  order: OrderByDirection;
}

export interface GetQuestionsByQuestionTitleSearchResponse {
  questions: DBQuestion[];
}

export const getQuestionsByQuestionTitleSearch = async (
  module_id: number,
  document_id: number,
  payload: GetQuestionsByQuestionTitleSearchPayload
): Promise<GetQuestionsByQuestionTitleSearchResponse> => {
  const response = await apiClient.post(`/search/module/${module_id}/document/${document_id}/question`, payload);
  return response.data;
};

export interface UpdateQuestionPayload {
  user_id: number;
  doc_id: number;
  question_title: string;
  question_blob: Uint8Array;
  question_type: string;
  module_id: number;
}

export interface UpdateQuestionResponse {
  message: string;
}

export const updateQuestion = async (
  module_id: number,
  document_id: number,
  question_id: number,
  payload: UpdateQuestionPayload
): Promise<UpdateQuestionResponse> => {
  const response = await apiClient.put(`/module/${module_id}/document/${document_id}/question/${question_id}`, payload);
  return response.data;
};

export interface DeleteQuestionPayload {
  user_id: number;
}

export interface DeleteQuestionResponse {
  message: string;
}

export const deleteQuestion = async (
  module_id: number,
  document_id: number,
  question_id: number,
  payload: DeleteQuestionPayload
): Promise<DeleteQuestionResponse> => {
  const response = await apiClient.delete(`/module/${module_id}/document/${document_id}/question/${question_id}`, {
    params: { user_id: payload.user_id },
  });
  return response.data;
};

// DBQuestion interface
export interface DBQuestion {
  question_id: number;
  user_id: number;
  doc_id: number;
  question_title: string;
  question_blob: Uint8Array;
  question_type: number;
  created_time: number;
  updated_time: number;
}


// QuestionReturn interface
export interface QuestionReturn {
  question_id: number;
  user_id: number;
  doc_id: number;
  question_title: string;
  question_serialized: string;
  question_type: QuestionType;
  created_time: number;
  updated_time: number;
}

// unserializedQuestion types
export interface MCQQuestion {
  choices: MCQChoice[];
}

export interface MCQChoice {
  choice: string;
  isCorrect: boolean;
}

export interface TextInputQuestion {
  answer: string;
}

export const unserializeQuestion = (question: QuestionReturn): MCQQuestion | TextInputQuestion | null => {
  let unserializedQuestion: MCQQuestion | TextInputQuestion | null = null;

  switch (question.question_type) {
    case QuestionType.QUESTION_TYPE_MCQ:
    case QuestionType.QUESTION_TYPE_MULTI_ANSWER_MCQ:
      // If the question type is MCQ or Multi-answer MCQ, parse it as MCQQuestion
      unserializedQuestion = JSON.parse(question.question_serialized) as MCQQuestion;
      break;

    case QuestionType.QUESTION_TYPE_OPEN_ENDED:
      // If the question type is Open Ended, parse it as TextInputQuestion
      unserializedQuestion = JSON.parse(question.question_serialized) as TextInputQuestion;
      break;

    default:
      throw new Error(`Unsupported question type: ${question.question_type}`);
  }

  // Return the original QuestionReturn object with the unserialized question as the serialized question
  return unserializedQuestion;
};
