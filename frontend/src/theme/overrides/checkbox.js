const checkbox = theme => {
    return {
      MuiCheckbox: {
        styleOverrides: {
          root: {
            width: 20,
            height: 20,
            border: "1px solid white", 
            borderRadius: 5,
            padding: 0, 
            opacity: 1,
            "& .MuiIconButton-label": {
              backgroundColor: "transparent", 
            },
            "& .MuiSvgIcon-root": {
              color: theme.palette.primary.dark,
              display: "none",
            },
            "&.Mui-checked": {
              backgroundColor: "white" + "!important",
              borderColor: theme.palette.primary.dark + "!important",
              "& .MuiSvgIcon-root": {
                display: "block",
              },
            },
            // disabled
            "&.Mui-disabled": {
              borderColor: theme.palette.action.disabled + "!important",
              "& .MuiSvgIcon-root": {
                color: theme.palette.action.disabled + "!important",
              },
            },
          },
        },
      },
    }
  }
  
  export default checkbox
  