package symbol

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/error"
	"GopherBook/chapter12/fina/web/blue"

	"github.com/graphql-go/graphql"
)

var Symbol = graphql.NewObject(graphql.ObjectConfig{
	Name: "Symbol",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolText": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolText, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolTextImage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolTextImage, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolTextShort": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolTextShort, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolDescription": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolDescription, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolDescriptionImage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolDescriptionImage, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolDescriptionShort": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolDescriptionShort, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolAnimalImage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolAnimalImage, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolAnimalDescription": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolAnimalDescription, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"symbolAnimalShort": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.SymbolAnimalShort, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"blueVersions": &graphql.Field{
			Type: graphql.NewList(blue.Blue),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if symbol, ok := p.Source.(models.SymbolSerializer); ok {
					return symbol.BlueVersions, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
