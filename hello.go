package main

import (
    "log"
    "net/http"
)

func main() {
    // h := indexHandler{}
    h := http.HandlerFunc(mux)
    err := http.ListenAndServe(":8900", h)
    log.Println(err)
}

func mux(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
        case "/":
            indexHandler(w, r)
        case "/about":
            aboutHandler(w, r)
        default:
            notFoundHandler(w, r)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("About Page"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("404 Page Not Found"))
}