package parser

import (
	"cloud_music/engine"
	"cloud_music/fetcher"
	"cloud_music/model"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var playerRe = regexp.MustCompile(`歌手：<span title="([^"]+)">`)
var albumRe = regexp.MustCompile(`所属专辑：<a href="/album\?id=[0-9]+" class="s-fc7">([^<]+)</a>`)

const commentApi = domain + `/api/v1/resource/hotcomments/R_SO_4_`
const pageSize = 100
const page = 0

// 解析歌曲信息
func parseSong(contents []byte, songName, songId string) engine.ParseResult {
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
func getComment(songId, songName, player, album string, pageSize, page int) []model.SongComment {
	url := commentApi + songId + "?limit=" + strconv.Itoa(pageSize) + "&offset=" + strconv.Itoa(page*pageSize)
	json, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", url, err)
	}
	result := model.RawData{}
	_ = jsoniter.Unmarshal(json, &result)
	var commentItem []model.SongComment
	var songComment model.SongComment
	songComment.Url = domain + "/song?id=" + songId
	songComment.SongName = songName
	songComment.Player = player
	songComment.Album = album
	if http.StatusOK == result.Code && 0 < len(result.HotComments) {
		for _, comment := range result.HotComments {
			word := strings.Replace(comment.Content, " ", "", -1)
			word = strings.Replace(comment.Content, "\n", "", -1)
			songComment.Nickname = comment.User.Nickname
			songComment.Time = time.Unix(comment.Time/1000, 0).Format("2006-01-02 15:04:05")
			songComment.Content = word
			songComment.Id = strconv.Itoa(comment.CommentId)
			songComment.LikedCount = comment.LikedCount
			commentItem = append(commentItem, songComment)
		}
	} else {
		songComment.Id = songName
		songComment.Content = songName
		commentItem = append(commentItem, songComment)
	}
	// 获取下一页评论
	if result.HasMore {
		getComment(songId, songName, player, album, pageSize, page+1)
	}
	return commentItem
}
