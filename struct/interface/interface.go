package _interface

import (
	"fmt"
)

type photo struct {
	name   string
	height int
	width  int
}

type photoInterface interface {
	getName()
	getHeight()
	getWidth()
	getTotalSize() int
}

func photoImpelements(p photoInterface) {
}

func (p photo) getHeight() {
	fmt.Print(p.height)
}

func (p photo) getWidth() {
	fmt.Print(p.width)
}

func (p photo) getName() {
	fmt.Print(p.name)
}

func (p photo) getTotalSize() int {
	return ((p.width * 2) + (p.height * 2))
}

func ExampleInterface() {
	photo := photo{"foto", 50, 100}
	fmt.Println(photo)
	photoImpelements(photo)
	fmt.Println(photo.getTotalSize(), "Total size")
	photo.getHeight()
	fmt.Println(" height")
	photo.getName()
	fmt.Println(" name")
	photo.getWidth()
	fmt.Println(" width")
}
