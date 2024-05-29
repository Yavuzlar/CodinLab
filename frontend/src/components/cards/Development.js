import { Box, Card, CardContent, Typography } from "@mui/material";

const Development = () => {
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
          <Typography variant="title">Development Components</Typography>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Development;
