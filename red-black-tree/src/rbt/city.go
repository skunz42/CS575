package rbt

import (
    "fmt"
    "strconv"
)

type City struct {
    name string
    lat float32
    lng float32
    id int
    population int
}

func stringToInt(s string) int {
    result, _ := strconv.Atoi(s)
    return result
}

func stringToFloat(s string) float32 {
    result, _ := strconv.ParseFloat(s, 32)
    return float32(result)
}

func cityFactory(row []string) *City {
    city := City{name: row[0], id: stringToInt(row[1]), population: stringToInt(row[2]), lat: stringToFloat(row[3]), lng: stringToFloat(row[4])}
    return &city
}

func Export(city City) []string {
    return []string{city.name, fmt.Sprintf("%d", city.id), fmt.Sprintf("%d", city.population), fmt.Sprintf("%f", city.lat), fmt.Sprintf("%f", city.lng)}
}
