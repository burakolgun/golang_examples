// Another way to synchronize access to a shared resource is by using a mutex. A mutex is named after the concept of mutual exclusion.
// A mutex is used to create a critical section around code that ensures only one goroutine at a time can execute that code section

//This sample program demonstrates how to use a mutex to define critical sections of code that need synchronous access
package mutual_exclusion

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func ExampleOne() {
	wg.Add(2)

	// Create two goroutines
	go incCounter(1)
	go incCounter(2)

	// Wait for the all goroutines to finish.
	wg.Wait()
	fmt.Printf("Final counter: %d\\n", counter)
}

func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//Only one goroutines through this
		//critical section at a time.
		mutex.Lock()
		// Capture the value of counter
		value := counter

		//Yield the thread and be placed back in queue
		fmt.Println("id  - -  value", id, " --- ", value)
		runtime.Gosched()

		// Increment our local value of counter
		value++
		fmt.Println("id  - -  value", id, " --- ", value)

		// Store the value back into counter.
		counter = value

		mutex.Unlock()
		// Release the lock and allow any waiting goroutine through
	}
}
