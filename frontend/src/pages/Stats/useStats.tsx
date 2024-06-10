import { useCallback, useEffect, useState } from "react";
import { HostStats, getHostStats } from "../../api/statsApi";
import { TimeSeriesData } from "../../components/LinearChart/TimeseriesChart";

export interface HostStatsTimeSeries {
  totalRequests: TimeSeriesData[];
  averageLatency: TimeSeriesData[];
  badRequests: TimeSeriesData[];
  serverErrors: TimeSeriesData[];
  [key: string]: TimeSeriesData[];
}

const useStats = () => {
  const [host, setHost] = useState<string>("localhost");
  const [stats, setStats] = useState<HostStatsTimeSeries | null>(null);

  const fetchStats = useCallback(async () => {
    try {
      if (!host) {
        return;
      }
      const stats = await getHostStats(host);
      const timeseries = mapHostDataToTimeSeries(stats);
      setStats(timeseries);
    } catch (error) {
      console.error(`Failed to fetch stats for host ${host}: ${error}`);
    }
  }, [host]);

  useEffect(() => {
    fetchStats();
  }, [host, fetchStats]);

  return { host, setHost, stats, setStats };
};

export const mapHostDataToTimeSeries = (
  data: HostStats,
): HostStatsTimeSeries => {
  return {
    totalRequests: data.totalRequests.map((entry) => ({
      timestamp: entry.timestamp * 1000,
      value: entry.value,
    })),
    averageLatency: data.averageLatency.map((entry) => ({
      timestamp: entry.timestamp * 1000,
      value: entry.value,
    })),
    badRequests: data.badRequests.map((entry) => ({
      timestamp: entry.timestamp * 1000,
      value: entry.value,
    })),
    serverErrors: data.serverErrors.map((entry) => ({
      timestamp: entry.timestamp * 1000,
      value: entry.value,
    })),
  };
};

export default useStats;
