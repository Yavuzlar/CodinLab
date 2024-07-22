import * as yup from "yup";
import { validation } from "src/utils/validation";
import Translations from "src/components/Translations";

export const registerValidation = async (values) => {
  const schema = yup.object().shape({
    name: yup.string().required(<Translations text="register.nameError" />),
    surname: yup
      .string()
      .required(<Translations text="register.surnameError" />),
    username: yup
      .string()
      .required(<Translations text="register.usernameError" />),
    githubProfile: yup
      .string()
      .required(<Translations text="register.githubProfileError" />),
    email: yup
      .string()
      .email(<Translations text="register.emailError" />)
      .required(<Translations text="register.emailRequiredError" />),
    password: yup
      .string()
      .required(<Translations text="register.passwordError" />)
      .min(8, <Translations text="register.passwordMinError" />),
    checkbox: yup
      .string()
      .required(<Translations text="register.checkboxError" />),
  });

  return await validation(schema, values);
};
