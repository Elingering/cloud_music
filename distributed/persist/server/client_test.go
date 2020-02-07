package main

import (
	"testing"
	"time"
	"yyy/distributed/config"
	"yyy/distributed/rpcsupport"
	"yyy/model"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "yyy_test", "comment")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := model.SongComment{
		Url:        "https://music.163.com/song?id=1336856777",
		Id:         "1336856777",
		Player:     "老樊",
		SongName:   "我曾",
		Album:      "我曾",
		Nickname:   "Elinger",
		Content:    "这是一个测试文件",
		Time:       "2020-01-01 00:00:01",
		LikedCount: 10000,
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
