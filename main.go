package main

import (
	"yyy/engine"
	"yyy/music/parser"
	"yyy/persist"
	"yyy/scheduler"
)

func main() {
	// 并发版
	itemChan, err := persist.ItemSaver("yyy", "comment")
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
