// ** React Imports
import { createContext, useEffect, useState } from "react";
// ** Next Import
import { useRouter } from "next/router";
// ** Axios
import authConfig from "src/configs/auth";
import axios from "axios";
import { showToast } from "src/utils/showToast";

// ** Defaults
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
};

const AuthContext = createContext(defaultProvider);

const AuthProvider = ({ children }) => {
  // ** States
  const [user, setUser] = useState(defaultProvider.user);
  const [loading, setLoading] = useState(defaultProvider.loading);
  const [isInitialized, setIsInitialized] = useState(
    defaultProvider.isInitialized
  );

  // ** Hooks
  const router = useRouter();

  const deleteStorage = () => {
    setUser(null);
    setLoading(false);
    setIsInitialized(false);

    // const firstPath = router.pathname.split('/')[1]
    // if (firstPath != 'login') window.location.href = '/login'
  };

  const handleLogout = () => {
    deleteStorage();
  };

  const initAuth = async () => {
    setIsInitialized(true);
    const userData = window.localStorage.getItem(authConfig.userDataName);

    if (userData) {
      setUser(userData);
      console.log(userData);
      console.log(user);
    } else {
      const user = { id: 1, name: "John Doe", role: "user" };
      setUser(user);
    }

    setLoading(false);
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
      }
    } catch (error) {
      showToast("dismiss");
      showToast("error", error.message);
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
        const user = response.data;
        window.localStorage.setItem(authConfig.userDataName, JSON.stringify(user));
        setUser(user);
        router.push("/home");
      } else {
        showToast("dismiss");
        showToast("error", response.data.message);
      }
    } catch (error) {
      showToast("dismiss");
      showToast("error", error.message);
    }
  }


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
