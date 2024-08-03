import { Box, Card, Grid, Typography,Button } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import {time, activity } from "src/data/admin";
import Translations from "src/components/Translations";
import Timestatistic from "src/components/cards/Timestatistic";
import Activity from "src/components/cards/Activity";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import settingsIcon from "../assets/icons/icons8-settings-128.png";




const Admin = () => {
  const theme = useTheme();
  return (
    <>
     <Grid container spacing={2}>
      {/* left */}
     <Grid item container xs={12} md={7} sx={{ pt: '0px !important' }}>
          <Grid item xs={12}>
            <Box sx={{ display: 'flex', gap: '1rem', flexDirection: 'column', height: '100%' }}>
            <Box sx={{height:'450px'}}>
                <Timestatistic {...time}/>
            </Box>
            <Box sx={{height:'450px'}}>
                <Activity {...activity}/>
            </Box>

            </Box>
          </Grid>
        </Grid>


{/* right */}
        <Grid item container xs={12} md={5} spacing={2} sx={{ pt: '0px !important' }}>
          <Grid item xs={12}>
         <Card sx={{  width: "100%",height:"660px" }}>
                <Typography
                  variant="h5"
                  sx={{
                    fontWeight: "bold",
                    padding: "20px",
                    textAlign: "center",
                  }}
                >
                  <Translations text="admin.language.rates" />
                </Typography>
              </Card>

              <Card sx={{  width: "100%",height:"235px",marginTop :"20px" }}>
                <Typography
                  variant="h5"
                  sx={{
                    fontWeight: "bold",
                    padding: "30px",
                    textAlign: "center",
                  }}
                >
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
      
    </>
  );
};

export default Admin;
