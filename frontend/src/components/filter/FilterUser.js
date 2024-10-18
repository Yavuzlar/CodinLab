import { Search } from "@mui/icons-material";
import {
  Box,
  FormControl,
  InputAdornment,
  TextField,
  Typography,
  useMediaQuery,
  useTheme,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { hexToRGBA } from "src/utils/hex-to-rgba";
import SortFilterUser from "./SortFilterUser";
import SortFilter from "./SortFilter";

const FilterUser = ({
  searchPlaceholder,
  filters = {
    search: "",
    sort: "",
  },
  setFilters = () => {},
}) => {
  const { t } = useTranslation();
  const theme = useTheme();

  const _sm = useMediaQuery((theme) => theme.breakpoints.down("sm"));

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "flex-start",
        flexDirection: _sm ? "column" : "row",
        gap: "1rem",
        width: "100%",
        justifyContent: "space-between",
      }}
    >
    

        <Box sx={{ width: _sm ? "100%" : "50%" }}>
          <FormControl fullWidth>
            <TextField
              name="search-in-labs"
              placeholder={t("users.search.placeholder")}
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
              }}
            />
          </FormControl>
        </Box>
    

      <Box sx={{ height: "44px", minWidth: "fit-content" }}>
        <SortFilter
          filters={filters}
          setFilters={setFilters}
          textKey="labs.sort_the_labs"
        />
      </Box>
    </Box>
  );
};

export default FilterUser;
