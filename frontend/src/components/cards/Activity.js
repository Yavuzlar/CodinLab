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
// import { data } from "../../data/activityDataExample";
import { useEffect, useState } from "react";
import styled from "@emotion/styled";
import Translations from "../Translations";
import { useDispatch, useSelector } from "react-redux";
import { fetchActivityByYear } from "src/store/admin/adminActivitySlice";

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
  const dispatch = useDispatch();
  const { activity: activityData } = useSelector((state) => state);

  const data = activityData.data.data;

  const currentYear = new Date().getFullYear();
  const [year, setYear] = useState(currentYear);

  const _smd = useMediaQuery((theme) => theme.breakpoints.down("smd"));

  const handleChange = (event) => {
    setYear(event.target.value);
  };

  useEffect(() => {
    dispatch(fetchActivityByYear({ year }));
  }, [year]);

  const years = Array.from(
    { length: new Date().getFullYear() - 2023 },
    (_, i) => 2024 + i
  ).reverse();

  // const years = [2024, 2023, 2022, 2021, 2020];

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
                    {years.map((y) => (
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
            {Array.isArray(data) && data.length !== 0 ? (
              <ActivityStatisticComponent activityData={data} />
            ) : (
              <Typography>
                <Translations text="admin_activity_calendar_no_data_error" />
              </Typography>
            )}
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default Activity;
