package cmd

import (
	"GopherBook/chapter11/pkg/database"
	"GopherBook/chapter11/pkg/middleware"
	"GopherBook/chapter11/web/model"
	"GopherBook/chapter11/web/vote"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

func Execute() {
	err := RootCMD.Execute()
	if err != nil {
		log.Println("FAIL")
		return
	}
}

var RootCMD = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("Start Execute Command")
		database.EngineGORMInit()
		database.EngineXORMInit()

	},
	Run: func(cmd *cobra.Command, args []string) {
		MigrateByGORM()
		MigrateByXORM()
		http.HandleFunc("/ping", middleware.Logger(func(writer http.ResponseWriter, request *http.Request) {
			var results = make(map[string]interface{})
			results["code"] = http.StatusOK
			results["data"] = "ping"
			writer.Header().Set("Content-type", "application/json;charset=UTF-8")
			enc := json.NewEncoder(writer)
			enc.SetIndent("", "")
			enc.Encode(results)

		}))
		prefix := "/v1/api"

		var v vote.ControllerVote
		http.HandleFunc(fmt.Sprintf("%s/votes", prefix), middleware.Logger(v.GetAllVote))
		http.HandleFunc(fmt.Sprintf("%s/vote", prefix), middleware.Logger(v.VoteHandler))

		//服务启动
		go func() {
			if err := http.ListenAndServe(":8888", nil); err != nil {
				log.Println(err)
			}
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		<-c

		_, cancel := context.WithTimeout(context.Background(), time.Hour)
		defer cancel()
		log.Println("shutting down")
		os.Exit(0)

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		database.EngineMySQLGORM.Close()
		database.EngineMySQLXORM.Close()
	},
}

func MigrateByGORM() {
	database.EngineMySQLGORM.AutoMigrate(model.Vote{}, model.Choice{})

}

func MigrateByXORM() {
	database.EngineMySQLXORM.CreateTables(model.ChoiceByXORM{}, model.VoteByXORM{})

}
