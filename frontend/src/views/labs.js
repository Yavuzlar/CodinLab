import { Box, Card, Grid, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import InfoCard from "src/components/cards/Info";
import LanguageProgress from "src/components/cards/LanguageProgress";
import Filter from "src/components/filter/Filter";
import { labs} from "src/data/home";
import { useTranslation } from "react-i18next";
import { useDispatch, useSelector } from "react-redux";
import { getUserLanguageLabStats } from "src/store/language/languageSlice";

const Labs = () => {
  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });
  const { t } = useTranslation();
  const searchPlaceholder = t("labs.search.placeholder");

  const dispatch = useDispatch();
  const { language: stateLanguage } = useSelector((state) => state);

  useEffect(() => {
    dispatch(getUserLanguageLabStats());
    console.log("stateLanguage", stateLanguage);
  }, [dispatch]);

  console.log("stateLanguage", stateLanguage.userLanguageLabStatsData.data);

  const labsStatsData =[
    {
      id: 1,
      totalLabs: stateLanguage.userLanguageLabStatsData.data?.totalLabs,
      completedLabs: stateLanguage.userLanguageLabStatsData.data?.completedLabs,
      percentage: stateLanguage.userLanguageLabStatsData.data?.percentage,
    },
  ]

  return (
    <>
      <Grid container spacing={2}>
        <Grid item container xs={12} md={7} sx={{ pt: "0px !important" }}>
          <Grid item xs={12}>
            <Box
              sx={{
                display: "flex",
                gap: "1rem",
                flexDirection: "column",
                height: "100%",
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
                  <Card sx={{ height: "100%" }}></Card>
                </Grid>

                <Grid item xs={12} md={6}>
                  <Card sx={{ height: "100%" }}>
                    <Typography variant="h6"></Typography>
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
            pt: "0px !important",
          }}
        >
          {labsStatsData.map((language, index) => (
            <Grid item xs={12} md={12} key={index}>
              <LanguageProgress
                language={language}
                type="lab"
              />
            </Grid>
          ))}
          <Box sx={{ width: "100%", height: "2px" }}></Box>
        </Grid>
      </Grid>
    </>
  );
};

export default Labs;
