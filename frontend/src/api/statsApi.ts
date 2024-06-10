import { proxyAxios } from "./proxyApi";

export interface StatsEntry {
  timestamp: number;
  value: number;
}

export interface HostStats {
  host: string;
  totalRequests: StatsEntry[];
  averageLatency: StatsEntry[];
  badRequests: StatsEntry[];
  serverErrors: StatsEntry[];
}

export const tryParseAsHostStats = (obj: unknown): HostStats => {
  if (typeof obj !== "object" || obj === null) {
    throw new Error("Invalid object");
  }

  const castedObj = obj as HostStats;
  if (typeof castedObj.host !== "string") {
    throw new Error("Invalid host");
  }
  if (!Array.isArray(castedObj.totalRequests)) {
    throw new Error("Invalid totalRequests");
  }
  if (!Array.isArray(castedObj.averageLatency)) {
    throw new Error("Invalid averageLatency");
  }
  if (!Array.isArray(castedObj.badRequests)) {
    throw new Error("Invalid badRequests");
  }
  if (!Array.isArray(castedObj.serverErrors)) {
    throw new Error("Invalid serverErrors");
  }

  return castedObj;
};

export const getHostStats = async (
  host: string,
  from?: number,
  to?: number,
  interval?: number,
) => {
  const response = await proxyAxios.get(
    `/stats?host=${host}&from=${from}&to=${to}&interval=${interval}`,
  );
  return tryParseAsHostStats(response.data);
};

export const getHosts = async () => {
  const response = await proxyAxios.get("/stats/hosts");
  return response.data as string[];
};
