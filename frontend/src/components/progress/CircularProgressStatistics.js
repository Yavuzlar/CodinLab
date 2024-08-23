import { Box, createTheme, ThemeProvider, Typography } from "@mui/material";
import CircularProgress, {
  circularProgressClasses,
} from "@mui/material/CircularProgress";

/*
Takes an array of object to shown as a prop.
progresses = [
  {
    name: "In progress", // String
    value: 90 // Number
    color: "#0A3B7A" // String
  }
]
*/

export const CircularProgressStatistics = ({ progresses }) => {
  if (!progresses) {
    return (
      <Box>
        <Typography variant="body1" color="error">
          No data provided for the CircularProgressStatistics
        </Typography>
      </Box>
    );
  }

  const theme = createTheme({
    palette: {
      first: {
        main: progresses[0]?.color ? progresses[0].color : "#0A3B7A",
      },
      second: {
        main: progresses[1]?.color ? progresses[1].color : "#8FDDFD",
      },
      third: {
        main: progresses[2]?.color ? progresses[2].color : "#000fff",
      },
    },
  });

  const sizeFirst = 160;
  const sizeSecond = sizeFirst - 35;
  const sizeThird = sizeSecond - 45;

  const thicknessFirst = 2.5;
  const thicknessSecond = 3.8;
  const thicknessThird = 5.5;

  return (
    <ThemeProvider theme={theme}>
      <Box
        sx={{
          position: "relative",
          display: "inline-flex",
        }}>
        <CircularProgress
          variant="determinate"
          sx={{
            color: "#fff",
          }}
          size={sizeFirst}
          thickness={thicknessFirst - 0.1}
          value={100}
        />
        <CircularProgress
          variant="determinate"
          value={progresses[0].value === 0 ? 1 : progresses[0].value}
          size={sizeFirst}
          thickness={thicknessFirst}
          color="first"
          sx={{
            position: "absolute",
            [`& .${circularProgressClasses.circle}`]: {
              strokeLinecap: "round",
            },
          }}
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
              sx={{
                color: "#fff",
              }}
              size={sizeSecond}
              thickness={thicknessSecond - 0.1}
              value={100}
            />
            <CircularProgress
              variant="determinate"
              value={progresses[1].value === 0 ? 1 : progresses[1].value}
              size={sizeSecond}
              thickness={thicknessSecond}
              color="second"
              sx={{
                position: "absolute",
                [`& .${circularProgressClasses.circle}`]: {
                  strokeLinecap: "round",
                },
              }}
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
                  sx={{
                    color: "#fff",
                  }}
                  size={sizeThird}
                  thickness={thicknessThird - 0.1}
                  value={100}
                />
                <CircularProgress
                  variant="determinate"
                  value={progresses[2].value === 0 ? 1 : progresses[2].value}
                  size={sizeThird}
                  thickness={thicknessThird}
                  color="third"
                  sx={{
                    position: "absolute",
                    [`& .${circularProgressClasses.circle}`]: {
                      strokeLinecap: "round",
                    },
                  }}
                />
              </Box>
            )}
          </Box>
        )}
      </Box>
    </ThemeProvider>
  );
};
