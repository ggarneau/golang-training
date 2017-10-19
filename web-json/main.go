package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seedboxtech/golang-training/web-json/handler"
)

func main() {

	h := &handler.JsonHandler{} // make handler
	http.Handle("/", h)         // add handler

	fmt.Println("Started server @ localhost:8000")
	err := http.ListenAndServe(":8000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
