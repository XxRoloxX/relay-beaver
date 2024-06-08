package request

import (
	"fmt"
	"github.com/rs/zerolog/log"
	httpmessage "proxy/internal/http_message"
	"strings"
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

func (p *SimpleRequestParser) ParseRequest(requestContent string, responseContent string, target string, startTime int64, endTime int64) {
	// requestInfo := strings.Split(strings.Split(requestContent, "\r\n\r\n")[0], "\r\n")
	// request := Request{}
	//
	// httpPart := requestInfo[0]
	// populateHttpInfo(httpPart, &request)
	//
	// headersPart := requestInfo[1:]
	// populateHttpHeaders(headersPart, &request)

	requestHttpMessage := httpmessage.FromString(requestContent)
	responseHttpMessage := httpmessage.FromString(responseContent)

	// println(requestHttpMessage.ToString())
	// println(responseHttpMessage.ToString())

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
