package net

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ExamplePostRequest() {
	baseUrl := "http://localhost:4000/api/recover"

	form := url.Values{
		"name":        {"burak"},
		"surname":     {"olgun"},
		"postRequest": {"Hello, World"},
		"from":        {"go"},
	}

	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(baseUrl, "application/x-www-form-urlencoded", body)

	if err != nil {
		panic(err)
	}

	defer rsp.Body.Close()

	bodyByte, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyByte))
}
