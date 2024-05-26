package target

type ProxyRuleEntryProvider struct {
	proxyRuleEntries map[string]ProxyRuleEntry
}

func NewProxyRuleEntryProvider(proxyRuleEntries map[string]ProxyRuleEntry) *ProxyRuleEntryProvider {
	return &ProxyRuleEntryProvider{
		proxyRuleEntries: proxyRuleEntries,
	}
}

func (m *ProxyRuleEntryProvider) GetProxyRuleEntry(host string) ProxyRuleEntry {
	return m.proxyRuleEntries[host]
}
