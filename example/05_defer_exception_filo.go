package main

import (
  "fmt"
)

/**
 * 包含defer的代码出现异常时,只有出现在异常之前的defer会被执行
 * 这个与filo略有出入
 */
func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in function f", r)
    }
  }()
  score_users := []int{234, 2393, 8993, 88973}
  printArr(score_users)
}

func printArr(score_users []int) {
  defer func() {
    fmt.Println()
    fmt.Println("ergodic slice before.")
  }()

  for i := 0; i < len(score_users) + 1; i++ {
    fmt.Print(score_users[i], " ")
  }

  defer func() {
    fmt.Println()
    fmt.Println("ergodic slice after.")
  }()
}
