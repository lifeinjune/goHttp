package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type NewsAggPag struct { //
	Title string
	News  string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/agg/", newsHandler)
	log.Fatal(http.ListenAndServe(":12345", mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	n := NewsAggPag{Title: "good title", News: "greatss news"}
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, n)
}
