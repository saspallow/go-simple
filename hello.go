package main

import (
    "log"
    "net/http"
)

func main() {
    // h := indexHandler{}
    // h := http.HandlerFunc(mux)
    // err := http.ListenAndServe(":8900", h)
    // log.Println(err)

    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/about", aboutHandler)
    err := http.ListenAndServe(":8900", mux)
    log.Println(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("About Page"))
}