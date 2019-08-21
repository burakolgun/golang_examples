package loop

//Go has just one loop word -> FOR
import (
	"fmt"
)

func ExampleFor() {
	for i := 1; i <= 25; i++ {
		fmt.Println(i)
	}
}
