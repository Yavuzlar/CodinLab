import { Editor } from "@monaco-editor/react";
import { Box, Typography, useMediaQuery } from "@mui/material";
import { useEffect, useRef, useState } from "react";
import StopIcon from "@mui/icons-material/Stop";
import NightsStayIcon from "@mui/icons-material/NightsStay";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import LightModeIcon from "@mui/icons-material/LightMode";
import Tooltip from "@mui/material/Tooltip";
// import PlayIcon from "src/assets/icons/play.svg";


const CodeEditor = ({ params, onRun, onStop, leng, defValue, title }) => {
  const [value, setValue] = useState("");
  const [theme, setTheme] = useState("vs-dark");
  const [editorActionsWidth, setEditorActionsWidth] = useState(0);
  const isMobile = useMediaQuery((theme) => theme.breakpoints.down("smd"));
  const editorRef = useRef(null);
  const editorActions = useRef(null);

  // here we will add the onMount function
  const onMount = (editor) => {
    editorRef.current = editor;
    editor.focus();
  };

  // here we will add the run calls
  const handleRun = () => {
    // in the future, we will add the run api call here

    setTimeout(() => {
      const response = "Output from backend";
      onRun(response);
    }, 2000);
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
          color: theme === "vs-dark" ? "white" : "black",
          borderBottom: theme === "vs-dark" ? "2px solid #DAF0FE" : "2px solid #3894D0",
          marginTop: "10px",
          paddingBottom: "10px",
          paddingLeft: "16px",
          fontSize: "18px",
          px: "26px",
          alignItems: "end",
        }}
      >
        <div
          style={{
            display: "flex",
            gap: "10px",
            lineHeight: "20px",
            width: `calc(100% - ${editorActionsWidth}px - 16px)`,
          }}
        >
          {/* hüseyin_selim_sürmelihhindi.js */}
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
            <div>
              <Tooltip title="More Options" placement="top" followCursor>
                <MoreVertIcon
                  sx={{ cursor: "pointer" }}
                  onClick={openMobileMenu}
                />
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
                    <PlayArrowIcon sx={{ cursor: "pointer" }} /> Run
                  </MenuItem>
                  <MenuItem
                    onClick={() => {
                      handleStop();
                      closeMobileMenu();
                    }}
                  >
                    <StopIcon sx={{ cursor: "pointer" }} /> Stop
                  </MenuItem>
                  <MenuItem
                    onClick={() => {
                      setTheme(theme === "vs-dark" ? "light" : "vs-dark");
                      closeMobileMenu();
                    }}
                  >
                    {theme === "vs-dark" ? (
                      <NightsStayIcon sx={{ cursor: "pointer" }} />
                    ) : (
                      <LightModeIcon sx={{ cursor: "pointer" }} />
                    )}
                    Change Theme
                  </MenuItem>
                </Menu>
              </Tooltip>
            </div>
          ) : (
            <div
              style={{
                display: "flex",
                justifyContent: "space-between",
                gap: "10px",
                color: theme === "vs-dark" ? "white" : "black",
              }}
            >
              <Tooltip title="Run" placement="top" followCursor>
                <Box>
                  <PlayArrowIcon
                    onClick={handleRun}
                    sx={{
                      cursor: "pointer",
                      width: "30px",
                      height: "30px",
                    }}
                  />
                </Box>
              </Tooltip>
              <Tooltip title="Stop" placement="top" followCursor>
                <StopIcon
                  onClick={handleStop}
                  sx={{ cursor: "pointer", width: "30px", height: "30px" }}
                />
              </Tooltip>
              {theme === "vs-dark" ? (
                <Tooltip title="Change Light Mode" placement="top" followCursor>
                  <NightsStayIcon
                    onClick={() => {
                      setTheme("light");
                    }}
                    sx={{ cursor: "pointer", width: "30px", height: "30px" }}
                  />
                </Tooltip>
              ) : (
                <Tooltip title="Change Dark Mode" placement="top" followCursor>
                  <LightModeIcon
                    onClick={() => {
                      setTheme("vs-dark");
                    }}
                    fontSize="medium"
                    sx={{ cursor: "pointer", width: "30px", height: "30px" }}
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
          // paddingBottom: "24px",
        }}
      >
        <Editor
          language={leng || "javascript"}
          defaultValue={defValue || "// Write your code here"}
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
