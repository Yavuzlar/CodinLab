import SettingsPage from "src/views/settings";

const Settings = () => <SettingsPage />;

Settings.acl = {
  action: "read",
  permission: "home",
};
export default Settings;
