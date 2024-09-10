import { Box, Card, CardContent, Typography } from "@mui/material";
import { CircularProgressStatistics } from "../progress/CircularProgressStatistics";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { getUserDevelopment } from "../../store/statistics/statisticsSlice";

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

  useEffect(() => {
    dispatch(getUserDevelopment());
  }, [dispatch]);

  const progresses = [
    {
      name: "Roads", //when de CicrularProgressStatistics.js is changed, this name should be changed too
      value: stateStatistics.data?.data?.roadPercentage,
    },
    {
      name: "Labs", //when de CicrularProgressStatistics.js is changed, this name should be changed too
      value: stateStatistics.data?.data?.labPercantage,
    },
  ];

  console.log(stateStatistics);

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
            textAlign: "center",
          }}
        >
          <Box sx={{ textAlign: "center" }}>
            <Typography variant="title">Development</Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              flexGrow: 1,
            }}
          >
            <CircularProgressStatistics progresses={progresses} />
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Development;
