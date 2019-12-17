package cmd

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/log"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/router"

	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := database.MySQL.Ping(); err != nil {
			log_for_project.Println(err.Error())
			return
		}
		router.StartWeb()

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.MySQL.Close()
	},
}

func init() {
	RootCMD.AddCommand(SyncCMD)
	RootCMD.AddCommand(ImportCMD)
}
func Execute() {
	if err := RootCMD.Execute(); err != nil {
		panic(err.Error())
	}
}
