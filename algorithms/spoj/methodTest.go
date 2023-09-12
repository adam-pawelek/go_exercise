package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) modifyValue() {
	r.width = 10
	r.height = -1
}

func (r rect) tryToModifyValue() {
	r.width = 10
	r.height = -1
}

func main() {

	r := rect{width: 10, height: 5}
	r.tryToModifyValue()

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	fmt.Println("\n\nNew code")

	r = rect{width: 10, height: 5}
	r.modifyValue()

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp = &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
