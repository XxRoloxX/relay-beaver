// import { LinearScale, TimeScale } from "chart.js";
import { Line } from "react-chartjs-2";
import "chartjs-adapter-moment";
import { useState } from "react";
import { ChartOptions } from "chart.js";
import "./TimeseriesChart.scss";

interface TimeSeriesChartProps {
  data: TimeSeriesData[];
  legend: string;
  title: string;
  from?: number; // Unix timestamp
  to?: number; // Unix timestamp
}

export interface TimeSeriesData {
  timestamp: number; // Unix timestamp
  value: number;
}

export const getChartOptions = (
  title: string,
  from?: number,
  to?: number,
): ChartOptions<"line"> => ({
  responsive: true,
  maintainAspectRatio: true,
  resizeDelay: 200,
  color: "white",
  backgroundColor: "none",
  scales: {
    y: {
      beginAtZero: true,
      ticks: {
        color: "white",
      },
    },
    x: {
      type: "time",
      time: {
        unit: "hour",
        displayFormats: {
          hour: "MMM D hA",
        },
      },
      ticks: {
        color: "white",
        maxTicksLimit: 2,
      },
      min: from,
      max: to,
    },
  },
  plugins: {
    legend: {
      display: true,
      position: "bottom",
    },
    title: {
      display: true,
      text: title,
      position: "top",
      color: "white",
      font: {
        size: 16,
      },
    },
  },
});

export const TimeSeriesChart = ({
  data,
  from,
  to,
  legend,
  title,
}: TimeSeriesChartProps) => {
  const [chartData] = useState({
    labels: [],
    datasets: [
      {
        label: legend,
        data: data.map((d) => ({
          x: d.timestamp,
          y: d.value,
        })),
        fill: true,
        borderColor: "rgb(255, 0, 0)",
        backgroundColor: "rgb(150, 0, 0)",
        tension: 0.1,
      },
    ],
  });

  const timeSeriesChartOptions = getChartOptions(title, from, to);

  return (
    <div className="timeseries-chart">
      <Line data={chartData} options={timeSeriesChartOptions} />
    </div>
  );
};
