package main

import (
  "fmt"
)

func main() {
  fmt.Println(c())
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}