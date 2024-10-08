// createSession Type
import { apiClient } from "./client";

export interface createSessionPayload {
  username: string;
  password: string;
}
export interface createSessionResponse {
  message: string;
}

export const createSession = async (
  payload: createSessionPayload
): Promise<createSessionResponse> => {
  const { username, password } = payload;

  const basicAuth = btoa(`${username}:${password}`);

  const response = await apiClient.post("/session", null, {
    headers: {
      Authorization: `Basic ${basicAuth}`,
    },
  });
  return response.data;
};

export interface createUserPayload {
  username: string;
  password: string;
}
export interface createUserResponse {
  message: string;
}

export const createUser = async (
  payload: createUserPayload
): Promise<createUserResponse> => {
  const { username, password } = payload;

  const basicAuth = btoa(`${username}:${password}`);

  const response = await apiClient.post("/user", null, {
    headers: {
      Authorization: `Basic ${basicAuth}`,
    },
  });
  return response.data;
};

export const deleteSession = async (): Promise<void> => {
  await apiClient.delete("/session");
}
