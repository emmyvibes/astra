package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	backend "github.com/emmyvibes/paul/backend/pubs"
	"github.com/gorilla/mux"
)

// Sample constants for testing
const sampleDB = "dbOne"
const sampleHash = "92387598ywefkhjabskleugfo8q2p7hb934q7gfp9q374eg0qwf"

// Row represents a single entry in the database
type Row struct {
	Hash   string
	HostId int
	Data   []byte
}

type Message struct {
	Content string `json:"content"`
}

func main() {

	//declare the mux Router r
	r := mux.NewRouter()

	//Assign API endpoints to router r
	r.HandleFunc("/getAllEntries", getAllEntriesHandler).Methods("GET")
	r.HandleFunc("/saveEntry", saveEntryHandler).Methods("GET")
	r.HandleFunc("/deleteEntry", deleteEntryHandler).Methods("GET")
	r.HandleFunc("/createEntry", createEntryHandler).Methods("GET")
	r.HandleFunc("/unsaveEntry", unsaveEntryHandler).Methods("GET")
	r.HandleFunc("/getDeviceList", getDeviceListHandler).Methods("GET")
	r.HandleFunc("/getPairCode", getPairCodeHandler).Methods("GET")
	r.HandleFunc("/usePairCode", usePairCodeHandler).Methods("GET")

	test := backend.GetAllEntries()
	fmt.Println(test)
	//establish server on localhost port 8080
	port := "8080"
	fmt.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

/**
*Handles API request for getting all entries in the database given a db name, response content is either a
*[]Row object containing all entries or an error message
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func getAllEntriesHandler(w http.ResponseWriter, r *http.Request) {

	//Database is fetched from the request query parameters
	db := r.URL.Query().Get("dbName")

	if db == "" {
		//error handler here
	}

	response := backend.GetAllEntries()
	fmt.Println(response)

	// byteString := []byte("abcdefg")
	// sampleRow := &Row{sampleHash, 1, byteString}
	// jsonString, err := json.Marshal(sampleRow)

	// if err == nil {
	// 	x := fmt.Sprintf("Connected to %s, Row: %s", sampleDB, jsonString)
	// 	//Construct and send the response
	// 	message := Message{Content: x}
	// 	responseJSON(w, http.StatusOK, message)
	// } else {
	// 	x := fmt.Sprintf("Connected to %s, Error: %s", sampleDB, err)
	// 	//Construct and send the response
	// 	message := Message{Content: x}
	// 	responseJSON(w, http.StatusOK, message)
	// }
}

/**
*Handles API request for downloading a specific entry from the database given the db name and a hash identifier for the entry.
*Response is either a []Byte object of the entry or an error message
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func saveEntryHandler(w http.ResponseWriter, r *http.Request) {

	hash := r.URL.Query().Get("hash")
	db := r.URL.Query().Get("db")

	if db == "" || hash == "" {
		//error handler here
	}

	hash = sampleHash

	x := fmt.Sprintf("Connected to %s, downloading object with hash %s", sampleDB, hash)
	//Construct and send the response
	message := Message{Content: x}
	responseJSON(w, http.StatusOK, message)

}

/**
*Handles API request for deleting a database entry, given the db name and hash identifier for the entry,
*response is an error (or success) message
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func deleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	db := r.URL.Query().Get("db")

	if db == "" || hash == "" {
		//error handler here
	}

	response := backend.DeleteEntry()
	fmt.Println(response)

	// x := fmt.Sprintf("Connected to %s, deleted file %s", db, sampleHash)

	// message := Message{Content: x}
	// responseJSON(w, http.StatusOK, message)
}

/**
*Handles API request for inserting a database entry, given the db name and hash identifier for the entry,
*response is an error (or success) message
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func createEntryHandler(w http.ResponseWriter, r *http.Request) {

	dsn := r.URL.Query().Get("dsn")
	hash := r.URL.Query().Get("hash")

	if dsn == "" || hash == "" {
		//error handler here
	}

	response := backend.CreateEntry(dsn, hash)
	fmt.Println(response)
	// x := fmt.Sprintf("Connected to %s, added file %s", db, sampleHash)

	// message := Message{Content: x}
	responseJSON(w, http.StatusOK, "response")
}

/**
*Handles API request for removing a local copy and updating the database that the file is no longer stored on that device
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func unsaveEntryHandler(w http.ResponseWriter, r *http.Request) {

	//hash := r.URL.Query().Get("hash")

	//hash = sampleHash

	response := backend.UnsaveEntry()
	fmt.Println(response)
	// x := fmt.Sprintf("unsaved object with hash %s", hash)
	// //Construct and send the response
	// message := Message{Content: x}
	// responseJSON(w, http.StatusOK, message)

}

/**
*Handles API request for getting a list of all devices paired with the current device
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func getDeviceListHandler(w http.ResponseWriter, r *http.Request) {

	response := backend.GetDeviceList()
	fmt.Println(response)
	//respond with all devices

}

/**
*Handles API request for generating a code to pair with another device
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func getPairCodeHandler(w http.ResponseWriter, r *http.Request) {

	//connect to yggdrasil and get a pair code

}

/**
*Handles API request for using a provided code to pair with another device
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Struct http.Request r ..............built in http request object, net/http package
*@Return:
*Void
**/
func usePairCodeHandler(w http.ResponseWriter, r *http.Request) {

	//use a pair code

}

/**
*Writes JSON data to the response writer
*@Params:
*Interface http.ResponseWriter w .....built in response interface, net/http package
*Int status ............................................................status code
*Interface payload ..........................the payload to be sent in the response
*@Return:
*Void
**/
func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
	// Set content type as JSON
	w.Header().Set("Content-Type", "application/json")
	// Write status code to header
	w.WriteHeader(status)

	// Convert payload to JSON and write to response
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
