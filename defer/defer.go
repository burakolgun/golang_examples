package main

// function bittikten sonra calisir
import(
	"fmt"
)

func main() {

	defer fmt.Println("Done")
	for i:=0;i < 100; i++ {
		fmt.Println(i)
	}
}