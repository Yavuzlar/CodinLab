import Register from "src/views/register";

const { default: BlankLayout } = require("src/layout/BlankLayout");

const RegisterPage = () => <Register />;

RegisterPage.guestGuard = true;
RegisterPage.getLayout = (page) => <BlankLayout> {page} </BlankLayout>;

export default RegisterPage;
