// Example of a very simple HTTP server
// Written by: Waqqas Sheikh (https://www.github.com/w-k-s)
// For: Dubai DevFest 2018 (https://www.meetup.com/en-AU/GDG-Dubai/events/253941428/)

package main

import (
	"fmt"
	"io"
	"log"
	"net/http" 
	"strings"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func greetHandler(w http.ResponseWriter, req *http.Request) {

	name := strings.TrimPrefix(req.URL.Path, "/greet/")

	var message = "You didn't tell me your name!\n"

	if len(name) > 0 {
		message = fmt.Sprintf("Hello, %s\n", name)
	}

	io.WriteString(w, message)
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/greet/", greetHandler)
	handler.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
