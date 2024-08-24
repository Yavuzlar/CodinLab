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
import TestTubeRed from "../../assets/icons/red.png";
import TestTubeOrgane from "../../assets/icons/orange.png";
import TestTubeGreen from "../../assets/icons/green.png";
import Image from "next/image";
import SortFilterLab from "./SortFilerLab";


const FilterLab = ({
  filters = {
    status: "all", // all, in-progress, completed
    difficulty: "all", // all, easy, medium, hard
    search: "",
    sort: "", // "", asc, desc
    languages: "",
  },
  setFilters = () => {},
}) => {
 

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

  const difficultyLevels = [
    {
      image: TestTubeGreen,
      difficulty: "easy",
    },
    {
      image: TestTubeOrgane,
      difficulty: "medium",
    },
    {
      image: TestTubeRed,
      difficulty: "hard",
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
      }}
    >
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          gap: "25px",
          marginRight: "20px",
        }}
      >
        {progressStatuses.map((item, index) => (
          <Typography
            key={index}
            sx={{
              cursor: "default",
              fontSize: "16px",
              color:
                filters.status === item.status
                  ? theme.palette.primary.dark
                  : hexToRGBA(theme.palette.primary.dark, 0.6),
              "&:hover": {
                textDecoration: "underline",
                cursor: filters.status !== item.status ? "pointer" : "default",
              },
            }}
            onClick={() => setFilters({ ...filters, status: item.status })}
          >
            {item.name}
          </Typography>
        ))}

        {difficultyLevels.map((item, index) => (
          <Box
            key={index}
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              cursor: "pointer",
            }}
          >
            <Image
              src={item.image}
              alt="test-tube"
              width={45}
              height={45}
              onClick={() =>
                setFilters({ ...filters, difficulty: item.difficulty })
              }
            />
          </Box>
        ))}
      </Box>

      <Box>
        <FormControl fullWidth>
          <TextField
            name="search-in-labs"
            placeholder={t("labs.search.placeholder")}
            variant="outlined"
            size="small"
            onChange={(e) => setFilters({ ...filters, search: e.target.value })}
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
      </Box>

      <Box sx={{ flex: 1, height: '44px', display: 'flex', justifyContent: 'flex-end',marginLeft:'32px' }}>
        <SortFilterLab
          filters={filters}
          setFilters={setFilters}
          textKey="labs.sort_the_labs"
        />  
      </Box>
    </Box>
  );
};

export default FilterLab;
