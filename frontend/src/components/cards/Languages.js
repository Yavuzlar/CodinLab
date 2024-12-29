import {
  Button,
  Card,
  Typography,
  CardContent,
  useMediaQuery,
} from "@mui/material";
import Image from "next/image";
import { useState } from "react";
import roadsIcon from "../../assets/icons/icons8-path-100.png";
import labsIcon from "../../assets/icons/icons8-test-tube-100.png";
import { useRouter } from "next/router";


const Languages = ({ language }) => {
  const [hovered, setHovered] = useState(false);
  const _md = useMediaQuery((theme) => theme.breakpoints.down("md"));
  const _xl = useMediaQuery((theme) => theme.breakpoints.down("xl"));
  const router = useRouter();


  const goToRoads = (id) => {
    router.push(`/roads/${id}`);
  };

  const goToLabs = (id) => {
    router.push(`/labs/${id}`);
  };

  return (
    //   {
    //     "id": 1,
    //     "name": "c++",
    //     "iconPath": "object/icons/c++.png"
    // }

    <Card
      // onMouseEnter={() => setHovered(true)}
      // onMouseLeave={() => setHovered(false)}
      sx={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "start",
        alignItems: "center",
        position: "relative",
        gap: "1.5rem",
        width: "100%",
        height: "6.5rem",
        ...(_md && { justifyContent: "center" }),
        ...(_xl && { justifyContent: "center" }),
        ".LanguageNameBox, .RoadsButton, .LabsButton": {
          transition: "opacity 600ms, width 400ms, left 200ms, right 200ms",
        },
        "&:hover": {
          ".LanguageNameBox": {
            opacity: 0,
          },
          ".RoadsButton, .LabsButton": {
            width: "50%",
            zIndex: 11,
            opacity: 1,
            "& .languageButtonIcon": {
              opacity: "1 !important",
              // display: 'block !important'
            },
          },
          ".RoadsButton": {
            left: 0,
          },
          ".LabsButton": {
            right: 0,
          },
        },
      }}
    >
      <CardContent
        className="LanguageNameBox"
        sx={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "start",
          alignItems: "center",
          gap: "0.5rem",
          zIndex: 10,
          opacity: 1,
        }}
      >
        <img
          src={"/api/v1/" + language.iconPath}
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

      {/* <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          width: hovered ? "100%" : "0%",
          height: "100%",
        }}
      > */}
      <Button
        onClick={()=> goToRoads(language.id)}
        className="RoadsButton"
        variant="dark"
        color="primary"
        sx={{
          position: "absolute",
          left: "-100%",
          height: "100%",
          opacity: 0,
          width: "0%",
          backgroundColor: "primary.main",
          // transition: "width 300ms",
          "&:hover": {
            zIndex: 1,
            width: "60% !important",
          },
        }}
      >
        <Image
          className="languageButtonIcon"
          src={roadsIcon}
          alt={"roadsIcon"}
          width={60}
          height={60}
          style={{ opacity: 0 }}
        />
      </Button>

      <Button
      onClick={()=> goToLabs(language.id)}
        className="LabsButton"
        variant="dark"
        sx={{
          position: "absolute",
          right: "-100%",
          height: "100%",
          opacity: 0,
          // width: "0%",
          backgroundColor: "primary.main",
          // transition: "width 300ms",
          "&:hover": {
            zIndex: 1,
            width: "60% !important",
          },
        }}
      >
        <Image
          className="languageButtonIcon"
          src={labsIcon}
          alt={"labsIcon"}
          width={60}
          height={60}
          style={{ opacity: 0 }}
        />
      </Button>
      {/* </Box> */}
    </Card>
  );
};

export default Languages;
