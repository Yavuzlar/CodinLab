import { useState, useEffect } from "react";
import dynamic from "next/dynamic";
import { theme } from "src/configs/theme";

const ReactApexChart = dynamic(() => import("react-apexcharts"), {
  ssr: false,
});

const DonotProggresStatistic = ({ data }) => {
  const [chartData, setChartData] = useState({
    series: data?.values,

    options: {
      stroke: {
        colors: [`${theme.palette.background.paper}`],
      },
      chart: {
        type: "donut",
      },

      labels: data?.labels,
      colors: [theme.palette.info.dark, theme.palette.primary.light],
      plotOptions: {
        pie: {
          donut: {
            labels: {
              show: false, 
            },
          },
        },
      },
      legend: {
        show: false, 
      },
      responsive: [
        {
          breakpoint: 480,
          options: {
            chart: {
              width: 200,
            },
            legend: {
              show: false, 
            },
          },
        },
      ],
    },
  });

  const [isClient, setIsClient] = useState(false);
  useEffect(() => {
    if (typeof window !== "undefined") {
      setIsClient(true);
    }
  }, []);

  return (
    <div>
      {isClient && (
        <div id="chart">
          <ReactApexChart
            options={chartData?.options}
            series={chartData?.series}
            type="donut"
            height={350}
          />
        </div>
      )}
      <div id="html-dist"></div>
    </div>
  );
};

export default DonotProggresStatistic;
