import { Editor } from "@monaco-editor/react";
import { Box, useMediaQuery } from "@mui/material";
import { useRef, useState } from "react";
import StopIcon from "@mui/icons-material/Stop";
import NightsStayIcon from "@mui/icons-material/NightsStay";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import LightModeIcon from '@mui/icons-material/LightMode';

const CodeEditor = ({ params, onRun, onStop, leng, defValue }) => {
  const [value, setValue] = useState("");
  const [theme, setTheme] = useState("vs-dark");
  const isMobile = useMediaQuery((theme) => theme.breakpoints.down("md"));
  const editorRef = useRef(null);

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


  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        padding: "10px",
        border: theme === "vs-dark" ? "2px solid #1E1E1E" : "2px solid #3894D0",
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
          fontWeight: "bold",
          alignItems: "center",
          borderBottom: theme === "vs-dark" ? "2px solid #ffff" : "2px solid #3894D0",
          marginTop: "10px",
          paddingBottom: "10px",
          paddingLeft: "16px",
          fontSize: "18px",
          fontWeight: "bold",
        }}
      >
        <div
          style={{
            display:"-webkit-box",
            WebkitBoxOrient: "vertical",
            WebkitLineClamp: 1,
            gap: "10px",
            maxWidth: "220px",
            overflow: "hidden",
            textOverflow: "ellipsis",
            lineHeight: "20px",
          }}
        >
          hüseyin_selim_sürmelihhindi.js
        </div>
        {isMobile ? (
          <div>
            <MoreVertIcon sx={{ cursor: "pointer" }} onClick={openMobileMenu} />
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
            <div>
              <PlayArrowIcon onClick={handleRun} sx={{ cursor: "pointer" }} fontSize="medium" />
            </div>
            <div>
              <StopIcon fontSize="medium" onClick={handleStop} sx={{ cursor: "pointer" }} />
            </div>
            <div>
              {theme === "vs-dark" ? (
                <NightsStayIcon
                  onClick={() => {
                    setTheme("light");
                  }}
                  fontSize="medium"
                  sx={{ cursor: "pointer" }}
                />
              ) : (
                <LightModeIcon
                  onClick={() => {
                    setTheme("vs-dark");
                  }}
                  fontSize="medium"
                  sx={{ cursor: "pointer" }}
                />
              )}
            </div>
          </div>
        )}
      </Box>
      <div
        style={{
          width: "100%",
          height: "100%",
          overflow: "hidden",
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
