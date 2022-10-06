package logging

import (
	"D7024E/log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Start logging middleware which logs start and end of request.
// Inspiration: https://stackoverflow.com/questions/64243247/go-gorilla-log-each-request-duration-and-status-code
func Start() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.INFO(
				"STATUS - [---] ROUTE - [%s]",
				r.RequestURI)

			startTime := time.Now()
			logRespWriter := newLogResponseWriter(w)

			next.ServeHTTP(logRespWriter, r)

			log.INFO(
				"STATUS - [%d] ROUTE - [%s] DURATION - [%s] ",
				logRespWriter.statusCode,
				r.RequestURI,
				time.Since(startTime).String())
		})
	}
}
