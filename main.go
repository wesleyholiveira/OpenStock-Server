package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/wesleyholiveira/handlers"
	"github.com/wesleyholiveira/middlewares"
)


func main() {
	// retrieve instance router of gorilla
	r := mux.NewRouter()

	// register the routes
	r.HandleFunc("/", handlers.MainHandler)

	// register the Handler
	http.Handle("/", middlewares.GeneralMiddleware(r))

	log.Fatal(http.ListenAndServe(":8080", nil))
}