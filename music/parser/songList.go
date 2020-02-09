package parser

import (
	"regexp"
	"yyy/engine"
)

var songRe = regexp.MustCompile(`<a href="/song\?id=([0-9]+)">([^<]+)</a>`)

// 解析每个歌手的歌曲列表
func ParseSongList(contents []byte, _ string) engine.ParseResult {
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
			Url:    domain + "/song?id=" + songId,
			Parser: NewSongParser(songName, songId),
		})
	}
	return result
}

type SongParser struct {
	songName string
	songId   string
}

func (s *SongParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseSong(contents, s.songName, s.songId)
}

func (s *SongParser) Serialize() (name string, args interface{}) {
	return "ParseSong", s.songName + "," + s.songId
}

func NewSongParser(name, id string) *SongParser {
	return &SongParser{
		songId:   id,
		songName: name,
	}
}
