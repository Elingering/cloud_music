package parser

import (
	"regexp"
	"yyy/engine"
)

const domain = "https://music.163.com"
const categoryListRe = `<a href="(/discover/artist/cat\?id=[0-9]{4}).*[^>]*>([^<]+)</a>`

func ParseCategoryList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(categoryListRe)
	matches := re.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", contents)
	result := engine.ParseResult{}
	// todo
	n := 0
	for _, m := range matches {
		n++
		if n > 3 {
			break
		}
		// 拼接完整地址 initial = [-10a-z]
		result.Items = append(result.Items, string(m[2])+"-1")
		result.Requests = append(result.Requests, engine.Request{
			Url:        domain + string(m[1]) + "&initial=-1",
			ParserFunc: ParsePlayerList,
		})
		// 数据太多，先测每个分类下的热门歌手
		//result.Items = append(result.Items, string(m[2]) + "0")
		//result.Requests = append(result.Requests, engine.Request{
		//	Url:        domain + string(m[1]) + "&initial=0",
		//	ParserFunc: ParsePlayerList,
		//})
		//for n := 65; n <= 90; n++ {
		//	result.Items = append(result.Items, string(m[2]) + string(n))
		//	result.Requests = append(result.Requests, engine.Request{
		//		Url:        domain + string(m[1]) + "&initial=" + string(n),
		//		ParserFunc: ParsePlayerList,
		//	})
		//}
		//fmt.Printf("Category: %s, Url: %s\n ", m[2], m[1])
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
