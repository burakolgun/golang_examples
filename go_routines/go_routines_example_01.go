package go_routines

import (
	"fmt"
	"time"
)

func sayHelloForMyName(name string) {
	for i := 0; i < 2000; i++ {
		time.Sleep(10000000)
		fmt.Println(i, "--", name, "--")
	}
}

func sleep(name string) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(name)
	}
}

func ExampleOne() {
	go sayHelloForMyName("--")
	go sayHelloForMyName("+++")
	go sleep("???????????????????????????????????????")
	go sayHelloForMyName("000")

	a := ""
	fmt.Scanln(&a)
	fmt.Println(a)
}
