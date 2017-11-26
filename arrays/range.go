package main

import(
	"fmt"
)
var  employee = map[string]int{"Burak Olgun" : 25, "Yasar Kemal" : 55}

func main() {
	var dizi = []string{"ali","ayse","fatma"}
	var assocArray = map[string]int{"ali" : 76, "veli" : 102}
	for key, value := range dizi {
			fmt.Println(key,value);
	}

	fmt.Println("map coming")

	for key, value := range assocArray {
		fmt.Println(key, value)
	}

	for key, value := range employee {
		fmt.Println(key, value)
	}
}