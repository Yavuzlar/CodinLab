import { Box, Button, Card, Typography } from "@mui/material";
import TestTubeRed from "../../assets/icons/red.png";
import TestTubeOrange from "../../assets/icons/orange.png";
import TestTubeGreen from "../../assets/icons/green.png";
import Image from "next/image";
import Translations from "../Translations";
const LabInfo = ({ programingId }) => {
  const labs = [
    {
      difficulty: "Easy",
      title: "Lab 1",
      finished: false,
    },
    {
      difficulty: "Medium",
      title: "Lab 2",
      finished: true,
    },
    {
      difficulty: "Easy",
      title: "Lab 3",
      finished: false,
    },
    {
      difficulty: "Hard",
      title: "Lab 4",
      finished: false,
    },
    {
      difficulty: "Medium",
      title: "Lab 5",
      finished: true,
    },
    {
      difficulty: "Hard",
      title: "Lab 6",
      finished: true,
    },
    {
      difficulty: "Easy",
      title: "Lab 7",
      finished: false,
    },
    {
      difficulty: "Medium",
      title: "Lab 8",
      finished: true,
    },
    {
      difficulty: "Easy",
      title: "Lab 9",
      finished: false,
    },
    {
      difficulty: "Hard",
      title: "Lab 10",
      finished: false,
    },
    {
      difficulty: "Medium",
      title: "Lab 11",
      finished: true,
    },
    {
      difficulty: "Hard",
      title: "Lab 12",
      finished: true,
    },
  ];

  return (
    <Box
      sx={{
        display: "flex",
        gap: "1.5rem",
        flexWrap: "wrap",
        justifyContent: "center",
      }}
    >
      {labs.map((lab, index) => (
        <Card
          sx={{
            width: "375px",
            height: "170px",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "column",
            gap: "1rem",
            padding: "1rem",
            borderRadius: "16px",
          }}
          key={index}
        >
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              flexDirection: "column",
            }}
          >
            <Image
              src={
                lab.difficulty === "Easy"
                  ? TestTubeGreen
                  : lab.difficulty === "Medium"
                  ? TestTubeOrange
                  : TestTubeRed
              }
              alt="difficulty"
              width={40}
              height={40}
            />

            <Typography
              variant="h5"
              sx={{
                fontFamily: "Outfit",
                fontWeight: "bold",
                fontSize: "25px",
                letterSpacing: "0",
                textAlign: "center",
                marginTop: "20px",
              }}
            >
              {lab.title}
            </Typography>

            <Button
              variant={lab.finished ? "dark" : "light"}
              sx={{
                borderRadius: "16px",
                marginTop: "20px",
                width: "140px",
                height: "40px",
                textTransform: "none",
                fontSize: "15px",
                fontWeight: "bold",
                fontFamily: "Outfit",
                letterSpacing: "0",
              }}
            >
              <Translations text={"lab.button.solve"} />
            </Button>
          </Box>
        </Card>
      ))}
    </Box>
  );
};

export default LabInfo;
