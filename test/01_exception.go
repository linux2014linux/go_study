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