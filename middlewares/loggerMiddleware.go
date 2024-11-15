package middlewares

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Incoming Request: ", r.URL.Path)
		fmt.Println("Request Method: ", r.Method)

		// reqBody, err := io.ReadAll(r.Body)
		// if err == nil || r.Method != "GET" {
		// 	fmt.Println("Request Body: ", string(reqBody))
		// }

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)

	})
}
