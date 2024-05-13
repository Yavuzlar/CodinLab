
const typography = theme => {
  return {
    MuiTypography: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          color: theme.palette.text.primary,
        }),
      }
    }
  }
}

export default typography
