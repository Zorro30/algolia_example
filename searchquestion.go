package main

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func SearchQuestion(keyword string) (res search.QueryRes) {
	client := search.NewClient(APP_ID, SEARCH_SECRET)
	index := client.InitIndex("questions")
	res, err := index.Search(keyword)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
