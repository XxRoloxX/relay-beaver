package logger

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

type HttpLogger struct {
	logger RequestLogger
}

type RequestLogger struct {
	Request *http.Request
}

func (logger HttpLogger) Request(request *http.Request) RequestLogger {
	logger.logger = RequestLogger{Request: request}
	return logger.logger
}

func (logger RequestLogger) formatMessage(message string) string {
	return fmt.Sprintf("Request %s %s %s %s",
		logger.Request.Method,
		logger.Request.URL.Path,
		logger.Request.RemoteAddr,
		message)
}

func (logger RequestLogger) Info(message string) {
	log.Info().Msg(logger.formatMessage(message))
}

func (logger RequestLogger) LogRequest() {
	log.Info().Msg(logger.formatMessage(""))
}

func (logger RequestLogger) Error(errorMessage interface{}) {
	log.Error().Msg(logger.formatMessage(fmt.Sprintf("%v", errorMessage)))
}

func (logger RequestLogger) Warn(message string) {
	log.Warn().Msg(logger.formatMessage(message))
}

func (logger RequestLogger) Debug(message string) {
	log.Debug().Msg(logger.formatMessage(message))
}
