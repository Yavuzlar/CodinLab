import { validation } from "src/utils/validation";
import * as yup from "yup";
import Translations from "src/components/Translations";

export const loginValidation = async (values) => {
  const schema = yup.object().shape({
    username: yup
      .string()
      .nullable()
      .required(<Translations text="login.usernameError" />),
    password: yup
      .string()
      .nullable()
      .required(<Translations text="login.passwordError" />),
  });

  return await validation(schema, values);
};
