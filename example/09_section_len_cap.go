package main

import (
  "fmt"
)

func main() {
  fmt.Println("1-------------------")
  stus := []int{1, 3, 5, 7, 9, 11, 13, 15}
  printSectionMeta(stus)

  // 此处将切片len收缩到0
  stus = stus[:0]
  printSectionMeta(stus)

  // 再将切片扩充到len为4
  stus = stus[:4]
  printSectionMeta(stus)

  // 因为之前切片的len已经扩充到4,此处[2:]==>[2:4]
  stus = stus[2:]
  printSectionMeta(stus)

  fmt.Println();
  fmt.Println("2-------------------")
  var test []int
  fmt.Println(test, len(test), cap(test))
  if test == nil {
    fmt.Println("test is nil!")
  }

  fmt.Println();
  fmt.Println("3-------------------")
  a := make([]int, 5)
  printSlice("a", a)

  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  // :后半部分不写,默认表示len,而非cap
  d := c[2:5]
  printSlice("d", d)

  fmt.Println();
  fmt.Println("4-------------------")
  twoDSection()
}

func twoDSection() {
  tbl := [][]string{
    []string{"a", "h", "i"},
    []string{"b", "g"},
    []string{"c", "d", "e", "g"},
  }

  for i := 0; i < len(tbl); i++ {
    var item []string = tbl[i];
    for j := 0; j < len(item); j++ {
      fmt.Printf("%s\n", item[j])
    }
  }

  // fmt.Println(tbl[0])
  // fmt.Println(tbl)
}

func printSectionMeta(s []int) {
  fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}