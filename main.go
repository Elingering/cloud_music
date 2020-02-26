package main

import (
	"cloud_music/engine"
	"cloud_music/music/parser"
	"cloud_music/persist"
	"cloud_music/scheduler"
)

func main() {
	// 并发版
	itemChan, err := persist.ItemSaver("cloud_music", "comment")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "https://music.163.com/discover/artist",
		ParserFunc: parser.ParseCategoryList,
	})
}
