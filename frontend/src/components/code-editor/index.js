import { Editor } from "@monaco-editor/react";
import { Box, Divider, useMediaQuery } from "@mui/material";
import { useRef, useState } from "react";
import StopIcon from "@mui/icons-material/Stop";
import NightsStayIcon from "@mui/icons-material/NightsStay";
import Brightness5Icon from "@mui/icons-material/Brightness5";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";

const CodeEditor = ({ params, onRun, onStop }) => {
  const [value, setValue] = useState("");
  const [theme, setTheme] = useState("vs-dark");
  const [language, setLanguage] = useState("javascript");
  const [output, setOutput] = useState("");
  const isMobile = useMediaQuery((theme) => theme.breakpoints.down("md"));
  const editorRef = useRef(null);
  const blackTheme = {
    base: "vs-dark",
    inherit: true,
    rules: [{ background: "#000000" }],
    colors: {
      "editor.background": "#000000",
    },
  };

  // here we will add the onMount function
  const onMount = (editor) => {
    editorRef.current = editor;
    editor.focus();
  };

  // here we will add the run calls
  const handleRun = () => {
    // in the future, we will add the run api call here

    setOutput("Running...");
    setTimeout(() => {
      const response = "Output from backend";
      setOutput(response);
      onRun(response);
    }, 2000);
  };

  // here we will add the stop api calls
  const handleStop = () => {
    // in the future, we will ad the stop api call here

    setOutput("Stopped");
    setTimeout(() => {
      const response = "Stopped from backend";
      setOutput(response);
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
        width: "50%",
        padding: "10px",
        border: theme === "vs-dark" ? "2px solid #1E1E1E" : "2px solid #3894D0",
        borderRadius: "30px",
        opacity: "1",
        backgroundColor: theme === "vs-dark" ? "#1E1E1E" : "white",
        color: theme === "vs-dark" ? "white" : "black",
      }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          color: theme === "vs-dark" ? "white" : "black",
          fontSize: "15px",
          fontWeight: "bold",
          alignItems: "center",
          paddingLeft: "10px",
          paddingRight: "10px",
        }}
      >
        <div
          sx={{
            display: "flex",
          }}
        >
          <div
            style={{
              display: "flex",
              gap: "10px",
            }}
          >
            Project Name
          </div>
        </div>
        {isMobile ? (
          <div>
            <MoreVertIcon onClick={openMobileMenu} />
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
                <PlayArrowIcon /> Run
              </MenuItem>
              <MenuItem
                onClick={() => {
                  handleStop();
                  closeMobileMenu();
                }}
              >
                <StopIcon /> Stop
              </MenuItem>
              <MenuItem
                onClick={() => {
                  setTheme(theme === "vs-dark" ? "light" : "vs-dark");
                  closeMobileMenu();
                }}
              >
                {theme === "vs-dark" ? (
                  <NightsStayIcon />
                ) : (
                  <Brightness5Icon />
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
              <PlayArrowIcon onClick={handleRun} fontSize="medium" />
            </div>
            <div>
              <StopIcon fontSize="medium" onClick={handleStop} />
            </div>
            <div>
              {theme === "vs-dark" ? (
                <NightsStayIcon
                  onClick={() => {
                    setTheme("light");
                  }}
                  fontSize="medium"
                />
              ) : (
                <Brightness5Icon
                  onClick={() => {
                    setTheme("vs-dark");
                  }}
                  fontSize="medium"
                />
              )}
            </div>
          </div>
        )}
      </Box>
      <Divider
        sx={{
          width: "103%",
          height: "2px",
          backgroundColor: theme === "vs-dark" ? "#1E1E1E" : "#3894D0",
          marginLeft: "-10px",
        }}
      />
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "10px",
        }}
      >
        <Editor
          height={params.height || "80vh"} // By default, it fully fits with its parent
          width={params.width || "100%"} // By default, it fully fits with its parent
          defaultLanguage={language} // The language of the editor we will take from the backend
          defaultValue="// Write your code here"
          value={value} // The value of the editor
          onChange={(newValue) => setValue(newValue)} // The change event of the editor
          onMount={onMount} // The mount event of the editor
          theme={theme === "vs-dark" ? theme : blackTheme} // The theme of the editor
          loading={<div>Loading...</div>} // The loading component+
        />
      </div>
    </Box>
  );
};

export default CodeEditor;
