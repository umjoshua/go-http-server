package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" || r.Method != "GET" {
		return
	}
	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/form.html")
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error")
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "name: %s\nemail: %s\n", name, email)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error")
	}
}
