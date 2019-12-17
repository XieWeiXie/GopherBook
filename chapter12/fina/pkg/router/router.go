package router

import (
	"context"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/middleware"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/web/query"
	"log"
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/graphql-go/graphql"
)

func RegisterSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: query.Query,
		//Mutation: mutation.Mutation,
	})
}

func RegisterHandler() *handler.Handler {
	schema, err := RegisterSchema()
	if err != nil {
		log.Println(err)
		return nil
	}
	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
}

func StartWeb() {
	h := RegisterHandler()
	ctx := context.TODO()
	http.HandleFunc("/graphql", middleware.Logger(ctx, h))
	log.Fatalln(http.ListenAndServe(":2345", nil))
}
