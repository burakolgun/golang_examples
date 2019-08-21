package handler

import (
	"fmt"
	"net/http"
)

type human int

//this is handler!
func (h human) ServeHTTP(writer http.ResponseWriter, response *http.Request) {
	fmt.Fprintf(writer, "Any code you want in this func %v", h)
}
func Handler() {
	var transporter = human(1)

	http.ListenAndServe(":8080", transporter)
}
