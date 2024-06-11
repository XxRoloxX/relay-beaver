import { ChangeEvent, ChangeEventHandler, useEffect, useState } from "react"
import { LoadBalancer } from "../../../configLogic"
import "./ProxyLoadBalancers.scss"

interface Props {
    loadBalancer: LoadBalancer,
    updateProxyLb: (lb: LoadBalancer) => void
}

const ProxyLoadBalancers: React.FC<Props> = ({ loadBalancer, updateProxyLb }) => {
    
    const [lb, setLb] = useState<LoadBalancer>(loadBalancer);
    console.log(lb); 
    
    function updateLb(e: ChangeEvent<HTMLSelectElement>) {
        const updatedLb = lb
        lb.Name = e.target.value;
        setLb(updatedLb);
        updateProxyLb(updatedLb)
    }

    return (
        <div>
            <h3 className="lb-header">Load Balancer</h3>
            <div className="load-balancers">
                <select className="load-balancers__config" onChange={updateLb} value={lb?.Name}>
                    <option value="round_robin">round robin</option>
                    <option value="least_connections">least connections</option>
                </select>
            </div>
        </div>
    )    
}

export default ProxyLoadBalancers