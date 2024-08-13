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
import roadsIcon from "../../assets/icons/icons8-path-100.png";
import labsIcon from "../../assets/icons/icons8-test-tube-100.png";

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
        height: "6.5rem",
        ...(_md && { justifyContent: "center" }),
        ...(_xl && { justifyContent: "center" }),
      }}
    >
      {!hovered ? (
        <CardContent
          sx={{
            display: "flex",
            flexDirection: "row",
            justifyContent: "start",
            alignItems: "center",
            gap: "0.5rem",
          }}
        >
          <Image
            src={language.image}
            alt={language.name}
            width={50}
            height={50}
          />
          <Typography
            variant="title"
            sx={{
              fontWeight: "bold",
              fontSize: "1.5rem",
            }}
          >
            {language.name}
          </Typography>
        </CardContent>
      ) : (
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            width: "100%",
            height: "100%",
          }}
        >
          <Button
            variant="dark"
            color="primary"
            sx={{
              height: "100%",
              width: "100%",
              backgroundColor: "primary.main",
              ":hover": {
                zIndex: 1,
                width: "100vh",
                transition: "width 0.3s",
              },
            }}
          >
            <Image src={roadsIcon} alt={"roadsIcon"} width={60} height={60} />
          </Button>

          <Button
            variant="dark"
            sx={{
              height: "100%",
              width: "100%",
              backgroundColor: "primary.main",
              ":hover": {
                zIndex: 1,
                width: "100vh",
                transition: "width 0.3s",
              },
            }}
          >
            <Image src={labsIcon} alt={"labsIcon"} width={60} height={60} />
          </Button>
        </Box>
      )}
    </Card>
  );
};

export default Languages;
