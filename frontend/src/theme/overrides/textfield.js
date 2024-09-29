const textfield = (theme) => {
  return {
    MuiTextField: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          "& fieldset": {
            border: "none",
            borderRadius: "1rem",
            backgroundColor: "#FFF",
          },
          "& .MuiInputBase-input": {
            color: theme.palette.common.black,
            zIndex: 1,
          },
          // variant={filled}
          ...(ownerState.variant === "filled" && {
            "& input, & textarea": {
              border: "unset !important",
            },
            "&:focus-within": {
              backgroundColor: theme.palette.primary.light,
              borderRadius: "15px",
            },
            "& .MuiInputBase-root": {
              backgroundColor: theme.palette.primary.main,
              borderRadius: "15px",
            },
            "& .MuiInputBase-root::before": {
              display: "none",
              border: "unset !important",
            },
            "& .MuiInputBase-root::after": {
              display: "none",
              border: "unset !important",
            },
          }),
          // variant={outlined}
          ...(ownerState.variant === "outlined" && {
            "& .MuiInputBase-root": {
              "&.Mui-focused fieldset": {
                border: "2px solid #0A3B7A",
              },
              "& input:-webkit-autofill": {
                borderRadius: "15px",
              },
              "& input:-webkit-autofill:focus": {
                borderRadius: "15px",
                border: "2px solid #0A3B7A",
              },
            },
          }),
        }),
      },
    },
  };
};

export default textfield;
