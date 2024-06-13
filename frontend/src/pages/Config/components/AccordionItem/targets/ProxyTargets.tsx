import "./ProxyTargets.scss"
import { Address } from "../../../configLogic"
import { useState } from "react"

interface Props {
    proxyTargets: Address[]
    updateTargets: (targets: Address[]) => void
}

const ProxyTargets: React.FC<Props> = ({ proxyTargets, updateTargets }) => {

    const [targets, setTargets] = useState<Address[]>(proxyTargets);

    function updateTargetHost(e: React.ChangeEvent<HTMLInputElement>, idx: number) {
        const newTargets = targets.slice()
        newTargets[idx].host = e.target.value
        setTargets(newTargets);
    }

    function updateTargetPort(e: React.ChangeEvent<HTMLInputElement>, idx: number) {
        const newTargets = targets.slice()
        newTargets[idx].port = Number(e.target.value)
        setTargets(newTargets)
    }

    function newTarget(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        e.preventDefault();
        setTargets((prev) => {
            const newTargets = prev.slice();
            newTargets[newTargets.length] = {host: "google.com", port: 80};
            updateTargets(newTargets);
            return newTargets;
        })
        // const newTargets = [...targets, {host: "google.com", port: 80}]
        // setTargets(newTargets)
        // updateTargets(newTargets);
    }

    function deleteTarget(idx: number) {
        setTargets((prev) => {
            const updatedArray = prev.slice();
            updatedArray.splice(idx, 1);
            // setTargets(updatedArray)
            updateTargets(updatedArray);
            return updatedArray;
        })
        // const updatedArray = targets.slice()
        // updatedArray.splice(idx, 1)
        // setTargets(updatedArray)
        // updateTargets(updatedArray);
    }

    return (
        <div>
            <h3 className="target-header">Targets</h3>
            <div className="targets">
                <div className="targets__container">
                    {
                        targets?.map((target, idx) => {
                            return (
                                <div className="targets__entry">
                                    <input className='targets__entry__host' value={target.host} type="text" onChange={(e) => updateTargetHost(e, idx)}></input>
                                    <span id="colon">:</span>
                                    <input className='targets__entry__port' value={target.port} type="number" onChange={(e) => updateTargetPort(e, idx)} min="1" max="65535" required/>
                                    <svg onClick={() => deleteTarget(idx)} xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                                        <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0z"/>
                                        <path d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4zM2.5 3h11V2h-11z"/>
                                    </svg>
                                    <div className="targets__status-indicator"></div>
                                </div>
                            )
                        })
                    }
                </div>
                <button className="targets__add-btn" onClick={(e) => newTarget(e)}>+</button>
            </div>
        </div>
    )
}

export default ProxyTargets