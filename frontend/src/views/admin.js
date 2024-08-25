import { Box, Card, Grid, Typography, Button } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import Translations from "src/components/Translations";
import Timestatistic from "src/components/cards/Timestatistic";
import Activity from "src/components/cards/Activity";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import settingsIcon from "../assets/icons/icons8-settings-128.png";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";
import cImg from "../assets/icons/c.png";
import cppImg from "../assets/icons/cpp.png";
import goImg from "../assets/icons/go.png";

const languageStatistics = [
  {
    image: cImg,
    process: "50",
  },
  {
    image: cppImg,
    process: "90",
  },
  {
    image: goImg,
    process: "80",
  },
];
const Admin = () => {
  const theme = useTheme();

  let Deneme = [
    {
      name: "In progress", // String
      value: 90, // Number
      color: "#0A3B7A", // String
    },
  ];

  return (
    <>
      <Grid container spacing={2}>
        {/* left */}
        <Grid item xs={12} md={7}>
          <Box
            sx={{
              display: "flex",
              gap: "1rem",
              flexDirection: "column",
              height: "100%",
            }}
          >
            <Box sx={{ height: "470px" }}>
              <Timestatistic />
            </Box>
          </Box>
        </Grid>

        {/* right */}
        <Grid item container xs={12} md={5}>
          <Grid item xs={12}>
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center",
                flexDirection: "row",
                mb: "0.5rem",
                gap: '16px'
              }}
            >
              <Button
                variant="dark"
                sx={{
                  textTransform: "none",
                  width: '100%',
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
          </Grid>
          <Grid item xs={12} md={12}>
            <Box
              sx={{
                height: "422px",
              }}
            >
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
                    justifyContent: "center",
                  }}
                >
                  <CircularProgressStatistics progresses={Deneme} />
                </Box>
                {languageStatistics.map((item, index) => (
                  <Box
                    sx={{
                      mt: "0.5rem",
                      display: "flex",
                      alignItems: "center",
                      flexWrap: "",
                      gap: 5,
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
                    <Image src={item.image} width={15} height={15} />
                    <Typography
                      sx={{ font: "normal normal normal 18px/23px Outfit" }}
                    >
                      %{item.process}
                    </Typography>
                  </Box>
                ))}
              </Card>
            </Box>
          </Grid>
        </Grid>
        <Grid item xs={12}>
          <Box sx={{ height: "450px" }}>
            <Activity />
          </Box>
        </Grid>
      </Grid>
    </>
  );
};

export default Admin;
