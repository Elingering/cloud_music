package parser

import (
	"cloud_music/engine"
	"regexp"
)

const domain = "https://music.163.com"

var categoryListRe = regexp.MustCompile(`<a href="(/discover/artist/cat\?id=[0-9]{4}).*[^>]*>([^<]+)</a>`)

// 解析歌手分类列表
func ParseCategoryList(contents []byte) engine.ParseResult {
	matches := categoryListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// todo
	n := 0
	for _, m := range matches {
		n++
		if n > 3 {
			break
		}
		// 拼接完整地址 initial = [-10a-z]
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
		// 按姓名首字母获取歌手列表
		//for n := 65; n <= 90; n++ {
		//	result.Items = append(result.Items, string(m[2]) + string(n))
		//	result.Requests = append(result.Requests, engine.Request{
		//		Url:        domain + string(m[1]) + "&initial=" + string(n),
		//		ParserFunc: ParsePlayerList,
		//	})
		//}
	}
	return result
}
