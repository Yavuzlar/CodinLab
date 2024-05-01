"use client";

// ** Util Imports
import { hexToRGBA } from "@/utils/hex-to-rgba"

const chip = theme => {
  return {
    MuiChip: {
      styleOverrides: {
        root: {
          '&.MuiChip-rounded': {
            borderRadius: '1.25rem'
          }
        },
        outlined: {
          '&.MuiChip-colorDefault': {
            borderColor: `rgba(${theme.palette.mainColor}, 0.22)`
          }
        },
        deleteIcon: {
          width: 18,
          height: 18
        },
        avatar: {
          color: theme.palette.text.primary
        },
        iconColorDefault: {
          color: theme.palette.text.primary
        },
        deletableColorPrimary: {
          '&.MuiChip-light .MuiChip-deleteIcon': {
            color: hexToRGBA(theme.palette.primary.main, 0.7),
            '&:hover': {
              color: theme.palette.primary.main
            }
          }
        },
        deletableColorSecondary: {
          '&.MuiChip-light .MuiChip-deleteIcon': {
            color: hexToRGBA(theme.palette.secondary.main, 0.7),
            '&:hover': {
              color: theme.palette.secondary.main
            }
          }
        },
        deletableColorSuccess: {
          '&.MuiChip-light .MuiChip-deleteIcon': {
            color: hexToRGBA(theme.palette.success.main, 0.7),
            '&:hover': {
              color: theme.palette.success.main
            }
          }
        },
        deletableColorError: {
          '&.MuiChip-light .MuiChip-deleteIcon': {
            color: hexToRGBA(theme.palette.error.main, 0.7),
            '&:hover': {
              color: theme.palette.error.main
            }
          }
        },
        deletableColorWarning: {
          '&.MuiChip-light .MuiChip-deleteIcon': {
            color: hexToRGBA(theme.palette.warning.main, 0.7),
            '&:hover': {
              color: theme.palette.warning.main
            }
          }
        },
        deletableColorInfo: {
          '&.MuiChip-light .MuiChip-deleteIcon': {
            color: hexToRGBA(theme.palette.info.main, 0.7),
            '&:hover': {
              color: theme.palette.info.main
            }
          }
        }
      }
    }
  }
}

export default chip
