package query

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/error"
	"GopherBook/chapter12/fina/web/blue"
	"GopherBook/chapter12/fina/web/competition"
	"GopherBook/chapter12/fina/web/country"
	"GopherBook/chapter12/fina/web/country_medal"
	"GopherBook/chapter12/fina/web/fifa"
	"GopherBook/chapter12/fina/web/history"
	"GopherBook/chapter12/fina/web/kind"
	"GopherBook/chapter12/fina/web/ping"
	"GopherBook/chapter12/fina/web/symbol"
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
	Query.AddFieldConfig("countryMedalRank", &graphql.Field{
		Name: "countryMedalRank",
		Type: graphql.NewList(country_medal.CountryMedal),
		Args: graphql.FieldConfigArgument{
			"year": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"sortBy": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var param country_medal.RankCountryMedalParam
			param.Year = p.Args["year"].(int)
			param.SortBy = p.Args["sortBy"].(string)
			controller := country_medal.Default
			return controller.Rank(param)
		},
	})
}

// history

func init() {
	Query.AddFieldConfig("history", &graphql.Field{
		Name: "history",
		Type: history.History,
		Args: graphql.FieldConfigArgument{
			"year": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var param history.GetHistoryParam
			param.Year = p.Args["year"].(int)
			controller := history.Default
			return controller.GetHistory(param)
		},
	})
	Query.AddFieldConfig("histories", &graphql.Field{
		Name: "histories",
		Type: graphql.NewList(history.History),
		Args: graphql.FieldConfigArgument{
			"orderBy": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var param history.GetAllHistoryParam
			param.OrderBy = p.Args["orderBy"].(string)
			controller := history.Default
			return controller.GetAll(param)
		},
	})
}

// fifa

func init() {
	Query.AddFieldConfig("fifa", &graphql.Field{
		Name: "fifa",
		Type: fifa.FiFa,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			controller := fifa.Default
			return controller.GetFiFa()
		},
	})
}

// blue

func init() {
	Query.AddFieldConfig("blues", &graphql.Field{
		Name: "blues",
		Type: graphql.NewList(blue.Blue),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			controller := blue.Default
			type result struct {
				data []models.BlueSerializer
				error
			}
			ch := make(chan result, 1)
			go func() {
				defer close(ch)
				data, err := controller.GetBlues()

				ch <- result{data: data, error: err}
			}()
			r := <-ch
			return r.data, r.error
		},
	})
}

// symbol

func init() {
	Query.AddFieldConfig("symbol", &graphql.Field{
		Name: "symbol",
		Type: symbol.Symbol,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			type result struct {
				data models.SymbolSerializer
				error
			}
			controller := symbol.Default
			ch := make(chan result, 1)
			go func() {
				defer close(ch)
				data, err := controller.GetSymbol()
				ch <- result{data: data, error: err}
			}()
			r := <-ch
			return r.data, r.error
		},
	})
}

// kind

func init() {
	Query.AddFieldConfig("kinds", &graphql.Field{
		Name: "kinds",
		Type: graphql.NewList(kind.Kind),
		Args: graphql.FieldConfigArgument{
			"class": &graphql.ArgumentConfig{
				Type: kind.KindEnum,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			type result struct {
				data []models.KindSerializer
				error
			}
			var param kind.GetKindParam
			param.Class = p.Args["class"].(int)
			controller := kind.Default
			ch := make(chan result, 1)
			go func() {
				defer close(ch)
				data, err := controller.GetKinds(param)
				ch <- result{data: data, error: err}
			}()
			r := <-ch
			return r.data, r.error
		},
	})
}

// competition

func init() {
	Query.AddFieldConfig("competitions", &graphql.Field{
		Name: "competitions",
		Type: competition.Competition,
		Args: graphql.FieldConfigArgument{
			"class": &graphql.ArgumentConfig{
				Type: competition.CompetitionEnum,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			type result struct {
				data []models.CompetitionSerializer
				error
			}
			controller := competition.Default
			var param competition.GetCompetitionParam
			param.Class = p.Args["class"].(int)
			ch := make(chan result, 1)
			go func() {
				defer close(ch)
				data, err := controller.GetCompetitions(param)
				ch <- result{data: data, error: err}
			}()
			r := <-ch
			return r.data, r.error
		},
	})
}
