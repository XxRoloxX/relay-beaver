package httpmessage

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type HttpMessage struct {
	requestLine string
	headers     []Header
	body        string
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func FromString(message string) HttpMessage {
	messagePartsArray := strings.Split(message, "\r\n\r\n")
	requestLineAndHeaders, body := messagePartsArray[0], messagePartsArray[1]
	requestLineAndHeadersPartsArray := strings.Split(requestLineAndHeaders, "\n")
	requestLine, HeadersString := requestLineAndHeadersPartsArray[0], requestLineAndHeadersPartsArray[1:]
	requestLine = strings.Replace(requestLine, "\r", "", 1)
	headers := make([]Header, 0)
	for _, header := range HeadersString {
		header = strings.Replace(header, "\r", "", 1)
		headerParts := strings.Split(header, ":")
		key := headerParts[0]
		v := strings.Join(headerParts[1:], ":")
		value := strings.Trim(v, " ")
		headers = append(headers, Header{Key: key, Value: value})
	}

	return HttpMessage{requestLine, headers, body}
}

func (m *HttpMessage) ToString() string {
	headers := make([]string, 0)
	for _, header := range m.headers {
		headers = append(headers, header.Key+":"+header.Value)
	}
	return m.requestLine + "\r\n" + strings.Join(headers, "\r\n") + "\r\n\r\n" + m.body
}

func (m *HttpMessage) GetResponseCode() (string, error) {
	if !m.IsResponse() {
		return "", errors.New(fmt.Sprintf("Message is not a response: %s", m.requestLine))
	}
	return strings.Split(m.requestLine, " ")[1], nil
}

func (m *HttpMessage) GetMethod() (string, error) {
	if m.IsResponse() {
		return "", errors.New(fmt.Sprintf("Message is not a request: %s", m.requestLine))
	}
	return strings.Split(m.requestLine, " ")[0], nil
}

func (m *HttpMessage) GetBody() string {
	return m.body
}

func (m *HttpMessage) GetPath() (string, error) {
	if m.IsResponse() {
		return "", errors.New(fmt.Sprintf("Message is not a request: %s", m.requestLine))
	}
	return strings.Split(m.requestLine, " ")[1], nil
}

func (m *HttpMessage) GetProtocol() (string, error) {
	if m.IsResponse() {
		return strings.Split(m.requestLine, " ")[0], nil
	} else {
		return strings.Split(m.requestLine, " ")[2], nil
	}
}

func (m *HttpMessage) GetStatusCode() (int, error) {
	if !m.IsResponse() {
		return 0, errors.New(fmt.Sprintf("Message is not a response: %s", m.requestLine))
	}
	statusCode, err := m.GetResponseCode()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(statusCode)
}

func (m *HttpMessage) GetHeaders() []Header {
	return m.headers
}

func (m *HttpMessage) IsResponse() bool {
	return strings.HasPrefix(m.requestLine, "HTTP/")
}

func (m *HttpMessage) IsRequest() bool {
	return !m.IsResponse()
}

func (m *HttpMessage) GetHeader(key string) string {
	for _, header := range m.headers {
		if header.Key == key {
			return header.Value
		}
	}
	return ""
}
func (m *HttpMessage) SetHeader(key string, value string) {
	for i, header := range m.headers {
		if header.Key == key {
			m.headers[i].Value = value
			return
		}
	}
	m.headers = append(m.headers, Header{Key: key, Value: value})
}
