package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
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
