import { Box, Card, Grid, Typography, Button } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import Translations from "src/components/Translations";
import Timestatistic from "src/components/cards/Timestatistic";
import Activity from "src/components/cards/Activity";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import settingsIcon from "../assets/icons/icons8-settings-128.png";
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { getLanguageUsageRates, getSolitionWeek } from "src/store/log/logSlice";
import { useEffect } from "react";
import DonotProggresStatistic from "src/components/progress/DonotProggresStatistic";
import CustomBreadcrumbs from "src/components/breadcrumbs";

const Admin = () => {
  const theme = useTheme();

  const router = useRouter();
  const dispatch = useDispatch();
  const { log: logStatistics } = useSelector((state) => state);

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
      <Grid item xs={12} sx={{ display: "flex", gap: 2 }}>
        {/* left */}
        <Grid item xs={12} md={7}>
          {/* <Box sx={{ minHeight: "500px" }}> */}
          <Timestatistic />
          {/* </Box> */}
        </Grid>

        {/* right */}
        <Grid item xs={12} md={5}>
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

            <Card
              sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                width: "100%",
                height: "100%",
              }}
            >
              <Typography
                variant="title"
                sx={{
                  fontWeight: "bold",
                  padding: "20px",
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                  textAlign: "center",
                }}
              >
                <Translations text="admin.language.rates" />
              </Typography>
              <Box
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  gap: 0.25,
                }}
              >
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "center",
                  }}
                >
                  {/* <CircularProgressStatistics progresses={progresses} /> */}
                  <DonotProggresStatistic data={progresses} />
                </Box>
                <Box
                  sx={{
                    mt: "0.5rem",
                    display: "flex",
                    flexWrap: "wrap",
                    gap: "2rem",
                  }}
                >
                  {logStatistics.data?.data?.map((item, index) => (
                    <Box
                      sx={{
                        mt: "0.5rem",
                        display: "flex",
                        alignItems: "center",
                        gap: "1rem",
                      }}
                      key={index}
                    >
                      <Box
                        sx={{
                          width: "15px",
                          height: "15px",
                          backgroundColor: backgroundColors[index],
                          borderRadius: "50%",
                        }}
                      />
                      <Box
                        sx={{
                          display: "flex",
                          gap: "0.5rem",
                          alignItems: "center",
                        }}
                      >
                        <img
                          src={"api/v1/" + item.iconPath}
                          width={30}
                          height={30}
                        />
                        <Typography
                          sx={{ font: "normal normal normal 18px/23px Outfit" }}
                        >
                          %{Math.round(item.usagePercentage)}
                        </Typography>
                      </Box>
                    </Box>
                  ))}
                </Box>
              </Box>
            </Card>
          </Box>
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
