import { TimeSeriesChart } from "../../components/LinearChart/TimeseriesChart";
import "./Stats.scss";
import useStats from "./useStats";

const STATS_CHARTS = [
  {
    title: "Bad Requests",
    key: "badRequests",
    legend: "Number of bad requests",
  },
  {
    title: "Server Errors",
    key: "serverErrors",
    legend: "Number of server errors",
  },
  {
    title: "Average Latency (s)",
    key: "averageLatency",
    legend: "Average latency (s)",
  },
  {
    title: "Total Requests",
    key: "totalRequests",
    legend: "Total requests",
  },
];

const Stats = () => {
  const { host, setHost, stats } = useStats();

  return (
    <div className="stats">
      <div className="stats__header">
        <h2>{host} Stats</h2>
        <select
          value={host!}
          onChange={(e) => setHost(e.target.value)}
          className="stats__header__select"
        >
          <option value="">Select a host</option>
          <option value="api1">api1</option>
          <option value="api2">api2</option>
        </select>
      </div>
      <div className="stats__charts">
        {stats &&
          STATS_CHARTS?.map((chart) => (
            <TimeSeriesChart
              key={chart.key}
              data={stats[chart.key]}
              title={chart.title}
              legend={chart.legend}
            />
          ))}
      </div>
    </div>
  );
};

export default Stats;
