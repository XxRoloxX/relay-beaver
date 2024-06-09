package backendprovider

import (
	"fmt"
	"proxy/internal/api"
	"proxy/internal/proxy_rule_entry"

	"github.com/rs/zerolog/log"
)

func getProxyRuleEntriesFromApi() map[string]proxyruleentry.ProxyRuleEntry {
	proxyRuleEntries := make(map[string]proxyruleentry.ProxyRuleEntry)
	proxyClient := api.NewProxyBackendClient()
	proxyRules, err := proxyClient.GetProxyRules()

	if err != nil {
		log.Error().Msg(fmt.Sprintf("error fetching proxy rules: %s", err.Error()))
		panic(err)
	}

	for _, rule := range proxyRules {
		proxyRuleEntries[rule.Host] = rule
	}

	return proxyRuleEntries
}

type RuleEntryProviderFactory struct{}

type ProxyRuleEntryProvider struct {
	proxyRuleEntries map[string]proxyruleentry.ProxyRuleEntry
}

func NewProxyRuleEntryProvider(proxyRuleEntries map[string]proxyruleentry.ProxyRuleEntry) *ProxyRuleEntryProvider {
	return &ProxyRuleEntryProvider{
		proxyRuleEntries: proxyRuleEntries,
	}
}

func (m *ProxyRuleEntryProvider) GetProxyRuleEntry(host string) proxyruleentry.ProxyRuleEntry {
	return m.proxyRuleEntries[host]
}

func (p *ProxyRuleEntryProvider) GetProxyRuleEntries() map[string]proxyruleentry.ProxyRuleEntry {
	return p.proxyRuleEntries
}

func (p *RuleEntryProviderFactory) BackendApiRuleEntryProvider() proxyruleentry.RuleEntryProvider {
	provider := NewProxyRuleEntryProvider(getProxyRuleEntriesFromApi())
	return provider
}

func (p *RuleEntryProviderFactory) ProxyTarget() proxyruleentry.RuleEntryProvider {
	// TODO - implement
	panic("not implemented yet")
}
