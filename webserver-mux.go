package main

import (
	"fmt"
	"github.com/gorilla/mux"
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

var articles = []Article{
	Article{Name: "Sample", Desc: "Sample Description"},
	Article{Name: "Sample2", Desc: "Sample Description 2"},
	Article{Name: "Sample3", Desc: "Sample Description 3"},
	Article{Name: "Sample4", Desc: "Sample Description 4"},
}

var articleMap map[string]Article = make(map[string]Article)

func getVersion(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%v\n", Version{Version: 1})
}
func getArticles(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%v\n", articles)
}

func insertArticlesIntoMap() {
	for _, article := range articles {
		articleMap[article.Name] = article
	}
}

func getArticleByName(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(response, "%v\n", articleMap[vars["articleName"]])
}

func registerHandlerForApplication() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/version", getVersion).Methods("GET")
	myRouter.HandleFunc("/articles", getArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{articleName}", getArticleByName)
	http.ListenAndServe(":8080", myRouter)
}

func main() {
	insertArticlesIntoMap()
	registerHandlerForApplication()
}
