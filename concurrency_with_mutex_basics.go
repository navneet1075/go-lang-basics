package main

import (
	"fmt"
	"sync"
)

var x int
var mut sync.Mutex

func incrementer(wg *sync.WaitGroup) {

	mut.Lock()
	x++
	mut.Unlock()
	wg.Done()
}
func main() {

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementer(&wg)
	}

	wg.Wait()

	fmt.Println("final value of x", x)

}
