import { createContext, useEffect, useState } from "react";
import { useRouter } from "next/router";
import authConfig from "src/configs/auth";
import axios from "axios";
import { showToast } from "src/utils/showToast";
import { t } from "i18next";
// ** Spinner Import
import Spinner from "src/components/spinner";

const defaultProvider = {
  user: null,
  loading: true,
  isInitialized: false,
  setUser: () => null,
  setLoading: () => Boolean,
  logout: () => Promise.resolve(),
  register: () => Promise.resolve(),
  initAuth: () => Promise.resolve(),
  login: () => Promise.resolve(),
};

const AuthContext = createContext(defaultProvider);

const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(defaultProvider.user);
  const [loading, setLoading] = useState(defaultProvider.loading);
  const [isInitialized, setIsInitialized] = useState(
    defaultProvider.isInitialized
  );

  const router = useRouter();

  const deleteStorage = () => {
    setUser(null);
    setLoading(false);
    // window.localStorage.removeItem(authConfig.userDataName);

    const firstPath = router.pathname.split("/")[1];
    if (firstPath !== "login") router.replace("/login");
  };

  const handleLogout = async () => {
    try {
      const response = await axios({
        url: authConfig.logout,
        method: "POST",
      });
      if (response.status === 200) {
        deleteStorage();
      } else {
        showToast("dismiss");
        showToast("error", response.data.message);
      }
    } catch (error) {
      showToast("dismiss");
      showToast("error", t(error.response.data.message));
    }
  };

  const initAuth = () => {
    setLoading(true);
    setIsInitialized(false)

    axios({
      url: authConfig.account,
      method: "GET",
    })
      .then(async (response) => {
        if (response.status === 200) {
          const user = response?.data?.data;

          if (user && user?.role) {
            setIsInitialized(true)
            setUser(user);

            if (router.pathname == "/login" || router.pathname == "/register") {
              router.push("/").then(() => router.reload())
            } else {
              setLoading(false);
            }

          } else {
            setLoading(false);
            handleLogout();
          }
        } else {
          setLoading(false);
          showToast("dismiss");
          showToast("error", response.data.message);
          handleLogout();
        }
      })
      .catch((error) => {
        setLoading(false);
        showToast("dismiss");
        showToast("error", t(error?.response?.data?.message ?? ""));
        handleLogout();
      })
  };

  const handleRegister = async (formData) => {
    try {
      const response = await axios({
        url: authConfig.register,
        method: "POST",
        data: formData,
      });
      if (response.status === 200) {
        showToast("dismiss");
        showToast("success", "Account created successfully");
        router.push("/login");
      } else {
        showToast("dismiss");
        showToast("error", response.data.message);
        handleLogout();
      }
    } catch (error) {
      showToast("dismiss");
      showToast("error", t(error.response.data.message));
      handleLogout();
    }
  };

  const handleLogin = async (formData) => {
    try {
      const response = await axios({
        url: authConfig.login,
        method: "POST",
        data: formData,
      });
      if (response.status === 200) {
        const user = response?.data?.data;
        // window.localStorage.setItem(
        //   authConfig.userDataName,
        //   JSON.stringify(user)
        // );
        setUser(user);
        router.push("/home");
      } else {
        showToast("dismiss");
        showToast("error", response.data.message);
      }
    } catch (error) {
      showToast("dismiss");
      showToast("error", t(error.response.data.message));
    }
  };

  useEffect(() => {

    initAuth();
  }, []);

  const values = {
    user,
    loading,
    setUser,
    setLoading,
    isInitialized,
    setIsInitialized,
    logout: handleLogout,
    register: handleRegister,
    initAuth,
    login: handleLogin,
  };

  if (!isInitialized && loading) return <Spinner />;
  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
