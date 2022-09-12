package loggingMiddleware

// https://stackoverflow.com/questions/64243247/go-gorilla-log-each-request-duration-and-status-code

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type LogMiddleware struct {
	InfoLogger *log.Logger
}

func NewLogMiddleware(InfoLogger *log.Logger) *LogMiddleware {
	return &LogMiddleware{InfoLogger: InfoLogger}
}

func (m *LogMiddleware) Start() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			logRespWriter := NewLogResponseWriter(w)

			m.InfoLogger.Printf(
				"STATUS - [---] ROUTE - [%s]",
				r.RequestURI)

			next.ServeHTTP(logRespWriter, r)

			m.InfoLogger.Printf(
				"STATUS - [%d] ROUTE - [%s] DURATION - [%s] ",
				logRespWriter.statusCode,
				r.RequestURI,
				time.Since(startTime).String())
		})
	}
}
