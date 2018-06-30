package main

import (
	"fmt"
)

func f(c chan string) {
	fmt.Println("start of goroutine")
	c <- "hello"
}
func main() {
	fmt.Println("start of main")
	c := make(chan string)
	go f(c)
	value := <-c
	fmt.Println(value)
}
