import { Box, Card, CardContent, Typography } from "@mui/material";
import LanguageProgress from "./LanguageProgress";
import LinearProgess from "../progress/LinearProgess";
import Image from "next/image";
import goImg from "../../assets/icons/go.png";
import cImg from "../../assets/icons/c.png";
import pythonImg from "../../assets/icons/python.png";
import { useTheme } from "@emotion/react";

const advancementLanguages = [
  {
    name: "Go",
    progress: 50,
    image: goImg,
  },
  {
    name: "C++",
    progress: 70,
    image: cImg,
  },
  {
    name: "Python",
    progress: 15,
    image: pythonImg,
  },
];
const Advancement = () => {
  const theme = useTheme();
  return (
    <Box
      sx={{
        width: "100%",
        height: "25rem",
      }}
    >
      <Card
        sx={{
          width: "100%",
          height: "100%",
        }}
      >
        <CardContent
          sx={{
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            height: "100%",
          }}
        >
          <Box sx={{ textAlign: "center" }}>
            <Typography variant="title">Advancement</Typography>
          </Box>
          <Box
            sx={{
              flexGrow: 1,
            }}
          >
            {advancementLanguages.map((languages, index) => (
              <Box
                sx={{
                  display: "flex",
                  px: "1rem",
                  flexDirection: "row",
                  alignContent: "center",
                  alignItems: "center",
                  height: "50px",
                  mt: "3rem",
                  mb: "2rem",
                }}
                key={index}
              >
                <Box sx={{ mr: "1rem" }}>
                  <Image src={languages.image} width={50} height={50} />
                </Box>
                <Box sx={{ width: "100%" }}>
                  <Typography sx={{ mt: "1rem" }}>{languages.name}</Typography>
                  <LinearProgess
                    progress={languages.progress}
                    backgroundColor={theme.palette.primary.dark}
                  />
                  <LinearProgess
                    progress={languages.progress}
                    backgroundColor={theme.palette.primary.light}
                  />
                </Box>
              </Box>
            ))}
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Advancement;
