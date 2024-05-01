
const typography = theme => {
  return {
    MuiTypography: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          fontFamily: 'Roboto, sans-serif',
          color: theme.palette.text.primary,
        }),
      }
    }
  }
}

export default typography
