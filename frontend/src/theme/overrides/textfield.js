const textfield = (theme) => {
  return {
    MuiTextField: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          "& fieldset": {
            border: "none",
            borderRadius: "15px",
            backgroundColor: "#FCFCFC",
          },
          // yazı rengi
          "& .MuiInputBase-input": {
            color: "#000",
            zIndex: 9,
          },

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
          ...(ownerState.variant === "outlined" && {
            "& .MuiInputBase-root": {
              "&.Mui-focused fieldset": {
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
