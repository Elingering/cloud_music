package worker

import (
	"cloud_music/distributed/config"
	"cloud_music/engine"
	"cloud_music/model"
	"cloud_music/music/parser"
	"errors"
	"fmt"
	"log"
	"strings"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items   []model.SongComment
	Request []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Request = append(result.Request, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	desParser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: desParser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Request {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializeing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCategoryList:
		return engine.NewFuncParser(parser.ParseCategoryList, config.ParseCategoryList), nil
	case config.ParsePlayerList:
		return engine.NewFuncParser(parser.ParsePlayerList, config.ParsePlayerList), nil
	case config.ParseSongList:
		return engine.NewFuncParser(parser.ParseSongList, config.ParseSongList), nil
	case config.ParseSong:
		args := strings.Split(p.Args.(string), ",")
		if "" != args[0] && "" != args[1] {
			return parser.NewSongParser(args[0], args[1]), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
