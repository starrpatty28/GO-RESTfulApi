package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 4. Define the Structure of my Articles
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 5. Define my Artiles as an Array
type Articles []Article

// 6. All Articles Function
func allArticles(w http.ResponseWriter, r *http.Request) {

	// Dummy Data
	articles := Articles{
		Article{Title: "Golang", Desc: "Google Golang Language", Content: "Hello World GoLang #GangGang"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// 1. Homepage Function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// 2. Handle Request Function
func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// 3. Main Function
func main() {
	handleRequest()
}
