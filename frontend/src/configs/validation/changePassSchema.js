import * as yup from "yup";
import { validation } from "@/utils/validation";
import Translations from "src/components/Translations";

export const changePasswordValidation = async (values) => {
    const schema = yup.object().shape({
      currentPassword: yup.string().required(<Translations text="changePassword.currentPasswordError" />),
      newPassword: yup.string()
        .required(<Translations text="changePassword.newPasswordError" />)
        .min(8, <Translations text="changePassword.passwordMinError" />),
      confirmPassword: yup.string().oneOf(
        [yup.ref("newPassword"), null],
        <Translations text="changePassword.confirmPasswordError" />
      ),
    });
  
    return await validation(schema, values);
  };