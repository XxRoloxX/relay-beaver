import { getProxyRules } from "../../api/proxyApi";

export type Header = {
  key: string;
  value: string;
};

export type Address = {
  host: string;
  port: number;
};

export type LoadBalancer = {
  Name: string;
  Params: Map<string, string>;
};

export type ProxyRule = {
  id: string;
  host: string;
  targets: Address[];
  headers: Header[];
  load_balancer: LoadBalancer;
};

export const fetchProxyRules = async () => {
  const res = await getProxyRules();
  return res as ProxyRule[];
};
