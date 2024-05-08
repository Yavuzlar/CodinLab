import { Box } from "@mui/material";

const Output = ({ value, params }) => {
  value = value || "Output will be displayed here";

  const width = params.width ? params.width : "50%";
  const height = params.height ? params.height : "50%";

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        width: width,
        height: height,
        padding: "12px",
        borderRadius: "30px",
        opacity: "1",
        backgroundColor: "black",
        color: "white",
        overflow: "auto", 
      }}
    >
      <Box
        sx={{
          borderBottom: "2px solid white",
          marginTop: "26px",
          paddingBottom: "10px",
          paddingLeft: "16px",
          fontSize: "18px",
          fontWeight: "bold",
        }}
      >
        <div
          sx={{
            display: "flex",
          }}
        >
          <div
            style={{
              textAlign: "left",
              letterSpacing: "0px",
              color: "#FFFFFF",
              opacity: "1",
            }}
          >
            Output
          </div>
        </div>
      </Box>
      <Box>
        <div
          style={{
            padding: "10px",
            fontSize: "20px",
            width: "100%",
            height: "100%",
            maxHeight: "100%",
            overflow: "auto",
          }}
        >
          {value}
        </div>
      </Box>
    </Box>
  );
};

export default Output;
