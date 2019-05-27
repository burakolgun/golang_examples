package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 5
	var b string = "burak"
	var c = 4
	var d = "olgun"

	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	fmt.Println(d, reflect.TypeOf(d))
}
