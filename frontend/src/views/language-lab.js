import { useState } from "react";
import { Box, Grid } from "@mui/material";
import Image from "next/image";
import C from "../assets/language/c.png";
import Cpp from "../assets/language/cpp.png";
import Go from "../assets/language/go.png";
import Js from "../assets/language/javascript.png";
import Python from "../assets/language/python.png";
import LabInfo from "../components/cards/LabInfo";
import SortFilter from "src/components/filter/SortFilter";
import SearchFilter from "src/components/filter/SearchFilter";
import PrgoressStatutesLabs from "src/components/filter/PrgoressStatutesLabs";

const languages = {
  c: { image: C, title: "C Language" },
  cpp: { image: Cpp, title: "C++ Language" },
  go: { image: Go, title: "Go Language" },
  javascript: { image: Js, title: "JavaScript" },
  python: { image: Python, title: "Python" },
};

const LanguageLab = ({ language = "" }) => {
  const [filters, setFilters] = useState({
    status: "all",
    difficulty: "all",
    search: "",
    sort: "",
  });

  const selectedLanguage = languages[language.toLowerCase()];

  return (
    <div>
      <Grid container spacing={2} direction="column">
        <Grid item xs={12}>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              gap: "1rem",
              flexDirection: { xs: "column", md: "row" },
            }}
          >
            {selectedLanguage && (
              <Box
                sx={{
                  display: "flex",
                  alignItems: "center",
                  gap: "1rem",
                  flexWrap: "wrap",
                }}
              >
                <Image
                  src={selectedLanguage.image}
                  height={65}
                  width={65}
                  alt={selectedLanguage.title}
                />
              </Box>
            )}

            <Grid
              container
              item
              xs={12}
              spacing={2}
              alignItems="center"
            >
              <Grid item xs={12} md={6} mdlg={5} sm={12}>
                <PrgoressStatutesLabs
                  filters={filters}
                  setFilters={setFilters}
                />
              </Grid>
              <Grid item xs={12} md={5} mdlg={5} sm={6}>
                <SearchFilter
                  searchKey="lab.search.placeholder"
                  filters={filters}
                  setFilters={setFilters}
                />
              </Grid>
              <Grid item xs={12} md={2} sm={6} mdlg={2}>
                <SortFilter
                  filters={filters}
                  setFilters={setFilters}
                  textKey="lab.sort_the_labs"
                />
              </Grid>
            </Grid>
          </Box>
        </Grid>

        <Grid item xs={12}>
          <LabInfo />
        </Grid>
      </Grid>
    </div>
  );
};

export default LanguageLab;
