import useFetch from "./useFetch";
import { useEffect } from "react";

interface PropsShape {
  username: string;
}
interface ReturnShape {
  isAvailable: boolean | null;
}

const useUsername = ({ username }: PropsShape) => {
  const { isLoading, error, data, send } = useFetch<ReturnShape>({
    url: `http://localhost:8080/check-username?username=${username}`,
  });

  useEffect(() => {
    send();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [username]);

  return {
    username: {
      isAvailable: username.length > 1 ? data?.isAvailable : null,
      isLoading,
      error,
    },
  };
};

export default useUsername;
