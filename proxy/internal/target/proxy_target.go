package target

type ProxyTarget struct {
	host string
	port string
}

func (p *ProxyTarget) GetURL() string {
	return p.host + ":" + p.port
}

func NewProxyTarget(host string, port string) *ProxyTarget {
	return &ProxyTarget{
		host: host,
		port: port,
	}
}
