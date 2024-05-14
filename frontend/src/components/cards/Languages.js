import {
  Box,
  Button,
  Card,
  Typography,
  Grid,
  CardContent,
  useMediaQuery,
} from "@mui/material";
import Image from "next/image";
import { useState } from "react";

const Languages = ({ language }) => {
  const [hovered, setHovered] = useState(false);
  const _md = useMediaQuery((theme) => theme.breakpoints.down("md"));
  const _xl = useMediaQuery((theme) => theme.breakpoints.down("xl"));

  return (
    <Card
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
      sx={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "start",
        alignItems: "center",
        gap: "1.5rem",
        width: "100%",
        height: "104px",
        position: "relative",
        ...(_md && {justifyContent:"center"}),

        ...(_xl && {justifyContent:"center"})

      }}
    >
      {!hovered ? (
        <CardContent
          sx={{
            display: "flex",
            flexDirection: "row",
            justifyContent: "start",
            alignItems: "center",
            gap: "1.5rem",
            ...(_md && {gap:"0.5rem"}),

          }}
        >
          <Image
            src={language.image}
            alt={language.name}
            width={60}
            height={60}
          />
          <Typography>{language.name}</Typography>
        </CardContent>
      ) : (
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            width: "100%",
            height: "100%",
            // position: "absolute",
            // top: 0,
            // left: 0,
            // backgroundColor: "white",
          }}
        >
          <Button
            variant="dark"
            color="primary"
            sx={{
              height: "100%",
              width: "100%",
              backgroundColor: "#3894d0",
              ":hover": {
                zIndex: 1,
                width: "70vh",
              },
            }}
          >
            Button 1
          </Button>

          <Button
            variant="dark"
            sx={{
              height: "100%",
              width: "100%",
              backgroundColor: "#3894d0",

              ":hover": {
                zIndex: 1,
                width: "70vh",
              },
            }}
          >
            Button 2
          </Button>
        </Box>
      )}
    </Card>
  );
};

export default Languages;
