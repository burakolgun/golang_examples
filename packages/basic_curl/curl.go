package basic_curl

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main.exe <url>")
		os.Exit(-1)
	}
}

func ExampleCurl() {
	//  Get a response from the web server.
	//  returns pointer of type http.Request after successfully communicates with the server
	// The http.Request type contains a field named Body, which is an interface value of type io.ReadCloser
	r, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	// Copies from the body to stdoud
	// The io.Copy function accepts values of interface type io.Reader for its second parameter
	// Body field implements the io.Reader interface, so we can pass the Body field into io.Copy and use a web server as our source.
	_, _ = io.Copy(os.Stdout, r.Body)

	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
