package records

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

var Records = graphql.NewObject(graphql.ObjectConfig{
	Name: "records",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"eventName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.EventName, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"record": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.Record, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"countryId": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.CountryId, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"countryName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.CountryName, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"date": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.Date, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"location": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.Location, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"competitionClass": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.CompetitionClass, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"competitionClassString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.CompetitionClassString, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"sportClass": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.SportClass, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"sportClassString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.SportClassString, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if record, ok := p.Source.(models.RecordsMaxSerializer); ok {
					return record.Name, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
