package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseSongList(t *testing.T) {
	contents, err := ioutil.ReadFile("SongList_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseSongList(contents)
	const resultSize = 3
	expectedUrl := []string{
		"https://music.163.com/song?id=1346104327",
		"https://music.163.com/song?id=1383927243",
		"https://music.163.com/song?id=1336856777",
	}
	expectedSong := []string{
		"多想在平庸的生活拥抱你",
		"这一生关于你的风景",
		"我曾",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("Result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrl {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
	if len(result.Requests) != resultSize {
		t.Errorf("Result should have %d items; but had %d", resultSize, len(result.Items))
	}
	for i, song := range expectedSong {
		if result.Items[i].(string) != song {
			t.Errorf("expected song #%d: %s; but was %s", i, song, result.Items[i].(string))
		}
	}
}
