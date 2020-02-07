package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCategoryList(t *testing.T) {
	contents, err := ioutil.ReadFile("categoryList_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCategoryList(contents, "")
	const resultSize = 3
	expectedUrl := []string{
		"https://music.163.com/discover/artist/cat?id=1001&initial=-1",
		"https://music.163.com/discover/artist/cat?id=1002&initial=-1",
		"https://music.163.com/discover/artist/cat?id=1003&initial=-1",
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
}
