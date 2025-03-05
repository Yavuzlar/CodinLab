let baseUrl = "/api/v1";

const authConfig = {
  account: baseUrl + "/private/user",
  login: baseUrl + "/public/login",
  logout: baseUrl + "/public/logout",
  register: baseUrl + "/public/register",
  sessionCookieName: "sessionID",

  userDataName: "userData",
};

export default authConfig;
