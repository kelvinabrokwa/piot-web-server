package main

import (
    "fmt"
    "log"
    "net/http"
    "math/rand"
    "time"
    "encoding/json"
    "database/sql"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "temp.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/temp", Temperature)
    router.HandleFunc("/temp/add/{temp}/{time}", AddTemp).Methods("POST")
    fmt.Printf("Listening on port: 80\n")
    log.Fatal(http.ListenAndServe(":80", cors.Default().Handler(router)))
}

func AddTemp(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    time := vars["time"]
    temp := vars["temp"]
    fmt.Println(time, ":", temp)

    // add to database
    _, err := db.Exec("insert into temp(time, temp) values (%s, %s)", time, temp)
    if err != nil { log.Fatal(err) }
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
