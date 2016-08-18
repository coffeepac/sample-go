package main

import (
  "os"
  "fmt"
  "time"
  ioutil "io/ioutil"
  http "net/http"
)

// HealthBadCount suck it
var HealthBadCount = 0

func statusServer() {
    http.HandleFunc("/ping", pingHandle)
    http.HandleFunc("/health", healthHandle)
    http.HandleFunc("/healthBad", healthBadHandle)
    http.HandleFunc("/pvDataReturn", pvDataReturn)
    http.HandleFunc("/pvDataSet", pvDataSet)

    //  create server that doesn't leave things open forever
    s := &http.Server{
            Addr:           ":8080",
            ReadTimeout:    10 * time.Second,
            WriteTimeout:   10 * time.Second,
        }
    s.ListenAndServe()
}

func pingHandle(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
        fmt.Fprintf(w, "PONG\n")
    }
}

func healthHandle(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
        fmt.Fprintf(w, "health\n")
    }
}

func pvDataReturn(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
        pvData, err := ioutil.ReadFile("/mnt/cephfs/pvData")
        if err != nil {
            fmt.Fprintf(w, "PV not present\n")
        } else {
            fmt.Fprintf(w, string(pvData[:]) + "\n")
        }
    }
}

func pvDataSet(w http.ResponseWriter, r *http.Request){
    if r.Method == "PUT" {
        bytes := make([]byte, 100)
        _, err := r.Body.Read(bytes)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)            
            fmt.Fprintf(w, "Failed to read PUT data\n")
        }
        
        err = ioutil.WriteFile("/mnt/cephfs/pvData", bytes, os.ModePerm)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "PV not present\n")
        } else {
            fmt.Fprintf(w, "Data saved")
        }
    }
}

func healthBadHandle(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
        HealthBadCount++
        if HealthBadCount > 15 {
            w.WriteHeader(http.StatusInternalServerError)
        }
        fmt.Fprintf(w, "health count %d next\n", HealthBadCount)
    }
}

func main() {
    statusServer()
}