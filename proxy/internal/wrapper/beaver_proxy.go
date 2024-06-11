package wrapper

import (
	"github.com/rs/zerolog/log"
	"proxy/internal/backend_provider"
	"proxy/internal/env"
	"proxy/internal/proxy"
	"proxy/internal/request"
	"proxy/internal/websocket"
	"time"
)

func RunProxy() {
	channel := make(chan request.ProxiedRequest)
	proxyPort := env.GetProxyServerPort()
	websocketPort, host, endpoint := env.GetWebsocketClientPort(), env.GetWebsocketServerHost(), env.GetWebsocketServerEndpoint()

	go startWebsocketClient(channel, host, websocketPort, endpoint)
	startProxy(channel, proxyPort)
}

func startProxy(channel chan request.ProxiedRequest, port int) {
	factory := backendprovider.RuleEntryProviderFactory{}
	proxyTarget := factory.BackendApiRuleEntryProvider()
	parser := request.NewSimpleRequestParser(channel)

	p := proxy.NewProxy(port, env.GetProxyCertPath(), env.GetProxyKeyPath(), proxyTarget, parser)
	err := p.Start()
	if err != nil {
		log.Error().Msg("error starting proxy")
		panic(err)
	}
}

func startWebsocketClient(channel chan request.ProxiedRequest, host string, port int, endpoint string) {
	ws := websocket.NewWebsocketServer(channel, host, port, endpoint)

	retries := 0
	for {
		if retries == 5 {
			panic("cannot connect to websocket server")
		}

		err := ws.Connect()
		if err != nil {
			time.Sleep(1 * time.Second)
			log.Info().Msg("sleeping for 1s...")
			retries++
		} else {
			log.Info().Msg("successfully connected to websocket server")
			break
		}
	}
}
