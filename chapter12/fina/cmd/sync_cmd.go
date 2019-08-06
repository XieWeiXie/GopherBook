package cmd

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/database"

	"github.com/spf13/cobra"
)

var SyncCMD = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s", "-s", "-S", "S"},
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, i := range tables() {
			database.MySQL.Sync2(i)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.MySQL.Close()
	},
}

func tables() []interface{} {
	return []interface{}{
		new(models.Symbol),
		new(models.Blue),
		new(models.FiFaChampionships),
		new(models.Kinds),
		new(models.FiNaHistory),
		new(models.FiNa),
		new(models.Sports),
		new(models.Competitions),
		new(models.CountryMedal),
		new(models.Country),
		new(models.RecordMax),
	}
}
