
import axios, { AxiosError } from 'axios';

// Define the structure of the error response
interface ApiErrorResponse {
  error: string;
}

export const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL, // using environment variable
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});

// Axios Interceptor for Global Error Handling
apiClient.interceptors.response.use(
  response => response, // Pass the response through if no error
  (error: AxiosError) => {
    if (error.response) {
      // Server responded with a status other than 2xx
      const errorData = error.response.data as ApiErrorResponse;  // Type assertion
      throw new Error(errorData.error || "An error occurred");
    } else if (error.request) {
      // Request was made but no response received
      throw new Error("No response from the server");
    } else {
      // Something else happened during the request setup
      throw new Error("An unexpected error occurred");
    }
  }
);