package main

import "fmt"

var num int = 9

func main() {
	switch {
	case num == 4:
		fmt.Println("num = 4")
	case num == 5:
		fmt.Println("num = 5")
	case num == 6:
		fmt.Println("num = 6")
	case num == 7:
		fmt.Println("num = 7")
	case num == 8:
		fmt.Println("num = 8")
	default:
		fmt.Println("hicbiri")
	}
}
