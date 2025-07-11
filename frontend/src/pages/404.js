// ** Next Import
// ** MUI Components
import { Box, Button, Typography } from "@mui/material";
import { styled } from "@mui/material/styles";
// ** Layout Import
import BlankLayout from "src/layout/BlankLayout";
import { useRouter } from "next/router"; // Router Import
import { hexToRGBA } from "src/utils/hex-to-rgba";

// ** Styled Components
const BoxWrapper = styled(Box)(({ theme }) => ({
  [theme.breakpoints.down("md")]: {
    width: "90vw",
  },
}));

const Error404 = () => {
  const router = useRouter();

  return (
    <Box
      sx={{
        height: "calc(100vh - 80px)",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        textAlign: "center",
        justifyContent: "center",
      }}
    >
      <BoxWrapper>
        <Typography
          variant="h1"
          sx={{
            fontSize: "230px",
            margin: "10vh 0 0",
            textAlign: "center",
            letterSpacing: "5px",
            backgroundColor: (theme) => theme.palette.secondary.main,
            color: "transparent",
            textShadow: (theme) =>
              `3px 3px 3px ${hexToRGBA(theme.palette.background.default, 0.5)}`,
            webkitBackgroundClip: "text",
            mozBackgroundClip: "text",
            backgroundClip: "text",
            fontWeight: 600,
          }}
        >
          404
        </Typography>
        <Typography variant="h5" sx={{ mb: 1, fontSize: "1.5rem !important" }}>
          Page Not Found ⚠️
        </Typography>
      </BoxWrapper>
      <Button
        variant="contained"
        sx={{ px: 5.5, mt: 3 }}
        onClick={() => {
          if (window.history.length > 2) {
            router.back();
            setTimeout(() => {
              if (router.push?.pathname?.includes("404")) {
                window.history.go(-2);
              }
            }, 100);
          } else {
            router.back();
          }
        }}
      >
        Go Back
      </Button>
    </Box>
  );
};

Error404.getLayout = (page) => <BlankLayout>{page}</BlankLayout>;
export default Error404;
