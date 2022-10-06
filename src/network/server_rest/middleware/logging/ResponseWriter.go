package logging

import "net/http"

// Inspiration: https://stackoverflow.com/questions/64243247/go-gorilla-log-each-request-duration-and-status-code
type logResponseWriter struct {
	http.ResponseWriter     // http response writer.
	statusCode          int // status code of request.
}

// Return new LogResponseWriter which include status code.
func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{ResponseWriter: w}
}

// Write header and update status code.
func (w *logResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
