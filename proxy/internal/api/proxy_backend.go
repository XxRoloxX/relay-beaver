package api

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"proxy/internal/env"
	"proxy/internal/proxy_rule_entry"
)

type ProxyBackendClient struct {
	Host string
	Port int
}

func NewProxyBackendClient() *ProxyBackendClient {
	return &ProxyBackendClient{
		Host: env.GetProxyBackendHost(),
		Port: env.GetProxyBackendPort(),
	}
}

func setAuthHeader(req *http.Request) {
	req.Header.Set(
		env.GetProxyBackendAuthHeader(),
		fmt.Sprintf("%s", env.GetProxyBackendAuthSecret()))
}

func (c *ProxyBackendClient) GetProxyRules() ([]proxyruleentry.ProxyRuleEntry, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/proxy-rules", c.Host, c.Port), nil)
	setAuthHeader(req)
	resp, err := client.Do(req)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("error fetching proxy rules: %s", err.Error()))
		return nil, err
	}

	defer resp.Body.Close()

	var proxyRules []proxyruleentry.ProxyRuleEntry

	err = json.NewDecoder(resp.Body).Decode(&proxyRules)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("error decoding proxy rules: %s", err.Error()))
		return nil, err
	}

	return proxyRules, nil
}
