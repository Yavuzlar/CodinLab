// ** Next Import
import Link from 'next/link'

// ** MUI Components
import {
  Button,
  Typography,
  Box,
} from '@mui/material'
import { styled } from '@mui/material/styles'

// ** Layout Import
import BlankLayout from '@/layout/BlankLayout'
import { hexToRGBA } from '@/utils/hex-to-rgba'

// ** Styled Components
const BoxWrapper = styled(Box)(({ theme }) => ({
  [theme.breakpoints.down('md')]: {
    width: '90vw'
  },
}))

const Error401 = () => {
  return (
    <Box className='content-center'>
      <Box sx={{ p: 5, display: 'flex', flexDirection: 'column', alignItems: 'center', textAlign: 'center' }}>
        <BoxWrapper>

          <Typography
            variant='h1'
            sx={{
              fontSize: "230px",
              margin: '10vh 0 0',
              textAlign: 'center',
              letterSpacing: '5px',
              backgroundColor: theme => theme.palette.error.main,
              color: 'transparent',
              textShadow: theme => `3px 3px 3px ${hexToRGBA(theme.palette.background.default, 0.5)}`,
              webkitBackgroundClip: 'text',
              mozBackgroundClip: 'text',
              backgroundClip: 'text',
              fontWeight: 600
            }}
          >
            401
          </Typography>

          <Typography variant='h5' sx={{ mb: 1, fontSize: '1.5rem !important' }}>
            Unauthorized! 🔐
          </Typography>

          <Typography variant='body2'>
            You are not authorized to access this page. Go back to the home page!
          </Typography>
        </BoxWrapper>

        <Link passHref href='/'>
          <Button variant='contained' sx={{ px: 5.5, mt: 3 }}>
            Go Home
          </Button>
        </Link>
      </Box>
    </Box>
  )
}
Error401.getLayout = page => <BlankLayout>{page}</BlankLayout>

export default Error401
