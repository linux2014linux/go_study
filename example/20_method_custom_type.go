package main

import (
    "fmt"
)

type MyFloat float64;

func (f MyFloat) FAbs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func (i int64) IAbs() int64 {
  if i < 0 {
    return int64(-i)
  }
  return int64(i)
}

func main() {
  f1 := MyFloat(-1.8923)
  f2 := MyFloat(3.14)
  fmt.Println(f1.FAbs())
  fmt.Println(f2.FAbs())
}
