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
  const { t } = useTranslation();
  const dispatch = useDispatch();

  const searchPlaceholder = t("roads.search.placeholder");
  const { language: stateLanguage, statistics: stateStatistics } = useSelector(
    (state) => state
  );

  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });

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

  useEffect(() => {
    dispatch(getUserLanguageRoadStats());
    dispatch(getRoadProgressStats());
  }, [dispatch]);

  // const Deneme = [
  //   {
  //     id: 1,
  //     name: "React",
  //     value: 20.2452,
  //     color: "#8FDDFD",
  //   },
  //   {
  //     id: 2,
  //     name: "Python",
  //     value: 40,
  //     color: "#0A3B7A",
  //   },
  // ];
  return (
    <>
      <Grid container spacing={2} gap={2}>
        <Grid item container xs={12} spacing={4} sx={{ pt: "0px !important" }}>
          <Grid item xs={12} md={7}>
            <InfoCard {...roads} />
          </Grid>

          <Grid item xs={12} md={5}>
            <Card sx={{ height: "18.5rem" }}>
              <CardContent
                sx={{
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <CircularProgressStatistics
                  progresses={roadProgressStats}
                  flexDirection={"column"}
                />
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
              <Grid item xs={12} md={12} key={index} sx={{ cursor: "pointer" }}>
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
