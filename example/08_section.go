package main

import (
  "fmt"
)

func main() {
  books := []string{
    "孔子",
    "孟子",
    "老子",
    "庄子",
  }
  fmt.Println(books)

  book_part1 := books[0:2]
  book_part2 := books[1:4]

  fmt.Println(book_part1)
  fmt.Println(book_part2)

  /**
   切片描述数组的一段,修改切片会修改底层的元素
   */
  book_part1[1] = "测试"
  fmt.Println(book_part1)
  fmt.Println(book_part2)
}