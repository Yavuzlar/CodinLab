import { Box, Typography } from "@mui/material";

const LinearProgess = ({ progress = 0, backgroundColor }) => {

  return (
    <Box
      sx={{
        display: "flex",
        width: "100%",
        alignItems: "center",
        gap: "0.5rem",
      }}
    >
      <Box
        sx={{
          display: "flex",
          position: "relative",
          width: "100%",
          height: "0.75rem",
          borderRadius: "0.5rem",
          backgroundColor: "white",
        }}
      >
        <Box
          sx={{
            display: "flex",
            position: "relative",
            width: `${progress}%`,
            height: "0.75rem",
            borderRadius: "0.5rem",
            backgroundColor: (theme) => theme.palette.primary.dark,
          }}
        ></Box>
      </Box>

      <Typography sx={{ width: "fit-content" }}>{progress}%</Typography>
    </Box>
  );
};

export default LinearProgess;
