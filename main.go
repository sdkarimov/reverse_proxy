package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
)

var mux *pat.Router = pat.New()

func init() {
	mux.Get("/api/reg", PostRegClient)
	mux.Get("/api/clients", GetRegClients)
	mux.Get("/api/client/{id}", GetClientInfo)
}

func main() {
	log.Printf("Starting Server on port 8888")
	log.Fatal(http.ListenAndServe(":8888", mux))

}
