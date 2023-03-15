package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/fizzbuzz", fizzBuzzHandler).Methods("POST")

    // Use cors middleware to allow requests from http://localhost:5000
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5000"},
    })
    handler := c.Handler(r)

    port := "3000"
    log.Printf("Starting server on port %s\n", port)
    err := http.ListenAndServe(":"+port, handler)
    if err != nil {
        log.Fatal(err)
    }
}

// define a struct to hold the response message
type Message struct {
    Value interface{} `json:"value"`
}

// handle POST requests to /fizzbuzz
func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
    // read the count parameter from the request body
    var data struct {
        Count int `json:"count"`
    }
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    count := data.Count

    // check if the number is divisible by 3 and/or 5
    var result string

    if count%3 == 0 {
        result += "FIZZ"
    }
    if count%5 == 0 {
        result += "BUZZ"
    }
    if result == "" {
        // if not divisible by 3 or 5, return an empty message
        json.NewEncoder(w).Encode(Message{Value: ""})
        return
    }
    // return the appropriate message
    json.NewEncoder(w).Encode(Message{Value: result})
}
