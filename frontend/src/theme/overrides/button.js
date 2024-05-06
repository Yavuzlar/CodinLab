const button = (theme) => {
  return {
    MuiButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          borderRadius: "0.938rem",
          fontSize: "1.125rem",
          opacity: 1,

          backgroundColor: theme.palette[ownerState.color || "primary"].main,
          color: theme.palette[ownerState.color || "primary"].contrastText,

          ...(ownerState.variant == "dark" && {
            backgroundColor: theme.palette.primary.dark,
            color: theme.palette.text.primary,
            "&:hover": {
              boxShadow: `0px 3px 12px ${theme.palette.common.black}`,
              backgroundColor: theme.palette.primary.dark,
              color: theme.palette.text.primary,
            },
          }),

          ...(ownerState.variant == "light" && {
            backgroundColor: theme.palette.divider,
            color: theme.palette.primary.dark,
            "&:hover": {
              boxShadow: `0px 3px 12px ${theme.palette.primary.dark}`,
              backgroundColor: theme.palette.divider,
              color: theme.palette.primary.dark,
            },
          }),
        }),
      },
    },
  };
};

export default button;
