package country

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

var Country = graphql.NewObject(graphql.ObjectConfig{
	Name: "Country",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if country, ok := p.Source.(models.CountrySerializer); ok {
					return country.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Name: "createdAt",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if country, ok := p.Source.(models.CountrySerializer); ok {
					return country.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Name: "updatedAt",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if country, ok := p.Source.(models.CountrySerializer); ok {
					return country.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if country, ok := p.Source.(models.CountrySerializer); ok {
					return country.Name, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"short": &graphql.Field{
			Name: "short",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if country, ok := p.Source.(models.CountrySerializer); ok {
					return country.Short, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
