import { validation } from "src/utils/validation";
import * as yup from "yup";
import Translations from "src/components/Translations";


export const loginValidation = async (values) => {
    const schema = yup.object().shape({
      email: yup.string().email(<Translations text="login.invalidError" />).nullable().required(<Translations text="login.emailError" />),
      password: yup.string().nullable().required(<Translations text="login.passwordError" />),
    });
  
    return await validation(schema, values);
  };