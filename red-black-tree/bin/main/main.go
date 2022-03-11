package main

import (
    "fmt"
    "os"
    "math/rand"
    "encoding/csv"
    "skunz42/red-black-tree/src/rbt"
)

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

    for i := range(rows) {
        j := rand.Intn(i+1)
        rows[i], rows[j] = rows[j], rows[i]
    }

    for r := range(rows) {
        fmt.Println(rows[r][0])
        rbt.Insert(&tree, rows[r])
    }
    fmt.Println("---------")
    rbt.PrintInorder(tree.Root)
}
