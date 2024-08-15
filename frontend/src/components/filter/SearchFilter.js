import { useTheme } from '@emotion/react';
import { Padding, Search } from '@mui/icons-material';
import { Box, FormControl, InputAdornment, TextField } from '@mui/material'
import React from 'react'
import { useTranslation } from 'react-i18next';

const SearchFilter = ({searchKey}) => {

    const { t } = useTranslation();
    const theme = useTheme()
  return (
    <Box sx={
        {
        
        }
    }>
    <FormControl fullWidth>
        <TextField
            name="search-in-labs"
            placeholder={t(searchKey)}
            variant="outlined"
            size="small"
            onChange={(e) => { setFilters({ ...filters, search: e.target.value }) }}
            InputProps={{
                startAdornment: (
                    <InputAdornment sx={{ zIndex: 10, mr: 1 }}>
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
  )
}

export default SearchFilter