package json

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
import "../../logging/customized_loggers"

type googleResult struct {
	GsearchResultClass string `json:"GsearchResultClass"`
	UnescapedURL       string `json:"unescapedUrl"`
	URL                string `json:"url"`
	VisibleURL         string `json:"visibleUrl"`
	CacheURL           string `json:"cacheUrl"`
	Title              string `json:"title"`
	TitleNoFormatting  string `json:"titleNoFormatting"`
	Content            string `json:"content"`
}

type googleResponse struct {
	ResponseDate struct{
		Result []googleResult `json:"results"`
	} `json:"responseData"`
	ResponseDetails string `json:"responseDetails"`
	Status int `json:"responseStatus"`
}

func ExampleTwo() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// Issue the search against Google.

	resp, err := http.Get(uri)

	if err != nil {
		customized_loggers.Error.Println("Error: ", err)
		return
	}

	defer func() {
		err := resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	var googleResponse googleResponse
	err = json.NewDecoder(resp.Body).Decode(&googleResponse)

	if err != nil {
		customized_loggers.Error.Println("Error: ", err)
		return
	}

	fmt.Println("Response: ", googleResponse)

}
