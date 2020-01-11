package parser

import (
	"io/ioutil"
	"testing"
)

func TestParsePlayerList(t *testing.T) {
	contents, err := ioutil.ReadFile("PlayerList_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParsePlayerList(contents)
	const resultSize = 3
	expectedUrl := []string{
		"https://music.163.com/artist?id=12429072",
		"https://music.163.com/artist?id=5781",
		"https://music.163.com/artist?id=4292",
	}
	expectedPlayer := []string{
		"隔壁老樊",
		"薛之谦",
		"李荣浩",
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
	for i, player := range expectedPlayer {
		if result.Items[i].(string) != player {
			t.Errorf("expected player #%d: %s; but was %s", i, player, result.Items[i].(string))
		}
	}
}
