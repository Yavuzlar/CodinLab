import { useState } from "react";
import { Box, Grid, useMediaQuery } from "@mui/material";
import Image from "next/image";
import C from "../assets/language/c.png";
import Cpp from "../assets/language/cpp.png";
import Go from "../assets/language/go.png";
import Js from "../assets/language/javascript.png";
import Python from "../assets/language/python.png";
import LabInfo from "../components/cards/LabInfo";
import SortFilter from "src/components/filter/SortFilter";
import SearchFilter from "src/components/filter/FilterLab";
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
  const lgmd_down = useMediaQuery((theme) => theme.breakpoints.down("lgmd"));
  
  return (
    <div>
      <Grid container spacing={2} direction="column" >
        <Grid item xs={12} >
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              gap: "1rem",
              flexDirection:  lgmd_down ? "column" : "row",


            }}
          >
            {selectedLanguage && (
              <Box
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
              justifyContent="center"
              sx={{
                flexDirection:  lgmd_down ? "column" : "row",
              }}
            >
              <Grid item xs={12} md={4} mdlg={5} sm={12}>
                <PrgoressStatutesLabs
                  filters={filters}
                  setFilters={setFilters}
                />
              </Grid>
              <Grid item xs={12} md={7} mdlg={7} sm={12}>
                <Grid container spacing={2}>
                  <Grid item xs={12} lgmd={8} md={6} smd = {12}

                  >
                    <SearchFilter
                      searchKey="lab.search.placeholder"
                      filters={filters}
                      setFilters={setFilters}
                    />
                  </Grid>
                  <Grid item xs={12} lgmd={4} md={6} smd = {12}>
                    <SortFilter
                  
                      filters={filters}
                      setFilters={setFilters}
                      textKey="lab.sort_the_labs"
                    />
                  </Grid>
                </Grid>
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
