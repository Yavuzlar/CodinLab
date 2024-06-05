import {
  Box,
  Card,
  CardContent,
  Grid,
  Typography,
  useMediaQuery,
} from "@mui/material";
import Image from "next/image";

const Welcome = ({ title = "", description = "", image = null }) => {
  const _lgPlus = useMediaQuery((theme) => theme.breakpoints.down("lgPlus"));
  const _mdlg = useMediaQuery((theme) => theme.breakpoints.down("mdlg"));
  const _smd = useMediaQuery((theme) => theme.breakpoints.down("smd"));
  const _lg = useMediaQuery((theme) => theme.breakpoints.down("lg"));

  return (
    <Box sx={{ display: "flex", position: "relative", width: "100%" }}>
      <Card sx={{ height: "20rem", width: "100%" }}>
        <CardContent sx={{ height: "calc(100% - 3rem)" }}>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              justifyContent: "center",
              height: "100%",
            }}
          >
            <Typography
              variant="title"
              sx={{ textAlign: "left", fontWeight: "bold" }}
            >
              {title}
            </Typography>
            <Typography
              sx={{
                textAlign: "left",
                paddingTop: "1.8rem",
                ...(_mdlg
                  ? { maxWidth: "40ch" }
                  : _lgPlus && { maxWidth: "60ch" }),
              }}
            >
              {description}
            </Typography>
          </Box>
        </CardContent>
      </Card>
      {image && (
        <Image
          style={{
            position: "absolute",
            right: _lg ? "1rem" : "8rem",
            opacity: _smd ? "0.25" : "1",
            bottom: "0",
          }}
          src={image}
          alt="Welcome"
          width={338}
          height={331}
        />
      )}
    </Box>
  );
};

export default Welcome;


