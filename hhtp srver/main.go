package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
	Structs are useful to group similar
	data together. Like a book or a person
	for example.

	This is called an Aggregate Data Type.

	The `json:"var_name"` will set the var_name
	in the json.  So the var 'title' will access
	the value
*/
type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// Setting content-type isn't always important
	// In this case the browser would automatically
	// know the content type as text/html
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: steelblue'>Hello</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Where text/html wasn't important to set
	// Setting the type for JSON is important
	w.Header().Set("Content-Type", "application/json")

	book := Book { Title: "The Gunslinger", Author: "Stephen King", Pages: 304 }

	// we pass the writer, because it will use our response writer to
	// send the json.
	// We encode the book to JSON with Encode()
	json.NewEncoder(w).Encode(book)

}

func main() {
	// The endpoints.  This is technically pattern matching
	// So when a client requests a URL, it checks if there
	// is any matching patterns that we have set here
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)

	log.Fatal(http.ListenAndServe(":5100", nil))
}
