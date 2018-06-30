package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func myFunction2(channel chan int) {

	channel <- 10
	value := <-channel
	fmt.Println(value)
	wg.Done()
}

func main() {
	wg.Add(1)

	ch := make(chan int)

	go myFunction2(ch)
	message := <-ch
	fmt.Println(message)
	wg.Wait()

}
