import { Box, Card, Grid, Typography, Button } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import Translations from "src/components/Translations";
import Timestatistic from "src/components/cards/Timestatistic";
import Activity from "src/components/cards/Activity";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import settingsIcon from "../assets/icons/icons8-settings-128.png";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { getLanguageUsageRates } from "src/store/log/logSlice";
import { useEffect } from "react";



const Admin = () => {
  const theme = useTheme();
  const dispatch = useDispatch();

  const {  log: logStatistics } = useSelector(
    (state) => state
  );

  useEffect(() => {
    dispatch(getLanguageUsageRates());
  }, [dispatch]);

  console.log(logStatistics);

  const progresses = logStatistics?.data?.data?.map((item) => ({
    name: item.name,
    value: item.usagePercentage,
  })) || [];


  const router = useRouter();
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

              <Button
                onClick={() => router.push("admin/settings")}
                variant="dark"
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
                <Image
                  src={settingsIcon}
                  alt="settingsIcon"
                  width={26}
                  height={26}
                />
                <Typography
                  variant="infoText"
                  sx={{
                    color: `${theme.palette.common.white} !important`,
                    fontWeight: "normal",
                  }}
                >
                  <Translations text="admin.settings.button" />
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
                  <CircularProgressStatistics progresses={progresses} />
                </Box>
                {logStatistics.data?.data?.map((item, index) => (
                  <Box
                    sx={{
                      mt: "0.5rem",
                      display: "flex",
                      alignItems: "center",
                      flexWrap: "",
                      gap: 15,
                    }}
                    key={index}
                  >
                    <Box
                      sx={{
                        width: "15px",
                        height: "15px",
                        backgroundColor: theme.palette.primary.light,
                        borderRadius: "50%",
                      }}
                    />
                    <img src={"api/v1/"+item.iconPath} width={30} height={30} />
                    <Typography
                      sx={{ font: "normal normal normal 18px/23px Outfit" }}
                    >
                      %{item.usagePercentage}
                    </Typography>
                  </Box>
                ))}
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
