package main

import (
    "encoding/json"
    "fmt"
	"net/http"
    "math/rand"
    "time"
)

type Cities struct {
	Start string `json:"start"`
	Destination string `json:"destination"`
}

var write_response Cities
var my_num int

// /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

// /endpoints
func getEndpointHandler(w http.ResponseWriter, r *http.Request) {
	myBytes, err := json.Marshal(write_response)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(myBytes)
}

// /endpoints
func createEndpointHandler(w http.ResponseWriter, r *http.Request) {
	endpoints := Cities{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	endpoints.Start = r.Form.Get("start")
	endpoints.Destination = r.Form.Get("destination")

    write_response = endpoints

	http.Redirect(w, r, "/results", http.StatusFound)
}

// /path
func getPathHandler(w http.ResponseWriter, r *http.Request) {
    myBytes, _ := json.Marshal(my_num)
    w.Write(myBytes)
}

// /results
func inputHandler(w http.ResponseWriter, r *http.Request) {
    //Blocking vs non-blocking?
    time.Sleep(10 * time.Second)

    //Replace rand n with path
    my_num = rand.Intn(100)
    http.ServeFile(w, r, "./site/map.html")
}
