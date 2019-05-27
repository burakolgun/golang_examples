package main

import (
	"fmt"
)

func main() {
	var num int = 2

	switch num {
	case 1:
		fmt.Println("num:", num)
	case 2:
		fmt.Println("num:", num)
	case 3:
		fmt.Println("num", num)
	default:
		fmt.Println("undefined case")
	}
}
