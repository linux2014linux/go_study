package main

import (
  "fmt"
)

/**
 * defer filo特征
 */
func main() {
  defer func() {
    fmt.Println("3st")
  }()

  defer func() {
    fmt.Println("2st")
  }()

  defer func() {
    fmt.Println("1st")
  }()
}