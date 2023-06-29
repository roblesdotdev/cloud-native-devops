package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}


func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press CTRL+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}