package proxy

import (
	"fmt"
	"net"
	"proxy/internal/env"
	"proxy/internal/http_message"
	"proxy/internal/proxy_rule_entry"
	request2 "proxy/internal/request"
	"proxy/internal/target"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
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
	provider proxyruleentry.RuleEntryProvider
	parser   request2.RequestParser
}

func NewProxy(port int, certPath string, keyPath string, provider proxyruleentry.RuleEntryProvider, parser request2.RequestParser) *Proxy {
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
	// TODO - maybe stick to no TLS for dev phase? - quite out of scope of the project
	//crt, err := tls.LoadX509KeyPair(p.certPath, p.keyPath)
	//if err != nil {
	//	log.Error().Msg(fmt.Sprintf("error loading key pair: %s", err))
	//return err
	//}

	//config := &tls.Config{Certificates: []tls.Certificate{crt}}
	//ln, err := tls.Listen(socket.String(), fmt.Sprintf(":%d", port), config)

	ln, err := net.Listen(socket.String(), fmt.Sprintf(":%d", port))
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

	if request != "" {
		startTime := time.Now().Unix()

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
		endTime := time.Now().Unix()

		go p.parser.ParseRequest(request, response, proxyAddress, startTime, endTime)

		bytesWritten, err := buffer.Write([]byte(response))

		if err != nil {
			log.Error().Msg(fmt.Sprintf("error writing response from destination %s, error: %s", proxyAddress, err))
			return
		}

		log.Info().Msg(fmt.Sprintf("successfuly wrote %d bytes from response", bytesWritten))
		return
	} else {
		badRequest := "HTTP/1.1 400 Bad Request\r\n\r\n"
		buffer.conn.Write([]byte(badRequest))
	}
}

func (p *Proxy) proxy(content string, proxyAddress string) (string, error) {
	conn, err := net.Dial("tcp", proxyAddress)
	defer conn.Close()

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

	r := httpmessage.FromString(response)
	r.SetHeader("Cache-Control", "no-cache") // TODO: demo reasons only (show that LB works)

	return replaceNewLinesWithCRLF(r.ToString()), nil
}

func replaceNewLinesWithCRLF(content string) string {
	return strings.Replace(strings.Replace(content, "\n\n", "\r\n\r\n", -1), "\n", "\r\n", 1)
}

func (p *Proxy) parseHttpRequest(request string) (string, string, error) {
	requestMessage := httpmessage.FromString(request)
	host := requestMessage.GetHeader("Host")

	var proxyTarget target.HostAddress
	var err error

	if host == "" {
		return "", "", fmt.Errorf("no host header found")
	}

	proxyRule := p.provider.GetProxyRuleEntry(host)
	proxyTarget, err = proxyRule.GetProxyTarget()

	if err != nil {
		log.Error().Msg(fmt.Sprintf("no targets found for host: %s", host))
		return "", "", err
	}

	forcedTargetHeader := requestMessage.GetHeader(env.GetTargetHeader())

	if forcedTargetHeader != "" {
		log.Info().Msg(fmt.Sprintf("Forced target header found: %s", forcedTargetHeader))
		proxyTarget, err = target.HostAddressFromString(forcedTargetHeader)
		if err != nil {
			return "", "", fmt.Errorf("error parsing forced target header: %s", err)
		}
	}

	requestMessage.SetHeader("Host", proxyTarget.GetURL())

	return requestMessage.ToString(), proxyTarget.GetURL(), nil
}
