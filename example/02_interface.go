package main

import (
    "fmt"
)

type Animal interface {
    run() string
    eat() string
}

type Cat struct {
}

func (cat Cat) run() string {
    return "cat run"
}

func (cat Cat) eat() string {
    return "cat eat"
}

func main() {

    var ani_if Animal
    ani_if = new(Cat)
    fmt.Println(ani_if.run())
    fmt.Println(ani_if.eat())

}
