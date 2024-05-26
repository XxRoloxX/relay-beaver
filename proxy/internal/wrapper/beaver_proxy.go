package wrapper

import (
	"github.com/rs/zerolog/log"
	"proxy/internal/env"
	"proxy/internal/proxy"
	"proxy/internal/request"
	"proxy/internal/target"
	"proxy/internal/websocket"
)

func RunProxy() {
	channel := make(chan request.Request)
	proxyPort := env.GetProxyServerPort()
	websocketPort, host, endpoint := env.GetWebsocketClientPort(), env.GetWebsocketServerHost(), env.GetWebsocketServerEndpoint()

	go startWebsocketClient(channel, host, websocketPort, endpoint)
	startProxy(channel, proxyPort)
}

func startProxy(channel chan request.Request, port int) {
	factory := target.RuleEntryProviderFactory{}
	proxyTarget := factory.MockProxyTarget()
	parser := request.NewSimpleRequestParser(channel)

	p := proxy.NewProxy(port, env.GetProxyCertPath(), env.GetProxyKeyPath(), proxyTarget, parser)
	err := p.Start()
	if err != nil {
		log.Error().Msg("error starting proxy")
		panic(err)
	}
}

func startWebsocketClient(channel chan request.Request, host string, port int, endpoint string) {
	ws := websocket.NewWebsocketServer(channel, host, port, endpoint)
	err := ws.Connect()
	if err != nil {
		panic(err)
	}
}
