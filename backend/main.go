package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emmyvibes/paul/backend/pubs"
	"github.com/gorilla/mux"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	fmt.Println("SUBPROCESS DESU")

	client := pubs.DBConnect("http://192.168.10.40:5984")
	fmt.Println(client)
	r := mux.NewRouter()

	//Assign API endpoints to router r
	r.HandleFunc("/addPeer", addPeerHandler).Methods("GET")
	r.HandleFunc("/getMyIP", getIPHandler).Methods("GET")

	//establish server on localhost port 8080
	port := "8081"
	fmt.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

func addPeerHandler(w http.ResponseWriter, r *http.Request) {

	//add a peer to the database

}

func getIPHandler(w http.ResponseWriter, r *http.Request) {

	//return my yggdrasil IP

}
