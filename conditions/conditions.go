package conditions

import (
	"fmt"
	"time"
)

func BetterIfThanElseExample() {
	var num = 9

	switch num < 1 {
	case num == 4:
		fmt.Println("num = 4")
	case num == 5:
		fmt.Println("num = 5")
	case num == 6:
		fmt.Println("num = 6")
	case num == 7:
		fmt.Println("num = 7")
	case num == 8:
		fmt.Println("num = 8")
	default:
		fmt.Println("None of above")
	}
}

func IfElseExample() {
	now := time.Now()
	hour := now.Hour()

	fmt.Println("now :", now)
	fmt.Println("hour :", hour)

	if hour <= 12 {
		fmt.Println("ogleden once")
	} else if hour > 12 && hour < 17 {
		fmt.Println("Ogleden Sonra")
	} else {
		fmt.Println(hour)
	}
}

func SwitchExample() {
	var num = 2

	switch num {
	case 1:
		fmt.Println("num:", num)
	case 2:
		fmt.Println("num:", num)
	case 3:
		fmt.Println("num", num)
	default:
		fmt.Println("undefined case")
	}
}

func WhileExample() {
	i := 1

	for i <= 50 {
		fmt.Println("i:", i)
		i++
	}
}
