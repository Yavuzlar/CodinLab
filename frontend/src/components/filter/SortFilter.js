import { PlayArrow } from "@mui/icons-material"
import { Box, Card, Typography, useTheme } from "@mui/material"
import { useEffect, useState } from "react"

const SortFilter = ({ filters, setFilters }) => {
    const theme = useTheme()

    const [sort, setSort] = useState(0) // 0,1,2

    useEffect(() => {
        setFilters({ ...filters, sort: sort == 2 ? "desc" : sort == 1 ? "asc" : "" })
    }, [sort])

    return (
        <Card sx={{
            width: '100%', height: 'calc(100% - 4px)', cursor: 'pointer',
            border: "2px solid " + theme.palette.primary.main,
            "&:hover": {
                border: "2px solid " + theme.palette.primary.dark,
            },

        }}
            onClick={() => { setSort((sort + 1) % 3) }}
        >
            <Box sx={{ display: 'flex', alignItems: 'center', p: '0px 16px 0px 8px', height: '100%' }}>
                <Box
                    sx={{
                        display: 'flex',
                        width: '28px',
                        height: '26px',
                        position: 'relative',
                        transform: "rotate(90deg)",
                        // color: theme.palette.primary.dark
                    }}
                >
                    <PlayArrow sx={{ color: sort == 1 && theme.palette.primary.dark, width: '18px', transform: "rotate(180deg)", position: 'absolute', left: 0 }} /> {/* ASC a-z */}
                    <PlayArrow sx={{ color: sort == 2 && theme.palette.primary.dark, width: '18px', transform: "rotate(0deg)", position: 'absolute', right: 0 }} /> {/* DESC z-a */}
                </Box>

                <Typography>
                    Sort the labs
                </Typography>
            </Box>
        </Card>
    )
}

export default SortFilter