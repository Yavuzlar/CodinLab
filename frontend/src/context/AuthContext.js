import { createContext, useEffect, useState } from "react";
import { useRouter } from "next/router";
import authConfig from "src/configs/auth";
import axios from "axios";
import { showToast } from "src/utils/showToast";
import { t } from "i18next";

const defaultProvider = {
  user: null,
  loading: true,
  setUser: () => null,
  setLoading: () => Boolean,
  isInitialized: false,
  setIsInitialized: () => Boolean,
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
    setIsInitialized(false);
    window.localStorage.removeItem(authConfig.userDataName);

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

  const initAuth = async () => {
    let userData = window.localStorage.getItem(authConfig.userDataName);
    userData = userData ? JSON.parse(userData) : null;

    if (userData) {
      setUser(userData);
      if (router.pathname === "/login" || router.pathname === "/register") {
        router.replace("/");
        setLoading(false);
      }
    } else if (router.pathname !== "/login" && router.pathname !== "/register") {
      try {
        const response = await axios({
          url: authConfig.account,
          method: "GET",
        });

        if (response.status === 200) {
          const user = response?.data?.data;

          if (user && user?.role) {
            window.localStorage.setItem(
              authConfig.userDataName,
              JSON.stringify(user)
            );
            setUser(user);
            router.push("/");
          } else {
            handleLogout();
          }
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
    }

    setLoading(false);
    setIsInitialized(true);
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
        window.localStorage.setItem(
          authConfig.userDataName,
          JSON.stringify(user)
        );
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

  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
