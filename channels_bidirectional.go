package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func myFunction(ch chan int) {
	ch <- 10
	//value := <-ch
	//fmt.Println(value)

}

func main() {

	fmt.Println("start of main")

	ch := make(chan int)

	go myFunction(ch)

	messages := <-ch

	fmt.Println(messages)

}
