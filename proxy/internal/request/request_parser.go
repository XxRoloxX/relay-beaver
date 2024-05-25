package request

import (
	"strings"
)

type SimpleRequestParser struct {
	RequestChannel chan Request
}

type RequestParser interface {
	ParseRequest(content string)
}

func NewSimpleRequestParser(requestChannel chan Request) RequestParser {
	return &SimpleRequestParser{RequestChannel: requestChannel}
}

func (p *SimpleRequestParser) ParseRequest(content string) {
	requestInfo := strings.Split(strings.Split(content, "\r\n\r\n")[0], "\r\n")
	request := Request{}

	httpPart := requestInfo[0]
	populateHttpInfo(httpPart, &request)

	headersPart := requestInfo[1:]
	populateHttpHeaders(headersPart, &request)

	p.RequestChannel <- request
}

func (p *SimpleRequestParser) populateHttpInfo(content string, request *Request) {
	info := strings.Split(content, " ")
	request.Method = info[0]
	request.Path = info[1]
	request.Protocol = info[2]
}

func (p *SimpleRequestParser) populateHttpHeaders(content []string, request *Request) {
	var headers []Header
	for _, header := range content {
		headerParts := strings.Split(header, ":")
		key := headerParts[0]
		value := strings.Trim(headerParts[1], " ")
		headers = append(headers, Header{Key: key, Value: value})
	}

	request.Headers = headers
}

//func ParseRequest(content string) {
//	requestInfo := strings.Split(strings.Split(content, "\r\n\r\n")[0], "\r\n")
//	request := Request{}
//
//	httpPart := requestInfo[0]
//	populateHttpInfo(httpPart, &request)
//
//	headersPart := requestInfo[1:]
//	populateHttpHeaders(headersPart, &request)
//
//	// TODO -> fix circular imports
//	channel <- request
//	//websocket.GetChannel() <- request
//}

func populateHttpInfo(content string, request *Request) {
	info := strings.Split(content, " ")
	request.Method = info[0]
	request.Path = info[1]
	request.Protocol = info[2]
}

func populateHttpHeaders(content []string, request *Request) {
	var headers []Header
	for _, header := range content {
		headerParts := strings.Split(header, ":")
		key := headerParts[0]
		value := strings.Trim(headerParts[1], " ")
		headers = append(headers, Header{Key: key, Value: value})
	}

	request.Headers = headers
}
