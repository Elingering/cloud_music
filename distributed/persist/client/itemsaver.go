package client

import (
	"log"
	"yyy/distributed/config"
	"yyy/distributed/rpcsupport"
	"yyy/model"
)

func ItemSaver(host string) (chan model.SongComment, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan model.SongComment)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, err
}
