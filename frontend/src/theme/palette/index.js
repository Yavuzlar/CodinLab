import { hexToRGBA } from "@/utils/hex-to-rgba";

const text = {
    primary: "#F0F3F6",
};

const border = {
    main: "#7A828E",
    light: "#525964",
    secondary: "#BDC4CC"
}

const palette = {
    action: {
        active: hexToRGBA(border.main, 0.12),
        disabled: border.dark,
        disabledBackground: border.light,
        focus: border.light,
        hover: hexToRGBA(border.main, 0.24),
        selected: border.light,
    },
    common: {
        black: "#0F0F0F",
        white: text.primary,
    },
    divider: border.main,
    primary: {
        main: "#409EFF",
        light: "#91CBFF",
        dark: "#318BF8",
        contrastText: text.primary,
    },
    secondary: {
        main: "#8997AC",
        light: "#9AA8BC",
        dark: "#647185",
        contrastText: text.primary,
    },
    success: {
        main: "#72F088",
        light: "#ACF7B6",
        dark: "#007728",
        contrastText: text.primary,
    },
    warning: {
        main: "#F0B72F",
        light: "#F7C843",
        dark: "#E09B13",
        contrastText: text.primary,
    },
    info: {
        main: "#CB9CFF",
        light: "#DBB7FF",
        dark: "#6921D7",
        contrastText: text.primary,
    },
    error: {
        main: "#FF4445",
        light: "#FF9492",
        dark: "#CC1421",
        contrastText: text.primary,
    },
    background: {
        default: "#0A0C10",
        paper: "#272B33",
    },
    border: {
        main: border.main,
        light: border.light,
        secondary: border.secondary,
    },
    text: {
        primary: text.primary,
        secondary: text.primary,
        disabled: border.dark,
    },
};

export default palette;