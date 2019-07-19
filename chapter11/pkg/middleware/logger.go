package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		format := fmt.Sprintf("[ http_log ]: %s | %s | %s | %s", request.URL, request.Host, request.Method, time.Now().Format(time.RFC3339))
		log.Printf("%s", Red(format))
		next.ServeHTTP(writer, request)
	}
}
func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}
