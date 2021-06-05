package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article ...
// Much more like an object in JS
type Article struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

// Articles ...
// Much more like array in JS
var Articles []Article

// Homepage function ...
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomoPage!")
	fmt.Println("Endpoint Hit: HomePage")
}

// Articles function ...
// to return all articles
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// Handle request function ...
func handleRequest() {
	http.HandleFunc("/", homePage)

	// add our articles route and map it to our
	// returnAllArticles function like so
	http.HandleFunc("/articles", returnAllArticles)

	// log http request to console
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	// create articles
	Articles = []Article{
		{
			Title:  "Hello World My Lord",
			Author: "Abrakdabra",
			Link:   "https://google.com",
		},
		{
			Title:  "Be My Lover and Cover Stricker at the Locker",
			Author: "Romansa Muda Terasa Lupa",
			Link:   "https://duckduckgo.com",
		},
	}

	// call handleRequest function
	handleRequest()
}
