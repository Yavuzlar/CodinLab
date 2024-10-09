import {
  Box,
  Card,
  CardContent,
  Typography,
  useMediaQuery,
} from "@mui/material";
import { CircularProgressStatistics } from "../progress/CircularProgressStatistics";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { getUserDevelopment } from "../../store/statistics/statisticsSlice";
import { useTranslation } from "react-i18next";

const Development = () => {
  // let Deneme = [
  //   {
  //     name: "In progress", // String
  //     value: 90, // Number
  //     color: "#0A3B7A", // String
  //   },
  // ];

  const dispatch = useDispatch();
  const { statistics: stateStatistics } = useSelector((state) => state);
  const { t } = useTranslation();
  const _sm = useMediaQuery((theme) => theme.breakpoints.down("sm"));

  useEffect(() => {
    dispatch(getUserDevelopment());
  }, [dispatch]);

  const progresses = [
    {
      name: t("home.development.roads"), //when de CicrularProgressStatistics.js is changed, this name should be changed too
      value: stateStatistics.data?.data?.roadPercentage,
    },
    {
      name: t("home.development.labs"), //when de CicrularProgressStatistics.js is changed, this name should be changed too
      value: stateStatistics.data?.data?.labPercantage,
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
              justifyContent: "center",
              flexGrow: 1,
            }}
          >
            <CircularProgressStatistics
              progresses={progresses}
              flexDirection={_sm ? "column" : "row"}
            />
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Development;
