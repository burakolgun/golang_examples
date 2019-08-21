package loop

import (
	"fmt"
	"time"
)

func Forever() {
	number := 0

	for {
		number++
		fmt.Println("forever", number)
		time.Sleep(time.Millisecond * 500)
	}
}
