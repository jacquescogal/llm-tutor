import { apiClient } from "./client";

// Interfaces for request payloads and responses
export interface CreateModulePayload {
  module_name: string;
  module_description: string;
  is_public: boolean;
}

export interface CreateModuleResponse {
    message: string;
}

export const createModule = async (
  payload: CreateModulePayload
): Promise<CreateModuleResponse> => {
  const response = await apiClient.post("/module", payload);
  return response.data;
};

export interface GetModulesRequest {
  page_number: number;
  sort_by: string;
  order: string;
}

export interface GetModulesResponse {
  modules: DBModule[];
}

export const getPublicModules = async (
  payload: GetModulesRequest
): Promise<GetModulesResponse> => {
  const response = await apiClient.get("/public/module", {
    params: payload,
  });
  return response.data;
};

export interface GetPrivateModulesRequest {
  page_number: number;
  
  sort_by: string;
  order: string;
}

export const getPrivateModulesByUserId = async (
  payload: GetPrivateModulesRequest
): Promise<GetModulesResponse> => {
  const response = await apiClient.get("/private/module", {
    params: payload,
  });
  return response.data;
};

export interface GetFavouriteModulesRequest {
  
  page_number: number;
  
  sort_by: string;
  order: string;
}

export const getFavouriteModulesByUserId = async (
  payload: GetFavouriteModulesRequest
): Promise<GetModulesResponse> => {
  const response = await apiClient.get("/favourite/module", {
    params: payload,
  });
  return response.data;
};

export interface GetModuleByIdRequest {
  
  module_id: number;
}

export interface GetModuleByIdResponse {
  module: DBModule;
}

export const getModuleById = async (
  payload: GetModuleByIdRequest
): Promise<GetModuleByIdResponse> => {
  const response = await apiClient.get(`/module/${payload.module_id}`, {
    params: {},
  });
  return response.data;
};

export interface GetModulesBySubjectIdRequest {
  
  subjectId: number;
  page_number: number;
  
  sort_by: string;
  order: string;
}

export const getModulesBySubjectId = async (
  payload: GetModulesBySubjectIdRequest
): Promise<GetModulesResponse> => {
  const response = await apiClient.get(`/subject/${payload.subjectId}/module`, {
    params: payload,
  });
  return response.data;
};

export interface GetModulesByNameSearchRequest {
  
  searchQuery: string;
  page_number: number;
  
  sort_by: string;
  order: string;
}

export const getModulesByNameSearch = async (
  payload: GetModulesByNameSearchRequest
): Promise<GetModulesResponse> => {
  const response = await apiClient.post("/search/module", payload);
  return response.data;
};

export interface UpdateModulePayload {
  
  module_id: number;
  module_name: string;
  module_description: string;
  is_public: boolean;
}

export interface UpdateModuleResponse {
    message: string;
}

export const updateModule = async (
  payload: UpdateModulePayload
): Promise<UpdateModuleResponse> => {
  const response = await apiClient.put(`/module/${payload.module_id}`, payload);
  return response.data;
};

export interface DeleteModulePayload {
  
  module_id: number;
}

export interface DeleteModuleResponse {
    message: string;
}

export const deleteModule = async (
  payload: DeleteModulePayload
): Promise<DeleteModuleResponse> => {
  const response = await apiClient.delete(`/module/${payload.module_id}`, {
    params: { },
  });
  return response.data;
};

// DBModule interface
export interface DBModule {
  module_id: number;
  module_name: string;
  module_description: string;
  is_public: boolean;
  createdTime: number;
  updatedTime: number;
}
