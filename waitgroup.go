package main

import (
	"fmt"
	"sync"
)

func myFunc(wg *sync.WaitGroup) {
	fmt.Println("Executing go routine myFunc")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Executing main goroutine")
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()
	fmt.Println("End of executing main go routine")
}
