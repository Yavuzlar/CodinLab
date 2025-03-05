import * as yup from "yup";
import { validation } from "src/utils/validation";
import Translations from "src/components/Translations";

export const profileSettingsValidation = async (values) => {
  const schema = yup.object().shape({
    username: yup
      .string()
      .required(<Translations text="settings.usernameError" />),
    name: yup.string().required(<Translations text="settings.nameError" />),
    surname: yup
      .string()
      .required(<Translations text="settings.surnameError" />),
    githubProfile: yup
      .string()
      .required(<Translations text="settings.githubError" />),
  });

  return await validation(schema, values);
};
