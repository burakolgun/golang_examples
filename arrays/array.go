package main

import ( 
	"fmt"
)

func main() {
	var arrayType [15]int
	var i int = 0

	for i<=14 {
		arrayType[i] = i+100
		fmt.Println(arrayType[i])
		i++
	}
}