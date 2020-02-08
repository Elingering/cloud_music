package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"yyy/model"
)

func ItemSaver(index, table string) (chan model.SongComment, error) {
	out := make(chan model.SongComment)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return out, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err = save(client, item, index, table)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, err
}

func save(client *elastic.Client, item model.SongComment, index, table string) error {
	_, err := client.Index().
		Index(index).
		Type(table).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
