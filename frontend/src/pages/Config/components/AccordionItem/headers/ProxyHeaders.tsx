import { useState } from "react"
import "./ProxyHeaders.scss"
import { Header } from "../../../configLogic"

const ProxyHeaders = () => {

    const [headers, setHeaders] = useState<Header[]>([{name: "test", value: "testowy"}, {name: "test2", value: "testowy3"}]);

    function addHeader() {
        setHeaders([...headers, {name: "sample", value: "header"}])
    } 

    function deleteHeader(idx: number) {
        const newHeaders = headers.slice();
        newHeaders.splice(idx, 1);
        setHeaders(newHeaders);
    }

    return (
        <div>
            <h3 className="headers-header">Headers</h3>
            <div className="proxy-headers">
                <div id="border">
                    <div className="grid-container">
                        <div>
                            <div className="grid-item">Name</div>
                            {
                                headers.map((header, idx) => {
                                    return <input className="input-left" value={header.name}/>
                                })
                            }
                        </div>
                        <div className="separator"></div>
                        <div>
                            <div className="grid-item">Value</div>
                            {
                                headers.map((header, idx) => {
                                    return (
                                        <>
                                            <input className="input-right" value={header.name}/>
                                            <svg onClick={() => deleteHeader(idx)} xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="trash-icon" viewBox="0 0 16 16">
                                                <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0z"/>
                                                <path d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4zM2.5 3h11V2h-11z"/>
                                            </svg>
                                        </>
                                    )
                                })
                            }
                        </div>
                    </div>

                    <button className="proxy-headers__add-btn" onClick={addHeader}>+</button>
                </div>
            </div>
        </div>
    )    
}

export default ProxyHeaders