package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "skunz42/shortest-path/src/inputs"
)

type Cities struct {
	Start string `json:"start"`
	Destination string `json:"destination"`
}

var write_response Cities

// /results
func resultHandler(w http.ResponseWriter, r *http.Request) {
    
    http.ServeFile(w, r, "./static/map.html")
}

// /endpoints
func endpointHandler(w http.ResponseWriter, r *http.Request) {
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

// /cities
func cityHandler(w http.ResponseWriter, r *http.Request) {
	myBytes, err := json.Marshal(write_response)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(myBytes)
}

// /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!\n")
}

func newRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/endpoints", endpointHandler).Methods("POST")
    r.HandleFunc("/cities", cityHandler).Methods("GET")
    r.HandleFunc("/results", resultHandler).Methods("GET")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

    return r
}

func main() {
//    all_cities := inputs.Read_From_Csv("../../data/cleancities.csv")
//    all_edges := inputs.Read_From_Csv("../../data/edges2.csv")

    r := newRouter()

    fmt.Println("Made router")

    http.ListenAndServe(":80", r)
}
