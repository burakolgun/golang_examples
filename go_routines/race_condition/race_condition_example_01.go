package race_condition

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func ExampleOne() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capture the value of counter
		value := counter

		//Yield the thread and be placed back in queue
		runtime.Gosched()

		//Increment our local value of Counter
		value++

		//Store the value back into Counter
		counter = value
	}
}
