package main

import (
    "fmt"
    "os"
    "strconv"
    "math/rand"
    "encoding/csv"
    "skunz42/red-black-tree/src/rbt"
)

//TODO 0. Insert
//TODO 1. Delete
//TODO 2. Option choice loop
//TODO 3. Insert city
//TODO 4. Remove city
//TODO 5. Export to CSV

func readFromCsvFile(filename string) ([][]string) {
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println("Unable to read file")
        os.Exit(1)
    }

    defer f.Close()
    csv_reader := csv.NewReader(f)
    records, err := csv_reader.ReadAll()
    if err != nil {
        fmt.Println("Unable to parse file")
        os.Exit(1)
    }

    return records
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please enter a filename");
        os.Exit(1);
    }

    csv_filename := os.Args[1]
    rows := readFromCsvFile(csv_filename)
    tree := rbt.Tree{Root : nil}

    city_id_map := make(map[string]int)

    for i := range(rows) {
        j := rand.Intn(i+1)
        rows[i], rows[j] = rows[j], rows[i]
    }

    for r := range(rows) {
//        fmt.Println(rows[r][0])
        rbt.Insert(&tree, rows[r])
        id_int, _ := strconv.Atoi(rows[r][1])
        city_id_map[rows[r][0]] = id_int
    }
//    fmt.Println("----------")
//    rbt.PrintLevelOrder(tree.Root)

    rbt.Delete(&tree, city_id_map, "Binghamton, NY")
    rbt.PrintLevelOrder(tree.Root)
}
