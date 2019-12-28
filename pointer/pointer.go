package pointer
// A pointer is a value that points to the memory address of another variable.

import (
	"fmt"
)

func point(p *int) {
	*p = 99
}
func ExamplePointer() {
	a := 5
	fmt.Println("a=", a)
	fmt.Println("a's pointer value=", &a)
	b := &a
	fmt.Println("b's value now=", *b)
	*b = 99
	fmt.Println("b's value now=", *b)
	fmt.Println("a's  value now=", a)

	fmt.Println()
	fmt.Println()
	p := 1
	fmt.Println("p =", p)
	point(&p)
	fmt.Println("p =", p)

	fmt.Println("Example two")

	o := 40
	l := &o
	*l++
	fmt.Println(p, "P")
}
