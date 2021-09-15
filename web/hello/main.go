package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World Quinn")
}

func main() {
	http.HandleFunc("/", home)

	log.Println("starting server on web port 8080")
	http.ListenAndServe(":8080", nil)
}
