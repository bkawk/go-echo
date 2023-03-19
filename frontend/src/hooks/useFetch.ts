import { useState, useCallback } from "react";

type HttpMethod = "GET" | "POST" | "DELETE" | "PATCH";

interface Options {
  endpoint: string;
}

interface ErrorShape {
  message?: string;
  error?: any;
}

interface FetchData<T> {
  data: T | null;
  isLoading: boolean;
  error: ErrorShape | null;
  send: (body?: object, method?: HttpMethod) => Promise<void>;
}

const useFetch = <T>({ endpoint }: Options): FetchData<T> => {
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<ErrorShape | null>(null);

  const send = useCallback(
    async (body?: object, method: HttpMethod = "GET") => {
      setIsLoading(true);
      try {
        const requestOptions: RequestInit = {
          method,
          headers: {
            "Content-Type": "application/json",
          },
        };

        if (
          body &&
          (method === "POST" || method === "PATCH" || method === "DELETE")
        ) {
          requestOptions.body = JSON.stringify(body);
        }

        const baseUrl = process.env.REACT_APP_API_BASE_URL;
        const response = await fetch(`${baseUrl}${endpoint}`, requestOptions);
        const json = await response.json();
        if (!response.ok) {
          setError(json);
          setIsLoading(false);
          return;
        }
        setData(json);
        setIsLoading(false);
      } catch (error: any) {
        setError({ message: error.message });
        setIsLoading(false);
      }
    },
    [endpoint]
  );

  return { data, isLoading, error, send };
};

export default useFetch;
