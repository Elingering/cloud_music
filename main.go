package main

import (
	"yyy/engine"
	"yyy/music/parser"
)

func main() {
	// seed 歌手分类列表
	engine.Run(engine.Request{
		Url:        "https://music.163.com/discover/artist",
		ParserFunc: parser.ParseCategoryList,
	})
}
