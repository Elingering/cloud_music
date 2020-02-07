package main

import (
	"fmt"
	"testing"
	"time"
	"yyy/distributed/config"
	"yyy/distributed/rpcsupport"
	"yyy/distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "https://music.163.com/artist?id=1195036",
		Parser: worker.SerializedParser{
			Name: config.ParseSongList,
			Args: nil,
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
