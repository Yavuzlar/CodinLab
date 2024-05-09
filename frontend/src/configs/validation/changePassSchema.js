import * as yup from "yup";
import { validation } from "@/utils/validation";

export const changePasswordValidation = async (values) => {
    const schema = yup.object().shape({
      currentPassword: yup.string().required("Current password is required"),
      newPassword: yup.string()
        .required("New password is required")
        .min(8, "Password must be at least 8 characters"),
      confirmPassword: yup.string().oneOf(
        [yup.ref("newPassword"), null],
        "Passwords must match"
      ),
    });
  
    return await validation(schema, values);
  };