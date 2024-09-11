import { Box, Card, CardContent, Typography } from "@mui/material";
import LinearProgess from "../progress/LinearProgess";
import { useTheme } from "@emotion/react";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { fetchAdvancement } from "src/store/statistics/statisticsSlice";

const Advancement = () => {
  const dispatch = useDispatch();
  const { statistics: stateStatistics } = useSelector((state) => state);

  useEffect(() => {
    dispatch(fetchAdvancement());
  }, [dispatch]);

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
            {advancementData.map((languages, index) => (
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
                  <img
                    src={"/api/v1/" + languages.iconPath}
                    width={50}
                    height={50}
                  />
                </Box>
                <Box sx={{ width: "100%" }}>
                  <Typography sx={{ mt: "1rem" }}>{languages.name}</Typography>
                  <LinearProgess
                    progress={languages.roadPercentage}
                    backgroundColor={theme.palette.primary.dark}
                  />
                  <LinearProgess
                    progress={languages.labPercentage}
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
