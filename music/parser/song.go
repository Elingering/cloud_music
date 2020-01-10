package parser

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"regexp"
	"time"
	"yyy/engine"
	"yyy/fetcher"
	"yyy/model"
)

//const songNameRe = `<em class="f-ff2">([^<]+)</em>`
const playerRe = `歌手：<span title="([^"]+)">`
const albumRe = `所属专辑：<a href="/album\?id=[0-9]+" class="s-fc7">([^<]+)</a>`
const commentApi = `https://music.163.com/api/v1/resource/hotcomments/R_SO_4_` //1415078941?limit=&offset=

func ParseSong(contents []byte, songName, songId string) engine.ParseResult {
	regPlayer := regexp.MustCompile(playerRe)
	matchPlayer := regPlayer.FindSubmatch(contents)
	regAlbum := regexp.MustCompile(albumRe)
	matchAlbum := regAlbum.FindSubmatch(contents)
	// 获取歌曲评论
	var hotComment []model.Comment
	url := commentApi + songId
	hotComment = getComment(url)
	// 填充数据
	data := model.Comment{
		SongName:   songName,
		Player:     string(matchPlayer[1]),
		Album:      string(matchAlbum[1]),
		HotComment: hotComment,
	}
	fmt.Printf("%v ", data)
	result := engine.ParseResult{}
	return result
}

func getComment(url string) []model.Comment {
	json, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", url, err)
	}
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", url, err)
	}
	result := map[string]interface{}{}
	_ = jsoniter.Unmarshal(json, &result)
	//result = result.(map[string]interface{})
	data := model.Comment{
		Item: func() []model.CommentItem {
			item := model.CommentItem{
				UserName:  "",
				CreatedAt: time.Time{},
				Content:   "",
				Like:      0,
			}
			item = append(item, item)
			return item
		},
		HasMore: result["hasMore"].(bool),
		Total:   result["total"].(int),
		Code:    result["code"].(int),
	}
	return data
}
