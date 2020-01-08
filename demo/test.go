package main

import (
	"fmt"
	"regexp"
)

const url = `<li class="sml">
<a href="/artist?id=12085016" class="nm nm-icn f-thide s-fc0" title="接个吻，开一枪的音乐">接个吻，开一枪</a>
<a class="f-tdn" href="/user/home?id=114305616" title="接个吻，开一枪的个人主页"><i class="u-icn u-icn-5"></i></a>
</li>`

const playerRe = `<a.*/artist\?id=[0-9]+.*[^>]*>([^<]+)</a>`

func main() {
	reg := regexp.MustCompile(playerRe)
	matches := reg.FindAllSubmatch([]byte(url), -1)
	for _, m := range matches {
		fmt.Printf("%s", m[1])
	}
	//body, _ := fetcher.Fetch("https://music.163.com/artist?id=12429072")
	//fmt.Printf("%s", body)
}
