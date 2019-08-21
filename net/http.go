package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExampleNet() {
	site := string("http://example.com")
	response, err := http.Get(site)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	responseHtml, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	fmt.Printf("%s \n --------- \n", responseHtml)
	fmt.Println(response)

}
