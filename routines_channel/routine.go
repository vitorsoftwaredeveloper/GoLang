package main

import (
	"fmt"
	"time"
)

func counter(currentType string) {
	for i := 0; i < 5; i++ {
		println(currentType, i)
		time.Sleep(time.Second)
	}
}

// func main is a thread in GO
func main() {
	// counter("Sem go routine")

	// // create thread
	// go counter("Com go routine")

	// fmt.Println("hello 1")
	// fmt.Println("hello 2")

	// time.Sleep(time.Second)
	// fmt.Println("fim")

	// -------------------------------------

	// go counter("a")
	// go counter("b")

	// time.Sleep(time.Second * 10)

	// -------------------------------------

	// runtime.GOMAXPROCS(1)
	// fmt.Println("Start")

	// // GO aborts this thread execution, because the main thread is blocked
	// // since it is waiting for the go routine to finish
	// go func() {
	// 	for {

	// 	}
	// }()

	// time.Sleep(time.Second)
	// fmt.Println("End")

	// ------------------------------------- CHANNEL

	// channel shared between threads
	hello := make(chan string)

	go func() {
		hello <- "Hello world"
	}()

	result := <-hello
	fmt.Println(result)

}
