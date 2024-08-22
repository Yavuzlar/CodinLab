import { Search } from "@mui/icons-material";
import {
  Box,
  FormControl,
  InputAdornment,
  TextField,
  Typography,
  useTheme,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { hexToRGBA } from "src/utils/hex-to-rgba";
import SortFilter from "./SortFilter";

const ProgressStatuses = ({
  filters = {
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
  },
  setFilters = () => {},
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
        gap: "20px",
        alignItems: "center",
        justifyContent: "center",
        
      }}
    >
      {progressStatuses.map((item, index) => {
        return (
          <Typography
            key={index}
            sx={{
              whiteSpace: "nowrap",
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
  );
};

export default ProgressStatuses;