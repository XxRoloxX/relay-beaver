import { useCallback, useEffect, useState } from "react";
import { HostStats, getHostStats, getHosts } from "../../api/statsApi";
import { TimeSeriesData } from "../../components/LinearChart/TimeseriesChart";

export interface HostStatsTimeSeries {
  totalRequests: TimeSeriesData[];
  averageLatency: TimeSeriesData[];
  badRequests: TimeSeriesData[];
  serverErrors: TimeSeriesData[];
  [key: string]: TimeSeriesData[];
}

export const countBadRequests = (stats: HostStatsTimeSeries) =>
  stats.badRequests.reduce((acc, entry) => acc + entry.value, 0);

export const countServerErrors = (stats: HostStatsTimeSeries) =>
  stats.serverErrors.reduce((acc, entry) => acc + entry.value, 0);

export const countGoodRequests = (stats: HostStatsTimeSeries) =>
  stats.totalRequests.reduce((acc, entry) => acc + entry.value, 0) -
  countBadRequests(stats) -
  countServerErrors(stats);

const useStats = () => {
  const [hosts, setHosts] = useState<string[]>([]);
  const [host, setHost] = useState<string | null>(null);
  const [stats, setStats] = useState<HostStatsTimeSeries | null>(null);
  const [showHosts, setShowHosts] = useState(false);

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

  const fetchHosts = async () => {
    const hosts = await getHosts();
    setHosts(hosts);
    if (!host) {
      setHost(hosts[0]);
    }
  };

  useEffect(() => {
    fetchStats();
  }, [host, fetchStats]);

  useEffect(() => {
    fetchHosts();
  }, []);

  return { host, setHost, stats, setStats, hosts, showHosts, setShowHosts };
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
