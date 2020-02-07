package main

import (
	"yyy/distributed/config"
	"yyy/engine"
	"yyy/music/parser"
	"yyy/persist"
	"yyy/scheduler"
)

func main() {
	// seed 歌手分类列表
	// 单机版
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "https://music.163.com/discover/artist",
	//	ParserFunc: parser.ParseCategoryList,
	//})

	// 并发版
	itemChan, err := persist.ItemSaver("yyy", "comment")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "https://music.163.com/discover/artist",
		Parser: engine.NewFuncParser(parser.ParseCategoryList, config.ParseCategoryList),
	})
}
