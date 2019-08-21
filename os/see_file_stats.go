package os

import (
	"fmt"
	"os"
)

func ExampleSeeFileStats() {
	file, err := os.Open("os/go.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	// Give file stats
	fileStats, err := file.Stat()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("file name  = ", fileStats.Name())
	fmt.Println("is directory?  = ", fileStats.IsDir())
	fmt.Println("file modes  = ", fileStats.Mode())
	fmt.Println("file size  = ", fileStats.Size())
	fmt.Println("change time  = ", fileStats.ModTime())
}
