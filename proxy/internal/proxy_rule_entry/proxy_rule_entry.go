package proxyruleentry

import (
	"proxy/internal/http_message"
	"proxy/internal/target"
)

type ProxyRuleEntry struct {
	Id           string                           `json:"id"`
	Host         string                           `json:"host"`
	Targets      []target.HostAddress             `json:"targets"`
	AddedHeaders []httpmessage.Header             `json:"headers"`
	LoadBalancer target.LoadBalancerConfiguration `json:"load_balancer"`
	LB           target.LoadBalancer              `json:"lb"`
}

func (p *ProxyRuleEntry) GetAddedHeaders() []httpmessage.Header {
	return p.AddedHeaders
}

func (p *ProxyRuleEntry) GetProxyTarget() (target.HostAddress, error) {
	return p.LB.NextTarget(p.Targets)
}
