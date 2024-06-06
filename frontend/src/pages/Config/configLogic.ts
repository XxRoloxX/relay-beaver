import { getProxyRules } from "../../api/proxyApi"

export type Header = {
    name: string
    value: string
}

export type Address = {
    host: string
    port: number
}

export type LoadBalancer = {
    name: string
}

export type ProxyRule = {
    id: string
    Destination: Address
    Targets: Address[]
    LoadBalancer: LoadBalancer
}

export const fetchProxyRules = async () => {
    const res = await getProxyRules();
    return res as ProxyRule[];
}
