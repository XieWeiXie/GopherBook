package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BasicAuth(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userName := request.Header.Get("username")
		password := request.Header.Get("password")
		if userName != "Go" && len(password) == 0 {
			var results = make(map[string]interface{})
			results["code"] = http.StatusBadRequest
			results["error"] = fmt.Sprintf("Add username and password in requests header")
			if err := json.NewEncoder(writer).Encode(results); err != nil {
				log.Println(err)
			}
			return
		}
		next.ServeHTTP(writer, request)
	}
}
