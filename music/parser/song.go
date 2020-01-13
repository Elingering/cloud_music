package parser

import (
	"bufio"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"yyy/engine"
	"yyy/fetcher"
	"yyy/model"
	"yyy/tool"
)

var playerRe = regexp.MustCompile(`歌手：<span title="([^"]+)">`)
var albumRe = regexp.MustCompile(`所属专辑：<a href="/album\?id=[0-9]+" class="s-fc7">([^<]+)</a>`)

const commentApi = domain + `/api/v1/resource/hotcomments/R_SO_4_`
const pageSize = 100
const page = 0

// 解析歌曲信息
func ParseSong(contents []byte, songName, songId string) engine.ParseResult {
	matchPlayer := playerRe.FindSubmatch(contents)
	matchAlbum := albumRe.FindSubmatch(contents)
	// 获取歌曲评论
	content := getComment(songId, songName, string(matchPlayer[1]), string(matchAlbum[1]), pageSize, page)
	result := engine.ParseResult{
		Items: content,
	}
	return result
}

// 记录爬取数据
func getComment(songId, songName, player, album string, pageSize, page int) []interface{} {
	url := commentApi + songId + "?limit=" + strconv.Itoa(pageSize) + "&offset=" + strconv.Itoa(page)
	json, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", url, err)
	}
	result := model.Song{}
	_ = jsoniter.Unmarshal(json, &result.HotComment)
	if http.StatusOK == result.HotComment.Code && 0 < len(result.HotComment.Content) {
		// 将数据记录文件
		name := strings.Replace(player, "/", "&", -1)
		// 创建目录 ~/data
		home, err := tool.Home()
		if err != nil {
			panic(err)
		}
		dataPath := home + "/data"
		err = tool.MakeDir(dataPath)
		if err != nil {
			log.Printf("dataPath: error "+"making dir %s: %v", dataPath, err)
			panic("making dir error: ~/data")
		}
		// Golang的相对路径是相对于执行命令时的目录，所以用绝对路径。否则执行测试文件会找不到文件
		file, err := os.Create(dataPath + "/" + name + "-" + songId + ".txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		defer writer.Flush()
		fmt.Fprintln(writer, "歌曲名称："+songName)
		fmt.Fprintln(writer, "歌手："+player)
		fmt.Fprintln(writer, "专辑："+album)
		fmt.Fprintln(writer, "")
		for _, comment := range result.HotComment.Content {
			word := strings.Replace(comment.Content, " ", "", -1)
			word = strings.Replace(comment.Content, "\n", "", -1)
			fmt.Fprintln(writer, "用户："+comment.User.Nickname)
			fmt.Fprintln(writer, "时间："+time.Unix(comment.Time/1000, 0).Format("2006-01-02 15:04:05"))
			fmt.Fprintln(writer, "评论："+word)
			fmt.Fprint(writer, "点赞：")
			fmt.Fprintln(writer, comment.LikedCount)
			fmt.Fprintln(writer, "")
		}
	}
	// 获取下一页评论
	if result.HotComment.HasMore {
		getComment(songId, songName, player, album, pageSize, (page+1)*pageSize)
	}
	var item []interface{}
	item = append(item, songName)
	return item
}
