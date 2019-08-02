package cmd

import (
	"GopherBook/chapter12/fina/cmd/data"
	"GopherBook/chapter12/fina/configs"
	"GopherBook/chapter12/fina/pkg/database"

	"github.com/spf13/cobra"
)

var ImportCMD = &cobra.Command{
	Use: "import",
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Step One
		data.RunForSymbol(configs.MatchSymbol)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.MySQL.Close()
	},
}
