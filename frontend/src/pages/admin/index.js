import Admin from "src/views/admin";

const AdminPage = () => <Admin />;

AdminPage.acl = {
  action: "read",
  permission: "admin",
};
export default AdminPage;
