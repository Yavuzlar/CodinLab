import { Grid } from "@mui/material";
import Welcome from "src/components/cards/Welcome";
import Languages from "src/components/cards/Languages";
import InfoCard from "src/components/cards/Info";
import { welcomeCard, languages, roads, labs } from "src/data/home";

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
      <Grid container spacing={4}>
        <Grid item xs={12}>
          <Welcome {...welcomeCard} />
        </Grid>
        <Grid item xs={12} container spacing={4}   sx={{
              display: "flex", justifyContent: "center", alignItems: "center" , flexDirection: "row"
            }}>
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
      </Grid>
    </>
  );
};

export default Home;
