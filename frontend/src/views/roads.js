import { Box, Card, Grid } from "@mui/material";
import React from "react";
import { useState } from "react";
import { roads, languages } from "src/data/home";
import InfoCard from "src/components/cards/Info";
import LanguageProgress from "src/components/cards/LanguageProgress";
import ProgressStatuses from "src/components/filter/ProgressStatuses";
import SearchFilter from "src/components/filter/SearchFilter";
import SortFilter from "src/components/filter/SortFilter";

const Roads = () => {
  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  });
  return (
    <>
      <Grid container spacing={2} gap={2} sx={{ px: "1rem" }}>
        <Grid item container xs={12} spacing={4} sx={{ pt: "0px !important" }}>
          <Grid item xs={12} md={8}>
            <InfoCard {...roads} sx={{ height: "212px" }} />
          </Grid>
          <Grid item xs={12} md={4}>
            <Card sx={{ height: "212px" }}></Card>
          </Grid>
        </Grid>
        
        <Grid item container xs={12} spacing={2} justifyContent="center" 
  alignItems="center"  >
  <Grid item xs={12} md={4}   >
    <ProgressStatuses filters={filters} setFilters={setFilters} />
  </Grid>
  <Grid item xs={12} md={5} sm={6}>
    <SearchFilter searchKey="roads.search.placeholder" />
  </Grid>
  <Grid item xs={12} md={3} sm={6}>
    <SortFilter filters={filters} setFilters={setFilters} textKey="roads.sort_the_labs" />
  </Grid>
</Grid>
      
<Grid item container xs={12} md={12} spacing={2} sx={{ maxHeight: 'calc(100vh - 143px)',  pt: '0px !important' }}>
          {languages.map((language, index) => (
            <Grid item xs={12} md={12} key={index}>
              <LanguageProgress language={language} />
            </Grid>
          ))}
          <Box sx={{ width: '100%', height: '2px' }}></Box>
        </Grid>
      </Grid>
    </>
  );
};

export default Roads;
