package request

import "proxy/internal/http_message"

type Request struct {
	Method   string               `json:"method"`
	Protocol string               `json:"protocol"`
	Path     string               `json:"path"`
	Headers  []httpmessage.Header `json:"headers"`
	Body     string               `json:"body"`
}

type Response struct {
	StatusCode int                  `json:"statusCode"`
	Protocol   string               `json:"protocol"`
	Headers    []httpmessage.Header `json:"headers"`
	Body       string               `json:"body"`
}

type ProxiedRequest struct {
	Request   Request  `json:"request"`
	Response  Response `json:"response"`
	Target    string   `json:"target"`
	StartTime int64    `json:"startTime"` // Start time as Unix timestamp
	EndTime   int64    `json:"endTime"`   // End time as Unix timestamp
}

func NewProxiedRequest(request Request, response Response, target string, startTime int64, endTime int64) ProxiedRequest {
	return ProxiedRequest{
		Request:   request,
		Response:  response,
		Target:    target,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func ProxiedRequestFromHttpMessages(
	requestMessage *httpmessage.HttpMessage,
	responseMessage *httpmessage.HttpMessage,
	target string,
	startTime int64,
	endTime int64,
) (ProxiedRequest, error) {
	request, err := RequestFromHttpMessage(requestMessage)
	if err != nil {
		return ProxiedRequest{}, err
	}
	response, err := ResponseFromHttpMessage(responseMessage)
	if err != nil {
		return ProxiedRequest{}, err
	}
	return NewProxiedRequest(request, response, target, startTime, endTime), nil
}

func RequestFromHttpMessage(message *httpmessage.HttpMessage) (Request, error) {
	method, err := message.GetMethod()
	if err != nil {
		return Request{}, err
	}
	protocol, err := message.GetProtocol()
	if err != nil {
		return Request{}, err
	}

	path, err := message.GetPath()
	if err != nil {
		return Request{}, err
	}

	return Request{
		Method:   method,
		Protocol: protocol,
		Path:     path,
		Headers:  message.GetHeaders(),
		Body:     message.GetBody(),
	}, nil
}

func ResponseFromHttpMessage(message *httpmessage.HttpMessage) (Response, error) {
	statusCode, err := message.GetStatusCode()
	if err != nil {
		return Response{}, err
	}
	protocol, err := message.GetProtocol()
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: statusCode,
		Protocol:   protocol,
		Headers:    message.GetHeaders(),
		Body:       message.GetBody(),
	}, nil
}
