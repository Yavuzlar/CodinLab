import { Box } from "@mui/material";
import { useTranslation } from "react-i18next";

const Output = ({ value, params }) => {

  const { t } = useTranslation();

  value =
    value ||
    t("codeEditor.output.display");

  // const width = params.width ? params.width : "50%";
  // const height = params.height ? params.height : "50%";

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        border: "2px solid white",
        borderRadius: "30px",
        opacity: "1",
        backgroundColor: "#1E1E1E",
        color: "white",
        height: params.height || "auto",
        width: params.width || "auto",
        
      }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          color: "white",
          borderBottom: "2px solid #DAF0FE",
          marginTop: "19px",
          paddingBottom: "10px",
          fontSize: "18px",
          px: "26px",
          alignItems: "end",
          fontWeight: "bold",
        }}
      >
        <Box
          style={{
            textAlign: "left",
            letterSpacing: "0px",
            color: "#FFFFFF",
            opacity: "1",
          }}
        >
          {t("codeEditor.output")}
        </Box>
      </Box>
      <Box
        sx={{
          maxHeight: "calc(100% - 38px)",
          overflowY: "hidden",
          overflowX: "hidden",
          pr: "10px",
          marginBottom: "30px",
          "&::-webkit-scrollbar": {
            width: "10px",
          },
          "&::-webkit-scrollbar-track": {
            background: "#333",
            borderRadius: "6px",
          },
          "&::-webkit-scrollbar-thumb": {
            background: "#666",
            borderRadius: "6px",
          },
          "&::-webkit-scrollbar-thumb:hover": {
            background: "#888",
            borderRadius: "6px",
          },
          "&:hover": {
            overflowY: "visible",
            pr: "0px",
          },
        }}
      >
        <Box
          style={{
            fontSize: "18px",
            paddingLeft: "26px",
            fontFamily: "Cascadia Code, regular",
            letterSpacing: "0px",
            paddingBottom: "24px",
            paddingTop: "8px",
            fontWeight: "normal",
            whiteSpace: 'pre-line'
          }}
        >
          {value?.output}
        </Box>
      </Box>
    </Box>
  );
};

export default Output;
