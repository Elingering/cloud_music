package parser

import (
	"regexp"
	"yyy/engine"
)

const domain = "https://music.163.com"
const categoryListRe = `<a href="(/discover/artist/cat\?id=[0-9]{4}).*[^>]*>([^<]+)</a>`

func ParserCategoryList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(categoryListRe)
	matches := re.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", contents)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        domain + string(m[1]),
			ParserFunc: engine.NilParser,
		})
		//fmt.Printf("Category: %s, Url: %s\n ", m[2], m[1])
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
