package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	now := time.Now()
	now2 := time.Now()
	elapsed := now.Sub(now2)
	fmt.Println("time = ", now, "type = ", reflect.TypeOf(now))
	fmt.Println("distane = ", elapsed)
	fmt.Println("add 2 hours", now.AddDate(1, 2, 3))
	fmt.Println("week day:", now.Weekday())
}
