import { useState } from "react";
import useUsername from "../hooks/useUsername";

interface Props {}

const UsernameInput: React.FC<Props> = () => {
  const [usernameVal, setUsernameVal] = useState("");

  const { username } = useUsername({ username: usernameVal });

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
        {username.isAvailable !== null &&
          (username.isAvailable ? (
            <div>Available</div>
          ) : (
            <div>Not Available</div>
          ))}
      </small>
    </div>
  );
};

export default UsernameInput;
