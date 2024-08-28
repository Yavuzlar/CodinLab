// ** Next Imports
import Head from "next/head";
import { Router } from "next/router";
// ** Store Imports
import { store } from "src/store";
import { Provider } from "react-redux";
// ** Loader Import
import NProgress from "nprogress";
// ** Config Imports
import { defaultACLObj } from "src/configs/acl";
import themeConfig from "src/configs/themeConfig";
// ** Third Party Import
import { Toaster } from "react-hot-toast";
// ** Spinner Import
import Spinner from "src/components/spinner";
// ** Contexts
import { AuthProvider } from "src/context/AuthContext";
import { NavProvider } from "src/context/NavContext";
// ** Styled Components
import ReactHotToast from "src/components/react-hot-toast";
// ** Component Imports
import Layout from "src/layout/Layout";
import GuestGuard from "src/layout/auth/GuestGuard";
import AuthGuard from "src/layout/auth/AuthGuard";
// ** Global css styles
import "../styles/main.css";
import ThemeComponent from "src/layout/ThemeComponent";
import WindowWrapper from "src/components/window-wrapper";
import AclGuard from "src/layout/auth/AclGuard";
import { appWithTranslation } from "next-i18next";
import "src/configs/i18n";

// ** Pace Loader
if (themeConfig.routingLoader) {
  Router.events.on("routeChangeStart", () => {
    NProgress.start();
  });
  Router.events.on("routeChangeError", () => {
    NProgress.done();
  });
  Router.events.on("routeChangeComplete", () => {
    NProgress.done();
  });
}

const Guard = ({ children, authGuard, guestGuard }) => {
  if (guestGuard) {
    return <GuestGuard fallback={<Spinner />}>{children}</GuestGuard>;
  } else if (!guestGuard && !authGuard) {
    return <>{children}</>;
  } else {
    return <AuthGuard fallback={<Spinner />}>{children}</AuthGuard>;
  }
};

// ** Configure JSS & ClassName
const App = (props) => {
  const { Component, pageProps } = props;

  // Variables
  const getLayout = Component.getLayout ?? ((page) => <Layout>{page}</Layout>);
  const authGuard = Component.authGuard ?? true;
  const guestGuard = Component.guestGuard ?? false;
  const aclAbilities = Component.acl ?? defaultACLObj;

  return (
    <Provider store={store}>
      <Head>
        <title>{`${themeConfig.projectName}`}</title>
        <meta name="description" content={`${themeConfig.projectName}`} />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link
          href="https://fonts.cdnfonts.com/css/outfit"
          rel="stylesheet"
        ></link>
        <link
          href="https://fonts.cdnfonts.com/css/cascadia-code"
          rel="stylesheet"
        ></link>
        <link
          href="https://fonts.cdnfonts.com/css/nunito"
          rel="stylesheet"
        ></link>
      </Head>

      <AuthProvider>
        <NavProvider>
        <ThemeComponent>
          <WindowWrapper>
            <Guard authGuard={authGuard} guestGuard={guestGuard}>
              <AclGuard aclAbilities={aclAbilities} guestGuard={guestGuard}>
                {getLayout(<Component {...pageProps} />)}
              </AclGuard>
            </Guard>
          </WindowWrapper>

          <ReactHotToast>
            <Toaster
              position={themeConfig.toastPosition}
              toastOptions={{ className: "react-hot-toast" }}
            />
          </ReactHotToast>
        </ThemeComponent>
        </NavProvider>
      </AuthProvider>
    </Provider>
  );
};

export default appWithTranslation(App);
