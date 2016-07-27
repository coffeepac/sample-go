package main

import (
  "fmt"
  "time"
  http "net/http"
)

func statusServer() {
    http.HandleFunc("/ping", pingHandle)
    http.HandleFunc("/health", healthHandle)

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

func main() {
    statusServer()
}