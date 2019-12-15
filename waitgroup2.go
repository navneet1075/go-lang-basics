package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{"https://google.com", "https://yahoo.com", "https://microsoft.com"}

func fetchStatus(response http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(response, "%v+n", err)
			}
			fmt.Fprintf(response, "%v\n", resp)
			wg.Done()
		}(url)
		wg.Wait()
	}

}

func main() {
	fmt.Println("Executing go routine  main")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
