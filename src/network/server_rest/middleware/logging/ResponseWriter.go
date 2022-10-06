package logging

import "net/http"

// Insperation: https://stackoverflow.com/questions/64243247/go-gorilla-log-each-request-duration-and-status-code
type LogResponseWriter struct {
	http.ResponseWriter     // http response writer.
	statusCode          int // status code of request.
}

/**
 * Return new LogResponseWriter which include status code.
 */
func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

/**
 * Write header and update status code
 */
func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
