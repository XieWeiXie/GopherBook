package country_medal

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/error"

	"github.com/graphql-go/graphql"
)

var CountryMedal = graphql.NewObject(graphql.ObjectConfig{
	Name: "CountryMedal",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.Id, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.CreatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.UpdatedAt, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"countryId": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.CountryId, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"countryName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.CountryName, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"gold": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.Gold, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"silver": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.Silver, nil
				}
				return nil, error_for_project.NotFound
			},
		},
		"bronze": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if countryMedal, ok := p.Source.(models.CountryMedalSerializer); ok {
					return countryMedal.Bronze, nil
				}
				return nil, error_for_project.NotFound
			},
		},
	},
})
