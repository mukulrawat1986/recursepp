package main

import (
	"log"
	"net/http"
	"github.com/mukulrawat1986/recursepp/db"
	"github.com/mukulrawat1986/recursepp/handlers"
)

func main(){

	db := db.NewMemoryDB()

	// Create a new http.ServeMux, used to register handlers to execute in response to routes
	mux := http.NewServeMux()

	// get the value of a key
	mux.Handle("/get", handlers.GetKey(db))

	// set the value of a key
	mux.Handle("/set", handlers.SetKey(db))

	log.Println("Serving at port 4000")
	
	err := http.ListenAndServe(":4000", mux)

	if err != nil{
		log.Fatal(err)
	}

}