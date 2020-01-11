package parser

import (
	"bufio"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"yyy/engine"
	"yyy/fetcher"
	"yyy/model"
)

//const songNameRe = `<em class="f-ff2">([^<]+)</em>`
const playerRe = `歌手：<span title="([^"]+)">`
const albumRe = `所属专辑：<a href="/album\?id=[0-9]+" class="s-fc7">([^<]+)</a>`
const commentApi = `https://music.163.com/api/v1/resource/hotcomments/R_SO_4_` //1415078941?limit=&offset=
const pageSize = 100
const page = 0

func ParseSong(contents []byte, songName, songId string) engine.ParseResult {
	regPlayer := regexp.MustCompile(playerRe)
	matchPlayer := regPlayer.FindSubmatch(contents)
	regAlbum := regexp.MustCompile(albumRe)
	matchAlbum := regAlbum.FindSubmatch(contents)
	// 获取歌曲评论
	getComment(songId, songName, string(matchPlayer[1]), string(matchAlbum[1]), pageSize, page)
	result := engine.ParseResult{}
	return result
}

// 记录爬取数据
func getComment(songId, songName, player, album string, pageSize, page int) {
	url := commentApi + songId + "?limit=" + strconv.Itoa(pageSize) + "&offset=" + strconv.Itoa(page)
	json, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", url, err)
	}
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", url, err)
	}
	result := model.Comment{}
	_ = jsoniter.Unmarshal(json, &result)
	if 0 < len(result.HotComments) {
		// 将数据记录文件
		name := strings.Replace(player, "/", "&", -1)
		file, err := os.Create("./data/" + name + "-" + songId + ".txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		defer writer.Flush()
		fmt.Fprintln(writer, "歌曲名称："+songName)
		fmt.Fprintln(writer, "歌手："+player)
		fmt.Fprintln(writer, "专辑："+album)
		fmt.Fprintln(writer, "\n")
		for _, comment := range result.HotComments {
			word := strings.Replace(comment.Content, " ", "", -1)
			word = strings.Replace(comment.Content, "\n", "", -1)
			fmt.Fprintln(writer, "用户："+comment.User.Nickname)
			fmt.Fprintln(writer, "时间："+time.Unix(comment.Time/1000, 0).Format("2006-01-02 15:04:05"))
			fmt.Fprintln(writer, "评论："+word)
			fmt.Fprint(writer, "点赞：")
			fmt.Fprintln(writer, comment.LikedCount)
			fmt.Fprintln(writer, "\n")
		}
	}
	// 获取下一页评论
	if result.HasMore {
		getComment(songId, songName, player, album, pageSize, (page+1)*pageSize)
	}
	return
}
