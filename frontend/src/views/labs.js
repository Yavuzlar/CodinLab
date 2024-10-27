import { Box, Card, CardContent, Grid } from "@mui/material";
import { useEffect, useState } from "react";
import InfoCard from "src/components/cards/Info";
import LanguageProgress from "src/components/cards/LanguageProgress";
import Filter from "src/components/filter/Filter";
import { labs } from "src/data/home";
import { useTranslation } from "react-i18next";
import { useDispatch, useSelector } from "react-redux";
import { getUserLanguageLabStats } from "src/store/language/languageSlice";
import {
  getDifficultyStatistics,
  getLabsProgressStats,
} from "src/store/statistics/statisticsSlice";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";
import { theme } from "src/configs/theme";

const Labs = () => {
  const { t } = useTranslation();
  const dispatch = useDispatch();

  const searchPlaceholder = t("labs.search.placeholder");

  const { language: stateLanguage, statistics: stateStatistics } = useSelector(
    (state) => state
  );

  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });

  const labsStatsData = stateLanguage.userLanguageLabStatsData?.data;

  const difficultyStats = [
    {
      id: 1,
      name: t("labs.difficulty.easy"),
      value: stateStatistics.difficultyStatsData.data?.easyPercentage,
      color: theme.palette.difficulty.easy,
    },
    {
      id: 2,
      name: t("labs.difficulty.medium"),
      value: stateStatistics.difficultyStatsData.data?.mediumPercentage,
      color: theme.palette.difficulty.medium,
    },
    {
      id: 3,
      name: t("labs.difficulty.hard"),
      value: stateStatistics.difficultyStatsData.data?.hardPercentage,
      color: theme.palette.difficulty.hard,
    },
  ];

  const labsProgressStats = [
    {
      id: 1,
      name: t("labs.progress.stats.progress"),
      value: stateStatistics.labsProgressStatsData.data?.progress,
      color: theme.palette.primary.dark,
    },
    {
      id: 2,
      name: t("labs.progress.stats.completed"),
      value: stateStatistics.labsProgressStatsData.data?.completed,
      color: theme.palette.primary.light,
    },
  ];

  useEffect(() => {
    dispatch(getUserLanguageLabStats());
    dispatch(getDifficultyStatistics());
    dispatch(getLabsProgressStats());
  }, [dispatch]);

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
                <Card sx={{ width: "100%", height: "100%" }}>
                  <CardContent>
                    <Box
                      sx={{
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",
                      }}
                    >
                      <CircularProgressStatistics
                        progresses={labsProgressStats}
                        flexDirection={"column"}
                      />
                    </Box>
                  </CardContent>
                </Card>
              </Grid>

              <Grid item xs={12} md={6}>
                <Card sx={{ width: "100%", height: "100%" }}>
                  <CardContent>
                    <Box
                      sx={{
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",
                      }}
                    >
                      <CircularProgressStatistics
                        progresses={difficultyStats}
                        flexDirection={"column"}
                      />
                    </Box>
                  </CardContent>
                </Card>
              </Grid>
            </Grid>
          </Box>
        </Grid>
      </Grid>

      <Grid
        item
        md={5}
        xs={12}
        e sx={{
          maxHeight: "calc(100vh - 143px)",
          overflow: "auto",
        }}
      >
        {labsStatsData?.map((language, index) => (
          <Grid item mb={2} xs={12} key={index} >
            <LanguageProgress language={language} type="lab" />
          </Grid>
        ))}
        <Box sx={{ width: "100%", height: "2px" }} />
      </Grid>
    </Grid>
  );
};

export default Labs;
