package main

import (
    "fmt"
    "log"
    "net/http"
    "math/rand"
    "time"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/temp", Temperature)
    fmt.Printf("Listening on port: 8080\n")
    log.Fatal(http.ListenAndServe(":8080", cors.Default().Handler(router)))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ayyy wuz good?!")
}

// temperature
type Temp struct {
    Temperature float64 `json:"temp"`
    Time time.Time `json:"time"`
}

type Temps []Temp

func Temperature(w http.ResponseWriter, r *http.Request) {
    var currTemp Temp;
    currTemp.Temperature = rand.Float64() * 100
    currTemp.Time = time.Now()
    json.NewEncoder(w).Encode(currTemp)
}
