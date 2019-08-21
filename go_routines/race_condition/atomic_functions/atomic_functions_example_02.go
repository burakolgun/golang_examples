//this sample program demonstrates how to use atomic package functions Store and Load to provide safe access to numeric types
package atomic_functions

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
)

func ExampleTwo() {

	wg.Add(2)

	//Create two goroutines

	go doWork("A")
	go doWork("B")

	// Give the goroutines time to run
	time.Sleep(1 * time.Second)

	//Safely flag it is time to shutdown.
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown,  1)

	wg.Wait()
}

// doWork simulates a goroutine performing work and checking the Shutdown flag to terminate early.

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

