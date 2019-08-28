// This sample program demonstrates how to use a buffered channel
// to work on multiple tasks with a predefined
// number of goroutines
package buffered_channels

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

const numberOfGoroutines = 4;
const taskLoad = 10; // Amount of work to process.

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}


func ExampleOne() {
    // Create a buffered channel to manage the taskLoad
	tasks := make(chan string, taskLoad)

	// Launch goroutines to handle the work
const numberOfGoroutines = 4;
	wg.Add(numberOfGoroutines)
	for gr := 1; gr <= numberOfGoroutines; gr++ {
		go worker(tasks, gr);
	}

	// Add bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// Close the channel so the goroutines will quit
	// when the all work is done.
	fmt.Println("Channel is closed")
	close(tasks)


	// Wait for all the work get done
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	// Report that we just returned
	defer wg.Done()

	for {
		// Wait for work to be assigned.
		task, ok := <-tasks

		if !ok {
			// This means channel is empty and closed
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// Display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond * 10)

		// Display the finished the work
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}


