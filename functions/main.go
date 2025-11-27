package main

import "fmt"

// method

type Car struct {
	Name string
}

func (c Car) Drive() {
	fmt.Println(c.Name, "is driving")
}

func main() {

	carro := Car{
		Name: "Ferrari",
	}

	carro.Drive()

	resultado := func(x ...int) func() int {
		res := 0

		for _, value := range x {
			res += value
		}

		return func() int {
			return res * res
		}
	}

	// resultado := soma(10, 20)
	// resultado := SomaTudo(10, 20, 4)
	fmt.Println(resultado(1, 2, 3)())
}

func soma(a int, b int) (result int) {
	result = a + b
	return
}

func SomaTudo(x ...int) int {
	resultado := 0

	for _, value := range x {
		resultado += value
	}

	return resultado
}
