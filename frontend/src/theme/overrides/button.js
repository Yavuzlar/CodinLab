
const container = theme => {
  return {
    MuiButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          borderRadius: "1.25rem",
          backgroundColor: ownerState.color === theme.palette.text.primary ? theme.palette.primary.main : theme.palette.common.white,

        }),
      },
    }
  }
}

export default container
