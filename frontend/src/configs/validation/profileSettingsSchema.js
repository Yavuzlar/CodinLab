import * as yup from "yup";
import { validation } from "src/utils/validation";
import Translations from "src/components/Translations";
import { current } from "@reduxjs/toolkit";

export const profileSettingsValidation = async (values) => {
  const schema = yup.object().shape({
    username : yup.string().required(<Translations text="settings.usernameError" />),
    email: yup.string().email(<Translations text={"settings.invalidError"} />).required(<Translations text="settings.emailError" />),
  });

  return await validation(schema, values);
};
