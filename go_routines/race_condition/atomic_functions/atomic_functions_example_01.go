package atomic_functions

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Atomic functions provide low level locking mechanism for synchronizing access to integers and pointers

//This sample program demonstrates how to use the atomic
// package to provide safe access to numeric types

var (
	counter int64
	wg sync.WaitGroup
)

func ExampleOne() {
	wg.Add(2)

	//Create two goroutine
	go incCounter(1)
	go incCounter(2)

	//wait for the goroutines to finish.
	wg.Wait()

	//Display the final value of counter.
	fmt.Println("Final counter: ", counter)
}

func incCounter(id int) {
	// Schedule the call to Done to tell main we are done
	defer wg.Done()

	for count := 0; count < 10; count++ {
		// Safely Add One To Counter.
		atomic.AddInt64(&counter, 1)
		fmt.Println("id: ", id)

		// Yield the thread and be placed back in queue
		runtime.Gosched()
	}
}
