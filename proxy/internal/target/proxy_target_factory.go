package target

type RuleEntryProvider interface {
	GetProxyRuleEntry(host string) ProxyRuleEntry
}

type RuleEntryProviderFactory struct{}

func (p *RuleEntryProviderFactory) MockProxyTarget() RuleEntryProvider {
	return NewProxyRuleEntryProvider(getMockedProxyRuleEntries())
}

func (p *RuleEntryProviderFactory) ProxyTarget() RuleEntryProvider {
	// TODO - implement
	panic("not implemented yet")
}

func getMockedProxyRuleEntries() map[string]ProxyRuleEntry {
	proxyRuleEntries := make(map[string]ProxyRuleEntry)
	proxyRuleEntries["Host: localhost"] = ProxyRuleEntry{host: "localhost:8080", targets: []ProxyTarget{*NewProxyTarget("localhost", "8081")}, addedHeaders: nil, loadBalancer: &RoundRobinLoadBalancer{}}
	return proxyRuleEntries
}
