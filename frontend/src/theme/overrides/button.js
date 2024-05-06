import { Opacity } from "@mui/icons-material"

const container = theme => {
  return {
    MuiButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          borderRadius: "0.938rem",
          fontSize: "1.125rem",
          Opacity: "1"
        }),
      },
    }
  }
}

export default container
