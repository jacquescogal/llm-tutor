import { apiClient } from "./client";


export interface CreateDocumentPayload {
    doc_name: string;
    doc_description: string;
  }
  
  export interface CreateDocumentResponse {
    message: string;
  }
  
  export const createDocument = async (
    module_id: number,
    file: File,
    jsonPayload: CreateDocumentPayload
  ): Promise<CreateDocumentResponse> => {
    const formData = new FormData();
  
    // Append the file and the JSON payload to formData
    formData.append("file", file);
    formData.append("json", JSON.stringify(jsonPayload));
  
    const response = await apiClient.post(`/module/${module_id}/document`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  
    return response.data;
  };


export interface GetDocumentByIdResponse {
  doc: DBDoc;
}

export const getDocumentById = async (
  module_id: number,
  document_id: number
): Promise<GetDocumentByIdResponse> => {
  const response = await apiClient.get(`/module/${module_id}/document/${document_id}`, {
  });
  return response.data;
};

export interface GetDocumentsByModuleIdPayload {
  page_number: number;
  sort_by: string;
  order: string;
}

export interface GetDocumentsByModuleIdResponse {
  docs: DBDoc[];
}

export const getDocumentsByModuleId = async (
  module_id: number,
  payload: GetDocumentsByModuleIdPayload
): Promise<GetDocumentsByModuleIdResponse> => {
  const response = await apiClient.get(`/module/${module_id}/document`, { params: payload });
  return response.data;
};

export interface GetDocumentsByNameSearchPayload {
  search_query: string;
  page_number: number;
  sort_by: string;
  order: string;
}

export interface GetDocumentsByNameSearchResponse {
  docs: DBDoc[];
}

export const getDocumentsByNameSearch = async (
  module_id: number,
  payload: GetDocumentsByNameSearchPayload
): Promise<GetDocumentsByNameSearchResponse> => {
  const response = await apiClient.post(`/search/module/${module_id}/document`, payload);
  return response.data;
};

export interface UpdateDocumentPayload {
  doc_name: string;
  doc_description: string;
  doc_summary: string;
  upload_status: string;
  s3_object_key: string;
}

export interface UpdateDocumentResponse {
  message: string;
}

export const updateDocument = async (
  module_id: number,
  document_id: number,
  payload: UpdateDocumentPayload
): Promise<UpdateDocumentResponse> => {
  const response = await apiClient.put(`/module/${module_id}/document/${document_id}`, payload);
  return response.data;
};

export interface DeleteDocumentResponse {
  message: string;
}

export const deleteDocument = async (
  module_id: number,
  document_id: number
): Promise<DeleteDocumentResponse> => {
  const response = await apiClient.delete(`/module/${module_id}/document/${document_id}`, {
  });
  return response.data;
};

// DBDoc interface
export interface DBDoc {
  doc_id: number;
  module_id: number;
  doc_name: string;
  doc_description: string;
  doc_summary: string;
  upload_status: string;
  s3_object_key: string;
  created_time: number;
  updated_time: number;
}
