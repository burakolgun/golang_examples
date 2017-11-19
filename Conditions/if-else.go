package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	hour := now.Hour()

	fmt.Println("now :", now)
	fmt.Println("hour :", hour)

	if hour <= 12 {
		fmt.Println("ogleden once");
	} else if hour > 12 && hour < 17 {
		fmt.Println("Ogleden Sonra")
	} else {
		fmt.Println(hour)
	}
}