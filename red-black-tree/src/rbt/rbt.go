package rbt

import (
    "fmt"
)

func Insert(row []string) {
    city := cityFactory(row)
    fmt.Println(city.name)
}
