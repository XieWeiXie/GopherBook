package sports

import (
	"GopherBook/chapter12/fina/models"

	"github.com/graphql-go/graphql"
)

var SportEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "sportEnum",
	Values: graphql.EnumValueConfigMap{
		"swimming": &graphql.EnumValueConfig{
			Value: models.SWIMMING,
		},
		"diving": &graphql.EnumValueConfig{
			Value: models.DIVING,
		},
		"highDiving": &graphql.EnumValueConfig{
			Value: models.HIGHDIVING,
		},
		"artisticSwimming": &graphql.EnumValueConfig{
			Value: models.ARTISTICSWIMMING,
		},
		"openWater": &graphql.EnumValueConfig{
			Value: models.OPENWATER,
		},
		"waterPolo": &graphql.EnumValueConfig{
			Value: models.WATERPOLO,
		},
	},
})
