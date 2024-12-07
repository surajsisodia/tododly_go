package middlewares

import (
	"context"
	"errors"
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
				username, _ := verifyToken(strings.Split(reqHeader, " ")[1])

				if username != "" {
					isValidToken = true

					ctx := r.Context()
					ctx = context.WithValue(ctx, "username", username)
					r = r.WithContext(ctx)
				}
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

func verifyToken(tokenString string) (string, error) {
	fmt.Println("Verifying token: ", tokenString)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) { return utils.SecretKey, nil })

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	fmt.Println("CLAIMS: ", claims["username"])
	var username string = claims["username"].(string)

	return username, nil
}
