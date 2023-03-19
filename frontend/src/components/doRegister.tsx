import React, { useState } from "react";
import useRegister from "../hooks/useRegister";
import useUsername from "../hooks/useUsername";

interface FormValues {
  name: string;
  username: string;
  email: string;
  password: string;
}

interface Props {}

const DoRegister: React.FC<Props> = () => {
  const [formData, setFormData] = useState<FormValues>({
    name: "",
    username: "",
    email: "",
    password: "",
  });

  const { response, isLoading, formError, serverError, register } =
    useRegister();

  const { isAvailable } = useUsername(formData.username);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    register(formData);
  };

  return (
    <div>
      <h1>Register</h1>
      <p>{JSON.stringify(isLoading)}</p>
      <form onSubmit={handleSubmit}>
        <label htmlFor="name">Name:</label>
        <input type="text" id="name" name="name" onChange={handleInputChange} />
        {formError?.name && <p>{formError?.name}</p>}
        <br />

        <label htmlFor="username">Username:</label>
        <input
          type="text"
          id="username"
          name="username"
          onChange={handleInputChange}
        />
        <small>
          {isAvailable !== null &&
            (isAvailable ? <div>Available</div> : <div>Not Available</div>)}
        </small>
        {formError?.username && <p>{formError?.username}</p>}
        <br />

        <label htmlFor="email">Email:</label>
        <input
          type="email"
          id="email"
          name="email"
          onChange={handleInputChange}
        />
        {formError?.email && <p>{formError?.email}</p>}
        <br />

        <label htmlFor="password">Password:</label>
        <input
          type="password"
          id="password"
          name="password"
          onChange={handleInputChange}
        />
        {formError?.password && <p>{formError?.password}</p>}
        <br />

        {serverError && <p>Server error: {serverError}</p>}
        {response && <p>Registration successful!</p>}

        <input type="submit" value="Register" />
      </form>
    </div>
  );
};

export default DoRegister;
