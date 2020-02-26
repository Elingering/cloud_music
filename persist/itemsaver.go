package persist

import (
	"cloud_music/model"
	"context"
	"gopkg.in/olivere/elastic.v5"
)

func Save(client *elastic.Client, item model.SongComment, index, table string) error {
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
