import { Box, CircularProgress } from "@mui/material";

export const CircularProgressStatistics = ({ progresses }) => {
  console.log(progresses);
  return (
    <Box
      sx={{
        position: "relative",
        display: "inline-flex",
        backgroundColor: "#aaaaaa",
      }}>
      <CircularProgress
        variant="determinate"
        value={progresses[0].value}
        size={120}
        thickness={4}
      />
      {progresses[1] && (
        <Box
          sx={{
            top: 0,
            left: 0,
            bottom: 0,
            right: 0,
            position: "absolute",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}>
          <CircularProgress
            variant="determinate"
            value={progresses[1].value}
            size={90}
            color="secondary"
          />
          {progresses[2] && (
            <Box
              sx={{
                top: 0,
                left: 0,
                bottom: 0,
                right: 0,
                position: "absolute",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
              }}>
              <CircularProgress
                variant="determinate"
                value={progresses[2].value}
                size={60}
              />
            </Box>
          )}
        </Box>
      )}
    </Box>
  );
};
