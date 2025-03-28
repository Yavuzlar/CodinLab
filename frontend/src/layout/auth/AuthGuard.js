// ** React Imports
import { useEffect } from "react";
// ** Next Imports
import { useRouter } from "next/router";
// ** Hooks Import
import { useAuth } from "src/hooks/useAuth";
import authConfig from "src/configs/auth";

const AuthGuard = (props) => {
  const { children, fallback } = props;
  const auth = useAuth();
  const router = useRouter();

  useEffect(
    () => {
      if (!router.isReady) {
        return;
      }

      if (
        auth.user === null &&
        !window.localStorage.getItem(authConfig.userDataName)
      ) {
        router.replace("/login");
      }
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [router.route]
  );
  if (auth.loading || auth.user == null) {
    return fallback;
  }

  return <>{children}</>;
};

export default AuthGuard;
