package hello_world

import (
	"../import_package"
	"fmt"
)

func ExampleHi() {
	fmt.Println("Hello, World")
	import_package.Now()
}
