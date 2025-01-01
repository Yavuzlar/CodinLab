import React, { useState, useEffect } from "react";
import ActivityCalendar from "react-activity-calendar";
import { Tooltip } from "@mui/material";
import { useTranslation } from "react-i18next";

const ActivityStatisticComponent = (props) => {
  const { activityData, selectedYear } = props;
  const { t } = useTranslation();

  const getLocalizedLabels = () => {
    const weekdays = t("admin_activity_calendar_weekdays", {
      returnObjects: true,
      defaultValue: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
    });

    const months = t("admin_activity_calendar_months", {
      returnObjects: true,
      defaultValue: [
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
    });

    return { weekdays, months };
  };

  const { weekdays, months } = getLocalizedLabels();

  const generateYearData = (year, data) => {
    const dates = new Map(
      data.map((item) => [item.date, { count: item.count, level: item.level }])
    );

    const yearData = [];
    const startDate = new Date(`${year}-01-01`);
    const endDate = new Date(`${year}-12-31`);

    let currentDate = new Date(startDate);

    while (currentDate <= endDate) {
      const formattedDate = currentDate.toISOString().split("T")[0];
      yearData.push({
        date: formattedDate,
        count: dates.get(formattedDate)?.count || 0,
        level: dates.get(formattedDate)?.level || 0,
      });
      currentDate.setDate(currentDate.getDate() + 1);
    }

    return yearData;
  };

  const formattedData = generateYearData(
    selectedYear,
    activityData.map((item) => ({
      date: new Date(item.date).toISOString().split("T")[0],
      count: parseInt(item.count, 10),
      level: item.level,
    }))
  );

  return (
    <ActivityCalendar
      data={formattedData}
      blockRadius={4}
      blockSize={16}
      showWeekdayLabels={true}
      theme={{
        dark: ["#DAF0FE", "#99D1F2", "#53B2F0", "#1E86E5", "#0A3B7A"],
        light: ["#0A3B7A", "#1561A5", "#1E86E5", "#53B2F0", "#99D1F2"],
      }}
      labels={{
        legend: {
          less: t("admin_activity_calendar_legend_less", {
            defaultValue: "Less",
          }),
          more: t("admin_activity_calendar_legend_more", {
            defaultValue: "High",
          }),
        },
        months,
        weekdays,
        totalCount: t("admin_activity_calendar_total_count", {
          defaultValue: "{{count}} katkı {{year}} yılında",
        }),
      }}
      renderBlock={(block, activity) => (
        <Tooltip
          title={
            activity && activity.count > 0
              ? t("admin_activity_calendar_logs_tooltip", {
                  count: activity.count,
                  date: activity.date,
                })
              : t("admin_activity_calendar_no_activity", {
                  defaultValue: "No Activity",
                })
          }
        >
          {block}
        </Tooltip>
      )}
    />
  );
};

export default ActivityStatisticComponent;
