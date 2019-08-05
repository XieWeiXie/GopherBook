package ping

import (
	"GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

type ResponseForPing struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

var Ping = graphql.NewObject(graphql.ObjectConfig{
	Name: "Ping",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Name: "code",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if ping, ok := p.Source.(ResponseForPing); ok {
					return ping.Code, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"data": &graphql.Field{
			Name: "data",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if ping, ok := p.Source.(ResponseForPing); ok {
					return ping.Data, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
