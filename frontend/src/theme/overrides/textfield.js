
const textfield = theme => {
  return {
    MuiTextField: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          // "& fieldset": {
          //     border: "1px solid #5C5C5C",
          //     borderRadius: "0.75rem",
          //     backgroundColor: "#FCFCFC",
          // },
          // // yazı rengi
          // "& .MuiInputBase-input": {
          //     color: "#000",
          //     zIndex: 9,
          // },

          // ...(ownerState.variant === "filled" && {
          //     borderRadius: "0.75rem",
          //     "& input, & textarea": {
          //         border: "unset !important",
          //     },
          //     "& .MuiInputBase-root": {
          //         backgroundColor: "#FCFCFC",
          //     },
          //     "& .MuiInputBase-root::before": {
          //         display: "none",
          //         border: "unset !important",
          //     },
          //     "& .MuiInputBase-root::after": {
          //         display: "none",
          //         border: "unset !important",
          //     },
          // }),
        }),
      },
    },
  }
}

export default textfield
