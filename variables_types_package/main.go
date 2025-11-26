package main

import (
	"fmt"
	"variables_types_package/math"
)

// func main() {
// 	a := 10
// 	b := 20
// 	c := a + b
// 	d := a - b

// 	fmt.Printf("%T \n", a)
// 	fmt.Printf("%T \n", b)
// 	fmt.Printf("%T \n", c)
// 	fmt.Printf("%T \n", d)

// }
// ----------------------------------------------------------------------

func main() {
	result := math.Soma(1, 1)

	fmt.Printf("%v \n", result)
	fmt.Printf("%v \n", math.A)
}
