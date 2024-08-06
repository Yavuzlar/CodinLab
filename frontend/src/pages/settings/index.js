import SettingsPage from "src/views/settings-user";

const SettingsUsers = () => <SettingsPage />;

SettingsUsers.acl = {
  action: "read",
  permission: "settings",
};
export default SettingsUsers;
