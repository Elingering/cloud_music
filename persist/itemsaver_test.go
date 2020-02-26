package persist

import (
	"cloud_music/model"
	"context"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestItemSaver(t *testing.T) {
	expected := model.SongComment{
		Url:        "https://music.163.com/song?id=1336856777",
		Id:         "1336856777",
		Player:     "老樊",
		SongName:   "我曾",
		Album:      "我曾",
		Nickname:   "Elinger",
		Content:    "这是一个测试文件",
		Time:       "2020-01-01 00:00:01",
		LikedCount: 10000,
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	index := "cloud_music_test"
	table := "comment"
	err = save(client, expected, index, table)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index(index).
		Type(table).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	//t.Logf("%s", resp.Source)
	var actual model.SongComment
	err = jsoniter.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	// 含有 slice 或 map 的结构体不能作比较
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
