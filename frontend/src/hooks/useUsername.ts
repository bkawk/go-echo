import useFetch from "./useFetch";
import { useEffect, useState } from "react";

interface ReturnShape {
  isAvailable: boolean | null;
}
interface FormValues {
  name: string;
  username: string;
  email: string;
  password: string;
}

const useUsername = (username: string) => {
  const [formError, setFormError] = useState<FormValues | null>(null);
  const [serverError, setServerError] = useState<string | null>(null);

  const { isLoading, error, data, send } = useFetch<ReturnShape>({
    endpoint: `/check-username?username=${username}`,
  });

  useEffect(() => {
    if (error) {
      error.message && setServerError(error.message);
      setFormError(error.error);
    }
  }, [error]);

  useEffect(() => {
    if (username !== "") {
      send();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [username]);

  const isAvailable = username.length > 1 ? data?.isAvailable : null;

  return {
    isAvailable,
    isLoading,
    formError,
    serverError,
  };
};

export default useUsername;
