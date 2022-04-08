package main

import (
    "fmt"
    "net/http"
    "skunz42/shortest-path/src/inputs"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    all_cities := inputs.Read_From_Csv("../../data/cleancities.csv")
    fmt.Println(all_cities[10][0])

    http.ListenAndServe(":80", nil)
}
