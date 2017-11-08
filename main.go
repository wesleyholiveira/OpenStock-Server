package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	. "github.com/wesleyholiveira/handlers"
	. "github.com/wesleyholiveira/middlewares"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	APP_ADDR = ""
	APP_PORT = 8080
)


func main() {
	// full address
	var fullAddr string = fmt.Sprintf("%s:%d", APP_ADDR, APP_PORT)

	// retrieve instance router of gorilla
	r := mux.NewRouter()

	log.Println("Iniciando o servidor (" + fullAddr + ")")

	db, err := gorm.Open("mysql", "root:zarp001@localhost/openstock")

	if err != nil {
		db.Close()
		log.Fatal("Ocorreu um problema ao conectar ao banco")
	}

	// register the routes
	r.HandleFunc("/", MainHandler)

	// register the Handler
	http.Handle("/", GeneralMiddleware(r))

	log.Fatal(http.ListenAndServe(fullAddr, nil))
}