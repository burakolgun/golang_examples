package main

import (
	"fmt"
	"os"
)

func main() {
	var totalFileSize int64
	var totalFileNumber int
	dir, err := os.Open("./")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dir.Close()

	//-1 like All
	files, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, file := range files {
		fmt.Println(file.Size())
		fmt.Println(file.Name())
		totalFileSize += file.Size()
	}

	totalFileNumber = len(files)

	fmt.Println("------------ Total File Size =", totalFileSize)
	fmt.Println("------------ Total File Number =", totalFileNumber)
}
