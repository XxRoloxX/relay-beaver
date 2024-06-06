import { LoadBalancer } from "../../../configLogic"
import "./ProxyLoadBalancers.scss"

interface Props {
    loadBalancer: LoadBalancer
}

const ProxyLoadBalancers: React.FC<Props> = ({ loadBalancer }) => {
    return (
        <div>
            <h3 className="lb-header">Load Balancer</h3>
            <div className="load-balancers">
                {/* <div className="load-balancers__config"> */}
                    <select className="load-balancers__config" name="lb">
                        <option value="round robin">round robin</option>
                        <option value="least connections">least connections</option>
                    </select>
                {/* </div> */}
            </div>
        </div>
    )    
}

export default ProxyLoadBalancers