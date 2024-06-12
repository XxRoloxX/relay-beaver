package request

import (
	"fmt"
	"proxy/internal/env"
	httpmessage "proxy/internal/http_message"
	"strings"

	"github.com/rs/zerolog/log"
)

type SimpleRequestParser struct {
	RequestChannel chan ProxiedRequest
}

type RequestParser interface {
	ParseRequest(
		requestContent string,
		responseContent string,
		target string,
		startTime int64,
		endTime int64)
}

func NewSimpleRequestParser(requestChannel chan ProxiedRequest) RequestParser {
	return &SimpleRequestParser{RequestChannel: requestChannel}
}
func (p *SimpleRequestParser) appendProxyPort(message httpmessage.HttpMessage) {
	host := message.GetHost()
	if len(strings.Split(host, ":")) >= 2 {
		return
	}

	message.SetHost(fmt.Sprintf("%s:%d", host, env.GetProxyServerPort()))
}

func (p *SimpleRequestParser) ParseRequest(requestContent string, responseContent string, target string, startTime int64, endTime int64) {
	requestHttpMessage := httpmessage.FromString(requestContent)
	responseHttpMessage := httpmessage.FromString(responseContent)

	p.appendProxyPort(requestHttpMessage)

	proxiedRequest, err := ProxiedRequestFromHttpMessages(
		&requestHttpMessage,
		&responseHttpMessage,
		target,
		startTime,
		endTime,
	)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("error creating proxied request: %s", err))
	}

	p.RequestChannel <- proxiedRequest
}

func (p *SimpleRequestParser) populateHttpInfo(content string, request *Request) {
	info := strings.Split(content, " ")
	request.Method = info[0]
	request.Path = info[1]
	request.Protocol = info[2]
}

func (p *SimpleRequestParser) populateHttpHeaders(content []string, request *Request) {
	var headers []httpmessage.Header
	for _, header := range content {
		headerParts := strings.Split(header, ":")
		key := headerParts[0]
		value := strings.Trim(headerParts[1], " ")
		headers = append(headers, httpmessage.Header{Key: key, Value: value})
	}

	request.Headers = headers
}

func populateHttpInfo(content string, request *Request) {
	info := strings.Split(content, " ")
	request.Method = info[0]
	request.Path = info[1]
	request.Protocol = info[2]
}

func populateHttpHeaders(content []string, request *Request) {
	var headers []httpmessage.Header
	for _, header := range content {
		headerParts := strings.Split(header, ":")
		key := headerParts[0]
		value := strings.Trim(headerParts[1], " ")
		headers = append(headers, httpmessage.Header{Key: key, Value: value})
	}

	request.Headers = headers
}
