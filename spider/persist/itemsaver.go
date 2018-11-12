package persist

import (
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"context"
	"log"
	"github.com/xxg3053/go-spider/spider/engine"
	"github.com/pkg/errors"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	client := getClient()
	go func() {
		itemCount := 0
		for {
			item := <- out
			fmt.Printf("Item Saver: Got item #%d: %v\n", itemCount, item)
			itemCount++
			err := save(*client, item)
			if err != nil{
				log.Printf("Item saver: error saveing item %v : %v", item, err)
			}

		}
	}()
	return out;
}

var elUrl = "http://172.16.92.218:9200"

func getClient()  *elastic.Client{
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(elUrl),
	)
	if err != nil{
		panic(err)
	}

	return client
}

func save(client elastic.Client, item engine.Item) error{

	if item.Type == ""{
		return  errors.New("must supply Type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).BodyJson(item)

	if item.Id != ""{
		indexService.Id(item.Id)
	}
	_,err := indexService.Do(context.Background())
	if err != nil{
		return err
	}
	//fmt.Printf("%+v", resp)

	return nil

}
