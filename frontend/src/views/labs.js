import { Box, Card, Grid } from "@mui/material";
import { useState } from "react";
import InfoCard from "src/components/cards/Info";
import LanguageProgress from "src/components/cards/LanguageProgress";
import Filter from "src/components/filter/Filter";
import { labs, languages } from "src/data/home";
import labsIcon from "src/assets/icons/icons8-test-tube-100.png";
import { useTranslation } from "react-i18next";

const Labs = () => {
  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });
  const { t } = useTranslation();
  const searchPlaceholder = t("labs.search.placeholder")

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
                <Filter filters={filters} setFilters={setFilters} searchPlaceholder={searchPlaceholder} />
              </Box>


              <InfoCard {...labs} />

              <Grid container spacing={2} sx={{ height: "100%" }}>
                <Grid item xs={12} md={6}>
                  <Card sx={{ height: "100%" }}></Card>
                </Grid>

                <Grid item xs={12} md={6}>
                  <Card sx={{ height: "100%" }}></Card>
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
          {languages.map((language, index) => (
            <Grid item xs={12} md={12} key={index}>
              <LanguageProgress language={language} icon={labsIcon} map={"20 Labs"} />
            </Grid>
          ))}
          <Box sx={{ width: "100%", height: "2px" }}></Box>
        </Grid>
      </Grid>
    </>
  );
};

export default Labs;
