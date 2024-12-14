package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
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
				username, user_id, _ := verifyToken(strings.Split(reqHeader, " ")[1])

				if username != "" || user_id != -1 {
					isValidToken = true

					ctx := r.Context()
					ctx = context.WithValue(ctx, "username", username)
					ctx = context.WithValue(ctx, "user_id", strconv.FormatFloat(user_id, 'f', 0, 64))
					fmt.Println(reflect.ValueOf(user_id))
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

func verifyToken(tokenString string) (string, float64, error) {
	fmt.Println("Verifying token: ", tokenString)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) { return utils.JWT_SECRET_KEY, nil })

	if err != nil || !token.Valid {
		return "", -1, errors.New("invalid token")
	}

	fmt.Println("CLAIMS: ", claims["username"])
	fmt.Println("CLAIMS: ", claims["user_id"])
	username := claims["username"].(string)
	user_id := claims["user_id"].(float64)

	return username, user_id, nil
}
