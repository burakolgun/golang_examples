package json

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleJson() {
	type Feed struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	const dataFile = "./toDos.json"

	file, err := os.Open(dataFile)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var feeds []*Feed

	err = json.NewDecoder(file).Decode(&feeds)

	for index, feed := range feeds {
		fmt.Println(index, "index")
		fmt.Println(feed.Completed, " is completed?")
		fmt.Println(feed.Id, " id")
		fmt.Println(feed.Title, " title")
		fmt.Println(feed.UserId, " userId")
	}
}
