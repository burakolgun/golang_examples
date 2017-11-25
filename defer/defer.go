package main

import(
	"fmt"
)
var i int = 1
func main() {
	defer fmt.Println("if this function done")	
	for i < 100 {
		fmt.Println(i)
		i++
		if i > 50 && i < 60 {
			defer fmt.Println("if done", i)
		}
	}
}