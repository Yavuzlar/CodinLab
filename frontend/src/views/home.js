import { Grid } from "@mui/material";
import Welcome from "src/components/cards/Welcome";
import Languages from "src/components/cards/Languages";
import InfoCard from "src/components/cards/Info";
import { welcomeCard, roads, labs } from "src/data/home";
import Development from "src/components/cards/Development";
import Advancement from "src/components/cards/Advancement";
import LevelStatistic from "src/components/statistics/LevelStatistic";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { getInventories } from "src/store/language/languageSlice";

const Home = () => {

  const dispatch = useDispatch();
  const { 
    language : stateLanguage,
   } = useSelector((state) => state);

   useEffect(() => {
    dispatch(getInventories());
  }, [dispatch]);

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
          {stateLanguage.data?.data?.map((language, index) => (
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
            <LevelStatistic levels={1} progress={20} />
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
