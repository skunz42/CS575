package main

import (
    "fmt"
    "os"
    "strconv"
    "bufio"
    "math/rand"
    "encoding/csv"
    "skunz42/red-black-tree/src/rbt"
)

func write_to_csv(root *rbt.Node) {
    in := bufio.NewReader(os.Stdin)

    fmt.Println("Enter a file name (no extension):")
    name, _ := in.ReadString('\n')
    fullname := name[:len(name)-1] + ".csv"

    f, err := os.Create(fullname)
    defer f.Close()

    if err != nil {
        fmt.Println("Error creating CSV")
        return
    }

    var cities []rbt.City
    rbt.WriteInorder(root, &cities)
    w := csv.NewWriter(f)
    defer w.Flush()

    for i := range(cities) {
        row := rbt.Export(cities[i])
        if err := w.Write(row); err != nil {
            fmt.Println("Error writing row to file")
        }
    }
}

func read_from_csv(filename string) ([][]string) {
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

func user_input_insert(tree *rbt.Tree, ids map[string]int) {
    var name string
    var id string
    var id_int int
    var population string
    var lat string
    var lng string
    var err error

    in := bufio.NewReader(os.Stdin)

    valid_input := false
    for !valid_input {
        err = nil
        fmt.Println("Enter a city name:")
        name, _ = in.ReadString('\n')
        name = name[:len(name)-1]

        fmt.Println("Enter an ID:")
        id, _ = in.ReadString('\n')
        id_int, err = strconv.Atoi(id[:len(id)-1])
        if err != nil {
            fmt.Println("Invalid ID")
            continue
        }

        fmt.Println("Enter a population:")
        population, _ = in.ReadString('\n')
        _, err = strconv.Atoi(population[:len(population)-1])
        if err != nil {
            fmt.Println("Invalid population")
            continue
        }

        fmt.Println("Enter a latitude:")
        lat, _ = in.ReadString('\n')
        _, err = strconv.ParseFloat(lat[:len(lat)-1], 32)
        if err != nil {
            fmt.Println("Invalid latitude")
            continue
        }

        fmt.Println("Enter a longitude:")
        lng, _ = in.ReadString('\n')
        _, err = strconv.ParseFloat(lng[:len(lng)-1], 32)
        if err != nil {
            fmt.Println("Invalid longitude")
            continue
        }
        valid_input = true
    }

    rows := []string{name, id, population, lat, lng}
    ids[rows[0]] = id_int
    rbt.Insert(tree, rows)
}

func user_input_delete(tree *rbt.Tree, city_id_map map[string]int) {
    var city string

    valid_input := false
    in := bufio.NewReader(os.Stdin)

    for !valid_input {
        fmt.Println("Enter a city to delete")
        city, _ = in.ReadString('\n')

        if _, ok := city_id_map[city[:len(city)-1]]; ok {
            rbt.Delete(tree, city_id_map, city[:len(city)-1])
            valid_input = true
        } else {
            fmt.Println("City not found in tree")
        }
    }
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please enter a filename");
        os.Exit(1);
    }

    csv_filename := os.Args[1]
    rows := read_from_csv(csv_filename)
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

//    rbt.Delete(&tree, city_id_map, "Danville, IL")
//    rbt.PrintLevelOrder(tree.Root)

    for {
        var user_input string

        fmt.Println("Please enter a choice: ")
        fmt.Println("1 - Print tree")
        fmt.Println("2 - Insert")
        fmt.Println("3 - Delete")
        fmt.Println("4 - Export to CSV")
        fmt.Println("5 - Quit")

        fmt.Scanf("%s", &user_input)
        fmt.Println("You chose: " + user_input)

        if user_input[0] == '1' {
            rbt.PrintLevelOrder(tree.Root)
        } else if user_input[0] == '2' {
            user_input_insert(&tree, city_id_map)
        } else if user_input[0] == '3' {
            user_input_delete(&tree, city_id_map)
        } else if user_input[0] == '4' {
            write_to_csv(tree.Root)
        } else if user_input[0] == '5' {
            break
        } else {
            fmt.Println("Unrecognized command")
        }
    }
    fmt.Println("Goodbye")
}
