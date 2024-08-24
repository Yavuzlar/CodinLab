import { Grid } from "@mui/material";
import Welcome from "src/components/cards/Welcome";
import Languages from "src/components/cards/Languages";
import InfoCard from "src/components/cards/Info";
import { welcomeCard, languages, roads, labs } from "src/data/home";
import Level from "src/components/cards/Level";
import Development from "src/components/cards/Development";
import Advancement from "src/components/cards/Advancement";
import LevelStatistic from "src/components/statistics/LevelStatistic";

const Home = () => {
  const progresses = [
    {
      name: "Easy",
      value: 50,
      color: "#39CE19",
    },
    {
      name: "Medium",
      value: 25,
      color: "#EE7A19",
    },
    {
      name: "Hard",
      value: 45,
      color: "#DC0101",
    },
  ];
  return (
    <>
      <Grid container spacing={4} sx={{ px: "1rem" }}>
        <Grid item xs={12}>
          <Welcome {...welcomeCard} />
        </Grid>
        <Grid
          item
          xs={12}
          container
          spacing={4}
          sx={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "row",
          }}
        >
          {languages.map((language, index) => (
            <Grid item xs={12} md={4} xl={2.4} key={index}>
              <Languages language={language} />
            </Grid>
          ))}
        </Grid>
        <Grid item xs={12} md={6}>
          <InfoCard {...roads} />
        </Grid>
        <Grid item xs={12} md={6}>
          <InfoCard {...labs} />
        </Grid>
        <Grid
          item
          container
          xs={12}
          spacing={4}
          sx={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "row",
          }}
        >
          <Grid item xs={12} md={6} xl={4}>
            <LevelStatistic levels={1} progress={90} />
          </Grid>
          <Grid item xs={12} md={6} xl={4}>
            <Development />
          </Grid>
          <Grid item xs={12} md={6} xl={4}>
            <Advancement />
          </Grid>
        </Grid>
      </Grid>
    </>
  );
};

export default Home;
