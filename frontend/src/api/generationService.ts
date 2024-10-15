// createSession Type
import { Chat, ID_TYPE } from "../types/chat";
import { apiClient } from "./client";

export interface createGenerationRequest {
  id: number;
  id_type: ID_TYPE
  chat_history: Chat[];
}
export interface createGenerationResponse {
  response: string;
}

export const createGeneration = async (
  payload: createGenerationRequest
): Promise<createGenerationResponse> => {
    const response = await apiClient.post(`/generate`, payload);
  return response.data;
};
