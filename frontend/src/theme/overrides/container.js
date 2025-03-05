import { Padding } from "@mui/icons-material";

const container = (theme) => {
  return {
    MuiContainer: {
      styleOverrides: {
        root: {
          background: "transparent !important",
          padding: "0 !important",
        },
      },
    },
  };
};

export default container;
