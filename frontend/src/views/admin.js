import { Box, Card, Grid, Typography, Button } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import Translations from "src/components/Translations";
import Timestatistic from "src/components/cards/Timestatistic";
import Activity from "src/components/cards/Activity";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { getLanguageUsageRates, getSolitionWeek } from "src/store/log/logSlice";
import { useEffect } from "react";
import DonotProggresStatistic from "src/components/progress/DonotProggresStatistic";

const Admin = () => {
  const theme = useTheme();

  const router = useRouter();
  const dispatch = useDispatch();
  // const { log: logStatistics } = useSelector((state) => state);
  const logStatistics = useSelector((state) => state.log);

  const progresses = {
    values: logStatistics.data?.data?.map((item) => item.usagePercentage),
    labels: logStatistics.data?.data?.map((item) => item.name),
  };

  const backgroundColors = [
    theme.palette.primary.light,
    theme.palette.info.dark,
  ];

  // const Deneme = {
  //   values: [10.5, 20.764674, 30, 40],
  //   labels: ["a", "b", "c", "d"],
  // };
  useEffect(() => {
    dispatch(getLanguageUsageRates());
  }, [dispatch]);

  return (
    <Grid container spacing={2}>
      <Grid
        item
        xs={12}
        sx={{ display: "flex", flexDirection: "column", gap: 2 }}
      >
        {/* top */}
        <Grid item xs={12} md={5} sx={{ alignSelf: "flex-start" }}>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              gap: "1rem",
              height: "100%",
            }}
          >
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center",
                flexDirection: "row",
                gap: "16px",
              }}
            >
              <Button
                variant="dark"
                onClick={() => router.push("admin/users")}
                sx={{
                  textTransform: "none",
                  width: "100%",
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                  gap: "16px",
                  borderRadius: "10px",
                }}
              >
                <Image src={userIcon} alt="userIcon" width={24} height={24} />
                <Typography
                  variant="infoText"
                  sx={{
                    color: `${theme.palette.common.white} !important`,
                    fontWeight: "normal",
                  }}
                >
                  <Translations text="admin.profile.button" />
                </Typography>
              </Button>
            </Box>
          </Box>
        </Grid>

        {/* bottom */}
        <Grid item xs={12} md={12}>
          {/* <Box sx={{ minHeight: "500px" }}> */}
          <Timestatistic />
          {/* </Box> */}
        </Grid>
      </Grid>
      {/* Activity Card */}
      <Grid item xs={12}>
        <Box sx={{ height: "auto", py: 2 }}>
          <Activity />
        </Box>
      </Grid>
    </Grid>
  );
};

export default Admin;
