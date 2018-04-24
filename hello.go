package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type JSONMessage struct {
	Message string
}

// Server loads on port 8080

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/message", JSONIndex)
	router.HandleFunc("/{urlMessage}", MessageIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", mux.Vars(r)["urlMessage"])
}

func JSONIndex(w http.ResponseWriter, r *http.Request) {
	messageJSON := JSONMessage{Message: "Hello"}

	json.NewEncoder(w).Encode(messageJSON)
}
