package engine

import "yyy/model"

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	//Items    []interface{}
	Items []model.SongComment
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type NilParser struct{}

func (n NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type ParserFunc func(contents []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
