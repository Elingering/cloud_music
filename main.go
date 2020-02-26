package main

import (
	"cloud_music/distributed/config"
	itemsaver "cloud_music/distributed/persist/client"
	"cloud_music/distributed/rpcsupport"
	worker "cloud_music/distributed/worker/client"
	"cloud_music/engine"
	"cloud_music/music/parser"
	"cloud_music/scheduler"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	// distributed
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)
	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://music.163.com/discover/artist",
		Parser: engine.NewFuncParser(parser.ParseCategoryList, config.ParseCategoryList),
	})
}

func createClientPool(host []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range host {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("Error connecting to %s: %v", h, err)
		} else {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
