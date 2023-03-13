import React, { useState } from "react";

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

  const { response, isLoading, error, register } = useRegister({ formData });

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  return (
    <div>
      <h1>Register</h1>

      <form onSubmit={register}>
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

        <input type="submit" value="Register" />
      </form>
    </div>
  );
};

export default DoRegister;
