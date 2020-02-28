package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
        "encoding/json"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome home!")
}

type event struct {
    ID          string `json:"id"`
    Title       string `json:"message"`
}

type allEvents []event

var events = allEvents{
    {
      ID:          "1",
      Title:       "Hello World",
    },
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(events[0])
}




func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
        router.HandleFunc("/message", getAllEvents).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

