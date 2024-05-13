import * as yup from "yup";
import { validation } from "@/utils/validation";

export const registerValidation = async (values) => {
    const schema = yup.object().shape({
      fullname: yup.string().required("Fullname is required"),
      username: yup.string().required("Username is required"),
      email: yup.string().email("Invalid email").required("Email is required"),
      password: yup.string()
        .required("Password is required")
        .min(8, "Password must be at least 8 characters"),
    });
  
    return await validation(schema, values);
  };