import * as yup from "yup";
import { validation } from "@/utils/validation";

export const loginValidation = async (values) => {
    const schema = yup.object().shape({
      email: yup.string().email("Invalid email").required("Email is required"),
      password: yup.string().required("Password is required"),
    });
  
    return await validation(schema, values);
  };