package external

import (
    "fmt"
//    "os"
    "skunz42/shortest-path/src/path"
    "skunz42/shortest-path/src/inputs"
    "skunz42/shortest-path/src/database"
)

func ShortestPath(start_city string, end_city string) [][]float32 {
    city_rows := inputs.Read_From_Csv("../../data/cleancities.csv")
    edge_rows := inputs.Read_From_Csv("../../data/edges2.csv")

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

    for i := range(all_cities) {
        path.Make_Adj_List(all_cities[i], all_edges)
    }

    if client != nil {
        float_path = database.Read(client, ctx, search_id)
        if float_path != nil {
            in_db = true
        }
    }

    if in_db == false {
        float_path = path.FindPath(start_city, end_city, all_cities, all_edges)
    }

    if (client != nil && in_db == false) {
        database.Write(search_id, float_path, client, ctx)
    }

    database.Disconnect(client, ctx)
    fmt.Println(float_path)
    return float_path
}
