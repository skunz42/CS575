package main

import (
    "fmt"
    "os"
    "skunz42/shortest-path/src/path"
    "skunz42/shortest-path/src/inputs"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Please enter a filename for cities and edges");
        os.Exit(1);
    }

    cities_csv := os.Args[1]
    edges_csv := os.Args[2]

    city_rows := inputs.Read_From_Csv(cities_csv)
    edge_rows := inputs.Read_From_Csv(edges_csv)

    all_cities := make([]*path.Node, 0)
    all_edges := make([]*path.Edge, 0)

    all_cities = path.Make_Cities(city_rows, all_cities)
    all_edges = path.Make_Edges(edge_rows, all_edges, all_cities)

//    for i := range(all_cities) {
//        path.Print_City(all_cities[i])
//    }

    for i := range(all_cities) {
        path.Make_Adj_List(all_cities[i], all_edges)
        path.Print_Adj_List(all_cities[i])
    }

//    path.FindPath()
}
