import { createContext, useEffect, useState, useRef } from "react";
import { useRouter } from "next/router";
import authConfig from "src/configs/auth";
import axios from "axios";
import { showToast } from "src/utils/showToast";
import { t } from "i18next";
// ** Spinner Import
import Spinner from "src/components/spinner";

const defaultProvider = {
  user: null,
  loading: false,
  isInitialized: false,
  containerLoading: false,
  setUser: () => null,
  setLoading: () => Boolean,
  logout: () => Promise.resolve(),
  register: () => Promise.resolve(),
  initAuth: () => Promise.resolve(),
  login: () => Promise.resolve(),
  setContainerLoading: () => Promise.resolve(),
};

const AuthContext = createContext(defaultProvider);

const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(defaultProvider.user);
  const [loading, setLoading] = useState(defaultProvider.loading);
  const [isInitialized, setIsInitialized] = useState(
    defaultProvider.isInitialized
  );
  const [containerLoading, setContainerLoading] = useState(
    defaultProvider.containerLoading
  );

  const ws = useRef(null);

  const router = useRouter();

  const webSocket = () => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      return;
    }
    ws.current = new WebSocket("ws://localhost/api/v1/private/socket/ws");

    ws.current.onopen = () => { };

    ws.current.onmessage = (e) => {
      const data = JSON.parse(e.data);

      if (data.Type === "Pull") {
        const message = data.Data.message;
        const programmingLanguage =
          data.Data.programminglanguage || "Unknown Language";

        const downloadedMessage = t("code.image.downloaded");
        const downloadingMessage = t("code.image.downloading");

        // $$$$'ı programmingLanguage ile değiştiriyoruz
        const formattedMessage = downloadedMessage.replace(
          "$$$$",
          programmingLanguage
        );
        const downloadingFormattedMessage = downloadingMessage.replace(
          "$$$$",
          programmingLanguage
        );

        if (message === "Started") {
          setContainerLoading(true);
          showToast("dismiss");
          showToast("loading", downloadingFormattedMessage);
        } else if (message === "Finished") {
          setContainerLoading(false);
          showToast("dismiss");
          showToast("success", formattedMessage);
        }
      }
    };

    ws.current.onclose = () => {
      setTimeout(webSocket, 5000);
    };

    ws.current.onerror = (error) => {
      console.error("WebSocket error observed:", error);
    };
  };

  const closeWebSocket = () => {
    if (ws.current) {
      ws.current.close();
      ws.current = null;
    }
  };

  const sendHistory = (userCode, programmingID, labPathID, labPathType) => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      const historyData = {
        type: "closed",
        data: {
          userCode: userCode,
          programmingID: programmingID,
          labPathID: labPathID,
          labPathType: labPathType,
        },
      };
      ws.current.send(JSON.stringify(historyData));
    } else {
      console.error("WebSocket is not connected");
    }
  };

  const createSession = (data) => {
    setUser(data); // Set the user data to the state
    webSocket();
    localStorage.setItem(authConfig.userDataName, JSON.stringify(data)); // Set the user data to the local storage
  };

  const restoreStorage = () => {
    setLoading(false);
    setUser(defaultProvider.user);
    localStorage.removeItem(authConfig.userDataName);
  };

  const handleLogout = async () => {
    try {
      const response = await axios({
        url: authConfig.logout,
        method: "POST",
      });
      if (response.status === 200) {
        restoreStorage();

        router.push("/login");
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
    const storedUserData = localStorage.getItem(authConfig.userDataName);
    
    if (["/login", "/register"].includes(router.pathname)) {
      if (storedUserData) {
        try {
          const userData = JSON.parse(storedUserData);
          setUser(userData);
          setLoading(false);
          setIsInitialized(true);
          router.push("/");
        } catch (error) {
          localStorage.removeItem(authConfig.userDataName);
          setLoading(false);
          setIsInitialized(true);
        }
      } else {
        setLoading(false);
        setIsInitialized(true);
      }
      return;
    }

    if (!storedUserData) {
      setLoading(false);
      setIsInitialized(true);
      return;
    }

    setLoading(true);
    setIsInitialized(false);

    axios({
      url: authConfig.account,
      method: "GET",
    })
      .then(async (response) => {
        if (response.status === 200) {
          const user = response?.data?.data;

          if (user && user?.role) {
            createSession(user); // Create a session for the user
            setLoading(false); // Set loading to false
            setIsInitialized(true);

            if (["/login", "/register"].includes(router.pathname)) {
              router.push("/");
            }
          } else {
            setLoading(false);
            handleLogout();
          }
        } else {
          setLoading(false);
          handleLogout();
        }
      })
      .catch((error) => {
        setLoading(false);
        showToast("dismiss");
        showToast("error", t(error?.response?.data?.message ?? ""));
        handleLogout();
      });
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

        createSession(user); // Create a session for the user
        router.push("/");
      } else {
        showToast("dismiss");
        showToast("error", response.data.message);
        restoreStorage(); // Delete the user data
      }
    } catch (error) {
      showToast("dismiss");
      showToast("error", t(error.response.data.message));
      restoreStorage(); // Delete the user data
    }
  };

  useEffect(() => {
    initAuth();
  }, [router.pathname]);

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
    containerLoading,
    sendHistory,
  };

  if (!isInitialized && loading) return <Spinner />;
  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
