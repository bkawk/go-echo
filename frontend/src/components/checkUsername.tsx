import { useState } from "react";
import useUsername from "../hooks/useUsername";

interface Props {}

const UsernameInput: React.FC<Props> = () => {
  const [usernameVal, setUsernameVal] = useState("");

  const { isAvailable, isLoading, formError, serverError } =
    useUsername(usernameVal);

  return (
    <div>
      <h1>Check Username</h1>
      <label>Enter your username: </label>
      <input
        type="text"
        value={usernameVal}
        onChange={(e) => {
          setUsernameVal(e.target.value);
        }}
      />
      <small>
        {isAvailable !== null &&
          (isAvailable ? <div>Available</div> : <div>Not Available</div>)}
      </small>
      {isLoading && <div>Loading...</div>}
      {formError?.username && <p>{formError?.username}</p>}
      {serverError && <p>Server error: {serverError}</p>}
    </div>
  );
};

export default UsernameInput;
