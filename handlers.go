package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	storage "reverse_proxy/storage"
	"strings"
)

func setBadResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func PostRegClient(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	log.Printf(r.RequestURI)
	if port, ok := query["port"]; ok {

		ip := strings.Split(r.RemoteAddr, ":")[0]
		host := ip + ":" + port[0]

		log.Print("New client reg " + host)
		storage.SetClient(host)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No port param"))
	}
}

func GetRegClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.GetClients())
}

func GetClientInfo(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get(":id")
	clientHost, ok := storage.GetClient(clientID)
	if !ok {
		setBadResponse(w, "There is no regged client with id "+clientID)
		return
	}
	log.Printf("http://" + clientHost)
	resp, err := http.Get("http://" + clientHost)
	if err != nil {
		setBadResponse(w, "Error from client with id "+clientID+" ; Error :"+err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
