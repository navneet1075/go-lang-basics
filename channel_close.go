package main

import (
	"fmt"
)

func producer(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

}
func main() {

	ch := make(chan int)
	go producer(ch)

	for {
		value, ok := <-ch

		if ok == false {
			fmt.Println("no more data to be received fron the channel")
			break

		}
		fmt.Println("receivced", value)

	}
}
