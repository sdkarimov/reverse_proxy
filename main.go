package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/reg", PostRegClient)
	http.HandleFunc("/clients", GetRegClients)
	log.Printf("Starting Server ")
	log.Fatal(http.ListenAndServe(":8888", nil))

}
