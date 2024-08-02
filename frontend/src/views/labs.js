import { Grid } from "@mui/material";
import { useState } from "react";
import InfoCard from "src/components/cards/Info";
import Filter from "src/components/filter/Filter";
import { labs } from "src/data/home";

const Labs = () => {

  const [filters, setFilters] = useState({
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  })

  return (
    <>
      <Grid container>
        <Grid item container xs={12} md={7}>
          <Grid item xs={12}>
            <Filter filters={filters} setFilters={setFilters} />
          </Grid>

          <Grid item xs={12}>
            <InfoCard {...labs} />
          </Grid>
        </Grid>

        <Grid item container xs={12} md={5}>

        </Grid>
      </Grid>
    </>
  );
};

export default Labs;
