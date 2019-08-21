package pointer

import "fmt"

type Vertex struct {
	X int
	Y int
	C string
}

func ExampleStructPointer() {
	v := Vertex{1, 2, "qwe"}
	p := &v
	p.C = "aaa"
	fmt.Println(v)
}
