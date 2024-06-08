package env

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func GetProxyServerPort() int {
	return getPortOrReturnDefault("PROXY_SERVER_PORT", 8000)
}

func GetWebsocketClientPort() int {
	return getPortOrReturnDefault("WEBSOCKET_SERVER_PORT", 8888)
}

func GetWebsocketServerHost() string {
	return getEnv("WEBSOCKET_SERVER_HOST", "localhost")
}

func GetWebsocketServerEndpoint() string {
	return getEnv("WEBSOCKET_SERVER_ENDPOINT", "/requests")
}

func GetProxyCertPath() string {
	return getEnv("PROXY_CERT_PATH", "proxy.crt")
}

func GetProxyKeyPath() string {
	return getEnv("PROXY_KEY_PATH", "proxy.key")
}

func GetProxyBackendAuthSecret() string {
	return getEnv("PROXY_AUTH_SECRET", "")
}

func GetProxyBackendAuthHeader() string {
	return getEnv("PROXY_AUTH_HEADER", "X-Auth-Secret")
}

func GetProxyBackendHost() string {
	return getEnv("PROXY_BACKEND_API_HOST", "localhost")
}
func GetProxyBackendPort() int {
	res := getEnv("PROXY_BACKEND_API_PORT", "8080")
	port, err := strconv.Atoi(res)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("error parsing port: %s", err.Error()))
		return 8080
	}
	return port
}

func getPortOrReturnDefault(env string, fallback int) int {
	portEnv := getEnv(env, "")

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("ENV: %S not found, defaulting to %d", env, fallback))
		return fallback
	}

	return port
}

func getEnv(name string, fallback string) string {
	env := os.Getenv(name)
	if env == "" {
		return fallback
	}
	return env
}
