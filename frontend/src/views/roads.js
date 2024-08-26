import { Box, Card, Grid, useMediaQuery } from "@mui/material";
import React, { use, useTransition } from "react";
import { useState } from "react";
import { roads, languages } from "src/data/home";
import InfoCard from "src/components/cards/Info";
import LanguageProgress from "src/components/cards/LanguageProgress";
import roadsIcon from "src/assets/icons/icons8-path-100.png";
import Filter from "src/components/filter/Filter";
import Translations from "src/components/Translations";
import { useTranslation } from "react-i18next";


const Roads = () => {
  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });

  const { t } = useTranslation();
  const searchPlaceholder = t("roads.search.placeholder")

  return (
    <>
      <Grid container spacing={2} gap={2}>
        <Grid item container xs={12} spacing={4} sx={{ pt: "0px !important" }}>
          <Grid item xs={12} md={8}>
            <InfoCard {...roads} sx={{ height: "212px" }} />
          </Grid>
          <Grid item xs={12} md={4}>
            <Card sx={{ height: "212px" }}></Card>
          </Grid>
        </Grid>

        <Grid
          item
          container
          xs={12}
        >
          <Grid item xs={12}>
            <Filter filters={filters} setFilters={setFilters} searchPlaceholder={searchPlaceholder}  />
          </Grid>
        </Grid>

        <Grid
          item
          container
          xs={12}
          spacing={2}
          sx={{ maxHeight: "calc(100vh - 143px)", pt: "0px !important" }}
        >
          {languages.map((language, index) => (
            <Grid item xs={12} md={12} key={index}>
              <LanguageProgress
                language={language}
                icon={roadsIcon}
                map={"20/40 Path"}
              />
            </Grid>
          ))}
          <Box sx={{ width: "100%", height: "2px" }}></Box>
        </Grid>
      </Grid>
    </>
  );
};

export default Roads;
