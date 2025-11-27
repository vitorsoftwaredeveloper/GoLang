package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// res, err := http.Get("http://google.com.br")

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println(res.Header)

	res, err := soma(10, 2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func soma(a, b int) (int, error) {
	res := a + b

	if res > 10 {
		return 0, errors.New("O resultado é maior que 10")
	}
	return res, nil
}
