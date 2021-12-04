package main

import (
        "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
        readings := readFile()
        fmt.Println(numberOfIncreases(readings))
        fmt.Println(numberOfIncreaseWindows(readings))
}

func numberOfIncreases(readings []int) int {
        increases := 0
        for i := 1; i < len(readings); i++ {
                if readings[i] > readings[i - 1] {
                        increases++
                }
        }
        return increases
}

func numberOfIncreaseWindows(readings []int) int {
        increases := 0
        for i := 3; i < len(readings); i++ {
                lastWindow := readings[i - 3] + readings[i - 2] + readings[i - 1]
                currentWindow := readings[i - 2] + readings[i - 1] + readings[i]
                if currentWindow > lastWindow {
                        increases++
                }
        }
        return increases
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func readFile() []int {
        data, err := os.ReadFile("./input.txt")
        check(err)
        trimmedData := strings.TrimSuffix(string(data), "\n")
        stringNumbers := strings.Split(trimmedData,"\n")
        numbers := make([]int, len(stringNumbers))
        for i, v := range stringNumbers {
                num, err := strconv.Atoi(v)
                check(err)
                numbers[i] = num
        }
        return numbers
}

