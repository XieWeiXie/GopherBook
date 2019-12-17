package history

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

var History = graphql.NewObject(graphql.ObjectConfig{
	Name: "History",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if history, ok := p.Source.(models.FiNaHistorySerializer); ok {
					return history.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if history, ok := p.Source.(models.FiNaHistorySerializer); ok {
					return history.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if history, ok := p.Source.(models.FiNaHistorySerializer); ok {
					return history.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"year": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if history, ok := p.Source.(models.FiNaHistorySerializer); ok {
					return history.Year, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"detail": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if history, ok := p.Source.(models.FiNaHistorySerializer); ok {
					return history.Detail, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
