import { Box, Card, CardContent, Divider, Grid } from "@mui/material";
import { useEffect, useState } from "react";
import { roads } from "src/data/home";
import InfoCard from "src/components/cards/Info";
import LanguageProgress from "src/components/cards/LanguageProgress";
import Filter from "src/components/filter/Filter";
import { useTranslation } from "react-i18next";
import { useDispatch, useSelector } from "react-redux";
import { getUserLanguageRoadStats } from "src/store/language/languageSlice";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";
import { getRoadProgressStats } from "src/store/statistics/statisticsSlice";

const Roads = () => {
  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });

  const { t } = useTranslation();
  const searchPlaceholder = t("roads.search.placeholder");

  const dispatch = useDispatch();
  const { language: stateLanguage, statistics: stateStatistics } = useSelector(
    (state) => state
  );

  useEffect(() => {
    dispatch(getUserLanguageRoadStats());
    dispatch(getRoadProgressStats());
  }, [dispatch]);

  const roadProgressStats = [
    {
      id: 1,
      name: t("labs.progress.stats.progress"),
      value: stateStatistics.roadProgressStatsData.data?.progress,
      color: "#8FDDFD",
    },
    {
      id: 2,
      name: t("labs.progress.stats.completed"),
      value: stateStatistics.roadProgressStatsData.data?.completed,
      color: "#0A3B7A",
    },
  ];

  return (
    <>
      <Grid container spacing={2} gap={2}>
        <Grid item container xs={12} spacing={4} sx={{ pt: "0px !important" }}>
          <Grid item xs={12} md={8}>
            <InfoCard {...roads} sx={{ height: "212px" }} />
          </Grid>

          <Grid item xs={12} md={4}>
            <Card sx={{ height: "212px" }}>
              <CardContent
                sx={{
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",

                }}
              >
              <CircularProgressStatistics progresses={roadProgressStats} />
              </CardContent>
            </Card>
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Filter
            filters={filters}
            setFilters={setFilters}
            searchPlaceholder={searchPlaceholder}
          />
        </Grid>

        <Grid
          item
          xs={12}
          sx={{ p: "0 !important", justifyContent: "center", display: "flex" }}
        >
          <Box sx={{ width: "95%" }}>
            <Divider
              sx={{ borderColor: (theme) => theme.palette.primary.dark }}
            />
          </Box>
        </Grid>

        <Grid
          item
          container
          xs={12}
          spacing={2}
          sx={{ maxHeight: "calc(100vh - 143px)", pt: "0px !important" }}
        >
          {stateLanguage.userLanguageRoadStatsData?.data?.map(
            (language, index) => (
              <Grid item xs={12} md={12} key={index}>
                <LanguageProgress language={language} type="road" />
              </Grid>
            )
          )}
          <Box sx={{ width: "100%", height: "2px" }}></Box>
        </Grid>
      </Grid>
    </>
  );
};

export default Roads;
