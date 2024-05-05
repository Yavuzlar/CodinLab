import { hexToRGBA } from "@/utils/hex-to-rgba";

const text = {
    primary: "#FFFFFF",
    secondary: "#9DB1CA",
};

const border = {
    main: "#0A3B7A",
    light: "#40C7FF",
    secondary: "#0a3b7a"
}

const palette = {
    action: {
        active: "#FFFFFF",
        disabled: "#9DB1CA", 
        disabledBackground: border.light,
        focus: "#FFFFFF",
        hover: hexToRGBA(text.primary, 0.24), // if dont like this color, change it for better readability
        selected: text.primary, 
    },
    common: {
        black: "#000000",
        white: text.primary,
    },
    divider: "#FFFFFF",
    primary: {
        main: "#3894D0",
        light: "#8FDDFD",
        dark: "#0A3B7A",
        contrastText: text.primary,
    },
    success: { 
        main: "#39CE19",
        light: "#BDEEAF",
        dark: "#20AC02",
        contrastText: text.primary,
    },
    warning: { 
        main: "#FFCA00",
        light: "#FDEDAE",
        dark: "#C0A600",
        contrastText: text.primary,
    },
    info: {
        main: "#40C7FF",
        light: "#8FDDFD",
        dark: "#1851B2",
        contrastText: text.primary,
    },
    error: { 
        main: "#DC0101",
        light: "#F3B3B3",
        dark: "#BA0000",
        contrastText: text.primary,
    },
    background: {
        default: "#DAF0FE",
        paper: "#3894d0",
    },
    border: {
        main: border.main,
        light: border.light,
        secondary: border.secondary,
    },
    text: {
        primary: text.primary,
        secondary: text.secondary,
        disabled: "#f00", // if we use this color in code , change it for better readability
    },
};

export default palette;