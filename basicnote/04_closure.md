#闭包
```go
/**
 官方文档对于闭包的描述是"函数被绑定在变量上" 
 */
package main

import (
    "fmt"
)

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func mian() {
    /**
     每次调用adder创建了独立的sum变量,并为之
     绑定了一个函数.对于该函数的调用,就是对sum
     值的处理.
     pos和neg有各自独立的sum.因此
     pos将输出序列0,1,3,6,10......
     neg将输出序列0,-2,-6,-20......
     */
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i)
        )
    }
}

```