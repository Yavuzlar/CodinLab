import Output from "src/components/output";
import { useState } from "react";
import CodeEditor from "src/components/code-editor";

const CodeEditorExample = () => {
  const [output, setOutput] = useState("");

  const handleRun = (outputData) => {
    setOutput(outputData);
  };

  const handleStop = (outputData) => {
    setOutput(outputData);
  };

  const params = {
    height: "60vh",
    width: "60vw",
  };

  return (
    <div
      style={{
        display: "flex",
      }}
    >
      <CodeEditor
        params={params}
        onRun={handleRun}
        onStop={handleStop}
        leng={"javascript"}
        defValue={"//deneme"}
        title={"deneme.js"}
      />
      <Output value={output} params={params} />
    </div>
  );
};

export default CodeEditorExample;
