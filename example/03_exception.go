package main

import (
  "fmt"
)

func main() {
  f()
  fmt.Println("Return normally from function f.")
}

func f() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in function f", r)
    }
  }()

  fmt.Println("Start calling function g")
  a := 10
  b := 0
  ret := g(a, b)
  fmt.Println("Finish function g. ret = ", ret)
}

func g(a, b int) (int) {
  return a / b;
}