import Home from "src/views/home";

const HomePage = () => <Home />;

HomePage.acl = {
  action: "read",
  permission: "home",
};
export default HomePage;
