// ** React Imports
import { useEffect } from "react";
// ** Next Imports
import { useRouter } from "next/router";
// ** Hooks Import
import { useAuth } from "src/hooks/useAuth";
import authConfig from "src/configs/auth";

const GuestGuard = (props) => {
  const { children, fallback } = props;
  const auth = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (!router.isReady) {
      return;
    }

    if (window.localStorage.getItem(authConfig.userDataName) && !auth.user) {
      router.replace("/");
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [router.route]);

  if (auth.loading) {
    return fallback;
  }

  return <>{children}</>;
};

export default GuestGuard;
