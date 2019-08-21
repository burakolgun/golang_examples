package _defer

//executes after the function
import (
	"fmt"
)

func Example() {

	defer fmt.Println("Done")
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
