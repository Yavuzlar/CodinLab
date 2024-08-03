import { Box, Button, Card, Grid, Typography } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import settingsIcon from "../assets/icons/icons8-settings-128.png";
import InfoCard from "src/components/cards/Info";
import { time, activity } from "src/data/admin";
import Translations from "src/components/Translations";

const Admin = () => {
  const theme = useTheme();
  return (
    <>
      <Grid container spacing={2}>
        <Grid item xs={12} md={8} sx={{ pt: "0px !important" }}>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <Card
                sx={{
                  height: "290px",
                  width: "100%",
                }}
              >
                <InfoCard {...time} />
              </Card>
            </Grid>
            <Grid item xs={12}>
              <Card
                sx={{
                  height: "290px",
                  width: "100%",
                }}
              >
                <InfoCard {...activity} />
              </Card>
            </Grid>
          </Grid>
        </Grid>

        <Grid item xs={12} md={4} sx={{ pt: "0px !important" }}>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <Card sx={{ height: "420px", width: "100%" }}>
                <Typography
                  variant="h5"
                  sx={{
                    fontWeight: "bold",
                    padding: "20px",
                    textAlign: "center",
                  }}
                >
                  <Translations text="admin.center.title" />
                </Typography>
              </Card>
            </Grid>
            <Grid item xs={12}>
              <Card
                sx={{
                  height: "160px",
                  width: "100%",
                  boxSizing: "border-box",
                  padding: "10px",
                  textAlign: "center",
                }}
              >
                <Typography variant="h5" sx={{ fontWeight: "bold" }}>
                  <Translations text="admin.center.title" />
                </Typography>
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                    flexDirection: "column",
                    gap: "10px",
                    marginTop: "20px",
                  }}
                >
                  <Button
                    variant="dark"
                    sx={{
                      textTransform: "none",
                      width: "250px",
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                      gap: "10px",
                      borderRadius: "10px",
                    }}
                  >
                    <Image
                      src={userIcon}
                      alt="userIcon"
                      width={24}
                      height={24}
                    />
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
                    variant="dark"
                    sx={{
                      textTransform: "none",
                      width: "250px",
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                      gap: "10px",
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
              </Card>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
};

export default Admin;
