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
	proxyPort, websocketPort := getPorts()
	go startWebsocket(channel, websocketPort)
	startProxy(channel, proxyPort)
}

func getPorts() (int, int) {
	proxyPort := env.GetProxyServerPort()
	wsPort := env.GetWebsocketServerPort()

	if proxyPort == wsPort {
		panic("proxy and websocket ports cannot be the same")
	}

	return proxyPort, wsPort
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

func startWebsocket(channel chan request.Request, port int) {
	ws := websocket.NewWebsocketServer(channel, port)
	err := ws.Start()
	if err != nil {
		panic(err)
	}
}
