package main

import(
	"fmt"
)

func main() {
	var dizi = []string{"ali","ayse","fatma"}
	var assocArray = map[string]int{"ali" : 76, "veli" : 102}

	for i, j := range dizi {
			fmt.Println(i,j);
	}

	fmt.Println("map coming")

	for i, j := range assocArray {
		fmt.Println(i,j)
	}
}