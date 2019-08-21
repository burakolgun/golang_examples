package os

import (
	"fmt"
	"os"
)

func ExampleReadFile() {
	file, err := os.Open("os/go.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()
	fileStats, err := file.Stat()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	CreateBytes := make([]byte, fileStats.Size())
	ReadByte, err := file.Read(CreateBytes)

	ReadText := string(CreateBytes)

	fmt.Println(ReadText)
	fmt.Println(ReadByte)
}
