import { useEffect, useState } from "react";
import "./Config.scss";
import { ProxyRule } from "./configLogic";
import AccordionItem from "./components/AccordionItem/AccordionItem";
import { fetchProxyRules } from "./configLogic";
import { deleteProxyRule } from "../../api/proxyApi";

const Config = () => {
  const [proxyRules, setProxyRules] = useState<ProxyRule[]>([]);

  useEffect(() => {
    console.log("test");
    const fetch = async () => {
      const rules = await fetchProxyRules();
      setProxyRules(rules);
    };
    fetch();
  }, []);

  function newProxyRule() {
    setProxyRules([
      ...proxyRules,
      {
        id: "",
        host: "google.com",
        targets: [],
        headers: [],
        load_balancer: { Name: "round robin", Params: new Map() },
      },
    ]);
  }

  function deleteRule(idx: number, id: string) {
    const newRules = proxyRules.slice();
    newRules.splice(idx, 1);
    setProxyRules(newRules);

    deleteProxyRule(id).then((response) => {
      console.log(response);
    });
  }

  return (
    <>
      <div className="config">
        <div>
          <h1>Configuration</h1>
        </div>
        <div className="config__content">
          <div className="header">
            <h2 className="header__heading--left">Source</h2>
            <h2 className="header__heading--right">Targets running</h2>
          </div>
          {proxyRules?.map((proxyRule, idx) => {
            return (
              <div className="accordion" key={idx}>
                <AccordionItem
                  proxyRule={proxyRule}
                  proxyRuleIdx={idx}
                  deleteProxyRule={deleteRule}
                />
              </div>
            );
          })}
          <div className="config__footer">
            <button className="config__new-btn" onClick={newProxyRule}>
              New
            </button>
          </div>
        </div>
      </div>
    </>
  );
};

export default Config;
