import { Box } from "@mui/material";

const Output = ({ value, params }) => {
  value = value || "Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd asOutput will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as  asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd asOutput will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as Output will be displayed here asdsad asdas asdasdas asd asd asdsa asd asdasdasd asdsadasd as ";

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
    overflow: "hidden", 
    
  }}
>
  <Box
    sx={{
      borderBottom: "2px solid white",
      marginTop: "18px",
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
  <Box
    sx={{
      maxHeight: "calc(100% - 38px)", // İçerik yüksekliği hesaplama
      overflowY: "auto", // Yalnızca dikey yönde scrollbar göster
      overflowX: "hidden", // Yatay scrollbar'ı gizle
      "&::-webkit-scrollbar": {
        width: "12px",
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
    }}
  >
    <div
      style={{
        padding: "10px",
        fontSize: "14px",
      }}
    >
      {value}
    </div>
  </Box>
</Box>
  );
};

export default Output;
