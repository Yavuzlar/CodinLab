import { Box, Card, CardContent, Typography } from "@mui/material";
import { CircularProgressStatistics } from "../progress/CircularProgressStatistics";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { getUserDevelopment } from "../../store/statistics/statisticsSlice";
import { useTranslation } from "react-i18next";
import { theme } from "src/configs/theme";

const Development = () => {
  const dispatch = useDispatch();
  const developmentData = useSelector((state) => state.statistics.developmentData);
  const { t } = useTranslation();

  useEffect(() => {
    dispatch(getUserDevelopment());
  }, [dispatch]);

  const progresses = [
    {
      name: t("home.development.roads"),
      value: developmentData?.data?.roadPercentage,
      color: theme.palette.primary.dark,
    },
    {
      name: t("home.development.labs"),
      value: developmentData?.data?.labPercantage,
      color: theme.palette.primary.light,
    },
  ];

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
            height: "calc(100% - 1.5rem)",
          }}
        >
          <Box sx={{ textAlign: "center" }}>
            <Typography variant="title">
              {t("home.development.title")}
            </Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              flexGrow: 1,
            }}
          >
            <CircularProgressStatistics
              progresses={progresses}
              flexDirection={"column"}
            />
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Development;
