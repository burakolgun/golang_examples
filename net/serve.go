package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/second", secondPage)
	http.ListenAndServe(":1983", nil)
}

func rootPage(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	io.WriteString(response, `
		<html>
			<body>
			Welcome Page
			</body>
		<html>
	`)
}

func secondPage(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	io.WriteString(response, `
		<html>
			<body>
			Second Page
			</body>
		<html>
	`)
}
