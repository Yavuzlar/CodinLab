import UsersPage from "src/views/admin-users";

const UsersAdmin = () => <UsersPage />;

UsersAdmin.acl = {
  action: "read",
  permission: "adminUserSettings",
};
export default UsersAdmin;
