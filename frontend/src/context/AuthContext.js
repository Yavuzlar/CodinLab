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
  loading: true,
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
      console.log("WebSocket is already connected");
      return;
    }
    console.log("Bağlanılıyor...");
    ws.current = new WebSocket("ws://localhost/api/v1/private/socket/ws");

    ws.current.onopen = () => {
      console.log("Connected to WebSocket");
    };

    ws.current.onmessage = (e) => {
      const data = JSON.parse(e.data);
      console.log("Message from server:", data);

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

      // this part is for get container id from websocket but not used in this project
      // const data = JSON.parse(e.data); //
      // if (data.Type === "container") {
      //   const containerId = data?.Data?.id;

      //   if (containerId) {
      //     localStorage.setItem('containerId', containerId);
      //   }
      // }
    };

    ws.current.onclose = () => {
      console.log("WebSocket disconnected. Attempting to reconnect...");
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
      console.log("WebSocket bağlantısı kapatıldı");
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
      console.log("History data sent to server:", historyData);
    } else {
      console.error("WebSocket is not connected");
    }
  };

  const deleteStorage = () => {
    setUser(null);
    setLoading(false);
    closeWebSocket();
    const firstPath = router.pathname.split("/")[1];
    if (firstPath !== "login" && firstPath !== "register") {
      router.replace("/login");
    }
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
    setIsInitialized(false);

    axios({
      url: authConfig.account,
      method: "GET",
    })
      .then(async (response) => {
        if (response.status === 200) {
          const user = response?.data?.data;

          if (user && user?.role) {
            setIsInitialized(true);
            setUser(user);

            if (
              router.pathname === "/login" ||
              router.pathname === "/register"
            ) {
              router.push("/").then(() => router.reload());
            } else {
              setLoading(false);
              webSocket();
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
        setUser(user);
        router.push("/home");
        webSocket();
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
    containerLoading,
    sendHistory,
  };

  if (!isInitialized && loading) return <Spinner />;
  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
