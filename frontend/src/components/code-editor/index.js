import { Editor } from "@monaco-editor/react";
import { Box, Divider } from "@mui/material";
import { useRef, useState } from "react";
import StopIcon from "@mui/icons-material/Stop";
import NightsStayIcon from "@mui/icons-material/NightsStay";
import Brightness5Icon from "@mui/icons-material/Brightness5";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";

const CodeEditor = ({ params, onRun,onStop }) => {
  const [value, setValue] = useState("");
  const [theme, setTheme] = useState("vs-dark");
  const [language, setLanguage] = useState("javascript");
  const [output, setOutput] = useState("");

  const editorRef = useRef(null);

  const blackTheme = {
    base: "vs-dark",
    inherit: true,
    rules: [{ background: "#000000" }],
    colors: {
      "editor.background": "#000000",
    },
  };

  const onMount = (editor) => {
    editorRef.current = editor;
    editor.focus();
  };

  const handleRun = () => {
    // in this function, you can run the code and set the output
    // we will add api call to stop the code execution
    setOutput("Running..."); 
    setTimeout(() => {
      const response = "Output from backend"; 
      setOutput(response);
      onRun(response); 
    }, 2000);
  };

  const handleStop = () => {
    // in this function, you can stop the running code and set the output.
    // we will add api call to stop the code execution
    setOutput("Stopped");
    setTimeout(() => {
      const response = "Stopped from backend"; 
      setOutput(response);
      onStop(response); 
    }, 2000);
  }

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
          height={params.height || "80vh"}
          width={params.width || "100%"}
          defaultLanguage={language}
          defaultValue="// Write your code here"
          value={value}
          onChange={(newValue) => setValue(newValue)}
          onMount={onMount}
          theme={theme === "vs-dark" ? theme : blackTheme}
          loading={<div>Loading...</div>}
        />
      </div>
    </Box>
  );
};

export default CodeEditor;
