package sports

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/error"
	"GopherBook/chapter12/fina/web/competition"

	"github.com/graphql-go/graphql"
)

var Sports = graphql.NewObject(graphql.ObjectConfig{
	Name: "Sports",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"total": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.Total, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"sportClass": &graphql.Field{
			Type: SportEnum,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.SportClass, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"sportClassString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.SportClassString, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"sportName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.SportName, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.Description, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"rule": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.Rule, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"competitions": &graphql.Field{
			Type: graphql.NewList(competition.Competition),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if sport, ok := p.Source.(models.SportSerializer); ok {
					return sport.Competitions, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
