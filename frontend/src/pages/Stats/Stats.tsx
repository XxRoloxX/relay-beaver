import { useCallback } from "react";
import { TimeSeriesChart } from "../../components/LinearChart/TimeseriesChart";
import PolarChart, {
  PolarChartEntry,
} from "../../components/PolarChart/PolarChart";
import "./Stats.scss";
import useStats, {
  countBadRequests,
  countGoodRequests,
  countServerErrors,
} from "./useStats";
import { HostSelector } from "./components/HostSelector/HostSelector";
import Shadow from "../../components/Shadow/Shadow";
import Popup from "../../components/Popup/Popup";

const STATS_CHARTS = [
  {
    title: "Bad Requests",
    key: "badRequests",
    legend: "Number of bad requests",
    color: "rgba(255, 0, 0, 0.6)",
    backgroundColor: "rgba(200, 0, 0, 0.6)",
  },
  {
    title: "Server Errors",
    key: "serverErrors",
    legend: "Number of server errors",
    color: "rgba(255, 159, 64, 0.6)",
    backgroundColor: "rgba(255, 159, 64, 0.6)",
  },
  {
    title: "Average Latency (s)",
    key: "averageLatency",
    legend: "Average latency (s)",
    color: "rgba(0,0,255)",
    backgroundColor: "rgba(0,0,200)",
  },
  {
    title: "Total Requests",
    key: "totalRequests",
    legend: "Total requests",
    color: "rgba(75, 192, 192, 0.6)",
    backgroundColor: "rgba(75, 192, 192, 0.6)",
  },
];

const POLAR_CHART_STATS = [
  {
    label: "Bad Requests",
    color: "rgba(255, 99, 132, 0.6)",
    accumulatorFn: countBadRequests,
  },
  {
    label: "Server Errors",
    color: "rgba(255, 159, 64, 0.6)",
    accumulatorFn: countServerErrors,
  },
  {
    label: "Good requests",
    color: "rgba(75, 192, 192, 0.6)",
    accumulatorFn: countGoodRequests,
  },
];

const Stats = () => {
  const { host, setHost, stats, hosts, showHosts, setShowHosts } = useStats();
  const getPolarChartData = useCallback(() => {
    const formatedData: PolarChartEntry[] = stats
      ? POLAR_CHART_STATS.map((stat) => ({
        color: stat.color,
        label: stat.label,
        value: stat.accumulatorFn(stats),
      }))
      : [];
    return formatedData;
  }, [stats]);

  return (
    <div className="stats">
      <div className="stats__header">
        <h2 className="stats__header__host" onClick={() => setShowHosts(true)}>
          {host}
        </h2>
        <Popup isDisplayed={showHosts} setIsDisplayed={setShowHosts}>
          <HostSelector host={host} hosts={hosts} setHost={setHost} />
        </Popup>
      </div>
      <div className="stats__charts--small">
        {stats &&
          STATS_CHARTS?.map((chart) => (
            <TimeSeriesChart
              key={chart.key}
              data={stats[chart.key]}
              title={chart.title}
              legend={chart.legend}
              color={chart.color}
              backgroundColor={chart.backgroundColor}
            />
          ))}
      </div>
      <div className="stats__charts--big">
        <PolarChart data={getPolarChartData()} label={"Requests summary"} />
      </div>
      {showHosts && <Shadow />}
    </div>
  );
};

export default Stats;
