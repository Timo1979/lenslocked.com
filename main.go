package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my cool site!</h1>")
}

func main() {
	fmt.Println("App starting")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
