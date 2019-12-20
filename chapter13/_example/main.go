package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type EsQueryByClientAction struct {
	client *elasticsearch.Client
}

type EsQueryByRequestAction struct {
	cat esapi.CatHealthRequest
}

var DefaultClient *elasticsearch.Client

func init() {
	DefaultClient, _ = elasticsearch.NewDefaultClient()
}

func Example1() {
	var es EsQueryByClientAction
	es.client, _ = elasticsearch.NewDefaultClient()
	res, _ := es.client.Cat.Health(es.client.Cat.Health.WithHuman())
	fmt.Println(res.String())
	// [200 OK] 1576825983 07:13:03 es-clustername green 3 3 24 11 0 0 0 0 - 100.0%
}

func Example2() {
	var esRequest EsQueryByRequestAction
	esRequest.cat = esapi.CatHealthRequest{}
	res, _ := esRequest.cat.Do(context.TODO(), DefaultClient)
	fmt.Println(res.String())
	// [200 OK] 1576825983 07:13:03 es-clustername green 3 3 24 11 0 0 0 0 - 100.0%

}
func main() {
	Example1()
	Example2()
}
