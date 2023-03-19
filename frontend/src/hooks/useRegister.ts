import { useState, useEffect } from "react";
import useFetch from "../hooks/useFetch";

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
  const [formError, setFormError] = useState<FormValues | null>(null);
  const [serverError, setServerError] = useState<string | null>(null);

  const { data, isLoading, error, send } = useFetch({ endpoint: `/register` });

  useEffect(() => {
    if (data) {
      setResponse(data);
    }
  }, [data]);

  useEffect(() => {
    if (error) {
      error.message && setServerError(error.message);
      setFormError(error.error);
    }
  }, [error]);

  const register = async (formData: T) => {
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
    } else {
      if (formData.password.length < 8) {
        errors.password = "Password must be at least 8 characters long";
      }
      if (!/[A-Z]+/.test(formData.password)) {
        errors.password = "Password must contain at least one uppercase letter";
      }
      if (!/[a-z]+/.test(formData.password)) {
        errors.password = "Password must contain at least one lowercase letter";
      }
      if (!/\d+/.test(formData.password)) {
        errors.password = "Password must contain at least one digit";
      }
      if (!/[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]+/.test(formData.password)) {
        errors.password =
          "Password must contain at least one special character";
      }
    }

    if (Object.values(errors).some((error) => error !== "")) {
      setFormError(errors);
      return;
    }

    await send(formData, "POST");
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
