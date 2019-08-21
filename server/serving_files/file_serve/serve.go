package file_serve

import (
	"log"
	"net/http"
)

func Serve() {
	http.Handle("/", http.FileServer(http.Dir("../../../")))
	log.Fatal(http.ListenAndServe(":3001", nil))
}
