import SettingsPage from "src/views/settings";

const Settings = () => <SettingsPage />;

Settings.acl = {
  action: "read",
  permission: "settings",
};
export default Settings;
