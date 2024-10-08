import { apiClient } from "./client";

// Payload and Response Interfaces

export interface createSubjectPayload {
  user_id: number;
  subject_name: string;
  subject_description: string;
  is_public: boolean;
}

export interface createSubjectResponse {
  message: string;
}

export const createSubject = async (
  payload: createSubjectPayload
): Promise<createSubjectResponse> => {
  const response = await apiClient.post("/subject", payload);
  return response.data;
};

export interface getPublicSubjectsPayload {
  page_number: number;
  
  sort_by: string;
  order: string;
}

export interface getPublicSubjectsResponse {
  subjects: DBSubject[];
}

export const getPublicSubjects = async (
  payload: getPublicSubjectsPayload
): Promise<getPublicSubjectsResponse> => {
  const response = await apiClient.get("/public/subject", { params: payload });
  return response.data;
};

export interface getPrivateSubjectsByUserIdPayload {
  user_id: number;
  page_number: number;
  
  sort_by: string;
  order: string;
}

export interface getPrivateSubjectsByUserIdResponse {
  subjects: DBSubject[];
}

export const getPrivateSubjectsByUserId = async (
  payload: getPrivateSubjectsByUserIdPayload
): Promise<getPrivateSubjectsByUserIdResponse> => {
  const response = await apiClient.get("/private/subject", { params: payload });
  return response.data;
};

export interface getFavouriteSubjectsByUserIdPayload {
  user_id: number;
  page_number: number;
  
  sort_by: string;
  order: string;
}

export interface getFavouriteSubjectsByUserIdResponse {
  subjects: DBSubject[];
}

export const getFavouriteSubjectsByUserId = async (
  payload: getFavouriteSubjectsByUserIdPayload
): Promise<getFavouriteSubjectsByUserIdResponse> => {
  const response = await apiClient.get("/favourite/subject", { params: payload });
  return response.data;
};

export interface getSubjectByIdPayload {
  user_id: number;
  subject_id: number;
}

export interface getSubjectByIdResponse {
  subject: DBSubject;
}

export const getSubjectById = async (
  payload: getSubjectByIdPayload
): Promise<getSubjectByIdResponse> => {
  const response = await apiClient.get(`/subject/${payload.subject_id}`, {
    params: { user_id: payload.user_id },
  });
  return response.data;
};

export interface getSubjectsByUserIdPayload {
  user_id: number;
  page_number: number;
  
  sort_by: string;
  order: string;
}

export interface getSubjectsByUserIdResponse {
  subjects: DBSubject[];
}

export const getSubjectsByUserId = async (
  payload: getSubjectsByUserIdPayload
): Promise<getSubjectsByUserIdResponse> => {
  const response = await apiClient.get("/user/subject", { params: payload });
  return response.data;
};

export interface getSubjectsByNameSearchPayload {
  user_id: number;
  search_query: string;
  page_number: number;
  
  sort_by: string;
  order: string;
}

export interface getSubjectsByNameSearchResponse {
  subjects: DBSubject[];
}

export const getSubjectsByNameSearch = async (
  payload: getSubjectsByNameSearchPayload
): Promise<getSubjectsByNameSearchResponse> => {
  const response = await apiClient.post("/search/subject", payload);
  return response.data;
};

export interface updateSubjectPayload {
  user_id: number;
  subject_id: number;
  subject_name: string;
  subject_description: string;
  is_public: boolean;
}

export interface updateSubjectResponse {
  message: string;
}

export const updateSubject = async (
  payload: updateSubjectPayload
): Promise<updateSubjectResponse> => {
  const response = await apiClient.put(`/subject/${payload.subject_id}`, payload);
  return response.data;
};

export interface deleteSubjectPayload {
  user_id: number;
  subject_id: number;
}

export interface deleteSubjectResponse {
  message: string;
}

export const deleteSubject = async (
  payload: deleteSubjectPayload
): Promise<deleteSubjectResponse> => {
  const response = await apiClient.delete(`/subject/${payload.subject_id}`, {
    params: { user_id: payload.user_id },
  });
  return response.data;
};

// DBSubject interface
export interface DBSubject {
  subject_id: number;
  subject_name: string;
  subject_description: string;
  is_public: boolean;
  created_time: number;
  updated_time: number;
}
