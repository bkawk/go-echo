import React from "react";
import useHealth from "../hooks/useHealth";

interface Props {}

const CheckHealth: React.FC<Props> = () => {
  const { health } = useHealth();

  return (
    <div>
      <h1>Health Check</h1>
      <pre>
        {health.result}
        {health.isLoading && <div>Loading</div>}
        {health.error && <div>Error</div>}
      </pre>
    </div>
  );
};

export default CheckHealth;
