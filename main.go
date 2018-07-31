package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Message struct {
	Body string `json:"body,omitempty"`
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var message = Message{Body: "Hello, world!"}

	json.NewEncoder(w).Encode(message)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/v1/message", GetMessage).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}
