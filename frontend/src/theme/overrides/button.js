import { hexToRGBA } from "src/utils/hex-to-rgba";

const button = (theme) => {
  return {
    MuiButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          borderRadius: "0.938rem",
          fontSize: "1.125rem",
          opacity: 1,
          textTransform: "none",

          "& .MuiTypography-root": {
            color: theme.palette[ownerState.color || "primary"].contrastText,
          },
          backgroundColor: theme.palette[ownerState.color || "primary"].main,

          ...(ownerState.variant == "dark" && {
            backgroundColor: theme.palette.primary.dark,
            "&:hover": {
              boxShadow: `0px 2px 6px ${theme.palette.common.black}`,
              backgroundColor: theme.palette.primary.dark,
              color: theme.palette.text.primary,
            },
          }),

          ...(ownerState.variant == "light" && {
            backgroundColor: theme.palette.divider,
            color : theme.palette.common.black,
            "&:hover": {
              boxShadow: `0px 2px 6px ${theme.palette.primary.dark}`,
              backgroundColor: theme.palette.divider,
              color: theme.palette.primary.dark,
            },
          }),

          ...(ownerState.variant == "outlined" && {
            backgroundColor: "transparent",
            border: `1px solid ${theme.palette.border.main}`,
            "&:hover": {
              border: `1px solid ${theme.palette.border.main}`,
              // border: "1px solid " + theme.palette.border.main,
              backgroundColor: hexToRGBA(theme.palette.border.main, 0.3),
              color: theme.palette.primary.dark,
            },
          }),
        }),
      },
    },
  };
};

export default button;
