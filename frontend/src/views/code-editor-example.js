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
    width: "50vw",
  };

  

  return (
    <div
      style={{
        display : "flex",
      }}
    >
      <CodeEditor params={params} onRun={handleRun} onStop={handleStop} leng={"javascript"} defValue={"//deneme"} /> 
      <Output value={output} params={params} /> 
    </div>
  );
};

export default CodeEditorExample;
