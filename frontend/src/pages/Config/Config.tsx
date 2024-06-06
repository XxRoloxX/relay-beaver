import { useEffect, useState } from 'react';
import './Config.scss'
import { fetchProxyRules } from './configLogic';
import { ProxyRule } from './configLogic';
import AccordionItem from './components/AccordionItem/AccordionItem';
import { deleteProxyRule } from '../../api/proxyApi';

const Config = () => {

  const [proxyRules, setProxyRules] = useState<ProxyRule[]>([])

  useEffect(() => {
    const fetch = async () => {
      const rules = await fetchProxyRules();
      setProxyRules(rules);
    }
    fetch();
  }, []);

  function newProxyRule() {
    setProxyRules([...proxyRules, {id: "", Destination: {host: "", port: 0}, Targets: [], LoadBalancer: {name: ""}}])
  }

  function deleteRule(idx: number, id: string) {
    const newRules = proxyRules.slice();
    newRules.splice(idx, 1);
    setProxyRules(newRules);
    deleteProxyRule(id)
    .then((response) => {
      console.log(response);
    });
  }

  return (
    <>
      <div className="config">
        <div>
          <h1>Configuration</h1>
        </div>
        <div>
            <div className="header">
              <h2 className="header__heading--left">Source</h2>
              <h2 className="header__heading--right">Targets running</h2>
            </div>
            {
              proxyRules.map((proxyRule, idx) => {
                return ( 
                  <div className="accordion">
                    <AccordionItem 
                      proxyRule={proxyRule}
                      proxyRuleIdx={idx}
                      deleteProxyRule={deleteRule}
                    />
                  </div>
                )
              })
            }
          <button className="config__new-btn" onClick={newProxyRule}>New</button>
        </div>
      </div>
    </>
  );
};
export default Config;
