package blue

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

var Blue = graphql.NewObject(graphql.ObjectConfig{
	Name: "blue",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.Description, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"short": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.Short, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"enName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.EnName, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"chName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.ChName, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if blue, ok := p.Source.(models.BlueSerializer); ok {
					return blue.Image, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
