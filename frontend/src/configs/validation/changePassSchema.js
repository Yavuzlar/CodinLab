import * as yup from "yup";
import { validation } from "src/utils/validation";
import Translations from "src/components/Translations";

export const changePasswordValidation = async (values) => {
  console.log(values);
    const schema = yup.object().shape({
      oldPassword: yup.string().required(<Translations text="changePassword.oldPasswordError" />),
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