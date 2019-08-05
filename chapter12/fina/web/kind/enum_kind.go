package kind

import (
	"GopherBook/chapter12/fina/models"

	"github.com/graphql-go/graphql"
)

var KindEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "kindClass",
	Values: graphql.EnumValueConfigMap{
		"venues": &graphql.EnumValueConfig{
			Value: models.VENUES,
		},
		"discipline": &graphql.EnumValueConfig{
			Value: models.DISCIPLINE,
		},
	},
})
