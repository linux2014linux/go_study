package main

import (
    "fmt"
)

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])

    var tmp = map[string]Vertex{
        "Bell Labs": Vertex{
            40, -74,
        },
        "Google": Vertex{
            37, -122,
        },
    }

    /**
    var tmp = map[string]Vertex{
        "Bell Labs": {40, -74,},
        "Google": {37, -122},
    }
     */
    fmt.Println(tmp)
}