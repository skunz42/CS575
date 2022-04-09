package main

import (
    "fmt"
    "os"
    "skunz42/shortest-path/src/path"
    "skunz42/shortest-path/src/inputs"
    "skunz42/shortest-path/src/database"
)

func main() {
    if len(os.Args) != 5 {
        fmt.Println("Please enter a filename for cities and edges");
        os.Exit(1);
    }

    cities_csv := os.Args[1]
    edges_csv := os.Args[2]

    start_city := os.Args[3]
    end_city := os.Args[4]

    city_rows := inputs.Read_From_Csv(cities_csv)
    edge_rows := inputs.Read_From_Csv(edges_csv)

    all_cities := make([]*path.Node, 0)
    all_edges := make([]*path.Edge, 0)

    all_cities = path.Make_Cities(city_rows, all_cities)
    all_edges = path.Make_Edges(edge_rows, all_edges, all_cities)

    start_id := path.NameId(start_city, all_cities)
    end_id := path.NameId(end_city, all_cities)

    var search_id string

    if end_id < start_id {
        search_id = end_id + start_id
    } else {
        search_id = start_id + end_id
    }

    float_path := make([][]float32, 0)
    in_db := false

    client, ctx := database.Connect()

//    for i := range(all_cities) {
//        path.Print_City(all_cities[i])
//    }

    for i := range(all_cities) {
        path.Make_Adj_List(all_cities[i], all_edges)
//        path.Print_Adj_List(all_cities[i])
    }

    if client != nil {
        float_path = database.Read(client, ctx, search_id)
        if float_path != nil {
            in_db = true
            fmt.Println(float_path)
        }
    }

    if in_db == false {
        float_path = path.FindPath(start_city, end_city, all_cities, all_edges)
//        if len(float_path) > 0 {
//            for i := len(float_path)-1; i >= 0; i-- {
//                fmt.Printf("%.6f, %.6f\n", float_path[i][0], float_path[i][1])
//            }
//        }
    }

    if (client != nil && in_db == false) {
        database.Write(search_id, float_path, client, ctx)
    }

    database.Disconnect(client, ctx)
}
