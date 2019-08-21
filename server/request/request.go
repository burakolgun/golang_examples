package request

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type serve int

var tpl *template.Template

func (s serve) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
		URL         *url.URL
		Header      map[string][]string
	}{
		request.Method,
		request.Form,
		request.URL,
		request.Header,
	}

	_ = tpl.ExecuteTemplate(writer, "index.gohtml", data)
}
func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func Request() {
	var serve serve
	_ = http.ListenAndServe(":3001", serve)
}
