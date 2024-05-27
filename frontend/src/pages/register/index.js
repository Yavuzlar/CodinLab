import React from "react";
import Register from "src/views/register";

const RegisterPage = () => <Register />;

RegisterPage.acl = {
  action: "read",
  permission: "home",
};

export default RegisterPage;
