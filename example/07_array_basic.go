package main

import (
    "fmt"
)

func main() {
    var greet [2]string;
    greet[0] = "hello";
    greet[1] = "world";
    fmt.Println(greet[0], greet[1]);
    fmt.Print(greet);

    primes := [6]int {2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}