import { useTheme } from "@emotion/react";
import React from "react";
import ActivityCalendar from "react-activity-calendar";
import { Tooltip } from "@mui/material";

const ActivityStatisticComponent = ({ activityData }) => {
  const theme = useTheme();
  return (
    <ActivityCalendar
      data={activityData}
      blockRadius={4}
      blockSize={16}
      theme={{
        light: ["#DAF0FE", "#99D1F2", "#53B2F0", "#1E86E5", "#0A3B7A"],
        dark: ["#0A3B7A", "#1561A5", "#1E86E5", "#53B2F0", "#99D1F2"],
      }}
      labels={{
        legend: {
          less: "Low",
          more: "High",
        },
        months: [
          "Jan",
          "Feb",
          "Mar",
          "Apr",
          "May",
          "Jun",
          "Jul",
          "Aug",
          "Sep",
          "Oct",
          "Nov",
          "Dec",
        ],
        totalCount: "{{count}} contributions in {{year}}",
        showWeekdayLabels: false,
      }}
      renderBlock={(block, activity) => (
        <Tooltip title={`${activity.count} activities on ${activity.date}`}>
          {block}
        </Tooltip>
      )}
    />
  );
};

export default ActivityStatisticComponent;
