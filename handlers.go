package main

import (
	"encoding/json"
	"log"
	"net/http"
	storage "reverse_proxy/storage"
	"strings"
)

func PostRegClient(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	log.Printf(r.RequestURI)
	if port, ok := query["port"]; ok {

		ip := strings.Split(r.RemoteAddr, ":")[0]
		host := ip + ":" + port[0]

		log.Print("New client reg " + host)
		storage.SetClient(host)
		log.Print(storage.GetClients())
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No port param"))
	}
}

func GetRegClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.GetClients())
}
