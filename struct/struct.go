package main

import (
	"fmt"
)

type human struct {
	name    string
	age     int
	agility int
	job     string
}

type circle struct {
	radius int
}

func (a *circle) area() int {
	return a.radius * a.radius * 3 // pi 3
}

func (a *circle) perimeter() int {
	return 2 * 3 * a.radius
}

func main() {
	firstHuman := human{}
	firstHuman.name = "John"
	firstHuman.agility = 4
	firstHuman.job = "waba laba dub dub"
	firstHuman.age = 88

	secondHuman := human{name: "mahmut", age: 25, agility: 4, job: "no job"}

	fmt.Println(secondHuman)

	circle := circle{4}

	fmt.Println("area = ", circle.area())
	fmt.Println("perimeter = ", circle.perimeter())

}
