import { useState } from "react";

interface Options {
  url: string;
}
interface FetchData<T> {
  data: T | null;
  isLoading: boolean;
  error: any;
  send: () => Promise<void>;
}

const useFetch = <T>({ url }: Options): FetchData<T> => {
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const send = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(url);
      const json = await response.json();
      setData(json);
    } catch (error: any) {
      setError(error);
    } finally {
      setIsLoading(false);
    }
  };

  return { data, isLoading, error, send };
};

export default useFetch;
