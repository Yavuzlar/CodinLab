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

  const userLanguageLabStatsData = useSelector((state) => state.language.userLanguageLabStatsData);
  const difficultyStatsData = useSelector((state) => state.statistics.difficultyStatsData);
  const labsProgressStatsData = useSelector((state) => state.statistics.labsProgressStatsData);

  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });

  const labsStatsData = userLanguageLabStatsData?.data;

  const difficultyStats = [
    {
      id: 1,
      name: t("labs.difficulty.easy"),
      value: difficultyStatsData.data?.easyPercentage,
      color: theme.palette.difficulty.easy,
    },
    {
      id: 2,
      name: t("labs.difficulty.medium"),
      value: difficultyStatsData.data?.mediumPercentage,
      color: theme.palette.difficulty.medium,
    },
    {
      id: 3,
      name: t("labs.difficulty.hard"),
      value: difficultyStatsData.data?.hardPercentage,
      color: theme.palette.difficulty.hard,
    },
  ];

  const labsProgressStats = [
    {
      id: 1,
      name: t("labs.progress.stats.progress"),
      value: labsProgressStatsData.data?.progress,
      color: theme.palette.primary.dark,
    },
    {
      id: 2,
      name: t("labs.progress.stats.completed"),
      value: labsProgressStatsData.data?.completed,
      color: theme.palette.primary.light,
    },
  ];

  useEffect(() => {
    dispatch(getUserLanguageLabStats());
    dispatch(getDifficultyStatistics());
    dispatch(getLabsProgressStats());
  }, [dispatch]);

  const { status, search, sort } = filters;
  const filterLabs = () => {
    let filteredLabs = labsStatsData;

    switch (status) {
      case "completed":
        filteredLabs = filteredLabs.filter(
          (lab) => lab.completedLabs == lab.totalLabs
        );
        break;
      case "in-progress":
        filteredLabs = filteredLabs.filter((lab) => lab.completedLabs > 0);
        break;
    }

    switch (sort) {
      case "desc":
        filteredLabs = [...filteredLabs].sort((a, b) => b.id - a.id);
        break;
      case "asc":
        filteredLabs = [...filteredLabs].sort((a, b) => a.id - b.id);
        break;
    }

    if (search != "") {
      filteredLabs = filteredLabs.filter(
        (road) =>
          road.name && road.name.toLowerCase().includes(search.toLowerCase())
      );
    }
    return filteredLabs;
  };

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
        sx={{
          maxHeight: "calc(100vh - 143px)",
          overflow: "auto",
        }}
      >
        {filterLabs()?.map((language, index) => (
          <Grid item mb={2} xs={12} key={index}>
            <LanguageProgress language={language} type="lab" />
          </Grid>
        ))}
        <Box sx={{ width: "100%", height: "2px" }} />
      </Grid>
    </Grid>
  );
};

export default Labs;
