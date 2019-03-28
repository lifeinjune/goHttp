package main

import (
		"fmt"
		"net/http"
		"log"
)

func main() {
	mux := NewServeMux()
	mux.HandleFunc("/",rootHandler)
	log.Fatal(http.ListenAndServe(":12345",mux))
}

func rootHandler(w http.ResponsWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello world")
}