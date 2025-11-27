package main

import "fmt"

type Car struct {
	Name string
}

func (c Car) Drive() {
	c.Name = "Ferrari"
	fmt.Println(c.Name, "is driving")
}

func (c *Car) DriveChangeValue() {
	c.Name = "Ferrari"
	fmt.Println(c.Name, "is driving")
}

func main() {
	// ---------------------------------------
	// a := 10
	// var pointer *int = &a

	// *pointer = 50

	// b := &a

	// *b = 60

	// fmt.Println(a)
	// ---------------------------------------

	// variavel := 10
	// abc(&variavel)
	// fmt.Println(variavel)

	// ---------------------------------------

	car := Car{
		Name: "Ford",
	}

	car.Drive()
	fmt.Println(car.Name)

	car.DriveChangeValue()
	fmt.Println(car.Name)
}

func abc(a *int) {
	*a = 200
}
