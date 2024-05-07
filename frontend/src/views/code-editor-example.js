import CodeEditor from "@/components/code-editor";
import Output from "@/components/output";
import { useState } from "react";

const CodeEditorExample = () => {
  const [output, setOutput] = useState(""); // we will store the output here

  const handleRun = (outputData) => { // this function will be called when the code is run
    setOutput(outputData);
  };

  const handleStop = (outputData) => { // this function will be called when the code is stopped
    setOutput(outputData);
  }
  
  const params = { // these are the parameters for the component settings.
    height: "50vh",
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        gap: "20px",
        justifyContent: "center",
        alignItems: "center",
        padding: "20px",
      }}
    >
      <CodeEditor params={params} onRun={handleRun} onStop={handleStop} /> // CodeEditor component
      <Output value={output} params={params} /> // Output component
    </div>
  );
};

export default CodeEditorExample;
