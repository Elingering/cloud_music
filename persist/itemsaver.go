package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"yyy/model"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			if s, ok := item.(model.SongComment); ok {
				log.Printf("Item saver: got item #%d: %v", itemCount, s)
				itemCount++
				_, err := save(s)
				if err != nil {
					log.Printf("Item saver: error saving item %v: %v", s, err)
				}
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().
		Index("yyy").
		Type("comment").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
