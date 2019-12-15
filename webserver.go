package main

import (
	"fmt"
	"net/http"
)

//Article describes a article entity
type Article struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

//Version defines the application current version
type Version struct {
	Version int `json:"version"`
}

func getVersion(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%v\n", Version{Version: 1})
}
func getArticles(response http.ResponseWriter, request *http.Request) {

	articles := []Article{
		Article{Name: "Sample", Desc: "Sample Description"},
	}
	fmt.Fprintf(response, "%v\n", articles)
}

func registerHandlerForApplication() {
	http.HandleFunc("/version", getVersion)
	http.HandleFunc("/articles", getArticles)
	http.ListenAndServe(":8080", nil)
}

func main() {
	registerHandlerForApplication()
}
