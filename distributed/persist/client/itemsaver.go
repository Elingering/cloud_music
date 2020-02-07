package client

import (
	"log"
	"yyy/distributed/config"
	"yyy/distributed/rpcsupport"
	"yyy/model"
)

func ItemSaver(host string) (chan interface{}, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			if s, ok := item.(model.SongComment); ok {
				log.Printf("Item saver: got item #%d: %v", itemCount, s)
				itemCount++
				result := ""
				err := client.Call(config.ItemSaverRpc, item, &result)
				if err != nil {
					log.Printf("Item saver: error saving item %v: %v", s, err)
				}
			}
		}
	}()
	return out, err
}
