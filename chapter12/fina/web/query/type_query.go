package query

import (
	"GopherBook/chapter12/fina/pkg/error"
	"GopherBook/chapter12/fina/web/country"
	"GopherBook/chapter12/fina/web/country_medal"
	"GopherBook/chapter12/fina/web/ping"
	"net/http"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Name: "ping",
			Type: ping.Ping,
			Args: graphql.FieldConfigArgument{
				"data": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "post data for heart beat",
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if p.Args["data"] == nil {
					return nil, error_for_project.NotFound
				}
				var response ping.ResponseForPing
				response = ping.ResponseForPing{
					Code: http.StatusOK,
					Data: p.Args["data"],
				}
				return response, nil
			},
		},
	},
})

// country
func init() {
	Query.AddFieldConfig("countries", &graphql.Field{
		Name: "countries",
		Type: graphql.NewList(country.Country),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"short": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"all": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var param country.GetCountryParam
			controller := country.Default
			if p.Args["all"].(bool) {
				return controller.AllList(param)
			} else {
				param.Name = p.Args["name"].(string)
				if p.Args["short"] != nil {
					param.Short = p.Args["short"].(string)
				}
				return controller.GetList(param)
			}

		},
	})
}

// country medal

func init() {
	Query.AddFieldConfig("countryMedal", &graphql.Field{
		Name: "countryMedal",
		Type: country_medal.CountryMedal,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"year": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var param country_medal.GetCountryMedalParam
			param = country_medal.GetCountryMedalParam{
				Name: p.Args["name"].(string),
				Year: p.Args["year"].(int),
			}
			controller := country_medal.Default
			return controller.GetCountryMedal(param)
		},
	})
}
