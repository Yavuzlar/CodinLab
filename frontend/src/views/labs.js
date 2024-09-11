import { Box, Card, CardContent, Grid, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import InfoCard from "src/components/cards/Info";
import LanguageProgress from "src/components/cards/LanguageProgress";
import Filter from "src/components/filter/Filter";
import { labs } from "src/data/home";
import { useTranslation } from "react-i18next";
import { useDispatch, useSelector } from "react-redux";
import { getUserLanguageLabStats } from "src/store/language/languageSlice";
import { getDifficultyStatistics } from "src/store/statistics/statisticsSlice";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";

const Labs = () => {
  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });
  const { t } = useTranslation();
  const searchPlaceholder = t("labs.search.placeholder");

  const dispatch = useDispatch();
  const { language: stateLanguage, statistics: stateStatistics } = useSelector(
    (state) => state
  );

  useEffect(() => {
    dispatch(getUserLanguageLabStats());
    dispatch(getDifficultyStatistics());
  }, [dispatch]);

  const labsStatsData = [
    {
      id: 1,
      totalLabs: stateLanguage.userLanguageLabStatsData.data?.totalLabs,
      completedLabs: stateLanguage.userLanguageLabStatsData.data?.completedLabs,
      percentage: stateLanguage.userLanguageLabStatsData.data?.percentage,
    },
  ];

  const difficultyStats = [
    {
      id: 1,
      name: t("labs.difficulty.easy"),
      value: stateStatistics.difficultyStatsData.data?.easyPercentage,
    },
    {
      id: 2,
      name: t("labs.difficulty.medium"),
      value: stateStatistics.difficultyStatsData.data?.mediumPercentage,
    },
    {
      id: 3,
      name: t("labs.difficulty.hard"),
      value: stateStatistics.difficultyStatsData.data?.hardPercentage,
    },
  ];

  const progresses = [
    {
      name: "Easy", // when CircularProgressStatistics.js is changed, this name should be changed too
      value: stateStatistics.difficultyStatsData.data?.easyPercentage,
      color: "#39CE19"
    },
    {
      name: "Medium", // when CircularProgressStatistics.js is changed, this name should be changed too
      value: stateStatistics.difficultyStatsData.data?.mediumPercentage,
      color: "#EE7A19"

    },
    {
      name: "Hard", // when CircularProgressStatistics.js is changed, this name should be changed too
      value: stateStatistics.difficultyStatsData.data?.hardPercentage,
      color: "#DC0101"

    },

  ];

  return (
    <Grid container spacing={2}>
      <Grid item container xs={12} md={7}>
        <Grid item xs={12}>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              height: "100%",
              gap: "1rem",
            }}
          >
            <Box>
              <Filter
                filters={filters}
                setFilters={setFilters}
                searchPlaceholder={searchPlaceholder}
              />
            </Box>

            <InfoCard {...labs} />

            <Grid container spacing={2} sx={{ height: "100%" }}>
              <Grid item xs={12} md={6}>
                <Card sx={{ height: "100%" }} />
              </Grid>

              <Grid item xs={12} md={6}>
                <Card sx={{ width: "100%" }}>
                  <CardContent>
                    <Box
                      sx={{
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",
                      }}
                    >
                      <CircularProgressStatistics progresses={progresses} />
                    </Box>

                    {difficultyStats.map((difficulty) => (
                      <Box
                        key={difficulty.id}
                        sx={{
                          display: "flex",
                          mt: "25px",
                          alignItems: "center",
                          justifyContent: "center", 
                        }}
                      >
                        <Box
                          sx={{
                            width: "15px",
                            height: "15px",
                            backgroundColor:
                              difficulty.name === t("labs.difficulty.easy")
                                ? "#39CE19"
                                : difficulty.name === t("labs.difficulty.medium")
                                ? "#EE7A19"
                                : "#DC0101",
                            borderRadius: "50%",
                            marginRight: "0.5rem",
                          }}
                        />
                        <Box
                        sx={{
                          display: "flex",
                          alignItems: "center",
                          justifyContent: "center",
                          width: "104px",
                        }}
                        >
                          <Typography
                          sx={{
                          
                            textAlign: "left",
                            width: "100%",
                            
                          }}
                          >
                            {difficulty.name}
                          </Typography>
                          <Typography
                          sx={{
                            textAlign: "left",
                          }}
                          >
                            {difficulty.value}%
                          </Typography>
                        </Box>
                      </Box>
                    ))}
                  </CardContent>
                </Card>
              </Grid>
            </Grid>
          </Box>
        </Grid>
      </Grid>

      <Grid
        item
        container
        xs={12}
        md={5}
        spacing={2}
        sx={{
          maxHeight: "calc(100vh - 143px)",
          overflow: "auto",
        }}
      >
        {labsStatsData.map((language, index) => (
          <Grid item xs={12} key={index}>
            <LanguageProgress language={language} type="lab" />
          </Grid>
        ))}
        <Box sx={{ width: "100%", height: "2px" }} />
      </Grid>
    </Grid>
  );
};

export default Labs;
