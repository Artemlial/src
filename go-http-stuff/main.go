package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "hello!")
}

func formHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Something wrong with your input: %v", err)
	}
	name := r.FormValue("Name")
	info := r.FormValue("Submit info")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Your Submit: %s", info)
}

func main() {
	fileServer := http.FileServer(http.Dir("./go-http-stuff/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/form", formHandle)

	fmt.Printf("starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
