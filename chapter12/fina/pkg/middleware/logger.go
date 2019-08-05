package middleware

import (
	"GopherBook/chapter12/fina/pkg/log"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/graphql-go/handler"
)

func Logger(ctx context.Context, h *handler.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		message := fmt.Sprintf("%s | %s | %s | %s", request.Method, request.Host, request.RequestURI, time.Now().Format(time.RFC3339))
		log_for_project.Println(message)
		bodyBytes, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()

		var opts handler.RequestOptions
		json.Unmarshal(bodyBytes, &opts)
		log_for_project.Println(opts.Query)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		h.ContextHandler(ctx, writer, request)
	}
}
