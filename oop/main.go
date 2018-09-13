package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type developer struct {
	person
	language string
}

type coffee struct {
	caffeineLevel int
}

type human interface {
	sayHello()
}

func (dev developer) coffeeToCode(coffee) string {
	return "coding ..."
}

func (p person) sayHello() {
	fmt.Println("Hello")
}

//interface ex
func saySomething(h human) {
	h.sayHello()
}

func (p person) getAge() {
	fmt.Println(p.age)
}

func (p person) getName() {
	fmt.Println(p.firstName)
}

func (dev developer) sayHello() {
	fmt.Println("Hello, World")
}

//oop ex
func main() {
	person := person{
		"Burak",
		"Olgun",
		27,
	}

	//polymorphism ex
	developer := developer{
		person,
		"NodeJs",
	}

	coffee := coffee{
		55,
	}

	person.sayHello()
	person.getName()
	person.getAge()

	fmt.Println(developer.coffeeToCode(coffee))
	developer.sayHello()
	developer.person.sayHello()
	saySomething(person)
	saySomething(developer)
}
