import { Box, Card, CardContent, Typography } from "@mui/material";
import LevelStatistic from "../statistics/LevelStatistic";

const Level = () => {
  return (
    <Box
      sx={{
        width: "100%",
        height: "20rem",
      }}
    >
      <Card
        sx={{
          width: "100%",
          height: "100%",
        }}
      >
        <CardContent sx={{ textAlign: "center" }}>
          <Typography variant="title">Level Components</Typography>
          <LevelStatistic />
        </CardContent>
      </Card>
    </Box>
  );
};

export default Level;
