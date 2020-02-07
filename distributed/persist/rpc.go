package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"yyy/model"
	"yyy/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
	Table  string
}

func (s *ItemSaverService) Save(item model.SongComment, result *string) error {
	err := persist.Save(s.Client, item, s.Index, s.Table)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
