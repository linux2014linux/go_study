package main

import (
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    var word_map = make(map[string]int)
    var word_slice = strings.Fields(s)
    for _, item := range word_slice {
        _, ok := word_map[item]
        if ok == true {
            word_map[item]++
        } else {
            word_map[item] = 1
        }
    }

    return word_map
}

func main() {
    var names = " zhangsan   lisi  wangwu    maliu    "
    fmt.Println(WordCount(names))
}