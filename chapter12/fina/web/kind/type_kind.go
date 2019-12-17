package kind

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

var Kind = graphql.NewObject(graphql.ObjectConfig{
	Name: "Kind",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if kind, ok := p.Source.(models.KindSerializer); ok {
					return kind.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if kind, ok := p.Source.(models.KindSerializer); ok {
					return kind.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if kind, ok := p.Source.(models.KindSerializer); ok {
					return kind.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if kind, ok := p.Source.(models.KindSerializer); ok {
					return kind.Name, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"class": &graphql.Field{
			Type: KindEnum,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if kind, ok := p.Source.(models.KindSerializer); ok {
					return kind.Class, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"classString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if kind, ok := p.Source.(models.KindSerializer); ok {
					return kind.ClassString, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
