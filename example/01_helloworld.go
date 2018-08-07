package main

import "fmt"

/* init在main之前执行 */
func init() {
    a, b := 1, 2
    var c = "test"
    // d = "test"
    fmt.Println("init func execute.", a, b, c)
}

func swaq(a *int32, b *int32) {
    var tmp int32 = *a;
    *a = *b;
    *b = tmp;
}

func main() {
    fmt.Println("hi，你好！")

    var student_num int32 = 23
    var charge int32 = 10
    fmt.Println("共计: ", student_num * charge)

    var idx int
    for idx = 0; idx < 1000; idx++ {
        if (idx % 8 == 0) {
            fmt.Print(idx, ", ")
        }
    }
    fmt.Println();

    var x, y int32 = 9, 323
    fmt.Println(x, y);
    swaq(&x, &y)
    fmt.Println(x, y);
}
