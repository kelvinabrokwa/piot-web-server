package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "math/rand"
    "time"
    "encoding/json"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/temp", Temperature)
    log.Fatal(http.ListenAndServe(":8080", router))
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
    temps := Temps{
        Temp{Temperature: rand.Float64()},
        Temp{Temperature: rand.Float64()},
    }
    json.NewEncoder(w).Encode(temps)
}
