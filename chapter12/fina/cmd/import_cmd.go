package cmd

import (
	"GopherBook/chapter12/fina/cmd/data"
	"GopherBook/chapter12/fina/configs"
	"GopherBook/chapter12/fina/pkg/database"
	"fmt"

	"github.com/spf13/cobra"
)

var ImportCMD = &cobra.Command{
	Use: "import",
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Step One
		fmt.Println(args)
		if len(args) == 0 {
			return
		}

		if args[0] == "symbol" {
			fmt.Println(data.RunForSymbol(configs.MatchSymbol))
		}
		if args[0] == "championship" || args[0] == "ch" {
			fmt.Println(data.RunChampionship(configs.MatchDescription))
		}
		if args[0] == "history" || args[0] == "his" {
			fmt.Println(data.RunFiFaHistory(configs.MatchHistory))
		}
		if args[0] == "brief" || args[0] == "brf" {
			fmt.Println(data.RunFiFaBrief(configs.MatchBrief))
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.MySQL.Close()
	},
}
