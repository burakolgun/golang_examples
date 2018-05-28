package main

import "fmt"

type Vertex struct {
	X int
	Y int
	C string
}

func main() {
	v := Vertex{1, 2, "qwe"}
	p := &v
	p.C = "aaa"
	fmt.Println(v)
}
