package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

var username string
var passwd string
var url string

func main() {

	flag.StringVar(&username, "u", "elastic", "username")
	flag.StringVar(&passwd, "p", "123456", "passwd")
	flag.StringVar(&url, "url", "http://127.0.0.1:9200", "url")
	flag.Parse()

	log.Println("username:", username, " passwd:", passwd, " url:", url)

	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(username, passwd),
		//elastic.SetURL("http://127.0.0.1:9200"),
		//elastic.SetBasicAuth("elastic", "123456"),
	)
	if err != nil {
		// Handle error
		panic(err)
	}

	log.Println("connect to es success!client:", client)

	p1 := Person{Name: "lmh", Age: 18, Married: false}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
