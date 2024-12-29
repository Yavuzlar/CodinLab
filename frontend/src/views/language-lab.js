import { useContext, useState } from "react";
import {
  Box,
  FormControl,
  Grid,
  InputAdornment,
  TextField,
  useMediaQuery,
} from "@mui/material";
import Image from "next/image";
import C from "../assets/language/c.png";
import Cpp from "../assets/language/cpp.png";
import Go from "../assets/language/go.png";
import Js from "../assets/language/javascript.png";
import Python from "../assets/language/python.png";
import LabInfo from "../components/cards/LabInfo";
import SortFilter from "src/components/filter/SortFilter";
import PrgoressStatutesLabs from "src/components/filter/PrgoressStatutesLabs";
import { Search } from "@mui/icons-material";
import { t } from "i18next";
import { theme } from "src/configs/theme";
import { useRouter } from "next/router";
import { AuthContext } from "src/context/AuthContext";



const LanguageLab = ({ language = "" }) => {
  const programingId = language;

  const lgmd_down = useMediaQuery((theme) => theme.breakpoints.down("lgmd"));
  const { containerLoading } = useContext(AuthContext)

  const [iconPath, setIconPath] = useState("");
  const [filters, setFilters] = useState({
    status: "all",
    difficulty: "all",
    search: "",
    sort: "",
  });

  return (
    <div>
      <Grid container spacing={2} direction="column">
        <Grid item xs={12}>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              gap: "1rem",
              flexDirection: lgmd_down ? "column" : "row",
            }}
          >
            <Box>
              <img
                src={"/api/v1/" + iconPath}
                height={65}
                width={65}
              // alt={selectedLanguage.title}
              />
            </Box>

            <Grid
              container
              item
              xs={12}
              spacing={2}
              alignItems="center"
              justifyContent="center"
              sx={{
                flexDirection: lgmd_down ? "column" : "row",
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
                  <Grid item xs={12} lgmd={8} md={6} smd={12}>
                    <FormControl fullWidth>
                      <TextField
                        name="search-in-labs"
                        placeholder={t("labs.search.placeholder")}
                        variant="outlined"
                        size="small"
                        onChange={(e) =>
                          setFilters({ ...filters, search: e.target.value })
                        }
                        InputProps={{
                          startAdornment: (
                            <InputAdornment sx={{ zIndex: 10, mr: 1 }}>
                              <Search />
                            </InputAdornment>
                          ),
                          style: { color: theme.palette.text.primary },
                        }}
                        sx={{
                          "& .MuiInputBase-input": {
                            color: theme.palette.text.primary,
                            zIndex: 9,
                            "&::placeholder": {
                              color: theme.palette.text.primary,
                              opacity: 0.7,
                            },
                          },
                          "& .MuiOutlinedInput-root": {
                            "& fieldset": {
                              backgroundColor: theme.palette.primary.main,
                            },
                          },
                          width: "100%",
                          height: "100%",
                          minWidth: "450px",
                        }}
                      />
                    </FormControl>
                  </Grid>
                  <Grid item xs={12} lgmd={4} md={6} smd={12}>
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
          <LabInfo containerLoading={containerLoading} setIconPath={setIconPath} filter={filters} programingId={programingId} />
        </Grid>
      </Grid>
    </div>
  );
};

export default LanguageLab;
