const custombutton = (theme) => {
  return {

    darkButton: {
      backgroundColor: `${theme.palette.primary.dark} !important`,
      color: `${theme.palette.text.primary} !important`,
    },
    lightButton: {
      backgroundColor: `${theme.palette.divider} !important`,
      color: `${theme.palette.primary.dark} !important`,
    },
  };
};

export default custombutton;
