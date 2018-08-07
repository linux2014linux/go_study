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

