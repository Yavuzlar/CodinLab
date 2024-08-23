let baseUrl = "/api"

const authConfig = {
  account: baseUrl + '/session',
  login: baseUrl + '/v1/public/login',
  logout: baseUrl + '/session',
  register: baseUrl + '/v1/public/register',
  sessionCookieName: 'sessionID',

  userDataName: 'userData',
};

export default authConfig;
