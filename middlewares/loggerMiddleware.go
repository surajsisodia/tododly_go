package middlewares

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Incoming Request: ", r.URL.Path)
		log.Println("Request Method: ", r.Method)

		if r.Body != nil {
			reqBody, err := io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

			if err == nil {
				log.Println("Request Body: ", string(reqBody))
			}
		}

		next.ServeHTTP(w, r)
	})
}
