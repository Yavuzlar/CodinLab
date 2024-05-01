// ** MUI Imports
import { styled } from '@mui/material/styles'
import Box from '@mui/material/Box'

const ReactHotToast = styled(Box)(({ theme }) => {
  return {
    '& > div': {
      left: `${theme.spacing(6)} !important`,
      right: `${theme.spacing(6)} !important`,
      bottom: `${theme.spacing(6)} !important`,
      top: '75px !important'
    },
    '& .react-hot-toast': {
      fontWeight: 400,
      fontSize: '1rem',
      borderRadius: '1.25rem',
      letterSpacing: '0.14px',
      zIndex: theme.zIndex.snackbar,
      color: theme.palette.text.primary,
      background: theme.palette.background.paper,
      boxShadow:
        theme.palette.mode === 'light'
          ? '0px 4px 10px -4px rgb(58 53 65 / 60%)'
          : '0px 8px 16px -4px rgb(19 17 32 / 65%)',
      '&>:first-of-type:not([role])>:first-of-type': {
        width: 14,
        height: 14
      }
    }
  }
})

export default ReactHotToast
