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

//    path.Print_City(all_cities[0])
//    path.Print_Edge(all_edges[0])

//    path.FindPath()
}
