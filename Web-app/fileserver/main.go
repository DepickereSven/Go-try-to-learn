package main

import (
	"net/http"
)

func main() {
	//http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}


