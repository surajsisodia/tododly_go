package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func RequestBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !(r.Method == http.MethodPost || r.Method == http.MethodPatch) {
			next.ServeHTTP(w, r)
			return
		}

		v, _ := io.ReadAll(r.Body)

		formatedJsonBytes, err := removeWhoColumns(v)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request Format"))
			return
		}

		r.Body = io.NopCloser(bytes.NewReader(formatedJsonBytes))
		next.ServeHTTP(w, r)
	})
}

func removeWhoColumns(orgReqBody []byte) ([]byte, error) {

	var mapBody map[string]interface{}
	err := json.Unmarshal(orgReqBody, &mapBody)

	if err != nil {
		fmt.Println("Request Body Bad Format")
		fmt.Println(err)
		return []byte{}, errors.New("bad format")
	}

	delete(mapBody, "created_at")
	delete(mapBody, "created_by")
	delete(mapBody, "last_updated_at")
	delete(mapBody, "last_updated_by")

	formatedJsonBytes, err := json.Marshal(mapBody)

	if err != nil {
		fmt.Println("Error in converting to JSON bytes")
		fmt.Println(err)
		return []byte{}, errors.New("bad format")
	}

	return formatedJsonBytes, nil
}
