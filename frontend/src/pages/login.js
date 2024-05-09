const { default: BlankLayout } = require("src/layout/BlankLayout");

const Login = () => {
  return <div>login</div>;
};

Login.guestGuard = true;
Login.getLayout = (page) => <BlankLayout>{page}</BlankLayout>;

export default Login;
