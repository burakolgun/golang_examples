package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (hotcat hotcat) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "cat")
}

func (hotdog hotdog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "dog")
}

func main() {
	 var d hotdog
	 var c hotcat

	 mux := http.NewServeMux()
	 mux.Handle("/dog/", d)
	 mux.Handle("/cat", c)

	 http.ListenAndServe(":3001", mux)
}
