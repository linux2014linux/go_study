package main

import (
    "fmt"
)

func fibonacci() func() int {
    prev, back := 0, 1
    return func () int {
        ret := prev
        prev, back = back, (prev + back)
        return ret
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
