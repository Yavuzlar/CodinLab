import { useTheme } from "@emotion/react";
import React, { useEffect, useState } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import Translations from "../Translations";
import { t } from "i18next";
import { useTranslation } from "react-i18next";

const GraphicalStatistics = ({ data }) => {
  const theme = useTheme();
  const { i18n } = useTranslation();
  const language = i18n.language;
  const [hoursLang, setHoursLang] = useState({ language });

  useEffect(() => {
    let storedLang = localStorage.getItem("i18nextLng");

    switch (storedLang) {
      case "tr":
        setHoursLang("saat");
        break;
      case "en":
        setHoursLang("hour");
        break;
      default:
        setHoursLang("hour");
        break;
    }
  }, [language, i18n]);

  return (
    <ResponsiveContainer width="100%" height={280}>
      <LineChart data={data}>
        <CartesianGrid
          stroke="#fff"
          strokeWidth={0.5}
          strokeDasharray="0"
          vertical={false}
        />
        <XAxis
          dataKey="name"
          padding={{ left: 30, right: 30 }}
          tick={{ fill: "#fff" }}
          tickLine={false}
          axisLine={{
            stroke: "#fff",
            strokeWidth: 2,
          }}
        />
        {/* <YAxis
          tick={{ fill: "#fff" }}
          axisLine={null}
          ticks={[0, 2, 4, 6, 8, 10, 12, 14, 16]}
          tickFormatter={(value) => `${value} ${hoursLang}`}
          tickLine={false}
        /> */}
        <Tooltip
          contentStyle={{ backgroundColor: theme.palette.background.default }}
          labelStyle={{ color: theme.palette.primary.main }}
        />
        <Line
          type="monotone"
          dataKey="Roads"
          stroke={theme.palette.primary.dark}
          strokeWidth={3}
          activeDot={{ r: 8 }}
        />
        <Line
          type="monotone"
          dataKey="Labs"
          stroke={theme.palette.info.main}
          strokeWidth={3}
        />
      </LineChart>
    </ResponsiveContainer>
  );
};

export default GraphicalStatistics;
