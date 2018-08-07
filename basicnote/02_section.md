#切片
切片像是数组的引用,切片本身不存储任何数据,至少描述底层数组的一段数据.更改切片的数据,与切片共享底层的切片数据都会发生变化.是吗?
```go
package main

import (
  "fmt"
)

func main() {
  books := []string{
    "测试",
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

  book_part1[1] = "测试"

  fmt.Println(book_part1)
  fmt.Println(book_part2)
}
```

#切片的元数据
```go
切片长度:切片包含元素的个数
切片容量:切片中第1个元素开始数,到其底层数组末尾元素的数量
```
