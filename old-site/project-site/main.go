package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func newRouter() *mux.Router {
    r := mux.NewRouter()
    staticFileDirectory := http.Dir("./site/")
    staticFileHandler := http.StripPrefix("/site", http.FileServer(staticFileDirectory))
    r.PathPrefix("/site").Handler(staticFileHandler).Methods("GET")

    r.HandleFunc("/hello", helloHandler).Methods("GET")
    r.HandleFunc("/endpoints", getEndpointHandler).Methods("GET")
	r.HandleFunc("/endpoints", createEndpointHandler).Methods("POST")
    r.HandleFunc("/results", inputHandler).Methods("GET")
    r.HandleFunc("/path", getPathHandler).Methods("GET")
    return r
}

func main() {
    r := newRouter()
    http.ListenAndServe(":8080", r)
}
