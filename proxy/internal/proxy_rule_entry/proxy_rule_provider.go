package proxyruleentry

type ProxyRuleEntryProvider struct {
	proxyRuleEntries map[string]ProxyRuleEntry
}

type RuleEntryProvider interface {
	GetProxyRuleEntry(host string) ProxyRuleEntry
	GetProxyRuleEntries() map[string]ProxyRuleEntry
}

func NewProxyRuleEntryProvider(proxyRuleEntries map[string]ProxyRuleEntry) *ProxyRuleEntryProvider {
	return &ProxyRuleEntryProvider{
		proxyRuleEntries: proxyRuleEntries,
	}
}
