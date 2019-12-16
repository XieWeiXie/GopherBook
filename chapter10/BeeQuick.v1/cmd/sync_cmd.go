package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"
)

var syncCMD = &cobra.Command{
	Use:   "sync2",
	Short: "xorm.Syn2(model)",
	Run:   sync2,
}

func sync2(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Panic("You should add one argument at least")
		return
	}
	database_v1.DataBaseInit()
	if args[0] == "db" {

		for _, i := range tables() {
			if err := database_v1.BeeQuickDatabase.Sync2(i); err != nil {
				fmt.Println(err)
			}
		}
	}
	if args[0] == "vip" {
		vipMember()
	}

	if args[0] == "drop" {
		database_v1.BeeQuickDatabase.DropTables(new(model_v1.Order))
	}
}

func tables() []interface{} {
	return []interface{}{
		new(model_v1.Account),
		new(model_v1.VipMember),
		new(model_v1.ExchangeCoupon),
		new(model_v1.Account2ExchangeCoupon),
		new(model_v1.RuleForExchangeOrCoupon),
		new(model_v1.Shop),
		new(model_v1.Province),
		new(model_v1.Activity),
		new(model_v1.Activity2Product),
		new(model_v1.Shop2Activity),
		new(model_v1.Product),
		new(model_v1.Product2Tags),
		new(model_v1.Tags),
		new(model_v1.Shop2Tags),
		new(model_v1.Brands),
		new(model_v1.Units),
		new(model_v1.Order),
	}
}

func vipMember() bool {
	if _, err := database_v1.BeeQuickDatabase.Insert(model_v1.DefaultVipMemberRecord()); err != nil {
		return false
	}
	return true
}
