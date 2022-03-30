package path

import (
    "os"
    "fmt"
    "container/heap"
)

func FindPath(start_city, end_city string, all_cities []*Node, all_edges []*Edge) ([][]float32) {
    pq := make(PriorityQueue, 1)
    pq_idx := 0

    start := find_node(start_city, all_cities)
    end := find_node(end_city, all_cities)

    if (start == nil || end == nil) {
        fmt.Println("City names not found")
        os.Exit(1)
    }

//    for i := range(all_cities) {
//        all_cities[i].index = i
//        pq[i] = all_cities[i]
//    }

    start.index = pq_idx
    pq[pq_idx] = start

    start.dist = 0.0

    heap.Init(&pq)

//    fmt.Println(pq_idx)

    for pq.Len() > 0 {
        min_node := heap.Pop(&pq).(*Node)
        if (min_node.name == end.name) {
            break
        }
        adj_list := min_node.adj_list
        for i := range(adj_list) {
            neighbor := find_node(adj_list[i].name, all_cities)
            edge := find_edge(min_node, neighbor, all_edges)
            alt := min_node.dist + edge.weight
            if alt < neighbor.dist {
                neighbor.dist = alt
                neighbor.prev = min_node
                if !contains(pq, neighbor) {
                    heap.Push(&pq, neighbor)
                }
                pq.update(neighbor)
            }
        }
    }

//    for i := range(all_cities) {
//        fmt.Printf("%s\t%f\n", all_cities[i].name, all_cities[i].dist)
//    }

    iter := end

    float_path := make([][]float32, 0)

    if end.prev == nil {
        fmt.Println("No path exists!")
        return float_path
    }

    if end != nil && iter != nil {
        path := make([]*Node, 0)
        for iter.name != start.name {
            path = append(path, iter)
            coords := []float32{iter.lat, iter.lng}
            float_path = append(float_path, coords)
            if iter != nil {
                iter = iter.prev
            } else {
                break
            }
        }
        path = append(path, start)
        coords := []float32{start.lat, start.lng}
        float_path = append(float_path, coords)
        for i := len(path)-1; i >= 0; i-- {
            fmt.Println(path[i].name)
        }
    }
    return float_path
}
