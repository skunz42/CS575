package path

import (
    "fmt"
    "strconv"
    "math"
)

type Node struct {
    name string
    lat float32
    lng float32
    id int
    population int
    dist float32
    prev *Node
    adj_list []*Node
}

type Edge struct {
    start *Node
    dest *Node
    distance float32
    population int
    weight float32
    start_string string
    dest_string string
}

func stringToInt(s string) int {
    result, _ := strconv.Atoi(s)
    return result
}

func stringToFloat(s string) float32 {
    result, _ := strconv.ParseFloat(s, 32)
    return float32(result)
}

func city_factory(row []string) *Node {
    city := Node{name: row[0], id: stringToInt(row[1]), population: stringToInt(row[2]), lat: stringToFloat(row[3]), lng: stringToFloat(row[4]),
    dist: math.MaxFloat32, prev: nil, adj_list: make([]*Node, 0)}
    return &city
}

func edge_factory(row []string) *Edge {
    edge := Edge{start: nil, dest: nil, distance: stringToFloat(row[2]), population: stringToInt(row[3]), weight: stringToFloat(row[4]),
    start_string: row[0], dest_string: row[1]}
    return &edge
}

func find_node(name string, cities []*Node) (*Node) {
    for i := range(cities) {
        if cities[i].name == name {
            return cities[i]
        }
    }
    return nil
}

func Make_Cities(rows [][]string, cities []*Node) ([]*Node) {
    for i := range(rows) {
        city := city_factory(rows[i])
        cities = append(cities, city)
    }
    return cities
}

func Make_Edges(rows [][]string, edges []*Edge, cities []*Node) ([]*Edge) {
    for i := range(rows) {
        edge := edge_factory(rows[i])
        start := find_node(edge.start_string, cities)
        dest := find_node(edge.dest_string, cities)
        edge.start = start
        edge.dest = dest
        edges = append(edges, edge)
    }
    return edges
}

func Make_Adj_List(city *Node, edges []*Edge) {
    for i := range(edges) {
        if edges[i].start_string == city.name {
            city.adj_list = append(city.adj_list, edges[i].dest)
        }
    }
}

func Print_City(city *Node) {
    fmt.Println("Name: ", city.name)
    /*fmt.Println("Lat: ", city.lat)
    fmt.Println("Lng: ", city.lng)
    fmt.Println("Id: ", city.id);
    fmt.Println("Population: ", city.population)*/
}

func Print_Edge(edge *Edge) {
    Print_City(edge.start)
    Print_City(edge.dest)
    fmt.Println("Distance: ", edge.distance)
    fmt.Println("Population: ", edge.population)
    fmt.Println("Weight: ", edge.weight)
    fmt.Println("Start: ", edge.start_string)
    fmt.Println("End: ", edge.dest_string)
}

func Print_Adj_List(city *Node) {
    fmt.Print("City: " + city.name + ", Adj: ")
    for i := range(city.adj_list) {
        fmt.Print(city.adj_list[i].name + ", ")
    }
    fmt.Println()
}
