import {
  Box,
  Card,
  CardContent,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Typography,
  useMediaQuery,
} from "@mui/material";
import ActivityStatisticComponent from "../statistics/ActivityStatisticComponent";
import { data } from "../../data/activityDataExample";
import { useState } from "react";
import styled from "@emotion/styled";
import Translations from "../Translations";

const CustomSelect = styled(Select)({
  backgroundColor: "#0A3B7A",
  color: "#FFFFFF",
  "& .MuiSvgIcon-root": {
    color: "#FFFFFF",
  },
  borderRadius: "0.5rem",
  height: "2.5rem",
});

const Activity = () => {
  const _smd = useMediaQuery((theme) => theme.breakpoints.down("smd"));

  const currentYear = new Date().getFullYear();
  const [year, setYear] = useState(currentYear);

  const handleChange = (event) => {
    setYear(event.target.value);
  };
  return (
    <Box
      sx={{
        display: "flex",
        width: "100%",
        height: "100%",
      }}
    >
      <Card
        sx={{
          width: "100%",
          minHeight: "12.5rem",
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between",
        }}
      >
        <CardContent sx={{ height: "calc(100% - 3rem)" }}>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              marginBottom: "1rem",
            }}
          >
            <Box sx={{ display: "flex", justifyContent: "space-between" }}>
              <Box>
                <Typography variant="title" sx={{ fontWeight: "bold" }}>
                  <Translations text="admin.activity.title" />
                </Typography>
              </Box>
              <Box
                sx={{
                  textAlign: "center",
                  minWidth: 90,
                }}
              >
                <FormControl fullWidth variant="standard">
                  <CustomSelect
                    labelId="year-select-label"
                    id="year-select"
                    value={year}
                    label="Year"
                    onChange={handleChange}
                    disableUnderline
                  >
                    {[2024, 2023, 2022, 2021].map((y) => (
                      <MenuItem key={y} value={y}>
                        {y}
                      </MenuItem>
                    ))}
                  </CustomSelect>
                </FormControl>
              </Box>
            </Box>
            <Typography
              sx={{
                maxWidth: "calc(100% - 9.625rem)",

                paddingTop: "0.8rem",
                ...(_smd && { maxWidth: "60ch" }),
              }}
            >
              <Translations text="admin.activity.content" />
            </Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
            }}
          >
            <ActivityStatisticComponent activityData={data} />
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Activity;
