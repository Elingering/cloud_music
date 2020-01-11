package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseSong(t *testing.T) {
	contents, err := ioutil.ReadFile("song_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseSong(contents, "这一生关于你的风景", "1383927243")
	expectedSong := []string{
		"这一生关于你的风景",
	}
	for i, songName := range expectedSong {
		if result.Items[i].(string) != songName {
			t.Errorf("expected song #%d: %s; but was %s", i, songName, result.Items[i].(string))
		}
	}
}
