import * as yup from "yup";
import { validation } from "src/utils/validation";

export const registerValidation = async (values) => {
  const schema = yup.object().shape({
    fullname: yup.string().required("Fullname is required"),
    username: yup.string().required("Username is required"),
    email: yup.string().email("Invalid email").required("Email is required"),
    password: yup
      .string()
      .required("Password is required")
      .min(8, "Password must be at least 8 characters"),
    checkbox: yup
      .string()
      .required("Please accept the policy & terms to proceed"),
  });

  return await validation(schema, values);
};
