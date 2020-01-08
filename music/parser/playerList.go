package parser

import (
	"regexp"
	"yyy/engine"
)

const playerRe = `<a .*(/artist\?id=[0-9]+).*[^>]*>([^<]+)</a>`

func ParsePlayerList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(playerRe)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// todo
	n := 0
	for _, m := range matches {
		n++
		if n > 3 {
			break
		}
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        domain + string(m[1]),
			ParserFunc: ParseSongList,
		})
		//fmt.Printf("Player: %s, Url: %s\n ", m[2], m[1])
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
