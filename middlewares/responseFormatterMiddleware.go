package middlewares

import (
	"bytes"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	buf        *bytes.Buffer
	statusCode int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.buf.Write(b)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func formatResponse(body []byte, status int) []byte {
	statusTest := "success"

	if status/100 != 2 {
		statusTest = "failure"
		body = []byte(`"` + string(body) + `"`)
	}

	formatted := `{"status":"` + statusTest + `","data":` + string(body) + `}`
	return []byte(formatted)
}

func ResponseBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		rw := &responseWriter{ResponseWriter: w, buf: new(bytes.Buffer)}

		next.ServeHTTP(rw, r)

		formatedResponse := formatResponse(rw.buf.Bytes(), rw.statusCode)
		w.Write(formatedResponse)

	})
}
