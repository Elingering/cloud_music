package parser

import (
	"regexp"
	"yyy/engine"
)

const songRe = `<a href="(/song\?id=[0-9]+)">([^<]+)</a>`

func ParseSongList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(songRe)
	matches := reg.FindAllSubmatch(contents, -1)
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
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: domain + string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseSong(bytes, songName, songId)
			},
		})
		//fmt.Printf("Song: %s, Url: %s\n ", m[2], m[1])
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
