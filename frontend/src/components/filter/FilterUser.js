import { Search } from "@mui/icons-material";
import { Box, FormControl, InputAdornment, TextField, Typography, useTheme } from "@mui/material"
import { useTranslation } from "react-i18next"
import { hexToRGBA } from "src/utils/hex-to-rgba"
import SortFilterUser from "./SortFilterUser"

const FilterUser = ({
    filters = {
        status: "all", 
        search: "",
        sort: "",
    },
    setFilters = () => { }
}) => {

    const { t } = useTranslation();
    const theme = useTheme()

    const progressStatuses = [
        {
            name: t("all"),
            status: "all"
        },
    ]

    return (
        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', minHeight: '44px', flexWrap: 'wrap' }}>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: '20px', marginRight:'32px' }}>
                {
                    progressStatuses.map((item, index) => {
                        return (
                            <Typography
                                key={index}
                                sx={{
                                    cursor: 'default',
                                    color: theme => filters.status == item.status
                                        ? theme.palette.primary.dark
                                        : hexToRGBA(theme.palette.primary.dark, 0.6),
                                    "&:hover": {
                                        textDecoration: 'underline',
                                        cursor: filters.status != item.status ? 'pointer' : 'default'
                                    }
                                }}
                                onClick={() => {
                                    setFilters({ ...filters, status: item.status })
                                }}
                            >
                                {item.name}
                            </Typography>
                        )
                    })
                }
            </Box>

            <Box sx={{ flex: 2 }}>
                <FormControl fullWidth>
                    <TextField
                        name="search-in-labs"
                        placeholder={t("users.search.placeholder")}
                        variant="outlined"
                        size="small"
                        onChange={(e) => { setFilters({ ...filters, search: e.target.value }) }}
                        InputProps={{
                            startAdornment: (
                                <InputAdornment sx={{ zIndex: 10, }}>
                                    <Search />
                                </InputAdornment>
                            ),
                            style: { color: theme.palette.text.primary }
                        }}
                        sx={{
                            "& .MuiInputBase-input": {
                                color: theme.palette.text.primary,
                                zIndex: 9,
                                "&::placeholder": {
                                    color: theme.palette.text.primary,
                                    opacity: 0.7
                                }
                            },
                            "& .MuiOutlinedInput-root": {
                                "& fieldset": {
                                    backgroundColor: theme.palette.primary.main
                                },
                            },
                            width: "100%",
                            height: "100%",
                        }}
                    />
                </FormControl>
            </Box>
            
            <Box sx={{ flex: 1, height: '44px', display: 'flex', justifyContent: 'flex-end',marginLeft:'32px' }}>
                <SortFilterUser filters={filters} setFilters={setFilters} />
            </Box>
        </Box>
    )
}

export default FilterUser;
