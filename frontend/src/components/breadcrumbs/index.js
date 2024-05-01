import { Breadcrumbs, Link, Typography } from '@mui/material'

const CustomBreadcrumbs = ({ titles }) => {
    return (
        <Breadcrumbs separator="›" aria-label="breadcrumb">
            {
                titles.map((item, index) => (
                    (index == titles.length - 1)
                        ? <Typography key={index} variant='caption' sx={{ display: 'flex', alignItems: 'center', gap: '8px', cursor: 'default', color: theme => `${theme.palette.text.primary} !important` }}>{item.icon}{item.title}</Typography>
                        : !item?.path
                            ? <Typography key={index} variant='caption' sx={{ display: 'flex', alignItems: 'center', gap: '8px', cursor: 'default' }}>{item.icon}{item.title}</Typography>
                            : <Link
                                underline="hover"
                                key={index}
                                color="inherit"
                                href={item?.path}
                            >
                                <Typography variant='caption' sx={{ display: 'flex', alignItems: 'center', gap: '8px' }}>{item.icon}{item.title}</Typography>
                            </Link>
                ))
            }
        </Breadcrumbs>
    )
}

export default CustomBreadcrumbs