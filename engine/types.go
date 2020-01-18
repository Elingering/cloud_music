package engine

import "yyy/model"

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	//Items    []interface{}
	Items []model.SongComment
}
