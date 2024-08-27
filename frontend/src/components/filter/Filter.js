import { Search } from "@mui/icons-material";
import {
  Box,
  FormControl,
  InputAdornment,
  TextField,
  Typography,
  useTheme
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { hexToRGBA } from "src/utils/hex-to-rgba";
import SortFilter from "./SortFilter";

const Filter = ({
  searchPlaceholder,
  filters = {
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  },
  setFilters = () => { },
}) => {
  // const [filters, setFilters] = useState({
  //     status: "all", // all, in-progress, completed
  //     search: "",
  //     sort: "", // "", asc, desc
  // })

  const { t } = useTranslation();
  const theme = useTheme();

  const progressStatuses = [
    {
      name: t("all"),
      status: "all",
    },
    {
      name: t("in_progress"),
      status: "in-progress",
    },
    {
      name: t("completed"),
      status: "completed",
    },
  ];

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        minHeight: "44px",
        flexWrap: "wrap",
        gap: '1rem',
        position: 'relative',

      }}
    >
      <Box sx={{ display: "flex", gap: "1rem", width: 'auto' }}>
        {progressStatuses.map((item, index) => {
          return (
            <Typography
              key={index}
              sx={{
                minWidth: 'fit-content',
                cursor: "default",
                color: (theme) =>
                  filters.status == item.status
                    ? theme.palette.primary.dark
                    : hexToRGBA(theme.palette.primary.dark, 0.6),
                "&:hover": {
                  textDecoration: "underline",
                  cursor: filters.status != item.status ? "pointer" : "default",
                },
              }}
              onClick={() => {
                setFilters({ ...filters, status: item.status });
              }}
            >
              {item.name}
            </Typography>
          );
        })}
      </Box>

      <Box sx={{ height: "44px", minWidth: 'fit-content' }}>
        <SortFilter
          filters={filters}
          setFilters={setFilters}
          textKey="labs.sort_the_labs"
        />
      </Box>

      <Box sx={{ width: '100%' }}>
        <FormControl fullWidth>
          <TextField
            name="search-in-labs"
            placeholder={searchPlaceholder}
            variant="outlined"
            size="small"
            onChange={(e) => {
              setFilters({ ...filters, search: e.target.value });
            }}
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
    </Box>
  );
};

export default Filter;
