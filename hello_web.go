package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, error := http.Get("http://www.google.com")
	if error == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			s := string(body)
			fmt.Fprintf(w, s)
		}
	}

}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
