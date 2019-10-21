#defer
```go
参考
https://blog.csdn.net/zzhongcy/article/details/50395415
```
##特性
###外层函数返回之后执行
```go
package main

import (
  "fmt"
)

func main() {
  defer fmt.Println("world!")
  fmt.Println("Hello")
}
```
输出
```go
Hello
world!
```  
###推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
```go
package main

import (
  "fmt"
)

func main() {
  var age int = 10
  defer func(param int) {
    fmt.Println(param)
  }(age)

  age = age + 1
  fmt.Println(age)
}
```
输出
```go
11
10
```

###推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

```go
package main

import (
  "fmt"
)

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
```
输出:
```go
1st
2nd
3rd
```
###所在方法出现异常依然执行
```go
package main

import (
  "fmt"
)

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
    fmt.Println("\n ergodic slice before.")
  }()

  for i := 0; i < len(score_users) + 1; i++ {
    fmt.Print(score_users[i], " ")
  }

  defer func() {
    fmt.Println("\n ergodic slice after.")
  }()
}
```
输出
```go
234 2393 8993 88973 
 ergodic slice before.
Recovered in function f runtime error: index out of range
```
解析在发生索引越界时,栈中的 defer 函数从下往上依次是: "Recovered in function f", "\n ergodic slice before.",此时 "\n ergodic slice after." 还未入栈.
defer必须在抛出异常之前出现,才会被执行.

##defer应用场景
###安全的资源回收
```go
func set(mu *sync.Mutex, arr []int, i, v int) {
  mu.Lock()
  // do something
  arr[i] = v
  // do something
  mu.Unlock()
}
```
如上代码,设置如果Lock和Unlock之间发生异常,那么Unlock将不会被执行到.使用defer将可以避免:
```go
func set(mu *sync.Mutex, arr []int, i, v int) {
  mu.Lock()
  defer mu.Unlock()
  // do something
  arr[i] = v
  // do something
}
```
无论函数内发生什么,最终锁依然会被打开.
###异常捕获
panic抛出异常,recover捕获异常.recover必须在defer语句中使用,直接调用recover无效.defer是异常框架的组成部分.
```
panic 是用来表示非常严重的不可恢复的错误的。在Go语言中这是一个内置函数，接收一个interface{}类型的值作为参数。panic 的作用就像我们平常接
触的异常。不过Go可没有try…catch，所以，panic一般会导致程序挂掉（除非recover）。所以，Go语言中的异 常，那真的是异常了。你可以试试，调用
panic看看，程序立马挂掉，然后Go运行时会打印出调用栈。但是，关键的一点是，即使函数执行的时候 panic了，函数不往下走了，运行时并不是立刻向上
传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时 候 defer 有点类似 try-catch-finally 中的 finally。
```
```
panic的函数并不会立刻返回，而是先defer，再返回，如果有办法将panic捕获到，并阻止panic传递，就正常处理，如果没有没有捕获，程序直接异常终止.
一旦panic，逻辑就会走到defer那，那我们就在defer那等着，调用recover函 数将会捕获到当前的panic（如果有的话），被捕获到的panic就不会向上传
递了.
```
```
不过要注意的是，recover之后，逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回。
```
```go
package main

import (
  "fmt"
)

func main() {
  f()
  fmt.Println("Return normally from function f.")
}

func f() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in function f", r)
    }
  }()

  fmt.Println("Start calling function g")
  g()
  // 上面g()抛出异常,因此下面不会继续执行
  fmt.Println("Finish function g")
}

func g() {
  panic("ERROR")
}
```
###修改函数返回值
```go
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```
如上c函数返回2.
