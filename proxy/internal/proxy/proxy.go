package proxy

import (
	"crypto/tls"
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
	request2 "proxy/internal/request"
	"proxy/internal/target"
	"strings"
)

type Socket int

const (
	Tcp Socket = iota
)

func (s Socket) String() string {
	switch s {
	case Tcp:
		return "tcp"
	}
	return "unknown"
}

type Proxy struct {
	port     int
	certPath string
	keyPath  string
	provider target.RuleEntryProvider
	parser   request2.RequestParser
}

func NewProxy(port int, certPath string, keyPath string, provider target.RuleEntryProvider, parser request2.RequestParser) *Proxy {
	return &Proxy{
		port:     port,
		certPath: certPath,
		keyPath:  keyPath,
		provider: provider,
		parser:   parser,
	}
}

func (p *Proxy) Start() error {
	log.Info().Msg(fmt.Sprintf("starting proxy server on port: %d", p.port))
	err := p.listenOnSocket(p.port, Tcp)

	if err != nil {
		return err
	}

	return nil
}

func (p *Proxy) listenOnSocket(port int, socket Socket) error {
	crt, err := tls.LoadX509KeyPair(p.certPath, p.keyPath)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("error loading key pair: %s", err))
		return err
	}

	config := &tls.Config{Certificates: []tls.Certificate{crt}}
	ln, err := tls.Listen(socket.String(), fmt.Sprintf(":%d", port), config)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("error binding tcp socket: %s", err))
		return err
	}

	// TODO - handle defer error
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Error().Msg(fmt.Sprintf("error accepting connection: %s", err))
			continue
		}

		go p.handleConnection(ConnectionBuffer{conn: conn})
	}
}

// TODO - return bad request / internal server error when parsing fails
func (p *Proxy) handleConnection(buffer ConnectionBuffer) {
	// TODO - handle defer error
	defer buffer.Close()

	request, err := buffer.ReadAll()
	if err != nil {
		log.Error().Msg(fmt.Sprintf("error reading request content, error: %s", err))
		return
	}

	req, proxyAddress, err := p.parseHttpRequest(request)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("error parsing request %s", err))
		return
	}

	response, err := p.proxy(req, proxyAddress)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("failed proxying to destination: %s, error: %s", proxyAddress, err))
		return
	}

	bytesWritten, err := buffer.Write([]byte(response))
	if err != nil {
		log.Error().Msg(fmt.Sprintf("error writing response from destination %s, error: %s", proxyAddress, err))
		return
	}

	log.Info().Msg(fmt.Sprintf("successfuly wrote %d bytes from response", bytesWritten))
	return
}

func (p *Proxy) proxy(content string, proxyAddress string) (string, error) {
	conn, err := net.Dial("tcp", proxyAddress)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error establishing TCP connection with %s", proxyAddress))
		return "", err
	}

	buffer := ConnectionBuffer{conn: conn}
	defer buffer.Close()

	_, err = buffer.Write([]byte(content))
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Proxying to destination %s failed, cannot write request content to buffer", proxyAddress))
		return "", err
	}

	response, err := buffer.ReadAll()
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error reading response from destination %s, reading from buffer failed", proxyAddress))
		return "", err
	}

	return response, nil
}

func (p *Proxy) parseHttpRequest(request string) (string, string, error) {
	res := strings.Split(request, "\r\n")
	go p.parser.ParseRequest(request)

	for i := 0; i < len(res); i++ {
		line := res[i]
		if strings.HasPrefix(line, "Host: ") {
			hostSplit := strings.Split(line, ":")
			host := strings.Join(hostSplit[:2], ":")

			proxyRule := p.provider.GetProxyRuleEntry(host)

			proxyTarget := proxyRule.GetProxyTarget()
			res[i] = fmt.Sprintf("Host: %s", proxyTarget.GetURL())
			return strings.Join(p.appendHeaders(res, proxyRule.GetAddedHeaders()), "\r\n") + "\r\n", proxyTarget.GetURL(), nil
		}
	}

	return "", "", fmt.Errorf("no host header found")
}

func (p *Proxy) appendHeaders(request []string, addedHeaders []request2.Header) []string {
	headers := make([]string, 0)
	for _, header := range addedHeaders {
		headers = append(headers, fmt.Sprintf("%s: %s", header.Key, header.Value))
	}

	res := make([]string, 1)
	res[0] = request[0]
	res = append(res, headers...)

	if len(request) > 1 {
		res = append(res, request[1:]...)
	}

	return res
}
