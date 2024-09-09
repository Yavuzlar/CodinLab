const typography = (theme) => {
  return {
    MuiTypography: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          fontFamily: "'Outfit', sans-serif",
          color: theme.palette.text.primary,
          maxWidth: "80ch",
        }),
      },
    },
  };
};

export default typography;
