package backendprovider

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"proxy/internal/api"
	"proxy/internal/proxy_rule_entry"
	"time"
)

func getProxyRuleEntriesFromApi() map[string]proxyruleentry.ProxyRuleEntry {
	proxyRuleEntries := make(map[string]proxyruleentry.ProxyRuleEntry)
	proxyClient := api.NewProxyBackendClient()

	var proxyRules []proxyruleentry.ProxyRuleEntry
	retries := 0
	for {
		if retries == 5 {
			panic("cannot fetch proxy rules from backend")
		}

		rules, err := proxyClient.GetProxyRules()
		if err != nil {
			log.Error().Msg(fmt.Sprintf("error fetching proxy rules from backend: %s", err.Error()))
			log.Info().Msg("sleeping for 1s...")
			time.Sleep(1 * time.Second)
			retries++
		} else {
			log.Info().Msg("successfully fetched proxy rules from backend")
			proxyRules = rules
			break
		}
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
