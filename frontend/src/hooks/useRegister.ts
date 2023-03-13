import { useState } from "react";

interface FormValues {
  name: string;
  username: string;
  email: string;
  password: string;
}

interface UseRegisterReturn<T extends FormValues> {
  response: any | undefined;
  isLoading: boolean;
  formError: FormValues | null;
  serverError: string | null;
  register: (formData: T) => Promise<void>;
}

export const useRegister = <T extends FormValues>(): UseRegisterReturn<T> => {
  const [response, setResponse] = useState<any>(undefined);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [formError, setFormError] = useState<FormValues | null>(null);
  const [serverError, setServerError] = useState<string | null>(null);

  const register = async (formData: T) => {
    setIsLoading(true);
    setServerError(null);
    setFormError(null);

    // Validation logic
    const errors: FormValues = {
      name: "",
      username: "",
      email: "",
      password: "",
    };

    if (!formData.name) {
      errors.name = "Name is required";
    }

    if (!formData.username) {
      errors.username = "Username is required";
    }

    if (!formData.email) {
      errors.email = "Email is required";
    } else if (!/\S+@\S+\.\S+/.test(formData.email)) {
      errors.email = "Email is invalid";
    }

    if (!formData.password) {
      errors.password = "Password is required";
    }

    if (Object.values(errors).some((error) => error !== "")) {
      setFormError(errors);
      setIsLoading(false);
      return;
    }

    try {
      const result = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      const data = await result.json();
      setResponse(data);
    } catch (error) {
      setServerError("Server error occurred");
    }

    setIsLoading(false);
  };

  return {
    response,
    isLoading,
    formError,
    serverError,
    register,
  };
};

export default useRegister;
