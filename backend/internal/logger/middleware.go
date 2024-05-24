package logger

import "net/http"

type LoggerMiddleware struct {
	Logger HttpLogger
}

func NewLoggerMiddleware() LoggerMiddleware {
	return LoggerMiddleware{
		Logger: HttpLogger{},
	}
}

func (l LoggerMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Logger.Request(r).LogRequest()
		next.ServeHTTP(w, r)
	})
}
