package danjuan

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter7/assistance"
	"log"
	"time"

	"github.com/tidwall/gjson"
)

// WEB 版

func ParseConf(url string) (Conf, bool) {
	var conf Conf
	content, err := assistance.GetContent(url)
	if err != nil {
		log.Println(err)
		return conf, false
	}
	doc := gjson.ParseBytes(content).Get("data")
	orders := doc.Get("order_model").Array()
	for _, order := range orders {
		var oneOrderModel OrderModel
		oneOrderModel = OrderModel{
			Name: order.Get("name").String(),
			Key:  order.Get("k").String(),
		}
		fmt.Println(oneOrderModel)
		conf.OrderModels = append(conf.OrderModels, oneOrderModel)

	}
	types := doc.Get("type_model").Array()
	for _, t := range types {
		var oneTypeModel TypeModel
		oneTypeModel = TypeModel{
			Name: t.Get("name").String(),
			Key:  t.Get("k").String(),
		}
		fmt.Println(oneTypeModel)
		conf.TypeModels = append(conf.TypeModels, oneTypeModel)
	}
	return conf, true
}

func ParseFund(url string) (FundResult, error) {
	var fr FundResult
	content, err := assistance.GetContent(url)
	if err != nil {
		log.Println(err)
		return fr, err
	}
	doc := gjson.ParseBytes(content).Get("data")
	fr.TotalItems = doc.Get("total_items").Int()
	for _, i := range doc.Get("items").Array() {
		var oneFund Fund
		oneFund = Fund{
			T:       i.Get("f_type").String(),
			Code:    i.Get("fd_code").String(),
			Name:    i.Get("fd_name").String(),
			SFType:  i.Get("sf_type").String(),
			UnitNav: i.Get("unit_nav").String(),
			Yield:   i.Get("yield").String(),
		}
		fmt.Println(oneFund)
		fr.Funds = append(fr.Funds, oneFund)
	}
	return fr, nil
}

func ParseEveryFund(url string) (FundEveryResult, error) {
	var fvr FundEveryResult
	content, err := assistance.GetContent(url)
	if err != nil {
		log.Println(err)
		return fvr, err
	}
	doc := gjson.ParseBytes(content).Get("data")
	fvr.TotalItems = doc.Get("total_items").Int()

	toTime := func(value int64) time.Time {
		return time.Unix(value, 0)
	}
	for _, i := range doc.Get("items").Array() {
		var one FundEvery
		one = FundEvery{
			BeginAt:      toTime(i.Get("begin_at").Int()),
			EvaType:      i.Get("eva_type").String(),
			IndexCode:    i.Get("index_code").String(),
			Name:         i.Get("name").String(),
			PB:           i.Get("pb").Float(),
			PBPercentile: i.Get("pb_percentile").Float(),
			PE:           i.Get("pe").Float(),
			PEPercentile: i.Get("pe_percentile").Float(),
			ROE:          i.Get("roe").Float(),
			CurrentDay:   toTime(i.Get("ts").Int()),
			Yield:        i.Get("yield").Float(),
		}
		fmt.Println(one)
		fvr.Funds = append(fvr.Funds, one)
	}
	return fvr, nil
}

// APP 版

func GetDanJuanResultWithAPP() {
	// API: JSON
}
