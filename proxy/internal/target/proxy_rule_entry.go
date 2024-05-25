package target

import (
	"proxy/internal/request"
)

type ProxyRuleEntry struct {
	host         string
	targets      []ProxyTarget
	addedHeaders []request.Header
	loadBalancer LoadBalancer
}

func (p *ProxyRuleEntry) GetAddedHeaders() []request.Header {
	return p.addedHeaders
}

func (p *ProxyRuleEntry) GetProxyTarget() ProxyTarget {
	return p.loadBalancer.NextTarget(p.targets)
}

// TODO - fetch configs from DB / CRUD
func NewProxyRuleEntry(host string, targets []ProxyTarget, additionalHeaders []request.Header, loadBalancer LoadBalancer) *ProxyRuleEntry {
	return &ProxyRuleEntry{
		host:         host,
		targets:      targets,
		addedHeaders: additionalHeaders,
		loadBalancer: loadBalancer,
	}
}
