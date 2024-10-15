import { OrderByDirection, OrderByField, UserSubjectRole } from "../types/enums";
import { apiClient } from "./client";

// Payload and Response Interfaces

export interface createSubjectPayload {
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
  sort_by?: OrderByField;
  order?: OrderByDirection;
}

export interface getPublicSubjectsResponse {
  subjects: FullSubject[];
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
  
  sort_by: OrderByField;
  order: OrderByDirection;
}

export interface getPrivateSubjectsByUserIdResponse {
  subjects: FullSubject[];
}

export const getPrivateSubjectsByUserId = async (
  payload: getPrivateSubjectsByUserIdPayload
): Promise<getPrivateSubjectsByUserIdResponse> => {
  const response = await apiClient.get("/private/subject", { params: payload });
  return response.data;
};

export interface getFavouriteSubjectsByUserIdPayload {
  page_number: number;
  
  sort_by: OrderByField;
  order: OrderByDirection;
}

export interface getFavouriteSubjectsByUserIdResponse {
  subjects: FullSubject[];
}

export const getFavouriteSubjectsByUserId = async (
  payload: getFavouriteSubjectsByUserIdPayload
): Promise<getFavouriteSubjectsByUserIdResponse> => {
  const response = await apiClient.get("/favourite/subject", { params: payload });
  return response.data;
};

export interface getSubjectByIdPayload {
  subject_id: number;
}

export interface getSubjectByIdResponse {
  subject: FullSubject;
}

export const getSubjectById = async (
  payload: getSubjectByIdPayload
): Promise<getSubjectByIdResponse> => {
  const response = await apiClient.get(`/subject/${payload.subject_id}`, {
  });
  return response.data;
};

export interface getSubjectsByUserIdPayload {
  user_id: number;
  page_number: number;
  
  sort_by: OrderByField;
  order: OrderByDirection;
}

export interface getSubjectsByUserIdResponse {
  subjects: FullSubject[];
}

export const getSubjectsByUserId = async (
  payload: getSubjectsByUserIdPayload
): Promise<getSubjectsByUserIdResponse> => {
  const response = await apiClient.get("/user/subject", { params: payload });
  return response.data;
};

export interface getSubjectsByNameSearchPayloadJSON {
  search_query: string;
}

export interface getSubjectsByNameSearchQuery {
    page_number: number;
  
  sort_by?: OrderByField;
  order?: OrderByDirection;
}

export interface getSubjectsByNameSearchResponse {
  subjects: FullSubject[];
}

export const getSubjectsByNameSearch = async (
  payload: getSubjectsByNameSearchPayloadJSON,
query: getSubjectsByNameSearchQuery
): Promise<getSubjectsByNameSearchResponse> => {
const response = await apiClient.post("/search/subject", payload, { params: query });
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


export interface setSubjectModuleMappingRequest {
  subject_id: number;
  module_ids: number[];
}

export interface setSubjectModuleMappingResponse {
  message: string;
}

export const setSubjectModuleMapping = async (
  payload: setSubjectModuleMappingRequest
): Promise<setSubjectModuleMappingResponse> => {
  const response = await apiClient.put(`/subject/${payload.subject_id}/module`, payload);
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

export interface setUserSubjectFavouritePayload {
    subject_id: number;
    is_favourite: boolean;
  }
  
  export interface setUserSubjectFavouriteResponse {
    message: string;
  }
  
  export const setUserSubjectFavourite = async (
    payload: setUserSubjectFavouritePayload
  ): Promise<setUserSubjectFavouriteResponse> => {
    const response = await apiClient.post(`favourite/subject/${payload.subject_id}`, { is_favourite: payload.is_favourite }
    );
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

// FullSubject interface
export interface FullSubject {
    subject: DBSubject;
    user_subject_role: UserSubjectRole;
    is_favourite: boolean;
  }
  
