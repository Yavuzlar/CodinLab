const typography = (theme) => {
  return {
    lineHeight: "1.5rem",
    fontWeight: 300,
    h1: {
      fontSize: "3.815rem",
    },
    h2: {
      fontSize: "3.052rem",
    },
    h3: {
      fontSize: "2.441rem",
    },
    h4: {
      fontSize: "1.953rem",
    },
    h5: {
      fontSize: "1.563rem",
    },
    h6: {
      fontSize: "1.25rem",
    },
    title: {
      fontSize: "2rem",
      letterSpacing: "0",
    },

    body1: {
      fontSize: "1.125rem",
      letterSpacing: "0",
    },

    caption: {
      fontSize: "1rem",
      color: `${theme.palette.border.secondary} !important`,
    },
    link: {
      //
      fontSize: "1rem",
      fontWeight: 400,
      color: "#121E35",
      fontStyle: "italic",
      cursor: "pointer",
      "&:hover": {
        color: theme.palette.primary.main,
      },
    },
    linklight: {
      //
      fontSize: "1rem",
      fontWeight: 400,
      color: "#F8F5EC",
      fontStyle: "italic",
      lineHeight: "1.875rem",
      cursor: "pointer",
      "&:hover": {
        color: "#fff",
      },
    },
    infoText: {
      color: theme.palette.primary.main,
      fontSize: "1rem",
    },
    infoText2: {
      color: theme.palette.primary.main,
      fontSize: "0.875rem",
    },
  };
};

export default typography;
