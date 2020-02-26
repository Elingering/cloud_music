package main

import (
	"cloud_music/distributed/config"
	"cloud_music/distributed/persist"
	"cloud_music/distributed/rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
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
