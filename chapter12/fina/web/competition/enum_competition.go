package competition

import (
	"GopherBook/chapter12/fina/models"

	"github.com/graphql-go/graphql"
)

var CompetitionEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "competitionEnum",
	Values: graphql.EnumValueConfigMap{
		"man": &graphql.EnumValueConfig{
			Value: models.MAN,
		},
		"woman": &graphql.EnumValueConfig{
			Value: models.WOMAN,
		},
		"team": &graphql.EnumValueConfig{
			Value: models.TEAM,
		},
	},
})
