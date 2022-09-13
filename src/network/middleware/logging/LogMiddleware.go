package loggingMiddleware

// https://stackoverflow.com/questions/64243247/go-gorilla-log-each-request-duration-and-status-code

import (
	"D7024E/log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			logRespWriter := NewLogResponseWriter(w)

			log.INFO(
				"STATUS - [---] ROUTE - [%s]",
				r.RequestURI)

			next.ServeHTTP(logRespWriter, r)

			log.INFO(
				"STATUS - [%d] ROUTE - [%s] DURATION - [%s] ",
				logRespWriter.statusCode,
				r.RequestURI,
				time.Since(startTime).String())
		})
	}
}
