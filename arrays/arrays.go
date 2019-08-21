package arrays

import (
	"fmt"
)

func ArrayExample() {
	var arrayType [15]int
	var i = 0

	for i <= 14 {
		arrayType[i] = i + 100
		fmt.Println(arrayType[i])
		i++
	}
}

func MapExample() {
	var assocArray = make(map[string]int)
	assocArray["burak"] = 25
	assocArray["baris"] = 16
	assocArray["nazim"] = 62
	assocArray["leyla"] = 44

	fmt.Println(assocArray)

	delete(assocArray, "nazim")
	fmt.Println("nazim deleted", assocArray)
}

func RangeExample() {
	var employee = map[string]int{"Burak Olgun": 25, "Yasar Kemal": 55}

	var dizi = []string{"ali", "ayse", "fatma"}
	var assocArray = map[string]int{"ali": 76, "veli": 102}
	for key, value := range dizi {
		fmt.Println(key, value)
	}

	fmt.Println("map coming")

	for key, value := range assocArray {
		fmt.Println(key, value)
	}

	for key, value := range employee {
		fmt.Println(key, value)
	}
}
