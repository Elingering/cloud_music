package worker

import (
	"cloud_music/engine"
	"fmt"
)

type CrawlService struct{}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	fmt.Printf("%v", engineReq)
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
