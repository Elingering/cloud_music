package main

import (
	"cloud_music/distributed/config"
	"cloud_music/distributed/rpcsupport"
	"cloud_music/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "cloud_music_test", "comment")
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
