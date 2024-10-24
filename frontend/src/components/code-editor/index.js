import { Editor } from "@monaco-editor/react";
import { Box, Typography, useMediaQuery } from "@mui/material";
import { useEffect, useRef, useState } from "react";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import Tooltip from "@mui/material/Tooltip";
import Image from "next/image";
import PlayIconBlack from "src/assets/icons/play-black.png";
import PlayIconWhite from "src/assets/icons/play-white.png";
import StopIconWhite from "src/assets/icons/stop-white.png";
import StopIconBlack from "src/assets/icons/stop-black.png";
import SunIcon from "src/assets/icons/sun.png";
import MoonIcon from "src/assets/icons/moon.png";
import MenuIconBlack from "src/assets/icons/menu-black.png";
import MenuIconWhite from "src/assets/icons/menu-white.png";
import axios from "axios";
import { t } from "i18next";
import { showToast } from "src/utils/showToast";
import toast from "react-hot-toast";

const CodeEditor = ({
  params,
  onRun,
  onStop,
  leng,
  defValue,
  title,
  apiData,
  val,
  editorRef
}) => {
  const [value, setValue] = useState(null);
  const [defaultValue, setDefaultValue] = useState(defValue);
  const [theme, setTheme] = useState("vs-dark");
  const [editorActionsWidth, setEditorActionsWidth] = useState(0);
  const isMobile = useMediaQuery((theme) => theme.breakpoints.down("smd"));
  const editorActions = useRef(null);

  // here we will add the onMount function
  const onMount = (editor) => {
    editorRef.current = editor;
    editor.focus();
  };

  // here we will add the run calls
  const handleRun = async () => {
    // const toastId = toast.loading(t("runCode")); 
    showToast("dismiss")
    showToast("loading", t("runCode"));

    try {
      const response = await axios({
        method: "POST",
        url: `/api/v1/private/${apiData.endPoint}/answer/${apiData.programmingId}/${apiData.pathId}`,
        data: { userCode: value },
        headers: {
          "Content-Type": "application/json",
        },
      });
      showToast("dismiss")
      showToast("success", t("succsesCode"));
      onRun(response.data);
    } catch (error) {
      toast.error(error.response?.data?.message || error.message, { id: 2 });
      onRun(error.response?.data?.message || error.message);
    }
  };


  // here we will add the stop api calls
  const handleStop = () => {
    // in the future, we will ad the stop api call here

    setTimeout(() => {
      const response = "Stopped from backend";
      onStop(response);
    }, 2000);
  };

  // for mobile menu options
  const [mobileMenuAnchor, setMobileMenuAnchor] = useState(null);
  const openMobileMenu = (event) => {
    setMobileMenuAnchor(event.currentTarget);
  };
  const closeMobileMenu = () => {
    setMobileMenuAnchor(null);
  };

  useEffect(() => {
    if (editorActions.current) {
      setEditorActionsWidth(editorActions.current.offsetWidth ?? 0);
    }
  }, [editorActions?.current?.offsetWidth]);

  useEffect(() => {
    setValue(val);
  }, [val]);

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        border: theme === "vs-dark" ? "2px solid #DAF0FE" : "2px solid #3894D0",
        borderRadius: "30px",
        opacity: "1",
        backgroundColor: theme === "vs-dark" ? "#1E1E1E" : "white",
        color: theme === "vs-dark" ? "white" : "black",
        height: params.height || "auto",
        width: params.width || "auto",
      }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          color: theme === "vs-dark" ? "white" : "black",
          borderBottom:
            theme === "vs-dark" ? "2px solid #DAF0FE" : "2px solid #3894D0",
          marginTop: "12px",
          paddingBottom: "10px",
          fontSize: "18px",
          px: "26px",
          fontWeight: "bold",
          letterSpacing: "0px",
          marginBottom: "0",
        }}
      >
        <div
          style={{
            display: "flex",
            gap: "10px",
            width: `calc(100% - ${editorActionsWidth}px - 4px)`,
          }}
        >
          <Tooltip title={title || "Untitled"} placement="top" followCursor>
            <Typography
              variant="span"
              sx={{
                display: "block",
                width: "100%",
                overflow: "hidden",
                textOverflow: "ellipsis",
                whiteSpace: "nowrap",
                maxWidth: "fit-content",
                letterSpacing: "0px",
                color: theme === "vs-dark" ? "white" : "black",
                fontWeight: "600px",
                cursor: "default",
              }}
            >
              {title || "Untitled"}
            </Typography>
          </Tooltip>
        </div>
        <Box ref={editorActions}>
          {isMobile ? (
            <div
              style={{
                display: "flex",
                justifyContent: "space-between",
                gap: "10px",
                color: theme === "vs-dark" ? "white" : "black",
                alignItems: "center",
                position: "relative",
              }}
            >
              <Tooltip title="More Options" placement="top" followCursor>
                <Image
                  src={theme === "vs-dark" ? MenuIconWhite : MenuIconBlack}
                  alt="My SVG"
                  width={30}
                  height={30}
                  onClick={openMobileMenu}
                />
              </Tooltip>
              <Menu
                anchorEl={mobileMenuAnchor}
                open={Boolean(mobileMenuAnchor)}
                onClose={closeMobileMenu}
              >
                <MenuItem
                  onClick={() => {
                    handleRun();
                    closeMobileMenu();
                  }}
                >
                  <Image
                    src={theme === "vs-dark" ? PlayIconWhite : PlayIconBlack}
                    alt="My SVG"
                    width={30}
                    height={30}
                    sx={{ cursor: "pointer" }}
                  />
                  <Typography variant="span" sx={{ ml: 1 }}>
                    Run
                  </Typography>
                </MenuItem>
                <MenuItem
                  onClick={() => {
                    handleStop();
                    closeMobileMenu();
                  }}
                >
                  <Image
                    src={theme === "vs-dark" ? StopIconWhite : StopIconBlack}
                    alt="My SVG"
                    width={30}
                    height={30}
                    sx={{ cursor: "pointer" }}
                  />
                  <Typography variant="span" sx={{ ml: 1 }}>
                    Stop
                  </Typography>
                </MenuItem>
                <MenuItem
                  onClick={() => {
                    setTheme(theme === "vs-dark" ? "light" : "vs-dark");
                    closeMobileMenu();
                  }}
                >
                  <Image
                    src={theme === "vs-dark" ? SunIcon : MoonIcon}
                    alt="My SVG"
                    width={30}
                    height={30}
                    sx={{ cursor: "pointer" }}
                  />

                  <Typography variant="span" sx={{ ml: 1 }}>
                    Change Theme
                  </Typography>
                </MenuItem>
              </Menu>
            </div>
          ) : (
            <div
              style={{
                display: "flex",
                justifyContent: "space-between",
                gap: "10px",
                color: theme === "vs-dark" ? "white" : "black",
                alignItems: "center",
                position: "relative",
              }}
            >
              <Tooltip title="Run" placement="top" followCursor>
                <Image
                  src={theme === "vs-dark" ? PlayIconWhite : PlayIconBlack}
                  alt="My SVG"
                  width={30}
                  height={30}
                  style={{ cursor: "pointer" }}
                  onClick={handleRun}
                />
              </Tooltip>
              <Tooltip title="Stop" placement="top" followCursor>
                <Image
                  src={theme === "vs-dark" ? StopIconWhite : StopIconBlack}
                  alt="My SVG"
                  width={30}
                  height={30}
                  style={{ cursor: "pointer" }}
                  onClick={handleStop}
                />
              </Tooltip>
              {theme === "vs-dark" ? (
                <Tooltip title="Change Light Mode" placement="top" followCursor>
                  <Image
                    src={SunIcon}
                    alt="My SVG"
                    width={30}
                    height={30}
                    onClick={() => {
                      setTheme("light");
                    }}
                    sx={{ cursor: "pointer" }}
                  />
                </Tooltip>
              ) : (
                <Tooltip title="Change Dark Mode" placement="top" followCursor>
                  <Image
                    src={MoonIcon}
                    alt="My SVG"
                    width={30}
                    height={30}
                    onClick={() => {
                      setTheme("vs-dark");
                    }}
                    sx={{ cursor: "pointer" }}
                  />
                </Tooltip>
              )}
            </div>
          )}
        </Box>
      </Box>
      <div
        style={{
          width: "100%",
          height: "100%",
          overflow: "hidden",
          borderRadius: "0px 0px 30px 30px",
          paddingTop: "10px",
          paddingBottom: "24px",
        }}
      >
        <Editor
          language={leng || "javascript"}
          // defaultValue={defaultValue || "// Write your code here"}
          defaultValue={"// Write your code here"}
          value={value}
          onChange={(newValue) => setValue(newValue)}
          onMount={onMount}
          theme={theme}
          loading={<div>Loading...</div>}
          options={{
            minimap: {
              enabled: false,
            },
          }}
        />
      </div>
    </Box>
  );
};

export default CodeEditor;