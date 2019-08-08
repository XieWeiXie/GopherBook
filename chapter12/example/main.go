package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func RegisterSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: Query,
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
func main() {
	h := RegisterHandler()
	http.Handle("/graphql", h)
	log.Fatalln(http.ListenAndServe(":9876", nil))
}

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Type: Ping,
			Args: graphql.FieldConfigArgument{
				"data": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				ch := make(chan Result, 1)
				var result Result
				go func() {
					defer close(ch)
					result.data = ResponsePing{
						Data: p.Args["data"].(string),
						Code: http.StatusOK,
					}
					ch <- result
				}()
				return func() (interface{}, error) {
					r := <-ch
					return r.data, r.error
				}, nil
			},
		},
	},
})

type ResponsePing struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

var Ping = graphql.NewObject(graphql.ObjectConfig{
	Name: "ping",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				ch := make(chan Result, 1)
				go func() {
					defer close(ch)
					if source, ok := p.Source.(ResponsePing); ok {
						ch <- Result{
							data:  source.Data,
							error: nil,
						}
					}
				}()
				return func() (interface{}, error) {
					r := <-ch
					return r.data, r.error
				}, nil

			},
		},
		"code": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				ch := make(chan Result, 1)
				go func() {
					defer close(ch)
					if source, ok := p.Source.(ResponsePing); ok {
						ch <- Result{
							data:  source.Code,
							error: nil,
						}
					}
				}()
				return func() (interface{}, error) {
					r := <-ch
					return r.data, r.error
				}, nil
			},
		},
	},
})

type Result struct {
	data interface{}
	error
}
