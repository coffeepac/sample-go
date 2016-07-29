package main

import (
  "fmt"
  "time"
  http "net/http"
)

var HealthBadCount = 0

func statusServer() {
    http.HandleFunc("/ping", pingHandle)
    http.HandleFunc("/health", healthHandle)
    http.HandleFunc("/healthBad", healthBadHandle)

    //  create server that doesn't leave things open forever
    s := &http.Server{
            Addr:           ":8080",
            ReadTimeout:    10 * time.Second,
            WriteTimeout:   10 * time.Second,
        }
    s.ListenAndServe()
}

func pingHandle(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "PONG\n")
}

func healthHandle(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "health\n")
}

func healthBadHandle(w http.ResponseWriter, r *http.Request){
    HealthBadCount += 1
    if HealthBadCount > 15 {
        w.WriteHeader(http.StatusInternalServerError)
    }
    fmt.Fprintf(w, "health count %d next\n", HealthBadCount)
}

func main() {
    statusServer()
}