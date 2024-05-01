"use client";
import { ThemeProvider, createTheme } from '@mui/material'
import { theme as defaultTheme } from "@/configs/theme";
import { trTR } from '@mui/material/locale'
import overrides from "@/theme/overrides";
import typography from "@/theme/typography";

const ThemeComponent = ({ children }) => {

    const theme = createTheme(
        defaultTheme,
        {
            components: { ...overrides(defaultTheme) },
            typography: { ...typography(defaultTheme) }
        },
        trTR
    )

    const style = {
        color: theme.palette.text.primary,
    }

    return (
        <ThemeProvider theme={theme}>
            <div style={style}>
                {children}
            </div>
        </ThemeProvider>
    )
}

export default ThemeComponent