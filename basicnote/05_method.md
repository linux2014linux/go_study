#方法
* Go没有类.可为结构定义指定的方法.方法是 **带有特殊接收者参数** 的函数.注意接收者是放在func关键字和函数名之间的.
```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
```

* 当然也可以为非结构体方法声明方法,但是注意 **只能为在同一个包内定义的类型的接收者声明方法,不能为其他包内定义的类型(包括int之类的内建类型)的接收者声明方法** --即接收者的类型定义和方法声明必须在同一包内,不能为内建类型声明方法.
如下代码为int64内建类型定义指定的方法,提示错误"cannot define new methods on non-local type int64"
```javascript
func (i int64) IAbs() int64 {
  if i < 0 {
    return int64(-i)
  }
  return int64(i)
}
```
* 指针接收者,可在方法内部修改接收者的值.go里面的指针的使用无疑会提高效率,避免无谓的拷贝.
```go
package main

import "fmt"

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale1(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
    // 接收者为指针,无论调用者是指针还是值，均会被go解释为指针，最终值都会发生变化
    v1 := Vertex{3, 4}
    v1.Scale1(2)
    // 6,8
    fmt.Println(v1)
    
    p1 := &Vertex{3, 4}
    p1.Scale1(2)
    // 6,8
    fmt.Println(p1)

    // 接收者为值,无论调用者是值还是指针，均会被go解释为值，最终值不会被修改
    v2 := Vertex{3, 4}
    v2.Scale2(2)
    // 3,4
    fmt.Println(v2)

    p2 := &Vertex{3, 4}
    p2.Scale2(2)
    // 3,4
    fmt.Println(p2)
}
```



