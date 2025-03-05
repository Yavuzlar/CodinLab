import {
  Box,
  Card,
  CardContent,
  Typography,
  useMediaQuery,
} from "@mui/material";
import GraphicalStatistics from "../graphical-statistics/GraphicalStatistics";
import { useTheme } from "@emotion/react";
import Translations from "../Translations";
import DateRangeIcon from "@mui/icons-material/DateRange";
import { useDispatch, useSelector } from "react-redux";
import { getSolitionWeek } from "src/store/log/logSlice";
import { useEffect, useState } from "react";

const Timestatistic = () => {
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  const today = new Date();
  const startOfWeek = new Date(
    today.setDate(today.getDate() - today.getDay() + 1)
  );
  const endOfWeek = new Date(today.setDate(today.getDate() + 6));

  const options = { day: "numeric", month: "long" };
  const startOfWeekFormatted = startOfWeek.toLocaleDateString("en-GB", options);
  const endOfWeekFormatted = endOfWeek.toLocaleDateString("en-GB", options);

  useEffect(() => {
    setStartDate(startOfWeekFormatted);
    setEndDate(endOfWeekFormatted);
  }, [startOfWeekFormatted, endOfWeekFormatted]);

  const dispatch = useDispatch();

  const { log: logStatistics } = useSelector((state) => state);

  useEffect(() => {
    dispatch(getSolitionWeek());
  }, [dispatch]);

  const _smd = useMediaQuery((theme) => theme.breakpoints.down("smd"));
  const theme = useTheme();

  return (
    <Box
      sx={{
        display: "flex",
        width: "100%",
        height: "100%",
      }}
    >
      <Card
        sx={{
          width: "100%",
          minHeight: "12.5rem",
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between",
        }}
      >
        <CardContent sx={{ height: "calc(100% - 3rem)" }}>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              // height: "100%",
              marginBottom: "1rem",
            }}
          >
            <Typography variant="title" sx={{ fontWeight: "bold" }}>
              <Translations text="admin.time.title" />
            </Typography>
            <Box sx={{ display: "flex", justifyContent: "space-between" }}>
              <Box>
                <Typography
                  sx={{
                    maxWidth: "calc(100% - 9.625rem)",

                    paddingTop: "0.8rem",
                    ...(_smd && { maxWidth: "60ch" }),
                  }}
                >
                  <Translations text="admin.time.content" />
                </Typography>
              </Box>
            </Box>
            {/* Week */}
            <Box
              sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "end",
              }}
            >
              <Box sx={{ display: "flex" }}>
                <Box sx={{ mr: "5px" }}>
                  <DateRangeIcon />
                </Box>
                <Box>
                  <Typography
                    sx={{ font: "normal normal bold 18px/23px Outfit;" }}
                  >
                    Week
                  </Typography>
                  <Typography
                    sx={{
                      font: "normal normal normal 16px/20px Outfit;",
                    }}
                  >
                    {startDate} - {endDate}
                  </Typography>
                </Box>
              </Box>
            </Box>
          </Box>
          <Box
            sx={{
              display: "flex",
              flexDirection: "row",
              alignItems: "center",
              justifyContent: "space-between",
            }}
          >
            <Box
              sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                gap: 15,
                mr: "3rem",
              }}
            ></Box>
            <Box sx={{ flexGrow: 1 }}>
              <GraphicalStatistics data={logStatistics.weekData.data} />
            </Box>
          </Box>
          <Box
            sx={{
              display: "flex",
              gap: 2,
              alignItems: "center",
              justifyContent: "center",
            }}
          >
            {/* Labs and Roads */}
            <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
              <Box
                sx={{
                  width: "15px",
                  height: "15px",
                  backgroundColor: theme.palette.primary.light,
                  borderRadius: "50%",
                }}
              />
              <Typography>Labs</Typography>
            </Box>
            <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
              <Box
                sx={{
                  width: "15px",
                  height: "15px",
                  backgroundColor: theme.palette.primary.dark,
                  borderRadius: "50%",
                }}
              />
              <Typography>Roads</Typography>
            </Box>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Timestatistic;
