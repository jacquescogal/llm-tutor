import { apiClient } from "./client";

// Payload and Response Interfaces

export interface CreateMemoryPayload {
  user_id: number;
  doc_id: number;
  memory_title: string;
  memory_content: string;
  is_public: boolean;
  module_id: number;
}

export interface CreateMemoryResponse {
  message: string;
}

export const createMemory = async (
  module_id: number,
  document_id: number,
  payload: CreateMemoryPayload
): Promise<CreateMemoryResponse> => {
  const response = await apiClient.post(`/module/${module_id}/document/${document_id}/memory`, payload);
  return response.data;
};

export interface GetMemoryByIdPayload {
  user_id: number;
}

export interface GetMemoryByIdResponse {
  memory: DBMemory;
}

export const getMemoryById = async (
  module_id: number,
  document_id: number,
  memory_id: number,
  payload: GetMemoryByIdPayload
): Promise<GetMemoryByIdResponse> => {
  const response = await apiClient.get(`/module/${module_id}/document/${document_id}/memory/${memory_id}`, {
    params: { user_id: payload.user_id },
  });
  return response.data;
};

export interface GetMemoriesByDocIdPayload {
  user_id: number;
  page_number: number;
  
  sort_by: string;
  order: string;
}

export interface GetMemoriesByDocIdResponse {
  memories: DBMemory[];
}

export const getMemoriesByDocId = async (
  module_id: number,
  document_id: number,
  payload: GetMemoriesByDocIdPayload
): Promise<GetMemoriesByDocIdResponse> => {
  const response = await apiClient.get(`/module/${module_id}/document/${document_id}/memory`, {
    params: payload,
  });
  return response.data;
};

export interface GetMemoriesByMemoryTitleSearchPayload {
  user_id: number;
  search_query: string;
  page_number: number;
  sort_by: string;
  order: string;
}

export interface GetMemoriesByMemoryTitleSearchResponse {
  memories: DBMemory[];
}

export const getMemoriesByMemoryTitleSearch = async (
  module_id: number,
  document_id: number,
  payload: GetMemoriesByMemoryTitleSearchPayload
): Promise<GetMemoriesByMemoryTitleSearchResponse> => {
  const response = await apiClient.post(`/search/module/${module_id}/document/${document_id}/memory`, payload);
  return response.data;
};

export interface UpdateMemoryPayload {
  user_id: number;
  doc_id: number;
  memory_title: string;
  memory_content: string;
  module_id: number;
}

export interface UpdateMemoryResponse {
  message: string;
}

export const updateMemory = async (
  module_id: number,
  document_id: number,
  memory_id: number,
  payload: UpdateMemoryPayload
): Promise<UpdateMemoryResponse> => {
  const response = await apiClient.put(`/module/${module_id}/document/${document_id}/memory/${memory_id}`, payload);
  return response.data;
};

export interface DeleteMemoryPayload {
  user_id: number;
}

export interface DeleteMemoryResponse {
  message: string;
}

export const deleteMemory = async (
  module_id: number,
  document_id: number,
  memory_id: number,
  payload: DeleteMemoryPayload
): Promise<DeleteMemoryResponse> => {
  const response = await apiClient.delete(`/module/${module_id}/document/${document_id}/memory/${memory_id}`, {
    params: { user_id: payload.user_id },
  });
  return response.data;
};

// DBMemory interface
export interface DBMemory {
  memory_id: number;
  user_id: number;
  doc_id: number;
  memory_title: string;
  memory_content: string;
  created_time: number;
  updated_time: number;
}
