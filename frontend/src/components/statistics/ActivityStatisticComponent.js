import React from "react";
import ActivityCalendar from "react-activity-calendar";
import { Tooltip } from "@mui/material";
import { useTranslation } from "react-i18next";

const ActivityStatisticComponent = ({ activityData }) => {
  const { t } = useTranslation();

  const formattedData = activityData.map((item) => ({
    date: new Date(item.date).toISOString().split("T")[0], // YYYY-MM-DD formatı
    count: parseInt(item.count, 10),
    level: item.level,
  }));

  return (
    <ActivityCalendar
      data={formattedData}
      blockRadius={4}
      blockSize={16}
      theme={{
        dark: ["#DAF0FE", "#99D1F2", "#53B2F0", "#1E86E5", "#0A3B7A"],
        light: ["#0A3B7A", "#1561A5", "#1E86E5", "#53B2F0", "#99D1F2"],
      }}
      labels={{
        legend: {
          less: t("admin_activity_calendar_legend_less", {
            defaultValue: "Az",
          }),
          more: t("admin_activity_calendar_legend_more", {
            defaultValue: "Çok",
          }),
        },
        months: t("admin_activity_calendar_months", {
          defaultValue: [
            "Oca",
            "Şub",
            "Mar",
            "Nis",
            "May",
            "Haz",
            "Tem",
            "Ağu",
            "Eyl",
            "Eki",
            "Kas",
            "Ara",
          ],
        }),
        totalCount: t("admin_activity_calendar_total_count", {
          defaultValue: "{{count}} katkı {{year}} yılında",
        }),
        showWeekdayLabels: false,
      }}
      renderBlock={(block, activity) => (
        <Tooltip
          title={
            activity
              ? t("admin_activity_calendar_logs_tooltip", {
                  count: activity.count,
                  date: activity.date,
                })
              : t("admin_activity_calendar_no_activity", {
                  defaultValue: "Etkinlik Yok",
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
