package main

import (
	"fmt"
	"net/http"
)

// creating a handler func to expose a web session with some html message
func handlerFunc(respWr http.ResponseWriter, req *http.Request) {
	fmt.Fprint(respWr, "<h1>Welcome to my website. This is just the beginning!</h1>")
}

// self explanatory
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}