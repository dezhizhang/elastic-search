package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	q := elastic.NewMatchQuery("address", "street")
	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	total := result.Hits.TotalHits.Value
	fmt.Printf("%d\n", total)
}
