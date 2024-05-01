let baseUrl = process.env.NEXT_PUBLIC_BASE_URL

const authConfig = {
  account: baseUrl + '/session',
  login: baseUrl + '/session',
  logout: baseUrl + '/session',
  register: baseUrl + '/sign-up',
  sessionCookieName: 'sessionID',

  userDataName: 'userData',
};

export default authConfig;
