import { useTheme } from "@emotion/react";
import React from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const GraphicalStatistics = ({ data }) => {
  const theme = useTheme();

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
        <YAxis
          tick={{ fill: "#fff" }}
          axisLine={null}
          ticks={[0, 2, 4, 6, 8, 10, 12, 14, 16]}
          tickFormatter={(value) => `${value} hours`}
          tickLine={false}
        />
        <Tooltip />
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
          stroke={theme.palette.primary.light}
          strokeWidth={3}
        />
      </LineChart>
    </ResponsiveContainer>
  );
};

export default GraphicalStatistics;
