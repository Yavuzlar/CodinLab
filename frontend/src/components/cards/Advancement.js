import { Box, Card, CardContent, Typography } from "@mui/material";
import LinearProgess from "../progress/LinearProgess";
import { useTheme } from "@emotion/react";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { fetchAdvancement } from "src/store/statistics/statisticsSlice";
import { useTranslation } from "react-i18next";

const Advancement = () => {
  const dispatch = useDispatch();
  const { statistics: stateStatistics } = useSelector((state) => state);

  useEffect(() => {
    dispatch(fetchAdvancement());
  }, [dispatch]);

  const theme = useTheme();

  const { t } = useTranslation();

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
            height: "calc(100% - 1.5rem)",
            display: "flex",
            flexDirection: "column",
          }}
        >
          <Box sx={{ textAlign: "center" }}>
            <Typography variant="title">
              {t("home.advancement.title")}
            </Typography>
          </Box>
          <Box
            sx={{
              width: "100%",
              height: "100%",
              overflow: "auto",
              "&::-webkit-scrollbar": {
                width: "0rem",
              },
              "&::-webkit-scrollbar-track": {
                background: "transparent",
              },
            }}
          >
            {stateStatistics.advancementData?.data?.map((languages, index) => (
              <Box
                sx={{
                  display: "flex",
                  px: ".5rem",
                  flexDirection: "row",
                  alignContent: "center",
                  alignItems: "center",
                  justifyContent: "center",
                  mt: "1rem",
                  mb: "1rem",
                }}
                key={index}
              >
                <Box sx={{ mr: "1rem" }}>
                  <img
                    src={"/api/v1/" + languages.iconPath}
                    width={50}
                    height={50}
                    alt={languages.name}
                    style={{ marginTop: "31px" }}
                  />
                </Box>
                <Box sx={{ width: "100%" }}>
                  <Typography sx={{}}>{languages.name}</Typography>
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
