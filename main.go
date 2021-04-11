package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/pat"
)

var mux *pat.Router = pat.New()
var routes map[string]string

func init() {
	routes = map[string]string{
		"reg":          "/api/reg",
		"clients":      "/api/clients",
		"client/stats": "/api/client/{id}/stats",
	}
	mux.Get(routes["reg"], PostRegClient)
	mux.Get(routes["clients"], GetRegClients)
	mux.Get(routes["client/stats"], GetClientInfo)

	mux.Get("/api_list", func(w http.ResponseWriter, r *http.Request) {
		var resp []string
		for _, v := range routes {
			resp = append(resp, v)
		}
		j, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintln(w, string(j))
	})
}

func main() {
	log.Printf("Starting Server on port 8888")
	log.Fatal(http.ListenAndServe(":8888", mux))

}
