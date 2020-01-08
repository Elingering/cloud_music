package main

import (
	"yyy/engine"
	"yyy/music/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://music.163.com/discover/artist",
		ParserFunc: parser.ParseCategoryList,
	})
}
