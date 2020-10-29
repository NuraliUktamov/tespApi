package main

import (
	"log"

	"net/http"
	"github.com/testApi/handlers"
	_ "github.com/lib/pq" // here

)
func main() {
	http.HandleFunc("/get", handlers.GETHandler)
	http.HandleFunc("/insert", handlers.POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
