package main

import (
	"fmt"
)

func main() {
	var assocArray = make(map[string]int)
	assocArray["burak"] = 25;
	assocArray["baris"] = 16;
	assocArray["nazim"] = 62;
	assocArray["leyla"] = 44;

	fmt.Println(assocArray);

	delete(assocArray,"nazim")
	fmt.Println("nazim deleted",assocArray);
	



}