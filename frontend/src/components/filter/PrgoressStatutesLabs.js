import {
  Box,
  Typography,
  useTheme
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { hexToRGBA } from "src/utils/hex-to-rgba";
import TestTubeRed from "../../assets/icons/red.png";
import TestTubeOrgane from "../../assets/icons/orange.png";
import TestTubeGreen from "../../assets/icons/green.png";
import Image from "next/image";

const ProgressStatusesLab = ({
  filters = {
    status: "all", // all, in-progress, completed
    search: "",
    sort: "", // "", asc, desc
    difficulty: "all",
  },
  setFilters = () => { },
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

  const handleDiffSet = (item) => {
    if (filters.difficulty === item.difficulty) {
      setFilters({ ...filters, difficulty: "all" });
    } else {
      setFilters({ ...filters, difficulty: item.difficulty });
    }
  }

  return (
    <Box
      sx={{
        display: "flex",
        gap: "15px",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      {progressStatuses.map((item, index) => (
        <Typography
          key={index}
          sx={{
            whiteSpace: "nowrap",
            cursor: "default",
            color:
              filters.status === item.status
                ? theme.palette.primary.dark
                : hexToRGBA(theme.palette.primary.dark, 0.6),
            "&:hover": {
              textDecoration: "underline",
              cursor: filters.status !== item.status ? "pointer" : "default",
            },
          }}
          onClick={() => {
            setFilters({ ...filters, status: item.status });
          }}
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
          onClick={() => {
            handleDiffSet(item)
          }}
        >
          <Image
            src={item.image}
            alt={`${item.difficulty}-level`}
            width={40}
            height={40}
          />
        </Box>
      ))}
    </Box>
  );
};

export default ProgressStatusesLab;
