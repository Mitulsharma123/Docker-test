package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Welcome to, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/Hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello")
	})

	log.Fatal(http.ListenAndServe(":8081",nil))
}