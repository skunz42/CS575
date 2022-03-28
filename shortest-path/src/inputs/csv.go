package inputs

import (
    "fmt"
    "os"
    "encoding/csv"
)

func Read_From_Csv(filename string) ([][]string) {
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
