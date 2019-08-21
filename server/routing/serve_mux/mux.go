package serve_mux

import (
	"io"
	"log"
	"net/http"
)

type hotDog int
type hotCat int

func (hotcat hotCat) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, _ = io.WriteString(writer, "cat")
}

func (hotdog hotDog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, _ = io.WriteString(writer, "dog")
}

func main() {
	var d hotDog
	var c hotCat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	log.Fatal(http.ListenAndServe(":3001", mux))
}
