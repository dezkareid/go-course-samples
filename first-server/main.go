package main

import (
	"fmt"
	"net/http"
)

//RootHandler struct
type RootHandler struct{}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello mux from handler")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello mux from function")
}

func main() {
	mux := http.NewServeMux()
	rootHandler := RootHandler{}
	mux.Handle("/", &rootHandler)
	mux.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", mux)
}
