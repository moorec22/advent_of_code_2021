package main

import (
        "fmt"
        "internal/file"
)

func main() {
        instructions := readFile("./sample.txt")
        fmt.Println(instructions)
}

