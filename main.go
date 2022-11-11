package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "This method is not supported in this page", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello Buddy")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/form" {
	// 	http.Error(w, "404 not found", http.StatusNotFound)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request Successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "Address = %s", address)
}

func main() {
	httpServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", httpServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting a Web server at port 7999\n")

	if err := http.ListenAndServe(":7999", nil); err != nil {
		log.Fatal(err)
	}
}
