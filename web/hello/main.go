package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	html := "<strong>Hello World</strong>"
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", home)

	log.Println("starting server on web port 8080")
	http.ListenAndServe(":8080", nil)
}
