package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("starting server on :5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}