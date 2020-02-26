package parser

import (
	"cloud_music/model"
	"io/ioutil"
	"testing"
)

func TestParseSong(t *testing.T) {
	contents, err := ioutil.ReadFile("song_test_data.html")
	if err != nil {
		panic(err)
	}
	// 热评是动态的，此测试不保证随时可用
	result := ParseSong(contents, "这一生关于你的风景", "1383927243")
	expected := model.SongComment{
		Url: "https://music.163.com/song?id=1383927243",
		//Id:         "1597116361",
		Player:   "隔壁老樊",
		SongName: "这一生关于你的风景",
		Album:    "这一生关于你的风景",
		//Nickname:   "隔壁老樊c",
		//Content:    "我在贩卖日光 你能不能来我身旁",
		//Time:       "2019-08-16 00:05:38",
		//LikedCount: 171742,
	}
	var actual model.SongComment
	actual.Url = result.Items[0].Url
	actual.Player = result.Items[0].Player
	actual.SongName = result.Items[0].SongName
	actual.Album = result.Items[0].Album
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
