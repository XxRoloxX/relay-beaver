import { useState } from "react"
import "./AccordionItem.scss"
import ProxyHeaders from "./headers/ProxyHeaders"
import ProxyTargets from "./targets/ProxyTargets"
import ProxyLoadBalancers from "./loadbalancers/ProxyLoadBalancers"
import { ProxyRule } from "../../configLogic"
import { createProxyRule, updateProxyRule } from "../../../../api/proxyApi"

interface Props {
    proxyRule: ProxyRule
    proxyRuleIdx: number
    deleteProxyRule: (idx: number, id: string) => void
}

const AccordionItem: React.FC<Props> = ({ proxyRule,  proxyRuleIdx, deleteProxyRule}) => {

    const [isActive, setIsActive] = useState(false)
    const [rule, setRule] = useState<ProxyRule>(proxyRule);

    function updateDestination(e: React.ChangeEvent<HTMLInputElement>) {
        const newHost = e.target.value;
        setRule(prevState => ({
            ...prevState,
            Destination: {
                ...prevState.Destination,
                host: newHost
            }
        }));
    }

    function apply(proxyRule: ProxyRule) {
        console.log(proxyRule.id)
        if(proxyRule.id === "") {
            createProxyRule(proxyRule)
            .then(response => {
                console.log(response);
            })
        } else {
            updateProxyRule(proxyRule)
            .then(response => {
                console.log(response);
            })
        }
    }

    function getAccordionHeader() {
        return (
            <div className="accordion-title">
                <input className="accordion-title__service-name" type="text" value={rule.Destination.host} onChange={updateDestination}/>
                <div className="accordion-title__info-icon">
                    <img width="35" height="35" src="https://img.icons8.com/ios/50/FFFFFF/info--v1.png" alt="info--v1"/>
                </div>
                <div className="accordion-title__active-targets">
                    <p className="accordion-title__targets-info">2/3</p>
                    <div className="accordion-title__active-targets__status-indicator"></div>
                </div>
                <div className="accordion-title__open" onClick={() => setIsActive(!isActive)}>
                    <p>{isActive ? '-' : '+'}</p>
                </div>
                <div className="accordion-title__delete">
                    <svg onClick={() => deleteProxyRule(proxyRuleIdx, proxyRule.id)} xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" viewBox="0 0 16 16"> 
                        <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0z"/>
                        <path d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4zM2.5 3h11V2h-11z"/>
                    </svg>
                </div>
            </div>
        )
    }
    
    function getAccordionContent(proxyRule: ProxyRule) {
        return (
            <div className="accordion-content">
                <hr/>
                <div className="accordion-content__column">
                    <ProxyHeaders/>
                </div>
                <div className="accordion-content__column">
                    <ProxyTargets proxyTargets={proxyRule.Targets}/>
                </div>

                <div className="accordion-content__column">
                    <ProxyLoadBalancers loadBalancer={proxyRule.LoadBalancer}/>
                    <button className="accordion-content__button--cancel" onClick={() => setIsActive(false)}>Cancel</button>  
                    <button className="accordion-content__button--apply" onClick={() => apply(proxyRule)}>Apply</button>
                </div>
            </div> 
        )
    }

    return (
        <div>
            { 
                getAccordionHeader() 
            }
            { isActive && getAccordionContent(proxyRule) }
        </div>
    )
}

export default AccordionItem