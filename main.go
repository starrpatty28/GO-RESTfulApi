package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 4. Define the Structure of my Song
type Song struct {
	Title  string `json:"Title"`
	Artist string `json:"Artist"`
	Song   string `json:"Song"`
}

// 5. Define my Songs as an Array
type Songs []Song

// 6. All Songs Function
func allSongs(w http.ResponseWriter, r *http.Request) {

	// Dummy Data
	songs := Songs{
		Song{Title: "Golanger", Artist: "Google Golang Language", Song: "GoLang #GangGang"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(songs)
}

func testPostSongs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoints work")
}

// 1. Homepage Function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// 2. Handle Request Function
func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/songs", allSongs).Methods("GET")
	myRouter.HandleFunc("/songs", testPostSongs).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// 3. Main Function
func main() {
	handleRequest()
}
