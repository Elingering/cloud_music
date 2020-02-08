package parser

import (
	"regexp"
	"yyy/engine"
)

var songRe = regexp.MustCompile(`<a href="/song\?id=([0-9]+)">([^<]+)</a>`)

// 解析每个歌手的歌曲列表
func ParseSongList(contents []byte) engine.ParseResult {
	matches := songRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// todo
	n := 0
	for _, m := range matches {
		n++
		if n > 3 {
			break
		}
		// 拷贝 songName songId 否则 ParserFunc 始终会是最后一个值
		songName := string(m[2])
		songId := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: domain + "/song?id=" + songId,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseSong(bytes, songName, songId)
			},
		})
	}
	return result
}
