package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"yyy/distributed/config"
	"yyy/distributed/persist"
	"yyy/distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex, config.ElasticTable))
}

func serveRpc(host, index, table string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
		Table:  table,
	})
}
