package main

import (
	"yyy/engine"
	"yyy/music/parser"
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
	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "https://music.163.com/discover/artist",
		ParserFunc: parser.ParseCategoryList,
	})
}
