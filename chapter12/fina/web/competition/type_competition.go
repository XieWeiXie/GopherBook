package competition

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

var Competition = graphql.NewObject(graphql.ObjectConfig{
	Name: "Competition",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if competition, ok := p.Source.(models.CompetitionSerializer); ok {
					return competition.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if competition, ok := p.Source.(models.CompetitionSerializer); ok {
					return competition.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if competition, ok := p.Source.(models.CompetitionSerializer); ok {
					return competition.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"competitionClass": &graphql.Field{
			Type: CompetitionEnum,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if competition, ok := p.Source.(models.CompetitionSerializer); ok {
					return competition.CompetitionClass, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"competitionClassString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if competition, ok := p.Source.(models.CompetitionSerializer); ok {
					return competition.CompetitionClassString, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"detail": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if competition, ok := p.Source.(models.CompetitionSerializer); ok {
					return competition.Detail, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
