import "./ProxyTargets.scss"
import { Address } from "../../../configLogic"
import { useEffect, useState } from "react"

interface Props {
    proxyTargets: Address[]
}

const ProxyTargets: React.FC<Props> = ({ proxyTargets }) => {

    const [targets, setTargets] = useState<Address[]>(proxyTargets);

    function updateTarget(e: React.ChangeEvent<HTMLInputElement>) {
        const newTarget = e.target.value;
        console.log(newTarget);      

    }

    function newTarget() {
        setTargets([...targets, {host: "google.com", port: 80}])
    }

    function deleteTarget(idx: number) {
        const updatedArray = targets.slice()
        updatedArray.splice(idx, 1)
        setTargets(updatedArray)
    }

    return (
        <div>
            <h3 className="target-header">Targets</h3>
            <div className="targets">
                {
                    targets.map((target, idx) => {
                        return (
                            <div className="targets__entry">
                                <input value={`${target.host}:${target.port}`} onChange={updateTarget}></input>
                                <svg onClick={() => deleteTarget(idx)} xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="trash-icon" viewBox="0 0 16 16">
                                    <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0z"/>
                                    <path d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4zM2.5 3h11V2h-11z"/>
                                </svg>
                                <div className="targets__status-indicator"></div>
                            </div>
                        )
                    })
                }
                <button className="targets__add-btn" onClick={newTarget}>+</button>
            </div>
        </div>
    )
}

export default ProxyTargets