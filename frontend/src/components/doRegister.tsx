import React, { useState } from "react";
import useRegister from "../hooks/useRegister";

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

      <form onSubmit={handleSubmit}>
        <input type="hidden" name="Content-Type" value="application/json" />

        <label htmlFor="name">Name:</label>
        <input
          type="text"
          id="name"
          name="name"
          required
          onChange={handleInputChange}
        />
        <br />

        <label htmlFor="username">Username:</label>
        <input
          type="text"
          id="username"
          name="username"
          required
          onChange={handleInputChange}
        />
        <br />

        <label htmlFor="email">Email:</label>
        <input
          type="email"
          id="email"
          name="email"
          required
          onChange={handleInputChange}
        />
        <br />

        <label htmlFor="password">Password:</label>
        <input
          type="password"
          id="password"
          name="password"
          required
          onChange={handleInputChange}
        />
        <br />

        {formError?.password && <p>Form error: {formError?.password}</p>}
        {formError?.email && <p>Form error: {formError?.email}</p>}
        {formError?.username && <p>Form error: {formError?.username}</p>}
        {formError?.name && <p>Form error: {formError?.name}</p>}
        {serverError && <p>Server error: {serverError}</p>}

        <input type="submit" value="Register" />
      </form>
    </div>
  );
};

export default DoRegister;
