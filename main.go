package coderesponsewriter

import "net/http"

type ResponseStatusWriter struct {
	origWriter http.ResponseWriter
	statusCode int
}

func (w *ResponseStatusWriter) WriteHeader(code int) {
	w.statusCode = code
	w.origWriter.WriteHeader(code)
}

func (w *ResponseStatusWriter) Header() http.Header {
	return w.origWriter.Header()
}

func (w *ResponseStatusWriter) Write(body []byte) (int, error) {
	return w.origWriter.Write(body)
}

func (w *ResponseStatusWriter) Code() int {
	return w.statusCode
}

func NewWriter(w http.ResponseWriter) *ResponseStatusWriter {
	return &ResponseStatusWriter{origWriter: w}
}
