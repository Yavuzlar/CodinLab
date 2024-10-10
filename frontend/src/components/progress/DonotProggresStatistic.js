import { useState, useEffect } from "react";
import dynamic from "next/dynamic";
import { theme } from "src/configs/theme";

const ReactApexChart = dynamic(() => import("react-apexcharts"), {
  ssr: false,
});

const DonutProgressStatistic = ({ data }) => {
  const [chartData, setChartData] = useState({
    series: [],
    options: {
      stroke: {
        colors: [`${theme.palette.background.paper}`],
      },
      chart: {
        type: "donut",
      },
      labels: [],
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

  useEffect(() => {
  
    if (data) {
      setChartData((prevState) => ({
        ...prevState,
        series: data?.values || [],
        options: {
          ...prevState.options,
          labels: data?.labels || [],
        },
      }));
    }
  }, [data]); 

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

export default DonutProgressStatistic;
