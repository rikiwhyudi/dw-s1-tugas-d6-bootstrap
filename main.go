package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/",  func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	route.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Hello gaes"))
	})

	route.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Testing...!"))
	})
	fmt.Println("Server running on port 5000");
	http.ListenAndServe("localhost:5000", route)
}