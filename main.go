package main

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Questions struct {
	ID          int64  `json:"objectID"`
	Question    string `json:"question"`
	Description string `json:"description"`
	CreatedAt   int    `json:"timestamp"`
	Bounty      bool   `json:"bounty"`
}

func main() {
	client := search.NewClient(APP_ID, ADMIN_SECRET)
	index := client.InitIndex("questions")

	res, err := index.SaveObjects([]Questions{
		{ID: 2, Question: "How to check payslip?", Description: "How check payslip",
			CreatedAt: 987656789, Bounty: false},
	})

	fmt.Println(res, err)

	// searching some keyword on algolia
	resp := SearchQuestion("Terminal")
	fmt.Println(resp)
}
