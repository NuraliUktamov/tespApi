package main

import (
	"log"

	"net/http"
	"github.com/testApi/handlers"

)
func main() {
	http.HandleFunc("/get", GETHandler)
	http.HandleFunc("/insert", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
