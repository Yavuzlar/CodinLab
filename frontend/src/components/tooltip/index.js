import { Tooltip, styled, tooltipClasses } from "@mui/material";

const CustomTooltip = styled(({ className, ...props }) => (
    <Tooltip {...props} classes={{ popper: className }} />
))(({ theme }) => ({
    zIndex: 999999,
    [`& .${tooltipClasses.arrow}`]: {
        "&::before": {
            backgroundColor: theme.palette.background.default,
            borderBottomRightRadius: "3px",
        },
    },
    [`& .${tooltipClasses.tooltip}`]: {
        backgroundColor: theme.palette.background.default,
        color: theme.palette.text.secondary,
        fontSize: theme.typography.pxToRem(12),
        padding: "0.5rem 0.75rem",
        fontFamily: "inherit",
        borderRadius: "1.25rem",
    },
}));

export default CustomTooltip;
