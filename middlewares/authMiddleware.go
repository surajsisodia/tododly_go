package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"tododly/utils"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqHeader := r.Header.Get("Authorization")

		isValidToken := false

		if reqHeader == "" {
			fmt.Println("Header is empty")
			isValidToken = false
		} else {
			strComp := strings.Split(reqHeader, " ")
			if strComp[0] != "Bearer" {
				fmt.Println("Not a bearer auth")
				isValidToken = false
			} else {
				isValidToken = verifyToken(strings.Split(reqHeader, " ")[1])
			}
		}

		if !isValidToken {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token Authorization Failed"))

		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func verifyToken(tokenString string) bool {
	fmt.Println("Verifying token: ", tokenString)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return utils.SecretKey, nil })

	if err != nil {
		return false
	}

	if token.Valid {
		return true
	}

	return false
}
