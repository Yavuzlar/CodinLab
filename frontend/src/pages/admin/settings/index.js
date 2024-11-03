import SettingsPage from "src/views/settings-admin";

const Settings = () => <SettingsPage />;

Settings.acl = {
  action: "read",
  permission: "adminSettings",
};
export default Settings;
