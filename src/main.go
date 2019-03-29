package main

import (
	//"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type newsAgg struct {
	Title string
	News  string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/agg/", newsHandler)
	mux.HandleFunc("/coin/", coinHandler)
	log.Fatal(http.ListenAndServe(":12345", mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	n := newsAgg{Title: "good title", News: "greatss news"}
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, n)
}

func coinHandler(w http.ResponseWriter, r *http.Request) {
	//var dat map[string]interface{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "AUD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "4476194f-c711-421d-8441-c4b5fc2c3bac")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(respBody))

}
