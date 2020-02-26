package parser

import (
	"cloud_music/engine"
	"regexp"
)

var playerListRe = regexp.MustCompile(`<a .*(/artist\?id=[0-9]+).*[^>]*>([^<]+)</a>`)

// 解析每个分类下的歌手列表
func ParsePlayerList(contents []byte) engine.ParseResult {
	matches := playerListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// todo
	n := 0
	for _, m := range matches {
		n++
		if n > 3 {
			break
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        domain + string(m[1]),
			ParserFunc: ParseSongList,
		})
	}
	return result
}
