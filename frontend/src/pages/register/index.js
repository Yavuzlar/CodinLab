import React from "react";
import Register from "src/views/register";

const { default: BlankLayout } = require("src/layout/BlankLayout");

const RegisterPage = () => <Register />;

RegisterPage.acl = {
  action: "read",
  permission: "home",
};

RegisterPage.getLayout = (page) => <BlankLayout> {page} </BlankLayout>;

export default RegisterPage;
