
const typography = theme => {
  return {
    MuiTypography: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          color: theme.palette.text.primary,
          maxWidth: '80ch',
        }),
      }
    }
  }
}

export default typography
