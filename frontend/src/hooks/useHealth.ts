import useFetch from "./useFetch";
import { useEffect } from "react";

interface PropsShape {
  awaitClick?: boolean;
}
interface ReturnShape {
  status: string;
}

const useHealth = ({ awaitClick = false }: PropsShape = {}) => {
  const { isLoading, error, data, send } = useFetch<ReturnShape>({
    url: "http://localhost:8080/health",
  });

  const result = data?.status;

  useEffect(() => {
    !awaitClick && send();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const click = async () => {
    await send();
  };

  return { health: { isLoading, error, result }, click };
};

export default useHealth;
