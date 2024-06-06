import axios from "axios";
import { ProxyRule } from "../pages/Config/configLogic";

export const proxyAxios = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_URL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

export const login = async (authCode: string) => {
  const response = await proxyAxios.post("/auth/login", { code: authCode });
  return response.data;
};

export const getTokenInfo = async () => {
  const response = await proxyAxios.get("/auth/profile");
  return response.data;
};

export const getProxyRules = async () => {
  const response = await proxyAxios.get("/proxy-rules");
  return response.data;
};

export const createProxyRule = async (proxyRule: ProxyRule) => {
  const response = await proxyAxios.post("/proxy-rules", JSON.stringify(proxyRule));
  return response.data;
}

export const updateProxyRule = async(proxyRule: ProxyRule) => {
  const response = await proxyAxios.put(`/proxy-rules/${proxyRule.id}`, JSON.stringify(proxyRule));
  return response.data;
}

export const deleteProxyRule = async(id: string) => {
  const response = await proxyAxios.delete(`/proxy-rules/${id}`);
  return response.data;
}

export const getRequests = async () => {
  const response = await proxyAxios.get("/requests");
  return response.data;
};
export const logout = async () => {
  document.cookie = "id_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
  document.location.href = "/";
};
