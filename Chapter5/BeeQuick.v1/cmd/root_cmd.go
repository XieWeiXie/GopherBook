package cmd

import (
	"os"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/router.v1"

	"log"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"

	"github.com/kataras/iris"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "root command",
	Short: "root command",
	Long:  "run web server",
	Run:   runRootCMD,
}

func runRootCMD(cmd *cobra.Command, args []string) {

	database_v1.DataBaseInit()
	iris.RegisterOnInterrupt(func() {
		database_v1.BeeQuickDatabase.Close()
	})
	app := router_v1.ApplyRouter()
	err := app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Execute() {
	rootCMD.AddCommand(syncCMD)
	if err := rootCMD.Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
